package main

import (
	"flag"
	"log"
	"fmt"
	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func twitterClientInitialize() *twitter.Client {

	flags := struct {
		consumerKey    string
		consumerSecret string
		accessToken	   string
		accessSecret   string
	}{}

	flag.StringVar(&flags.accessToken, "access-token", "", "Twitter Access Token")
	flag.StringVar(&flags.accessSecret, "access-secret", "", "Twitter Access Secret")
	flag.StringVar(&flags.consumerKey, "consumer-key", "", "Twitter Consumer Key")
	flag.StringVar(&flags.consumerSecret, "consumer-secret", "", "Twitter Consumer Secret")
	flag.Parse()
	flagutil.SetFlagsFromEnv(flag.CommandLine, "TWITTER")

	if flags.consumerKey == "" || flags.consumerSecret == "" || flags.accessToken == "" || flags.accessSecret == ""{
		log.Fatal("Application & User Access Token required")
	}

	config := oauth1.NewConfig(flags.consumerKey, flags.consumerSecret)
	token := oauth1.NewToken(flags.accessToken, flags.accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	return twitter.NewClient(httpClient)
}

func tweet(data string) bool {
	// Initialize HTTP client
	client := twitterClientInitialize()
	// Send a Tweet
	_, _, err := client.Statuses.Update(data, nil)
	if err != nil {
		fmt.Println("Error encountered while sending tweet", err)
		return false
	}
	return true
}