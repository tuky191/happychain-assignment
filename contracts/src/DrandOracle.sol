// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;
// import "hardhat/console.sol";
import  {Ownable} from "../node_modules/@openzeppelin/contracts/access/Ownable.sol";

contract DrandOracle is Ownable {
    uint constant TIMEOUT = 10; // seconds, for testing purposes

    struct DrandEntry {
        bytes32 randomness;
        uint timestamp;
        bool filled;
    }
    constructor() Ownable(msg.sender) {}

    mapping(uint => DrandEntry) public drandEntries;

    event DrandUpdated(uint indexed T, bytes32 randomness);
    
    function postDrandRandomness(uint T, bytes32 randomness) external onlyOwner{
    // console.log("block.timestamp: %s, T: %s, TIMEOUT: %s", block.timestamp, T, TIMEOUT);
        require(block.timestamp <= T + TIMEOUT, "Update period has expired");

        DrandEntry storage entry = drandEntries[T];
        require(!entry.filled, "Drand entry already filled");

        entry.randomness = randomness;
        entry.timestamp = block.timestamp;
        entry.filled = true;

        emit DrandUpdated(T, randomness);
    }

    function getDrand(uint T) external view returns (bytes32) {
        DrandEntry storage entry = drandEntries[T];
        require(entry.filled, "Drand entry not available");

        return entry.randomness;
    }

    function isDrandAvailable(uint T) external view returns (bool) {
        return drandEntries[T].filled;
    }

    function hasUpdatePeriodExpired(uint T) external view returns (bool) {
        return block.timestamp > T + TIMEOUT;
    }
}
