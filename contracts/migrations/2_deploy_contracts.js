const DrandOracle = artifacts.require("DrandOracle");
const SequencerRandomOracle = artifacts.require("SequencerRandomOracle");
const RandomnessOracle = artifacts.require("RandomnessOracle");

const fs = require("fs");
const path = require("path");

module.exports = async function (deployer) {
  await deployer.deploy(DrandOracle);
  const drandOracle = await DrandOracle.deployed();

  await deployer.deploy(SequencerRandomOracle);
  const sequencerRandomOracle = await SequencerRandomOracle.deployed();

  await deployer.deploy(
    RandomnessOracle,
    drandOracle.address,
    sequencerRandomOracle.address
  );
  const randomnessOracle = await RandomnessOracle.deployed();

  const contractAddresses = {
    drandOracleAddress: drandOracle.address,
    sequencerRandomOracleAddress: sequencerRandomOracle.address,
    randomnessOracleAddress: randomnessOracle.address,
  };

  const outputPath = path.resolve(__dirname, "../shared/addresses.json");
  fs.writeFileSync(
    outputPath,
    JSON.stringify(contractAddresses, null, 2),
    "utf-8"
  );

  console.log(`Contract addresses saved to ${outputPath}`);
};
