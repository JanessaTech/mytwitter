pragma solidity ^0.8.19;

import "./TwitterUser.sol";

// SPDX-License-Identifier: UNLICENSED
// Perform the basic functions for twitter
contract Twitter {
    struct tweet {
        uint timestamp;
        string tweetMessage;
    }

    mapping (address => tweet[]) tweets;

    TwitterUser twitterUsers;
    
    constructor(address _TwitterUsers) {
        twitterUsers = TwitterUser(_TwitterUsers);
    }

    event postTracker(address indexed from, string who, string content);

    modifier isValidUser() {
         require(twitterUsers.checkIfUserExists(msg.sender), "Unregister user");
         _;
    }

    function post(string memory content) public isValidUser {
        tweets[msg.sender].push(tweet(block.timestamp, content));
        string memory name = twitterUsers.checkUserName(msg.sender);
        emit postTracker(msg.sender, name, content);
    }

    function readPost(uint index) public view isValidUser returns(string memory, string memory, uint) {
        if (index >= tweets[msg.sender].length) {
            return ("", "", 0);
        }
        tweet storage tw = tweets[msg.sender][index];
        string memory name = twitterUsers.checkUserName(msg.sender);

        return (name, tw.tweetMessage, tw.timestamp);
    }

    function postNumbers() public view returns(uint) {
        return tweets[msg.sender].length;
    }
}