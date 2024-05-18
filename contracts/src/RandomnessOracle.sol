// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./DrandOracle.sol";
import "./SequencerRandomOracle.sol";

contract RandomnessOracle {
    DrandOracle public drandOracle;
    SequencerRandomOracle public sequencerRandomOracle;

    constructor(address _drandOracle, address _sequencerRandomOracle) {
        drandOracle = DrandOracle(_drandOracle);
        sequencerRandomOracle = SequencerRandomOracle(_sequencerRandomOracle);
    }

    function getRandomness(uint256 time) public view returns (bytes32) {
        bytes32 drandValue = drandOracle.getRandomness(time - 9); // Use DELAY = 9
        bytes32 commitmentValue = sequencerRandomOracle.getSequencerRandom(time);
        return keccak256(abi.encodePacked(drandValue, commitmentValue));
    }

    function safeGetRandomness(uint256 time) public view returns (bytes32) {
        bytes32 randomness = getRandomness(time);
        require(randomness != bytes32(0), "RandomnessOracle: Randomness not available");
        return randomness;
    }

    function isRandomnessAvailable(uint256 time) public view returns (bool) {
        return drandOracle.hasRandomness(time - 9) && sequencerRandomOracle.revealedValues(time) != bytes32(0);
    }
}
