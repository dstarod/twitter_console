package main

import(
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"io/ioutil"
	"encoding/json"
)

type Keys struct{
	ConsumerKey string `json:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret"`
	AccessToken string `json:"access_token"`
	AccessSecret string `json:"access_secret"`
}

func main(){
	// Command line arguments
	tracking := ""
	if len(os.Args)>1{
		tracking = os.Args[1]
	}
	
	// File keys.json must exists
	j, err := ioutil.ReadFile("keys.json")
	if err != nil{
		fmt.Println("Need file keys.json")
		os.Exit(0)
	}
	
	// Format of keys.json must be similar with pattern
	keys := &Keys{}
	json.Unmarshal(j, &keys)
	
	if keys.ConsumerKey == "" || keys.ConsumerSecret == "" || keys.AccessToken == "" || keys.AccessSecret == "" {
		fmt.Println("Expected keys.json with format like:")
		fmt.Println(`
{
	"consumer_key": "your-consumer-key",
	"consumer_secret" : "your-consumer-secret",
	"access_token": "your-access-token",
	"access_secret": "your-access-secret"
}
		`)
		os.Exit(0)
	}
		
	// Make twitter API connection	
	anaconda.SetConsumerKey(keys.ConsumerKey)
	anaconda.SetConsumerSecret(keys.ConsumerSecret)
	api := anaconda.NewTwitterApi(keys.AccessToken, keys.AccessSecret)

	// By default read sample feed
	stream := api.PublicStreamSample(nil)
	if tracking != "" {
		// Make conditions for tracking
		v := url.Values{}
		v.Set("track", tracking)
		stream = api.PublicStreamFilter(v)	
	}
	
	for tweet := range stream.C{
		t, ok := tweet.(anaconda.Tweet)
		if ok{
			fmt.Printf("%v : %v\n", t.User.ScreenName, t.Text)			
		}
	}
	
}