// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

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

    error SequencerEntryAlreadyCommitted();
    error SequencerRandomnessNotCommitted();
    error PrecommitDelayNotPassed();
    error SequencerRandomnessAlreadyRevealed();
    error InvalidRandomnessReveal(bytes32 expected, bytes32 actual);

    event SequencerRandomnessPosted(uint indexed T, bytes32 randomnessHash);
    event SequencerRandomnessRevealed(uint indexed T, bytes32 randomness);

    function postRandomnessCommitment(uint T, bytes32 randomnessHash) external {
        SequencerEntry storage entry = sequencerEntries[T];
        if (entry.committed) {
            revert SequencerEntryAlreadyCommitted();
        }

        entry.randomnessHash = randomnessHash;
        entry.blockNumber = block.number;
        entry.committed = true;
        entry.revealed = false;

        emit SequencerRandomnessPosted(T, randomnessHash);
    }

    function revealSequencerRandomness(uint T, bytes32 randomness) external {
        SequencerEntry storage entry = sequencerEntries[T];
        if (!entry.committed) {
            revert SequencerRandomnessNotCommitted();
        }
        if (block.number <= entry.blockNumber + PRECOMMIT_DELAY) {
            revert PrecommitDelayNotPassed();
        }
        if (entry.revealed) {
            revert SequencerRandomnessAlreadyRevealed();
        }

        bytes32 computedHash = keccak256(abi.encodePacked(randomness));
        if (entry.randomnessHash != computedHash) {
            revert InvalidRandomnessReveal(entry.randomnessHash, computedHash);
        }

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
