package api

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "github.com/codegangsta/cli"
)

const (
  apiUrl = "https://www.cloudflare.com/api_json.html"
)

type Client struct {
  Email, Token, Zone string
}

func BuildClient(c *cli.Context) (client *Client) {
  return &Client{
    Email: c.String("email"),
    Token: c.String("token"),
    Zone: c.String("zone"),
  }
}

func (client *Client) Post(action string, params map[string]string) (resp *http.Response, err error) {
  postParms := url.Values{}
  postParms.Add("a", action)
  postParms.Add("tkn", client.Token)
  postParms.Add("email", client.Email)
  postParms.Add("z", client.Zone)

  if len(params) > 0 {
    for k, v := range params {
      postParms.Add(k, v)
    }
  }

  return checkResp(http.PostForm(apiUrl, postParms))
  // // u.RawQuery = p.Encode()
  // // req, err := http.NewRequest("POST", u.String(), nil)
  // // resp, err := checkResp(http.DefaultClient.Do(req))
  // assert(err)
}

func (client *Client) PostAndParse(action string, params map[string]string, out interface{}) (err error) {
  resp, err := client.Post(action, params)
  return decodeResponse(resp, &out, err)
}

func checkResp(resp *http.Response, err error) (*http.Response, error) {
  // Already had an error, just spit it back
  if err != nil {
    return resp, err
  }

  switch i := resp.StatusCode; {
  case i == 200:
    return resp, nil
  default:
    return nil, fmt.Errorf("API Error: %s", resp.Status)
  }
}

func decodeResponse(resp *http.Response, out interface{}, err error) error {
  // Already had an error, just spit it back
  if err != nil {
    return err
  }

  contents, err := ioutil.ReadAll(resp.Body)
  if err = json.Unmarshal(contents, &out); err != nil {
    return err
  }

  return nil
}
