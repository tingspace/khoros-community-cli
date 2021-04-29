package app

import (
	"fmt"
	"io/ioutil"
	"khoros-community-cli/plumbing"
	"log"
	"net/http"
	"net/url"
)

// Env Keys
const (
	BaseURL    = "COMMUNITY_DOMAIN"
	APIVersion = "API_VERSION"
	APIKey     = "API_KEY"
)

// Available Commands
const (
	CmdKey  = "Command"
	CmdList = "list"
)

// Run determines Fn to call based on the supplied Command
func Run(cmn *plumbing.Common) {
	switch cmn.Commands[CmdKey] {
	case CmdList:
		getList(cmn)
		break
	default:
		defaultRequest(cmn)
		break
	}
}

type ListCollection struct {
	Results []interface{}
	Command string
	Domain  string
}

type ListFn func(*plumbing.Common) *ListCollection

// Available Collections
const (
	MessageCollection = "messages"
	BoardCollection   = "board"
	UserCollection    = "user"
)

func getList(cmn *plumbing.Common) {
	fn := findListFn(cmn.Commands["Collection"])
	cc := fn(cmn)
	fmt.Println(cc)
}

func findListFn(collection string) ListFn {
	switch collection {
	case MessageCollection:
		return GetMessageList
	default:
		return defaultListFn
	}

}

func defaultListFn(cmn *plumbing.Common) *ListCollection {
	cc := ListCollection{}
	return &cc
}

func defaultRequest(cmn *plumbing.Common) {
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
