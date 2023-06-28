package services

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hi-supergirl/mytwitter/config"
	twitterapi "github.com/hi-supergirl/mytwitter/twitterApi"
)

type TwitterService interface {
	Post(who string, message string)
	PostNumbers(who string)
	ReadPost(from string, index int)
}

func NewTwitterService(config *config.Config, client *ethclient.Client, twitter *twitterapi.Twitter) TwitterService {
	return &twitterService{config: config, client: client, twitter: twitter}
}

type twitterService struct {
	client  *ethclient.Client
	config  *config.Config
	twitter *twitterapi.Twitter
}

func (ts *twitterService) Post(who string, message string) {

}

func (ts *twitterService) PostNumbers(who string) {

}

func (ts *twitterService) ReadPost(from string, index int) {

}
