// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DrandOracle {
    uint256 public constant TIMEOUT = 10;
    mapping(uint256 => bytes32) public randomness;
    mapping(uint256 => bool) public hasRandomness;

    event RandomnessUpdated(uint256 indexed time, bytes32 randomness);

    function updateRandomness(uint256 time, bytes32 value) external {
        require(block.timestamp <= time + TIMEOUT, "DrandOracle: TIMEOUT exceeded");
        require(!hasRandomness[time], "DrandOracle: Randomness already set");

        randomness[time] = value;
        hasRandomness[time] = true;

        emit RandomnessUpdated(time, value);
    }

    function getRandomness(uint256 time) external view returns (bytes32) {
        return randomness[time];
    }
}
