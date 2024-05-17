#!/bin/bash

# Read contract addresses from JSON file in the shared directory and set environment variables
addresses=$(cat /app/shared/addresses.json)
export DRAND_ORACLE_ADDRESS=$(echo $addresses | jq -r .drandOracleAddress)
export SEQUENCER_RANDOM_ORACLE_ADDRESS=$(echo $addresses | jq -r .sequencerRandomOracleAddress)
export RANDOMNESS_ORACLE_ADDRESS=$(echo $addresses | jq -r .randomnessOracleAddress)

# Run the Go server
./main