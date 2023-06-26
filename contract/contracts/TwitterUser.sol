pragma solidity ^0.8.19;

// SPDX-License-Identifier: UNLICENSED
// Manage twitter users
contract TwitterUser {
    struct User {
        string name;
        string password;
        bool isLogin;
    }
    mapping (address => User) users;
    
    uint userNumber;
    address admin;

   constructor() {
    admin = msg.sender;
   }

   event userTracker(address indexed userAddress, string action, string name, string password, uint userNumber);
   event userLoginTracker(address indexed userAddress, string pswd1, string pswd2);
   event userUnloginTracker(address indexed userAddress);

   modifier isAdmin() {
    require(msg.sender == admin, "Only admin can do this operation");
    _;
   }

   modifier isAddressNotRegistered(address userAddress) {
    require(bytes(users[userAddress].name).length == 0, "address is registered");
    _;
   }

   modifier isAddressRegistered(address userAddress) {
    require(bytes(users[userAddress].name).length > 0, "address is not registered");
    _;
   }

   modifier validateNameAndPassword(string memory name, string memory password) {
    require(bytes(name).length > 0, "name cannot be empty");
    require(bytes(password).length > 0, "password cannot be empty");
    _;
   }

   function register(address userAddress, string memory name, string memory password) public isAdmin isAddressNotRegistered(userAddress) validateNameAndPassword(name, password) {
    User memory user = User(name, password, false);
    users[userAddress] = user;
    userNumber++;
    emit userTracker(userAddress, "register", name, password, userNumber);
   }

   function unregister(address userAddress) public isAdmin {
    if (bytes(users[userAddress].name).length > 0 ) {
        users[userAddress].name = "";
        users[userAddress].password = "";
        users[userAddress].isLogin = false;
        userNumber--;
        emit userTracker(userAddress, "unregister", "", "", userNumber);
    }
   }

   function login(address userAddress, string memory password) public isAddressRegistered(userAddress) {
    if (bytes(users[userAddress].name).length > 0 ) {
        User memory user = users[userAddress];
        if (keccak256(abi.encodePacked(user.password)) == keccak256(abi.encodePacked(password))) {
            users[userAddress].isLogin = true;
            emit userLoginTracker(userAddress, password, user.password);
        } else {
            revert("failed to login because of the mistched password");
        }
    }
   }

   function unlogin(address userAddress) public isAddressRegistered(userAddress){
    if (!checkIfUserLogined(userAddress)) {
        revert("failed to unlogin");
    }
    users[userAddress].isLogin = false;
    emit userUnloginTracker(userAddress);
   }

   function checkIfUserLogined(address userAddress) public view isAddressRegistered(userAddress) returns(bool) {
    return users[userAddress].isLogin;
   }

   function checkUserInfo(address userAddress) public view returns(string memory, string memory, bool) {
    if (bytes(users[userAddress].name).length > 0 ) {
        User memory user = users[userAddress];
        return (user.name, user.password, user.isLogin);
    }
    return ("", "", false);
   }

   function checkUserNumbers() public view returns(uint) {
    return userNumber;
   }
}