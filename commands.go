package main

import (
  "fmt"
  "os"

  "github.com/bkono/dokku-cloudflare-src/command"
  "github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
  {
    Name:   "add",
    Usage:  "",
    Action: command.CmdAdd,
    Flags:  []cli.Flag{
      cli.StringFlag{
        Name: "email",
        Value: "",
        Usage: "Cloudflare account email",
        EnvVar: "CLOUDFLARE_EMAIL",
      },
      cli.StringFlag{
        Name: "token, t",
        Value: "",
        Usage: "Cloudflare API token",
        EnvVar: "CLOUDFLARE_TOKEN",
      },
      cli.StringFlag{
        Name: "zone, z",
        Value: "",
        Usage: "Zone to add the records",
        EnvVar: "CLOUDFLARE_ZONE",
      },
      cli.StringFlag{
        Name: "domain, d",
        Value: "",
        Usage: "Domain record to add",
      },
      cli.StringFlag{
        Name: "ip",
        Value: "",
        Usage: "IP address to set. Will default to current host's public ip",
      },
    },
  },
  {
    Name:   "list",
    Usage:  "",
    Action: command.CmdList,
    Flags:  []cli.Flag{
      cli.StringFlag{
        Name: "email",
        Value: "",
        Usage: "Cloudflare account email",
        EnvVar: "CLOUDFLARE_EMAIL",
      },
      cli.StringFlag{
        Name: "token, t",
        Value: "",
        Usage: "Cloudflare API token",
        EnvVar: "CLOUDFLARE_TOKEN",
      },
      cli.StringFlag{
        Name: "zone, z",
        Value: "",
        Usage: "Zone to to scope the list",
        EnvVar: "CLOUDFLARE_ZONE",
      },
    },
  },
  {
    Name:   "delete",
    Usage:  "",
    Action: command.CmdDelete,
    Flags:  []cli.Flag{},
  },
}

func CommandNotFound(c *cli.Context, command string) {
  fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
  os.Exit(2)
}
