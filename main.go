package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hi-supergirl/mytwitter/server"
	twitterapi "github.com/hi-supergirl/mytwitter/twitterApi"
	"github.com/spf13/cobra"
)

var (
	firstAccountKey  = "707a20975dc5af51c3d95b57c20d32e95d8f607f19a4e966461416c4394f08b8"
	secondAccountKey = "2a6b6829dd97420c19fa534cf064e8e810609758aa9140672bd43ed2eab4a931"
	thirdAccountKey  = "cf5f924d716d4e2e3439312dc81205d01a1059e15c74933bb020041dc5041761"

	twitterUserAddress = "0xCfe30d5B4774048c1Dd231714B0Ac33960cd955f" // the address of the contract TwitterUser
	twitterAddress     = "0x899299896C52797623201347Cfd4a99B6297e200" // the address of the contract Twitter
	//firstAccountAddress  = "0x4c4Db58D4B7DC8268781a23C309d016EDa18FbB8"
	secondAccountAddress = "0x5A144C215331e498E0AEBA71E358DC924BcA59D1"
	thirdAccountAddress  = "0x77793955d13aE7F9199FB252256595b1a1631075"
)

// mnemonic is "cake attract curve erase army talent parade patrol dumb cruel hotel sauce" when start ganache
func CreateConn() *ethclient.Client {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	//client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("We have a connection")

	return client // we'll use this in the upcoming sections
}

func getAuth(key string, client *ethclient.Client) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	_ = gasPrice
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	//auth.GasPrice = gasPrice
	return auth
}

func LoadTwitterUser() *twitterapi.TwitterUser {
	client := CreateConn()
	address := common.HexToAddress(twitterUserAddress)

	instance, err := twitterapi.NewTwitterUser(address, client)

	if err != nil {
		log.Fatal(err)
	}
	return instance
}

func Register(userAddress common.Address, name string, password string) {
	client := CreateConn()
	twitterUserInstance := LoadTwitterUser()

	firstAccountAuth := getAuth(firstAccountKey, client)
	tx, err := twitterUserInstance.Register(firstAccountAuth, userAddress, name, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s for register(%s, %s, %s)\n", tx.Hash().Hex(), userAddress.Hex(), name, password)
}

func CheckUserInfo(address common.Address) {
	twitterUserInstance := LoadTwitterUser()
	name, password, isLogin, err := twitterUserInstance.CheckUserInfo(nil, address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User info for %s : name=%s, password=%s, isLogin=%t\n", address.Hex(), name, password, isLogin)
}

func Login(address common.Address, password string) {
	client := CreateConn()
	twitterUserInstance := LoadTwitterUser()
	firstAccountAuth := getAuth(firstAccountKey, client)
	tx, err := twitterUserInstance.Login(firstAccountAuth, address, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s for login(%s, %s)\n", tx.Hash().Hex(), address.Hex(), password)
}

func Unregister(address common.Address) {
	client := CreateConn()
	twitterUserInstance := LoadTwitterUser()
	firstAccountAuth := getAuth(firstAccountKey, client)
	tx, err := twitterUserInstance.Unregister(firstAccountAuth, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s for unregister(%s)\n", tx.Hash().Hex(), address.Hex())
}

func CheckUserNumbers() {
	twitterUserInstance := LoadTwitterUser()
	cnt, err := twitterUserInstance.CheckUserNumbers(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d users", cnt)
}

func Unlogin(userAddress common.Address) {
	client := CreateConn()
	twitterUserInstance := LoadTwitterUser()
	firstAccountAuth := getAuth(firstAccountKey, client)
	tx, err := twitterUserInstance.Unlogin(firstAccountAuth, userAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s for unlogin(%s)\n", tx.Hash().Hex(), userAddress.Hex())

}

func LoadTwitter() *twitterapi.Twitter {
	client := CreateConn()
	address := common.HexToAddress(twitterAddress)
	instance, err := twitterapi.NewTwitter(address, client)

	if err != nil {
		log.Fatal(err)
	}
	return instance
}

func Post(content string, accountKey string, accountAddress string) {
	client := CreateConn()
	twitterInstance := LoadTwitter()
	secondAccountAuth := getAuth(accountKey, client)
	tx, err := twitterInstance.Post(secondAccountAuth, content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s for Post(%s) from %s\n", tx.Hash().Hex(), content, accountAddress)
}

func PostNumbers(from common.Address) {
	twitterInstance := LoadTwitter()
	posts, err := twitterInstance.PostNumbers(&bind.CallOpts{From: from})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("There are %d posts from %s", posts, from)
}

func ReadPost(from common.Address, index *big.Int) {
	twitterInstance := LoadTwitter()
	name, message, timestamp, err := twitterInstance.ReadPost(nil, from, index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %dth post from %s: name=%s, message=%s, timestamp=%d", index.Int64(), from.Hex(), name, message, timestamp)
}

func test() {
	// callings for TwitterUser
	Register(common.HexToAddress(secondAccountAddress), "user1", "123456")
	Register(common.HexToAddress(thirdAccountAddress), "user2", "78910")
	CheckUserInfo(common.HexToAddress(secondAccountAddress))
	CheckUserInfo(common.HexToAddress(thirdAccountAddress))
	Login(common.HexToAddress(secondAccountAddress), "12345")
	CheckUserInfo(common.HexToAddress(secondAccountAddress))
	CheckUserNumbers()

	//callings for Twitter
	Post("message1-1", secondAccountKey, secondAccountAddress)
	Post("message1-2", secondAccountKey, secondAccountAddress)
	PostNumbers(common.HexToAddress(secondAccountAddress))
	ReadPost(common.HexToAddress(secondAccountAddress), big.NewInt(0))
	ReadPost(common.HexToAddress(secondAccountAddress), big.NewInt(1))
	ReadPost(common.HexToAddress(secondAccountAddress), big.NewInt(2)) // name=, message=, timestamp=0

	Post("message2-1", thirdAccountKey, thirdAccountAddress) //exception: revert It is not logined
	Login(common.HexToAddress(thirdAccountAddress), "7891")  // exception: failed to login because of the mistched password
	Login(common.HexToAddress(thirdAccountAddress), "78910")
	CheckUserInfo(common.HexToAddress(thirdAccountAddress))
	Post("message2-1", thirdAccountKey, thirdAccountAddress)
	PostNumbers(common.HexToAddress(thirdAccountAddress))
	ReadPost(common.HexToAddress(thirdAccountAddress), big.NewInt(0))

	Unlogin(common.HexToAddress(secondAccountAddress))
	CheckUserInfo(common.HexToAddress(secondAccountAddress))
	Post("message1-3", secondAccountKey, secondAccountAddress) //exception:It is not logined
	Unregister(common.HexToAddress(secondAccountAddress))
	Login(common.HexToAddress(secondAccountAddress), "12345") //exception:address is not registered
	CheckUserNumbers()
}

var configFile string

var rootCmd = &cobra.Command{
	Use:   "go-microservice-template.exe",
	Short: "A template for gin based micro-service project",
	Long: `
	This template will use the following tools:
	- gin : Http web framework
	- fx : A dependency injection system
	- gorm : Database access solution
	- koanf : a simple, extremely lightweight, extensible, configuration management library
	- cobra : A CLI application
	- zap :  A fast, structured, leveled logging
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start web server from rootCmd ...")
		fmt.Println("configFile =", configFile)
		server.StartApplication(configFile)
	},
}

var subCommand = &cobra.Command{
	Use:     "server",
	Short:   "start web server",
	Aliases: []string{"s"},
	Args:    cobra.ExactArgs(0), // only 0 parameter for command1
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start web server from subCommand ...")
		fmt.Println("configFile =", configFile)
		server.StartApplication(configFile)
	},
}

func init() {
	rootCmd.AddCommand(subCommand)
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "c", "./config/properties.json", "config file path")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
