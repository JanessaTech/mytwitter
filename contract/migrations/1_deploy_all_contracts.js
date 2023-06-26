const TwitterUser = artifacts.require("TwitterUser");

module.exports = function(deployer,  network, accounts) {
  deployer.deploy(TwitterUser, {from: accounts[0]});
}