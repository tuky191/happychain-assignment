// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DrandOracle {
    mapping(uint256 => bytes32) public drandValues;

    function addDrandValue(uint256 round, bytes32 value) public {
        drandValues[round] = value;
    }
}
