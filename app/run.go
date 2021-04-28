package app

import (
	"fmt"
	"io/ioutil"
	"khoros-community-cli/plumbing"
	"log"
	"net/http"
	"net/url"
)

// Keys in Config Map
const (
	BaseURL    = "COMMUNITY_DOMAIN"
	APIVersion = "API_VERSION"
	APIKey     = "API_KEY"
)

// Run is the start of Client
func Run(cmn *plumbing.Common) {
	url := constructURL(cmn.Keys)
	params := constructParams()

	resp, err := http.Get(url + params)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	fmt.Println(sb)
}

func constructURL(k map[string]string) string {
	return k[BaseURL] + "/api/" + k[APIVersion] + "/search?q="
}

func constructParams() string {
	params := "SELECT count(*) FROM messages WHERE author.login='tingspace'"
	return url.QueryEscape(params)
}
