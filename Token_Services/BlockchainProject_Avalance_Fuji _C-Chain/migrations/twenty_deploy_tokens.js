const SimpleToken = artifacts.require("TokenTwenty");

module.exports = function (deployer) {
  deployer.deploy(TokenTwenty);
};