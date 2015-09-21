package command

import (
  "fmt"
  "os"

  "github.com/bkono/dokku-cloudflare-src/api"
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

func validateRequiredArgs(c *cli.Context) {
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
}

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
