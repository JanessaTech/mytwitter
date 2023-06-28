package services

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hi-supergirl/mytwitter/config"
	"github.com/hi-supergirl/mytwitter/helper"
	"github.com/hi-supergirl/mytwitter/logging"
	twitterapi "github.com/hi-supergirl/mytwitter/twitterApi"
)

type TwitterUserService interface {
	Register(ctx context.Context, name string, password string) (string, error)
	Unregister(ctx context.Context, name string) (string, error)
	Login(ctx context.Context, name string, password string) (string, error)
	Unlogin(ctx context.Context, name string) (string, error)
	CheckUserInfo(ctx context.Context, name string) (string, bool, error)
	CheckUserNumbers(ctx context.Context) (int, error)
}

func NewTwitterUserService(client *ethclient.Client, config *config.Config, twitterUser *twitterapi.TwitterUser) TwitterUserService {
	return &twitterUserService{client: client, config: config, twitterUser: twitterUser}
}

type twitterUserService struct {
	client      *ethclient.Client
	config      *config.Config
	twitterUser *twitterapi.TwitterUser
}

func (tws *twitterUserService) Register(ctx context.Context, name string, password string) (string, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("TwitterUserService.Register", "name", name, "password", password)
	_, addr, err := helper.GetAccountByName(name, tws.config)
	if err != nil {
		return "", err
	}
	adminKey, _, err := helper.GetAccountByName("admin", tws.config)
	if err != nil {
		return "", err
	}

	adminAuth, err := helper.GetAuth(adminKey, tws.client)
	if err != nil {
		return "", err
	}
	tx, err := tws.twitterUser.Register(adminAuth, common.HexToAddress(addr), name, password)
	if err != nil {
		return "", err
	}
	logger.Infoln(fmt.Sprintf("tx sent: %s for register(%s, %s, %s)\n", tx.Hash().Hex(), addr, name, password))

	return tx.Hash().Hex(), nil
}

func (tws *twitterUserService) Unregister(ctx context.Context, name string) (string, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("TwitterUserService.Unregister", "name", name)
	_, addr, err := helper.GetAccountByName(name, tws.config)
	if err != nil {
		return "", err
	}
	adminKey, _, err := helper.GetAccountByName("admin", tws.config)
	if err != nil {
		return "", err
	}

	adminAuth, err := helper.GetAuth(adminKey, tws.client)
	if err != nil {
		return "", err
	}
	tx, err := tws.twitterUser.Unregister(adminAuth, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}

	logger.Infoln(fmt.Printf("tx sent: %s for unregister(%s)\n", tx.Hash().Hex(), addr))
	return tx.Hash().Hex(), nil
}

func (tws *twitterUserService) Login(ctx context.Context, name string, password string) (string, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("TwitterUserService.Login", "name", name, "password", password)
	_, addr, err := helper.GetAccountByName(name, tws.config)
	if err != nil {
		return "", err
	}
	adminKey, _, err := helper.GetAccountByName("admin", tws.config)
	if err != nil {
		return "", err
	}

	adminAuth, err := helper.GetAuth(adminKey, tws.client)
	if err != nil {
		return "", err
	}
	tx, err := tws.twitterUser.Login(adminAuth, common.HexToAddress(addr), password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s for login(%s, %s)\n", tx.Hash().Hex(), addr, password)
	return tx.Hash().Hex(), nil
}

func (tws *twitterUserService) Unlogin(ctx context.Context, name string) (string, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("TwitterUserService.Unlogin", "name", name)
	_, addr, err := helper.GetAccountByName(name, tws.config)
	if err != nil {
		return "", err
	}
	adminKey, _, err := helper.GetAccountByName("admin", tws.config)
	if err != nil {
		return "", err
	}

	adminAuth, err := helper.GetAuth(adminKey, tws.client)
	if err != nil {
		return "", err
	}
	tx, err := tws.twitterUser.Unlogin(adminAuth, common.HexToAddress(addr))
	if err != nil {
		return "", err
	}

	logger.Infoln(fmt.Printf("tx sent: %s for Unlogin(%s)\n", tx.Hash().Hex(), addr))
	return tx.Hash().Hex(), nil
}

func (tws *twitterUserService) CheckUserInfo(ctx context.Context, name string) (string, bool, error) {
	logger := logging.FromContext(ctx)
	logger.Infow("TwitterUserService.CheckUserInfo", "name", name)
	_, addr, err := helper.GetAccountByName(name, tws.config)
	if err != nil {
		return "", false, err
	}
	name, _, isLogin, err := tws.twitterUser.CheckUserInfo(nil, common.HexToAddress(addr))
	if err != nil {
		return "", false, err
	}

	logger.Infow("TwitterUserService.CheckUserInfo Result:", "address", addr, "name", name, "isLogin", isLogin)
	return name, isLogin, nil
}

func (tws *twitterUserService) CheckUserNumbers(ctx context.Context) (int, error) {
	logger := logging.FromContext(ctx)
	logger.Infoln("TwitterUserService.Login")
	cnt, err := tws.twitterUser.CheckUserNumbers(nil)
	if err != nil {
		return 0, err
	}
	logger.Infoln(fmt.Printf("There are %d users", cnt))
	return int(cnt.Int64()), nil
}
