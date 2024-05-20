// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
import  {Ownable} from "../node_modules/@openzeppelin/contracts/access/Ownable.sol";
contract SequencerRandomOracle is Ownable {
    uint constant PRECOMMIT_DELAY = 10; // blocks, for testing purposes

    struct SequencerEntry {
        bytes32 randomnessHash;
        bytes32 randomness;
        uint blockNumber;
        bool committed;
        bool revealed;
    }
    constructor() Ownable(msg.sender) {}

    mapping(uint => SequencerEntry) public sequencerEntries;

    error SequencerEntryAlreadyCommitted();
    error SequencerRandomnessNotCommitted(uint T);
    error PrecommitDelayNotPassed(uint T, uint currentBlock, uint requiredBlock, uint committedBlock);
    error SequencerRandomnessAlreadyRevealed(uint T);
    error InvalidRandomnessReveal(bytes32 expectedHash, bytes32 computedHash);

    event SequencerRandomnessPosted(uint indexed T, bytes32 randomnessHash);
    event SequencerRandomnessRevealed(uint indexed T, bytes32 randomness);

    function postRandomnessCommitment(uint T, bytes32 randomnessHash) external onlyOwner {
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

    function revealSequencerRandomness(uint T, bytes32 randomness) external onlyOwner {
        SequencerEntry storage entry = sequencerEntries[T];
        if (!entry.committed) {
            revert SequencerRandomnessNotCommitted(T);
        }
        uint requiredBlock = entry.blockNumber + PRECOMMIT_DELAY;
        if (block.number <= requiredBlock) {
            revert PrecommitDelayNotPassed(T, block.number, requiredBlock, entry.blockNumber);
        }
        if (entry.revealed) {
            revert SequencerRandomnessAlreadyRevealed(T);
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
