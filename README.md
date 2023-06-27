# mytwitter
Build a distirbuted twitter using go and blockchain

mnemonic: cake attract curve erase army talent parade patrol dumb cruel hotel sauce

# How to build solidity files to go files:
```
solcjs --abi TwitterUser.sol -o build
solcjs --bin TwitterUser.sol -o build
abigen --abi ./build/TwitterUser_sol_TwitterUser.abi --pkg twitter --type TwitterUser --out TwitterUser.go --bin ./build/TwitterUser_sol_TwitterUser.bin

solcjs --abi Twitter.sol -o build
solcjs --bin Twitter.sol -o build
abigen --abi ./build/Twitter_sol_Twitter.abi --pkg twitter --type Twitter --out Twitter.go --bin ./build/Twitter_sol_Twitter.bin
```

# download dependences:
```
go get github.com/ethereum/go-ethereum
go get github.com/ethereum/go-ethereum/accounts/abi/bind 
```

