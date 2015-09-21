package command

import (
  "fmt"
  "os"

  "github.com/bkono/dokku-cloudflare-src/api"
  "github.com/codegangsta/cli"
)

func CmdList(c *cli.Context) {
  validateRequiredArgs(c)
  client := api.BuildClient(c)
  params := make(map[string]string)
  var response api.CloudflareListResponse
  err := client.PostAndParse("rec_load_all", params, &response)

  if err != nil || response.Result == "error" {
    fatal("Error while retrieving records: " + response.Msg)
    os.Exit(2)
  }

  fmt.Printf("%+v\n", response)
}
