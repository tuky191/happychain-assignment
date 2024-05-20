package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestGetDrandValue(t *testing.T) {

// 	tests := []struct {
// 		name           string
// 		timestamp      int64
// 		responseStatus int
// 		responseBody   string
// 		expectedValue  string
// 		expectError    bool
// 		mockHTTPError  bool
// 	}{
// 		{
// 			name:           "Valid response",
// 			timestamp:      1609459260,
// 			responseStatus: http.StatusOK,
// 			responseBody:   `{"round":53648642262,"randomness":"abc123"}`,
// 			expectedValue:  "abc123",
// 			expectError:    false,
// 			mockHTTPError:  false,
// 		},
// 		{
// 			name:           "Invalid JSON response",
// 			timestamp:      1609459260,
// 			responseStatus: http.StatusOK,
// 			responseBody:   `{"round":53648642262,"randomness":}`,
// 			expectedValue:  "",
// 			expectError:    true,
// 			mockHTTPError:  false,
// 		},
// 		{
// 			name:           "Non-200 HTTP status",
// 			timestamp:      1609459260,
// 			responseStatus: http.StatusInternalServerError,
// 			responseBody:   `Internal Server Error`,
// 			expectedValue:  "",
// 			expectError:    true,
// 			mockHTTPError:  false,
// 		},
// 		{
// 			name:           "HTTP request error",
// 			timestamp:      1609459260,
// 			responseStatus: http.StatusOK,
// 			responseBody:   `{"round":53648642262,"randomness":"abc123"}`,
// 			expectedValue:  "",
// 			expectError:    true,
// 			mockHTTPError:  true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Create a mock server
// 			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				w.WriteHeader(tt.responseStatus)
// 				w.Write([]byte(tt.responseBody))
// 			}))
// 			defer server.Close()

// 			// Replace the httpGet function to use the mock server URL
// 			oldHttpGet := httpGet
// 			httpGet = func(url string) (*http.Response, error) {
// 				if tt.mockHTTPError {
// 					return nil, fmt.Errorf("mock HTTP error")
// 				}
// 				round := (tt.timestamp - drandGenesis) / drandInterval
// 				mockURL := fmt.Sprintf("%s/public/%d", server.URL, round)
// 				return http.Get(mockURL)
// 			}
// 			defer func() { httpGet = oldHttpGet }()

// 			// Call the function
// 			value, err := getDrandValue(tt.timestamp)

// 			// Check the result
// 			if (err != nil) != tt.expectError {
// 				t.Errorf("Expected error: %v, got: %v", tt.expectError, err)
// 			}
// 			if value != tt.expectedValue {
// 				t.Errorf("Expected value: %s, got: %s", tt.expectedValue, value)
// 			}
// 		})
// 	}
// }

func TestGetRevertReason(t *testing.T) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	txHash := common.HexToHash("0xfdb1cf0628337db9814b88cf95b1ee376ed53fdfbea5ae11ea6e24771eae58aa")
	// Get transaction receipt
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Failed to get transaction by hash: %v", err)
	}
	spew.Dump(tx)
	spew.Dump(isPending)
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

	bytes32Value := stringToBytes32(input)
	bytes32Hex := bufferToHex(bytes32Value[:])
	assert.Equal(t, expectedBytes32Hex, bytes32Hex, "Bytes32 value should match expected")

	types := []string{"bytes32"}
	values := []interface{}{bytes32Value}
	contractABI, err := abi.JSON(strings.NewReader(`[{"constant":false,"inputs":[{"name":"T","type":"uint256"},{"name":"randomnessHash","type":"bytes32"}],"name":"postRandomnessCommitment","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]`))
	fmt.Print(contractABI.Methods)
	hash, err := solidityPackedKeccak256(types, values)
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
