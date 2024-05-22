package utils

import (
	servertypes "server/v0/types/server"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadContractAddresses(t *testing.T) {
	tests := []struct {
		name           string
		filename       string
		expectedResult *servertypes.ContractAddresses
		expectedError  bool
	}{
		{
			name:     "Valid JSON file",
			filename: "../../fixtures/addresses.json",
			expectedResult: &servertypes.ContractAddresses{
				DrandOracleAddress:           "0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9",
				SequencerRandomOracleAddress: "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9",
				RandomnessOracleAddress:      "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
			},
			expectedError: false,
		},
		{
			name:           "Non-existent file",
			filename:       "testdata/non_existent.json",
			expectedResult: nil,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := LoadContractAddresses(tt.filename)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
