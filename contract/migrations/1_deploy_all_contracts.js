const TwitterUser = artifacts.require("TwitterUser");
const Twitter = artifacts.require("Twitter");

module.exports = function(deployer,  network, accounts) {
  deployer.deploy(TwitterUser, {from: accounts[0]}).then(function() {
    return deployer.deploy(Twitter, TwitterUser.address);
  });
}