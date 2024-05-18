// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DrandOracle {
    uint constant TIMEOUT = 10; // seconds, for testing purposes

    struct DrandEntry {
        bytes32 randomness;
        uint timestamp;
        bool filled;
    }

    mapping(uint => DrandEntry) public drandEntries;

    event DrandUpdated(uint indexed T, bytes32 randomness);
    
    function postDrandRandomness(uint T, bytes32 randomness) external {
        require(block.timestamp <= T + TIMEOUT, "Update period has expired");

        DrandEntry storage entry = drandEntries[T];
        require(!entry.filled, "Drand entry already filled");

        entry.randomness = randomness;
        entry.timestamp = block.timestamp;
        entry.filled = true;

        emit DrandUpdated(T, randomness);
    }

    function getDrand(uint T) external view returns (bytes32) {
        DrandEntry storage entry = drandEntries[T];
        require(entry.filled, "Drand entry not available");

        return entry.randomness;
    }

    function isDrandAvailable(uint T) external view returns (bool) {
        return drandEntries[T].filled;
    }

    function hasUpdatePeriodExpired(uint T) external view returns (bool) {
        return block.timestamp > T + TIMEOUT;
    }
}
