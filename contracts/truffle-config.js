const path = require("path");

module.exports = {
  contracts_build_directory: path.join(__dirname, "build"),
  contracts_directory: path.join(__dirname, "src"),
  networks: {
    development: {
      host: "anvil",
      port: 8545,
      network_id: "*", // Match any network id
    },
  },
  compilers: {
    solc: {
      version: "0.8.24", // Fetch exact version from solc-bin
    },
  },
};
