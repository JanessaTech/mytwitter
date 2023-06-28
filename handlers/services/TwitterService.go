package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hi-supergirl/mytwitter/config"
	"github.com/hi-supergirl/mytwitter/helper"
	"github.com/hi-supergirl/mytwitter/logging"
	twitterapi "github.com/hi-supergirl/mytwitter/twitterApi"
)

type TwitterService interface {
	Post(context context.Context, who string, message string) (string, error)
	PostNumbers(context context.Context, who string) (int, error)
	ReadPost(context context.Context, from string, index uint) (string, string, *big.Int, error)
}

func NewTwitterService(config *config.Config, client *ethclient.Client, twitter *twitterapi.Twitter) TwitterService {
	return &twitterService{config: config, client: client, twitter: twitter}
}

type twitterService struct {
	client  *ethclient.Client
	config  *config.Config
	twitter *twitterapi.Twitter
}

func (ts *twitterService) Post(context context.Context, who string, message string) (string, error) {
	logger := logging.FromContext(context)
	logger.Infow("TwitterService.Post", "who", who, "message", message)
	key, addr, err := helper.GetAccountByName(who, ts.config)
	if err != nil {
		return "", err
	}
	accountAuth, err := helper.GetAuth(key, ts.client)
	if err != nil {
		return "", err
	}
	tx, err := ts.twitter.Post(accountAuth, message)
	if err != nil {
		return "", err
	}
	logger.Infoln(fmt.Printf("tx sent: %s for Post(%s) from %s\n", tx.Hash().Hex(), message, addr))
	return tx.Hash().Hex(), nil
}

func (ts *twitterService) PostNumbers(context context.Context, who string) (int, error) {
	logger := logging.FromContext(context)
	logger.Infow("TwitterService.Post", "who", who)
	_, addr, err := helper.GetAccountByName(who, ts.config)
	if err != nil {
		return 0, err
	}

	posts, err := ts.twitter.PostNumbers(&bind.CallOpts{From: common.HexToAddress(addr)})
	if err != nil {
		return 0, err
	}
	logger.Infoln(fmt.Printf("There are %d posts from %s", posts, addr))
	return int(posts.Int64()), nil
}

func (ts *twitterService) ReadPost(context context.Context, from string, index uint) (string, string, *big.Int, error) {
	logger := logging.FromContext(context)
	_, addr, err := helper.GetAccountByName(from, ts.config)
	if err != nil {
		return "", "", nil, err
	}
	logger.Infow("TwitterService.ReadPost", "from", from, "index", index)
	name, message, timestamp, err := ts.twitter.ReadPost(nil, common.HexToAddress(addr), big.NewInt(int64(index)))
	if err != nil {
		return "", "", nil, err
	}

	logger.Infoln(fmt.Printf("Read %dth post from %s: name=%s, message=%s, timestamp=%d", index, from, name, message, timestamp))
	return name, message, timestamp, nil
}
