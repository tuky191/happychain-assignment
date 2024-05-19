package main

import (
	"context"
	"encoding/hex"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func TestGetTransactionReceipt(t *testing.T) {
	// Create Ethereum client
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Transaction hash to check
	txHash := "0xdd8345f303b887d27b88896beaa94444cfba8c5f8e575dfdc4e2ac156e42dc09"

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
	input := "randomness"
	expectedBytes32 := "72616e646f6d6e65737300000000000000000000000000000000000000000000"
	expectedHash := "0xc8fcc37de2a2c63bda685088cc2cbc983e9210f4a2c941de0ac6f35a4d09efce" // Change this to match the expected hash

	bytes32Value := stringToBytes32(input)
	bytes32Hex := hex.EncodeToString(bytes32Value[:])
	assert.Equal(t, expectedBytes32, bytes32Hex, "Bytes32 value should match expected")

	types := []string{"bytes32"}
	values := []interface{}{bytes32Value}

	hash, err := solidityPackedKeccak256(types, values)
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, expectedHash, hash, "Hash should match expected")
}
