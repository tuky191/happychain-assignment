const { ethers } = require("ethers");
const fs = require("fs");
require("dotenv").config();

const addresses = JSON.parse(
  fs.readFileSync("/app/shared/addresses.json", "utf8")
);
require("dotenv").config();
const mnemonic = process.env.MNEMONIC;
const rpcUrl = process.env.ANVIL_URL;

if (!mnemonic || !rpcUrl) {
  console.error(
    "Please set MNEMONIC and RPC_URL in the environment variables."
  );
  process.exit(1);
}
async function main() {
  const mnemonic =
    "test test test test test test test test test test test junk";
  const provider = new ethers.JsonRpcProvider(rpcUrl);
  const wallet = ethers.Wallet.fromPhrase(mnemonic);
  const signer = wallet.connect(provider);

  const drandOracleAddress = addresses.drandOracleAddress;
  const sequencerRandomOracleAddress = addresses.sequencerRandomOracleAddress;
  const randomnessOracleAddress = addresses.randomnessOracleAddress;

  const drandOracleABI = [
    "function getDrand(uint256 T) external view returns (bytes32)",
    "function isDrandAvailable(uint256 T) external view returns (bool)",
  ];

  const sequencerRandomOracleABI = [
    "function revealSequencerRandomness(uint256 T, bytes32 randomness) external",
    "function getSequencerRandomness(uint256 T) external view returns (bytes32)",
  ];

  const randomnessOracleABI = [
    "function computeRandomness(uint256 T) public view returns (bytes32)",
    "function isRandomnessEverAvailable(uint256 T) external view returns (bool)",
    "function simpleGetRandomness(uint256 T) external view returns (bytes32)",
    "function unsafeGetRandomness(uint256 T) external view returns (bytes32)",
  ];
  const drandOracle = new ethers.Contract(
    drandOracleAddress,
    drandOracleABI,
    signer
  );
  const sequencerRandomOracle = new ethers.Contract(
    sequencerRandomOracleAddress,
    sequencerRandomOracleABI,
    signer
  );
  const randomnessOracle = new ethers.Contract(
    randomnessOracleAddress,
    randomnessOracleABI,
    signer
  );

  let T = Math.floor(Date.now() / 1000) - 40;

  if (T % 2 !== 0) {
    T -= 1;
  }
  while (true) {
    T = T + 2;
    const currentTime = Math.floor(Date.now() / 2000);

    console.log(`Fetching randomness for T=${T}`);

    await new Promise((resolve) => setTimeout(resolve, 2000));

    // Get randomness from DrandOracle
    let drandRandomness;
    try {
      drandRandomness = await drandOracle.getDrand(T);
      console.log(`Drand randomness for T=${T}: ${drandRandomness}`);
    } catch (error) {}

    // Get randomness from SequencerRandomOracle
    let sequencerRandomness;
    try {
      sequencerRandomness = await sequencerRandomOracle.getSequencerRandomness(
        T
      );
      console.log(`Sequencer randomness for T=${T}: ${sequencerRandomness}`);
    } catch (error) {}

    // Compute combined randomness using RandomnessOracle
    try {
      const computedRandomness = await randomnessOracle.computeRandomness(T);
      console.log(`Computed randomness for T=${T}: ${computedRandomness}`);
    } catch (error) {}

    // Check if randomness is available
    try {
      const isRandomnessEverAvailable =
        await randomnessOracle.isRandomnessEverAvailable(T);
      console.log(
        `Is randomness ever available for T=${T}? ${isRandomnessEverAvailable}`
      );
    } catch (error) {}

    // Get randomness using simpleGetRandomness
    try {
      const simpleRandomness = await randomnessOracle.simpleGetRandomness(T);
      console.log(`Simple get randomness for T=${T}: ${simpleRandomness}`);
    } catch (error) {}

    // Get randomness using unsafeGetRandomness
    try {
      const unsafeRandomness = await randomnessOracle.unsafeGetRandomness(T);
      console.log(`Unsafe get randomness for T=${T}: ${unsafeRandomness}`);
    } catch (error) {}
  }
}

main().catch(console.error);
