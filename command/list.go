package command

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "os"

  "github.com/bkono/dokku-cloudflare-src/types"
  "github.com/codegangsta/cli"
)

func assert(err error) {
  if err != nil {
    println("!!", err.Error())
    os.Exit(2)
  }
}

func fatal(msg string) {
  println("!!", msg)
  os.Exit(2)
}

func checkResp(resp *http.Response, err error) (*http.Response, error) {
  // If the err is already there, there was an error higher
  // up the chain, so just return that
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

func CmdList(c *cli.Context) {
  email, token, zone := c.String("email"), c.String("token"), c.String("zone")
  if zone == "" {
    fatal("A zone is required")
  }
  if email == "" {
    fatal("An email is required")
  }
  if token == "" {
    fatal("A token is required")
  }

  u, err := url.Parse("https://www.cloudflare.com/api_json.html")
  assert(err)
  p := url.Values{}
  p.Add("a", "rec_load_all")
  p.Add("tkn", token)
  p.Add("email", email)
  p.Add("z", zone)

  u.RawQuery = p.Encode()
  req, err := http.NewRequest("POST", u.String(), nil)
  resp, err := checkResp(http.DefaultClient.Do(req))
  assert(err)
  contents, err := ioutil.ReadAll(resp.Body)

  var response cloudflare.CloudflareListResponse
  err = json.Unmarshal(contents, &response)
  fmt.Printf("%+v\n", response)
}
