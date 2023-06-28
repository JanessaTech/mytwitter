package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hi-supergirl/mytwitter/config"
	"github.com/hi-supergirl/mytwitter/handlers"
	"github.com/hi-supergirl/mytwitter/handlers/services"
	"github.com/hi-supergirl/mytwitter/logging"
	twitterapi "github.com/hi-supergirl/mytwitter/twitterApi"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// Connect to the blockchain network.
// We will use this connection heavily to send transaction or query data on blockchain
func CreateConn(config *config.Config, logger *zap.Logger) *ethclient.Client {
	client, err := ethclient.Dial(config.Connection)
	if err != nil {
		log.Fatal(err)
	}
	logger.Sugar().Infoln("server is connected to the blockchain network")

	return client
}

func LoadTwitterUser(client *ethclient.Client, config *config.Config) *twitterapi.TwitterUser {

	address := common.HexToAddress(config.TwitterUserAddr)

	instance, err := twitterapi.NewTwitterUser(address, client)

	if err != nil {
		log.Fatal(err)
	}
	return instance
}

func LoadTwitter(client *ethclient.Client, config *config.Config) *twitterapi.Twitter {
	address := common.HexToAddress(config.TwitterAddr)
	instance, err := twitterapi.NewTwitter(address, client)

	if err != nil {
		log.Fatal(err)
	}
	return instance
}

func printBanner(logger *zap.Logger) {
	logger.Sugar().Infoln(`
****************************************************************
****************************************************************
***************JanessaTech's distributed Twitter**************
****************************************************************
****************************************************************`)
}

func Server(lc fx.Lifecycle, logger *zap.Logger) *gin.Engine {
	r := gin.Default()

	srv := &http.Server{Addr: ":8080", Handler: r} // define a web server
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				logger.Sugar().Infoln("Failed to start HTTP server at", srv.Addr)
				return err
			}
			go srv.Serve(ln)
			logger.Sugar().Infoln("Succeeded to start HTTP server at", srv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			srv.Shutdown(ctx)
			logger.Sugar().Infoln("HTTP server is stopped")
			return nil
		},
	})
	return r
}

func StartApplication(configFile string) {
	config, err := config.NewConfig(configFile)
	if err != nil {
		fmt.Println("Failed to read config file. Exit")
		return
	}

	app := fx.New(
		fx.Supply( // set initial variables needed by funcs defined in fx.Provide()
			logging.GetLogger(config),
			config,
		),
		fx.Invoke(printBanner),
		fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: logger.Named("JanessaTech's distributed twitter")}
		}),
		fx.Provide(
			CreateConn,
			LoadTwitterUser,
			LoadTwitter,
			Server,
			services.NewTwitterUserService,
			services.NewTwitterService,
			handlers.NewTwitterUserHandler,
			handlers.NewTwitterHander,
		),
		fx.Invoke(
			func(*gin.Engine) {},
			handlers.TwitterUserRoute,
			handlers.TwitterRoute,
		),
	)
	app.Run()
}
