#!/bin/bash

# Directory containing the ABI files
ABI_DIR=/app/abis
OUTPUT_DIR=/app/pkg/contracts

# Create output directory if it doesn't exist
mkdir -p $OUTPUT_DIR

# List of contract names to process
CONTRACTS=("DrandOracle" "RandomnessOracle" "SequencerRandomOracle")

for contract in "${CONTRACTS[@]}"; do
    # ABI file path
    abi_file=$ABI_DIR/${contract}.abi
    
    # Generate Go bindings
    abigen --abi "$abi_file" --pkg contracts --type "$contract" --out "$OUTPUT_DIR/${contract}.go"
done
