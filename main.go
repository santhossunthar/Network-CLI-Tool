package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network CLI Tool"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "example.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "nslookup",
			Usage: "Look up the name servers",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name: "cname",
			Usage: "Look up the CNAME records",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cnames, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return nil
				}

				fmt.Println(cnames)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
