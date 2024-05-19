// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

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
        bytes32 drandRandomness;
        bytes32 sequencerRandomness;

        try drandOracle.getDrand(T) returns (bytes32 randomness) {
            drandRandomness = randomness;
        } catch {
            return 0;
        }

        try sequencerRandomOracle.getSequencerRandomness(T) returns (bytes32 randomness) {
            sequencerRandomness = randomness;
        } catch {
            return 0;
        }

        if (drandRandomness == 0 || sequencerRandomness == 0) {
            return 0;
        }

        return keccak256(abi.encodePacked(drandRandomness, sequencerRandomness));
    }

    function isRandomnessEverAvailable(uint T) external view returns (bool) {
        bool drandAvailable;
        bool sequencerAvailable;

        try drandOracle.isDrandAvailable(T) returns (bool available) {
            drandAvailable = available;
        } catch {
            drandAvailable = false;
        }

        try drandOracle.hasUpdatePeriodExpired(T) returns (bool expired) {
            drandAvailable = drandAvailable || expired;
        } catch {
            drandAvailable = drandAvailable;
        }

        try sequencerRandomOracle.getSequencerRandomness(T) returns (bytes32 randomness) {
            sequencerAvailable = randomness != 0;
        } catch {
            sequencerAvailable = false;
        }

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
