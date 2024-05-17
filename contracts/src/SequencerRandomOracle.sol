// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract SequencerRandomOracle {
    mapping(uint256 => bytes32) public commitments;
    mapping(uint256 => bytes32) public revealedValues;

    function commit(uint256 round, bytes32 commitment) public {
        commitments[round] = commitment;
    }

    function reveal(uint256 round, string memory value) public {
        revealedValues[round] = keccak256(abi.encodePacked(value));
    }
}