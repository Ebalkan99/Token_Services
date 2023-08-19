const MyERC1155Token = artifacts.require("MyERC1155Token");

module.exports = function (deployer) {
  deployer.deploy(MyERC1155Token);
};