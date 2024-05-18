// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SequencerRandomOracle {
    uint constant PRECOMMIT_DELAY = 10; // blocks, for testing purposes

    struct SequencerEntry {
        bytes32 randomnessHash;
        bytes32 randomness;
        uint blockNumber;
        bool committed;
        bool revealed;
    }

    mapping(uint => SequencerEntry) public sequencerEntries;

    event SequencerRandomnessPosted(uint indexed T, bytes32 randomnessHash);
    event SequencerRandomnessRevealed(uint indexed T, bytes32 randomness);

    function postRandomnessCommitment(uint T, bytes32 randomnessHash) external {
        SequencerEntry storage entry = sequencerEntries[T];
        require(!entry.committed, "Sequencer entry already committed");

        entry.randomnessHash = randomnessHash;
        entry.blockNumber = block.number;
        entry.committed = true;
        entry.revealed = false;

        emit SequencerRandomnessPosted(T, randomnessHash);
    }

    function revealSequencerRandomness(uint T, bytes32 randomness) external {
        SequencerEntry storage entry = sequencerEntries[T];
        require(entry.committed, "Sequencer randomness not committed");
        require(block.number > entry.blockNumber + PRECOMMIT_DELAY, "Precommit delay not passed");
        require(!entry.revealed, "Sequencer randomness already revealed");
        require(entry.randomnessHash == keccak256(abi.encodePacked(randomness)), "Invalid randomness reveal");

        entry.randomness = randomness;
        entry.revealed = true;

        emit SequencerRandomnessRevealed(T, randomness);
    }

    function getSequencerRandomness(uint T) external view returns (bytes32) {
        SequencerEntry storage entry = sequencerEntries[T];
        require(entry.revealed, "Sequencer randomness not available");

        return entry.randomness;
    }
}
