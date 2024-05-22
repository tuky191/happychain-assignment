#!/bin/bash

# Directory containing the JSON files
JSON_DIR=/app/build
ABI_DIR=/app/build/abis

# Create ABI output directory if it doesn't exist
mkdir -p $ABI_DIR

# List of contract names to process
CONTRACTS=("DrandOracle" "RandomnessOracle" "SequencerRandomOracle")

for contract in "${CONTRACTS[@]}"; do
    # JSON file path
    json_file=$JSON_DIR/${contract}.json
    
    # ABI file path
    abi_file=$ABI_DIR/${contract}.abi
    
    # Extract ABI and save to .abi file
    jq -r .abi "$json_file" > "$abi_file"
done
