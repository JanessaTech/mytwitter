package services

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hi-supergirl/mytwitter/config"
	"github.com/hi-supergirl/mytwitter/helper"
	"github.com/hi-supergirl/mytwitter/logging"
	twitterapi "github.com/hi-supergirl/mytwitter/twitterApi"
)

type TwitterUserService interface {
	Register(ctx context.Context, name string, password string) (string, error)
	Unregister(ctx context.Context, name string) error
	Login(ctx context.Context, name string, password string) error
	Unlogin(ctx context.Context, name string) error
	CheckUserInfo(ctx context.Context, name string) error
	CheckUserNumbers(ctx context.Context) int
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

func (tws *twitterUserService) Unregister(ctx context.Context, name string) error {
	return nil
}

func (tws *twitterUserService) Login(ctx context.Context, name string, password string) error {
	return nil
}

func (tws *twitterUserService) Unlogin(ctx context.Context, name string) error {
	return nil
}

func (tws *twitterUserService) CheckUserInfo(ctx context.Context, name string) error {
	return nil
}

func (tws *twitterUserService) CheckUserNumbers(ctx context.Context) int {
	return 0
}
