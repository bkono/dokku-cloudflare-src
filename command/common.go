package command

import (
  "os"
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

