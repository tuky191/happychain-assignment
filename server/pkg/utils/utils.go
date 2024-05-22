package utils

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	servertypes "server/v0/types/server"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	bip39 "github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

var abiErrorMappings = map[string]string{
	"SequencerEntryAlreadyCommitted":     `[{"inputs":[],"name":"SequencerEntryAlreadyCommitted","type":"error"}]`,
	"SequencerRandomnessNotCommitted":    `[{"inputs":[{"internalType":"uint256","name":"T","type":"uint256"}],"name":"SequencerRandomnessNotCommitted","type":"error"}]`,
	"PrecommitDelayNotPassed":            `[{"inputs":[{"internalType":"uint256","name":"T","type":"uint256"},{"internalType":"uint256","name":"currentBlock","type":"uint256"},{"internalType":"uint256","name":"requiredBlock","type":"uint256"},{"internalType":"uint256","name":"committedBlock","type":"uint256"}],"name":"PrecommitDelayNotPassed","type":"error"}]`,
	"SequencerRandomnessAlreadyRevealed": `[{"inputs":[{"internalType":"uint256","name":"T","type":"uint256"}],"name":"SequencerRandomnessAlreadyRevealed","type":"error"}]`,
	"InvalidRandomnessReveal":            `[{"inputs":[{"internalType":"bytes32","name":"expectedHash","type":"bytes32"},{"internalType":"bytes32","name":"computedHash","type":"bytes32"}],"name":"InvalidRandomnessReveal","type":"error"}]`,
}

func GenerateCustomHash(methodName string, timestamp int64, value string) string {
	data := fmt.Sprintf("%s:%d:%s", methodName, timestamp, value)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func DerivePrivateKeyFromMnemonic(mnemonic string) (*ecdsa.PrivateKey, error) {
	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Derive the master key from the seed
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	// Derive the account key from the master key
	purpose, _ := masterKey.Child(44 + hdkeychain.HardenedKeyStart) // Purpose: BIP44
	coinType, _ := purpose.Child(60 + hdkeychain.HardenedKeyStart)  // CoinType: 60 for Ethereum
	account, _ := coinType.Child(0 + hdkeychain.HardenedKeyStart)   // Account: 0

	// Derive the first address key
	change, _ := account.Child(0) // External chain
	addressIndex, _ := change.Child(0)

	// Get the private key
	privateKey, err := addressIndex.ECPrivKey()
	if err != nil {
		return nil, err
	}

	return crypto.ToECDSA(privateKey.Serialize())
}

func Keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

func CalculateSequencerRandomness(drandValue, commitmentValue string) (string, error) {
	drandBytes32 := StringToBytes32(drandValue)
	commitmentBytes32 := StringToBytes32(commitmentValue)

	types := []string{"bytes32", "bytes32"}
	values := []interface{}{drandBytes32, commitmentBytes32}

	return SolidityPackedKeccak256(types, values)
}

func CalculateSequencerRandomnessHash(randomness string) (string, error) {
	randomnessBytes32 := StringToBytes32(randomness)
	types := []string{"bytes32"}
	values := []interface{}{randomnessBytes32}

	return SolidityPackedKeccak256(types, values)
}

func GenerateSequencerRandom() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(h.Sum(nil))
}

func SolidityPackedKeccak256(types []string, values []interface{}) (string, error) {
	if len(types) != len(values) {
		return "", fmt.Errorf("types and values must have the same length")
	}

	packed, err := solidityPacked(types, values)
	if err != nil {
		return "", err
	}

	hash := crypto.Keccak256(packed)
	return hexutil.Encode(hash), nil
}

// solidityPacked packs the values according to Solidity's abi.encodePacked
func solidityPacked(types []string, values []interface{}) ([]byte, error) {
	arguments := make(abi.Arguments, len(types))
	for i, typ := range types {
		argType, err := abi.NewType(typ, "", nil)
		if err != nil {
			return nil, err
		}
		arguments[i] = abi.Argument{Type: argType}
	}

	packed, err := arguments.PackValues(values)
	if err != nil {
		return nil, err
	}

	return packed, nil
}

// func stringToBytes32(s string) [32]byte {
// 	var b [32]byte
// 	copy(b[:], s)
// 	return b
// }

func To4Bytes(data []byte) [4]byte {
	var b [4]byte
	copy(b[:], data)
	return b
}

// Function to hash a string to [32]byte using SHA-256
func StringToBytes32(inputString string) [32]byte {

	hash := sha256.New()
	hash.Write([]byte(inputString))
	hashBytes := hash.Sum(nil)

	var hashArray [32]byte
	copy(hashArray[:], hashBytes)

	return hashArray
}

func BufferToHex(buffer []byte) string {
	return "0x" + hex.EncodeToString(buffer)
}

func DecodeCustomError(data string) (string, error) {
	dataBytes, err := hex.DecodeString(data[2:]) // assuming data starts with "0x"
	if err != nil {
		return "", err
	}
	methodID := dataBytes[:4]
	dataBytes = dataBytes[4:]
	for name, abiJSON := range abiErrorMappings {
		contractABI, err := abi.JSON(bytes.NewReader([]byte(abiJSON)))
		if err != nil {
			return "", err
		}
		v := make(map[string]interface{})
		abiErr, err := contractABI.ErrorByID(To4Bytes(methodID))
		if err != nil {
			continue
		}

		err = abiErr.Inputs.UnpackIntoMap(v, dataBytes)
		if err != nil {
			continue
		}

		return convertReadableJson(v, name)
	}

	return "Unknown custom error", nil
}

func convertReadableJson(m map[string]interface{}, name string) (string, error) {
	for key, value := range m {
		// Check if the value is an array of 32 bytes (uint8)
		if reflect.TypeOf(value) == reflect.ArrayOf(32, reflect.TypeOf(uint8(0))) {
			byteArray := value.([32]uint8)
			byteSlice := byteArray[:]
			// Convert byte slice to string and trim null bytes
			m[key] = strings.TrimRight(string(byteSlice), "\x00")
		}
	}
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("Error: %s, Values: %s", name, string(jsonBytes)), nil
}

func LoadServerConfig() (*servertypes.Config, error) {

	anvilURL := getEnv("ANVIL_URL", "http://anvil:8545")
	mnemonic := getEnv("MNEMONIC", "")
	if mnemonic == "" {
		return nil, fmt.Errorf("MNEMONIC environment variable is not set")
	}
	privateKey, err := DerivePrivateKeyFromMnemonic(mnemonic)
	if err != nil {
		return nil, fmt.Errorf("Failed to derive private key: %v", err)
	}
	log.Println("Derived private key from mnemonic successfully.")
	delay := getEnvAsInt("DELAY", 9)
	precommitDelay := getEnvAsInt("PRECOMMIT_DELAY", 10)
	drandInterval := getEnvAsInt("DRAND_INTERVAL", 3)
	drandGenesis := getEnvAsInt("DRAND_GENESIS", 1609459200)
	drandChainHash := getEnv("DRAND_CHAIN_HASH", "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971")
	drandURL := getEnv("DRAND_URL", "https://api.drand.sh")
	drandChainHashBytes, err := hex.DecodeString(drandChainHash)
	if err != nil {
		log.Fatalf("Failed to decode Drand chain hash: %v", err)
	}

	contractAddresses, err := LoadContractAddresses("/app/shared/addresses.json")
	if err != nil {
		log.Fatalf("Failed to load contract addresses: %v", err)
	}
	log.Println("Loaded contract addresses successfully.")

	return &servertypes.Config{
		AnvilURL:            anvilURL,
		Delay:               delay,
		PrecommitDelay:      precommitDelay,
		DrandInterval:       drandInterval,
		DrandGenesis:        drandGenesis,
		DrandURL:            drandURL,
		DrandChainHashBytes: drandChainHashBytes,
		ContractAddresses:   *contractAddresses,
		PrivateKey:          privateKey,
	}, nil
}

func LoadContractAddresses(filename string) (*servertypes.ContractAddresses, error) {
	contracts := servertypes.ContractAddresses{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &contracts)
	if err != nil {
		return nil, fmt.Errorf("Unable to load contract addresses: %v", err)
	}
	return &contracts, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultValue int64) int64 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}
