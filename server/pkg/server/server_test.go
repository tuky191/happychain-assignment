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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

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

func TestCreateTransactionReveal(t *testing.T) {
	// Setup
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err, "Failed to generate private key")

	client, err := ethclient.Dial("http://localhost:8545")
	assert.NoError(t, err, "Failed to connect to the Ethereum client")

	contractAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_value","type":"bytes32"}],"name":"revealSequencerRandomness","outputs":[],"payable":false,"stateMutability":"view","type":"function"}]`))
	assert.NoError(t, err, "Failed to parse contract ABI")

	timestamp := int64(1234567890)
	value := "0xaff07feb058c3908335b80f307a0f0b6e45a45ff17437ea0e2ceec940ce9dd65"

	// Execute
	tx, err := createTransaction(client, privateKey, contractAddress, "revealSequencerRandomness", contractABI, timestamp, value)
	assert.NoError(t, err, "Failed to create transaction")

	// Validate
	assert.NotNil(t, tx, "Transaction should not be nil")
	assert.Equal(t, uint64(3000000), tx.Gas())
	assert.Equal(t, common.HexToAddress(contractAddress), *tx.To())
}

func TestCreateTransactionCommit(t *testing.T) {
	// Setup
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err, "Failed to generate private key")

	client, err := ethclient.Dial("http://localhost:8545")
	assert.NoError(t, err, "Failed to connect to the Ethereum client")

	contractAddress := "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":true,"inputs":[{"name":"_timestamp","type":"uint256"},{"name":"_value","type":"bytes32"}],"name":"postRandomnessCommitment","outputs":[],"payable":false,"stateMutability":"view","type":"function"}]`))
	assert.NoError(t, err, "Failed to parse contract ABI")

	timestamp := int64(1234567890)
	value := "0xa9cc7948e824bf6db251ba79694ea3ff4953d52f66ac2dea7b431f8bab2d070c"

	// Execute
	tx, err := createTransaction(client, privateKey, contractAddress, "postRandomnessCommitment", contractABI, timestamp, value)
	assert.NoError(t, err, "Failed to create transaction")

	// Validate
	assert.NotNil(t, tx, "Transaction should not be nil")
	assert.Equal(t, uint64(3000000), tx.Gas())
	assert.Equal(t, common.HexToAddress(contractAddress), *tx.To())
}

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
		{
			name:     "SequencerEntryAlreadyCommitted",
			data:     "0x00002495",
			expected: `Error: SequencerEntryAlreadyCommitted, Values: {}`,
		},
		{
			name:     "SequencerRandomnessNotCommitted",
			data:     "0x38401d6c0000000000000000000000000000000000000000000000000000000000000002",
			expected: `Error: SequencerRandomnessNotCommitted, Values: {"T":2}`,
		},
		{
			name:     "PrecommitDelayNotPassed",
			data:     "0xda70d14500000000000000000000000000000000000000000000000000000000664c429600000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000000016000000000000000000000000000000000000000000000000000000000000000c",
			expected: `Error: PrecommitDelayNotPassed, Values: {"T":1716273814,"committedBlock":12,"currentBlock":18,"requiredBlock":22}`,
		},
		{
			name:     "SequencerRandomnessAlreadyRevealed",
			data:     "0xc6b108280000000000000000000000000000000000000000000000000000000000000001",
			expected: `Error: SequencerRandomnessAlreadyRevealed, Values: {"T":1}`,
		},
		{
			name:     "InvalidRandomnessReveal",
			data:     "0x0267bee865787065637465644861736856616c7565000000000000000000000000000000636f6d70757465644861736856616c7565000000000000000000000000000000",
			expected: `Error: InvalidRandomnessReveal, Values: {"computedHash":"computedHashValue","expectedHash":"expectedHashValue"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := utils.DecodeCustomError(tc.data)
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
	drandOracleABI := contracts.DrandOracleMetaData.ABI
	spew.Dump(drandOracleABI)
	//oracleABI, err := contracts.DrandOracleMetaData.GetAbi()

}
