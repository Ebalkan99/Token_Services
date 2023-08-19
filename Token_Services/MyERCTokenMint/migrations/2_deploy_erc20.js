const MyERC20Token = artifacts.require("MyERC20Token");

module.exports = function (deployer) {
  deployer.deploy(MyERC20Token);
};
