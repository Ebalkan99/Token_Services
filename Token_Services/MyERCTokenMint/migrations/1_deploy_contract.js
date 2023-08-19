const MyERC721Token = artifacts.require("MyERC721Token");

module.exports = function (deployer) {
  deployer.deploy(MyERC721Token);
};

