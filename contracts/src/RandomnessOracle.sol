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

    function computeRandomness(uint T) public view returns (bytes32) {
        bytes32 drandRandomness = drandOracle.getDrand(T);
        bytes32 sequencerRandomness = sequencerRandomOracle.getSequencerRandomness(T);
        if (drandRandomness == 0 || sequencerRandomness == 0) {
            return 0;
        }
        return keccak256(abi.encodePacked(drandRandomness, sequencerRandomness));
    }

    function isRandomnessEverAvailable(uint T) external view returns (bool) {
        bool drandAvailable = drandOracle.isDrandAvailable(T) || drandOracle.hasUpdatePeriodExpired(T);
        bool sequencerAvailable = sequencerRandomOracle.getSequencerRandomness(T) != 0;
        return drandAvailable && sequencerAvailable;
    }

    function simpleGetRandomness(uint T) external view returns (bytes32) {
        bytes32 randomness = computeRandomness(T);
        require(randomness != 0, "Randomness is not available");
        return randomness;
    }

    function unsafeGetRandomness(uint T) external view returns (bytes32) {
        return computeRandomness(T);
    }
}
