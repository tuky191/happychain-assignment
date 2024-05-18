// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SequencerRandomOracle {
    uint256 public constant PRECOMMIT_DELAY = 10;
    mapping(uint256 => bytes32) public commitments;
    mapping(uint256 => bytes32) public revealedValues;

    event CommitmentPosted(uint256 indexed time, bytes32 commitment);
    event ValueRevealed(uint256 indexed time, bytes32 value);

    function postCommitment(uint256 time, bytes32 commitment) external {
        require(block.timestamp <= time - PRECOMMIT_DELAY, "SequencerRandomOracle: Commitment too late");
        require(commitments[time] == bytes32(0), "SequencerRandomOracle: Commitment already posted");

        commitments[time] = commitment;

        emit CommitmentPosted(time, commitment);
    }

    function revealValue(uint256 time, bytes32 value) external {
        require(block.timestamp >= time, "SequencerRandomOracle: Too early to reveal");
        require(commitments[time] != bytes32(0), "SequencerRandomOracle: No commitment found");
        require(keccak256(abi.encodePacked(value)) == commitments[time], "SequencerRandomOracle: Invalid value");

        revealedValues[time] = value;

        emit ValueRevealed(time, value);
    }

    function getSequencerRandom(uint256 time) external view returns (bytes32) {
        return revealedValues[time];
    }
}
