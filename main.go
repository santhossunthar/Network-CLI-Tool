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
					return err
				}

				fmt.Println(cnames)
				return nil
			},
		},
		{
			Name: "iplookup",
			Usage: "Look up the IP Addresses",
			Flags: myFlags,
			Action: func (c *cli.Context) error  {
				ips, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ips); i++ {
					fmt.Println(ips[i])
				}

				return nil
			},
		},
		{
			Name: "mxlookup",
			Usage: "Look up the MX records",
			Flags: myFlags,
			Action: func (c *cli.Context) error  {
				mxs, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(mxs); i++ {
					fmt.Println(mxs[i].Host, mxs[i].Pref)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
