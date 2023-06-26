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

    TwitterUser userManager;
    
    constructor(address _userManager) {
        userManager = TwitterUser(_userManager);
    }

    event postTracker(address indexed from, string who, string content);

    modifier isLogin() {
        require(userManager.checkIfUserLogined(msg.sender), "It is not logined");
        _;
    }

    function post(string memory content) public isLogin {
        tweets[msg.sender].push(tweet(block.timestamp, content));
        (string memory name, , ) = userManager.checkUserInfo(msg.sender);
        emit postTracker(msg.sender, name, content);
    }

    function readPost(address readFrom, uint index) public view returns(string memory, string memory, uint) {
        if (index >= tweets[readFrom].length) {
            return ("", "", 0);
        }
        tweet memory tw = tweets[readFrom][index];
        (string memory name, , ) = userManager.checkUserInfo(readFrom);

        return (name, tw.tweetMessage, tw.timestamp);
    }

    function postNumbers() public view returns(uint) {
        return tweets[msg.sender].length;
    }
}