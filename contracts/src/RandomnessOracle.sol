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

    function getRandomness(uint256 round) public view returns (bytes32) {
        bytes32 drandValue = drandOracle.drandValues(round);
        bytes32 sequencerValue = sequencerRandomOracle.revealedValues(round);
        return keccak256(abi.encodePacked(drandValue, sequencerValue));
    }
}