package command

import (
  "io/ioutil"
  "net/http"
  "os"
  "regexp"

  "github.com/bkono/dokku-cloudflare-src/api"
  "github.com/codegangsta/cli"
)

func CmdAdd(c *cli.Context) {
	// Write your code here
  validateRequiredArgs(c)
  domain := c.String("domain")
  ip := c.String("ip")
  var err error

  if ip == "" {
    ip, err = getPublicIp()
  }
  assert(err)

  client := api.BuildClient(c)
  params := make(map[string]string)
  params["type"] = "A"
  params["name"] = domain
  params["content"] = ip
  params["ttl"] = "1"

  var response api.CloudFlareRecordResponse
  err = client.PostAndParse("rec_new", params, &response)

  if err != nil || response.Result == "error" {
    fatal("Error while creating a record: " + response.Msg)
    os.Exit(2)
  }

  println("Successfully created record name=" + response.Response.Rec.Record.Name + " ip=" + response.Response.Rec.Record.Content)

}

func getPublicIp() (ip string, err error) {
  resp, err := http.Get("http://icanhazip.com")
  assert(err)

  contents, err := ioutil.ReadAll(resp.Body)
  ip = string(contents)
  ip = findIp(ip)
  return ip, err
}

func findIp(input string) string {
  numBlock := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
  regexPattern := numBlock + "\\." + numBlock + "\\." + numBlock + "\\." + numBlock

  regEx := regexp.MustCompile(regexPattern)
  return regEx.FindString(input)
}
