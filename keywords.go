package go_papi

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  simplejson "github.com/bitly/go-simplejson"
)

var p = fmt.Println

func (client *Client) KeywordsPost(keyword string) (*simplejson.Json, error) {
  resp, err := http.PostForm("http://api.authoritylabs.com/keywords",
    url.Values{"keyword": {keyword}, "auth_token": {client.ApiKey}})
  if err != nil {
    p("ERROR")
  }
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  return simplejson.NewJson(bodyBytes)
}

func (client *Client) KeywordsPriorityPost(keyword string) (*simplejson.Json, error) {
  resp, err := http.PostForm("http://api.authoritylabs.com/keywords/priority",
    url.Values{"keyword": {keyword}, "auth_token": {client.ApiKey}})
  if err != nil {
    p("ERROR")
  }
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  return simplejson.NewJson(bodyBytes)
}

func (client *Client) KeywordsGet(keyword string) (*simplejson.Json, error) {
  getParams := url.Values{"keyword": {keyword}, "auth_token": {client.ApiKey}}.Encode()
  p(getParams)
  resp, err := http.Get(fmt.Sprintf("http://api.authoritylabs.com/keywords/get?%s", getParams))
  if err != nil {
    p("ERROR")
  }
  defer resp.Body.Close()
  bodyBytes, _ := ioutil.ReadAll(resp.Body)
  return simplejson.NewJson(bodyBytes)
}
