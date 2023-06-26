pragma solidity ^0.8.19;

// SPDX-License-Identifier: UNLICENSED
// Manage twitter users
contract TwitterUser {
    mapping (address => string) users;
    uint userNumber;
    address admin;

   constructor() {
    admin = msg.sender;
   }

   event userTracker(address indexed userAddress, string action, string name, uint userNumber);

   modifier isAdmin() {
    require(msg.sender == admin, "Only admin can do this operation");
    _;
   }

   function register(address userAddress, string memory name) public isAdmin {
    users[userAddress] = name;
    userNumber++;
    emit userTracker(userAddress, "register", name, userNumber);
   }

   function unregister(address userAddress) public isAdmin {
    if (bytes(users[userAddress]).length > 0 ) {
        users[userAddress] = "";
        userNumber--;
        emit userTracker(userAddress, "unregister", "", userNumber);
    }
   }

   function checkIfUserExists(address userAddress) public view returns(bool){
    if (bytes(users[userAddress]).length > 0 ) {
        return true;
    }
    return false;
   }

   function checkUserName(address userAddress) public view returns(string memory ) {
    if (bytes(users[userAddress]).length > 0 ) {
        return users[userAddress];
    }
    return "";
   }

   function checkUserNumbers() public view returns(uint) {
    return userNumber;
   }
}