package server

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"server/v0/pkg/contracts"
	"server/v0/pkg/utils"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

type MockJsonError struct {
	err     string
	errCode int
	errData string
}

func (m *MockJsonError) Error() string {
	return m.err
}

func (m *MockJsonError) ErrorCode() int {
	return m.errCode
}

func (m *MockJsonError) ErrorData() interface{} {
	return m.errData
}

func TestGetTransactionReceipt(t *testing.T) {
	// Create Ethereum client
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Transaction hash to check
	txHash := "0x0dd8cf485b7461c66503cf8a6347d5114841be69135eaee9cfa53a430cb839fd"

	// Get transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		t.Fatalf("Failed to get transaction receipt: %v", err)
	}

	if receipt == nil {
		t.Fatalf("Transaction receipt is nil")
	}
	log.Printf("Transaction receipt: %+v", receipt)
}

// solidityPackedKeccak256 hashes the packed Solidity values using Keccak-256

func TestSolidityPackedKeccak256(t *testing.T) {
	input := "0xa8774c23cf38bbba005f2a4318cb56aba6a08c7b4e75dd8b6827e10b188e5400"
	expectedBytes32Hex := "0x767b8f9294d1ba38d120c16ee7c354ec58fd89d8d42cc4a9a3c70e74ed04f0cd"
	expectedHash := "0x30d109baad1c2d0db758043170cee40f9306645f6b8d75355b22a5e62845035f" // Change this to match the expected hash

	bytes32Value := utils.StringToBytes32(input)
	bytes32Hex := utils.BufferToHex(bytes32Value[:])
	assert.Equal(t, expectedBytes32Hex, bytes32Hex, "Bytes32 value should match expected")

	types := []string{"bytes32"}
	values := []interface{}{bytes32Value}
	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomnessHash","type":"bytes32"}],"name":"postRandomnessCommitment","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	fmt.Print(contractABI.Methods)
	hash, err := utils.SolidityPackedKeccak256(types, values)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, expectedHash, hash, "Hash should match expected")
}

// func TestCreateTransactionReveal(t *testing.T) {
// 	// Setup
// 	privateKey, err := crypto.GenerateKey()
// 	assert.NoError(t, err, "Failed to generate private key")

// 	client, err := ethclient.Dial("http://localhost:8545")
// 	assert.NoError(t, err, "Failed to connect to the Ethereum client")

// 	contractAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
// 	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_value","type":"bytes32"}],"name":"revealSequencerRandomness","outputs":[],"payable":false,"stateMutability":"view","type":"function"}]`))
// 	assert.NoError(t, err, "Failed to parse contract ABI")

// 	timestamp := int64(1234567890)
// 	value := "0xaff07feb058c3908335b80f307a0f0b6e45a45ff17437ea0e2ceec940ce9dd65"

// 	// Execute
// 	tx, err := createTransaction(client, privateKey, contractAddress, "revealSequencerRandomness", contractABI, timestamp, value)
// 	assert.NoError(t, err, "Failed to create transaction")

// 	// Validate
// 	assert.NotNil(t, tx, "Transaction should not be nil")
// 	assert.Equal(t, uint64(3000000), tx.Gas())
// 	assert.Equal(t, common.HexToAddress(contractAddress), *tx.To())
// }

// func TestCreateTransactionCommit(t *testing.T) {
// 	// Setup
// 	privateKey, err := crypto.GenerateKey()
// 	assert.NoError(t, err, "Failed to generate private key")

// 	client, err := ethclient.Dial("http://localhost:8545")
// 	assert.NoError(t, err, "Failed to connect to the Ethereum client")

// 	contractAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
// 	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_value","type":"bytes32"}],"name":"postRandomnessCommitment","outputs":[],"payable":false,"stateMutability":"view","type":"function"}]`))
// 	assert.NoError(t, err, "Failed to parse contract ABI")

// 	timestamp := int64(1234567890)
// 	value := "0xa9cc7948e824bf6db251ba79694ea3ff4953d52f66ac2dea7b431f8bab2d070c"

// 	// Execute
// 	tx, err := createTransaction(client, privateKey, contractAddress, "postRandomnessCommitment", contractABI, timestamp, value)
// 	assert.NoError(t, err, "Failed to create transaction")

// 	// Validate
// 	assert.NotNil(t, tx, "Transaction should not be nil")
// 	assert.Equal(t, uint64(3000000), tx.Gas())
// 	assert.Equal(t, common.HexToAddress(contractAddress), *tx.To())
// }

func TestGetRevertReason(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	txHash := common.HexToHash("0xcbfc5a38002fe09909596d813aa9637b62cf217bac4a9439bc454ea0a610f81e")
	// Get transaction receipt
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Failed to get transaction by hash: %v", err)
	}
	block := big.NewInt(8339)
	revertReason, err := getRevertReason(client, tx, block)
	fmt.Printf("%s", revertReason)
}

func TestDecodeCustomError(t *testing.T) {
	// Example error data for the known errors
	testCases := []struct {
		name      string
		data      string
		expected  string
		expectErr bool
	}{
		// {
		// 	name:     "SequencerEntryAlreadyCommitted",
		// 	data:     "0x00002495",
		// 	expected: `Error: SequencerEntryAlreadyCommitted, Values: {}`,
		// },
		// {
		// 	name:     "SequencerRandomnessNotCommitted",
		// 	data:     "0x38401d6c0000000000000000000000000000000000000000000000000000000000000002",
		// 	expected: `Error: SequencerRandomnessNotCommitted, Values: {"T":2}`,
		// },
		// {
		// 	name:     "PrecommitDelayNotPassed",
		// 	data:     "0xda70d14500000000000000000000000000000000000000000000000000000000664c429600000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000c",
		// 	expected: `Error: PrecommitDelayNotPassed, Values: {"T":1716273814,"committedBlock":12,"currentBlock":18,"requiredBlock":22}`,
		// },
		// {
		// 	name:     "SequencerRandomnessAlreadyRevealed",
		// 	data:     "0xc6b108280000000000000000000000000000000000000000000000000000000000000001",
		// 	expected: `Error: SequencerRandomnessAlreadyRevealed, Values: {"T":1}`,
		// },
		// {
		// 	name:     "InvalidRandomnessReveal",
		// 	data:     "0x0267bee865787065637465644861736856616c7565000000000000000000000000000000636f6d70757465644861736856616c7565000000000000000000000000000000",
		// 	expected: `Error: InvalidRandomnessReveal, Values: {"computedHash":"computedHashValue","expectedHash":"expectedHashValue"}`,
		// },
		{
			name:     "InvalidRandomnessReveal",
			data:     "0x0267bee8fde81b86ee24263afe72e784f41707b090eae5e3bda3a621403b2ed5f2f265350558579641e802271b46ee67dccbccf98571d1fb18a4e1dda9b342bc8a25934c",
			expected: `Error: InvalidRandomnessReveal, Values: {"computedHash":"computedHashValue","expectedHash":"expectedHashValue"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockErr := &MockJsonError{
				err:     "mock error",
				errCode: 400,
				errData: tc.data,
			}
			result, err := utils.DecodeCustomError(mockErr)
			if (err != nil) != tc.expectErr {
				t.Fatalf("expected error: %v, got: %v", tc.expectErr, err)
			}
			if result != tc.expected {
				spew.Dump(tc.expected)
				spew.Dump(result)
				t.Errorf("expected result: %s, got: %s", tc.expected, result)
			}
		})
	}
}

func TestContracts(t *testing.T) {
	// drandOracleABI := contracts.DrandOracleMetaData.ABI
	// spew.Dump(drandOracleABI)
	// oracleABI, err := contracts.DrandOracleMetaData.GetAbi()
	// assert.NoError(t, err, "Failed to get ABI")
	// //spew.Dump(oracleABI)

	client, err := ethclient.Dial("http://localhost:8545")
	assert.NoError(t, err, "Failed to connect to the Ethereum client")
	sequencerRandomOracle, err := contracts.NewSequencerRandomOracle(common.HexToAddress("0x79F72425DA9AeEF0E7DEAcA120EcA45cad20D9DB"), client)

	// privateKey, err := utils.DerivePrivateKeyFromMnemonic("test test test test test test test test test test test junk")
	// assert.NoError(t, err, "Failed to get private key")
	// opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	// assert.NoError(t, err, "Failed to send opts")
	// // value := hex.EncodeToString(valueBytes[:])
	// // log.Printf("%s", value)
	var hashBytes32 [32]byte
	copy(hashBytes32[:], common.HexToHash("0x3fd54831f488a22b28398de0c567a3b064b937f54f81739ae9bd545967f3abab").Bytes())
	// _, err = sequencerRandomOracle.SequencerRandomOracleTransactor.PostRandomnessCommitment(opts, big.NewInt(31330), hashBytes32)
	// if err != nil {
	// 	decodedError, err := utils.DecodeCustomError(err)
	// 	log.Printf("custom error: %s", decodedError)
	// 	assert.NoError(t, err, "Failed to decode custom error")

	// }
	// assert.NoError(t, err, "Failed to send tx")
	// log.Printf("%v", tx)

	privateKey, err := utils.DerivePrivateKeyFromMnemonic("test test test test test test test test test test test junk")
	assert.NoError(t, err, "Failed to get private key")
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	assert.NoError(t, err, "Failed to send opts")
	_, err = sequencerRandomOracle.SequencerRandomOracleTransactor.RevealSequencerRandomness(opts, big.NewInt(31333), utils.StringToBytes32("randomnessHash"))
	if err != nil {
		decodedError, err := utils.DecodeCustomError(err)
		log.Printf("custom error: %s", decodedError)
		assert.NoError(t, err, "Failed to decode custom error")
	}
	assert.NoError(t, err, "Failed to send tx")

	// === RUN   TestContracts
	// 7d45432822a73a87cff134ddc0ccf06c
	// 7d45432822a73a87cff134ddc0ccf06cd38d747c7dce8f5f1b8deca0c113029e 2024/05/22 23:03:23 7d45432822a73a87cff134ddc0ccf06cd38d747c7dce8f5f1b8deca0c113029e
	// --- PASS: TestContracts (0.03s)
	// PASS
	// ok  	server/v0/pkg/server	0.326s
	result, err := sequencerRandomOracle.SequencerRandomOracleCaller.SequencerEntries(nil, big.NewInt(31330))
	assert.NoError(t, err, "Failed to retrieve")
	spew.Dump(result.RandomnessHash)

	// assert.NoError(t, err, "Failed to bind contract")
	// result, err := oracleDrandContract.DrandEntries(nil, big.NewInt(1716388692))
	// assert.NoError(t, err, "Failed to get entries")
	// spew.Dump(result.Randomness)
	// privateKey, err := utils.DerivePrivateKeyFromMnemonic("test test test test test test test test test test test junk")
	// assert.NoError(t, err, "Failed to get private key")

	// opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337))
	// assert.NoError(t, err, "Failed to send opts")

	// tx, err := oracleDrandContract.DrandOracleTransactor.PostDrandRandomness(opts, big.NewInt(1716588692), result.Randomness)
	// assert.NoError(t, err, "Failed to send tx")
}
