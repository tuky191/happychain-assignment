package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/drand/drand/client"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	bip39 "github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
)

func generateCustomHash(methodName string, timestamp int64, value string) string {
	data := fmt.Sprintf("%s:%d:%s", methodName, timestamp, value)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func derivePrivateKeyFromMnemonic(mnemonic string) (*ecdsa.PrivateKey, error) {
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

func loadContractAddresses(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &contractAddresses)
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
func getDrandValue(drandClient client.Client, timestamp int64) ([]byte, error) {
	round := (timestamp - drandGenesis) / drandInterval
	log.Printf("Fetching Drand value for round: %d", round)

	result, err := drandClient.Get(context.Background(), uint64(round))
	if err != nil {
		return nil, err
	}

	return result.Randomness(), nil
}

func keccak256(data []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	return hash.Sum(nil)
}

func calculateSequencerRandomness(drandValue, commitmentValue string) (string, error) {
	drandBytes32 := stringToBytes32(drandValue)
	commitmentBytes32 := stringToBytes32(commitmentValue)

	types := []string{"bytes32", "bytes32"}
	values := []interface{}{drandBytes32, commitmentBytes32}

	return solidityPackedKeccak256(types, values)
}

func calculateSequencerRandomnessHash(randomness string) (string, error) {
	randomnessBytes32 := stringToBytes32(randomness)
	types := []string{"bytes32"}
	values := []interface{}{randomnessBytes32}

	return solidityPackedKeccak256(types, values)
}

func generateSequencerRandom() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(h.Sum(nil))
}

func solidityPackedKeccak256(types []string, values []interface{}) (string, error) {
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

// Function to hash a string to [32]byte using SHA-256
func stringToBytes32(inputString string) [32]byte {

	hash := sha256.New()
	hash.Write([]byte(inputString))
	hashBytes := hash.Sum(nil)

	var hashArray [32]byte
	copy(hashArray[:], hashBytes)

	return hashArray
}

func bufferToHex(buffer []byte) string {
	return "0x" + hex.EncodeToString(buffer)
}
