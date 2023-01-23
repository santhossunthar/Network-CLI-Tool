package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
	"main/port"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network CLI Tool"

	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "example.com",
		},
		cli.StringFlag{
			Name: "ip",
			Value: "127.0.0.1",
		},
		cli.StringFlag{
			Name: "port",
			Value: "80",
		},
		cli.StringFlag{
			Name: "protocol",
			Value: "tcp",
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
		{
			Name: "addrlookup",
			Usage: "Look up the address of the IP address",
			Flags: myFlags,
			Action: func (c *cli.Context) error  {
				addr, err := net.LookupAddr(c.String("ip"))
				if err != nil {
					return err
				}

				for i := 0; i < len(addr); i++ {
					fmt.Println(addr[i])
				}

				return nil
			},
		},
		{
			Name: "dnstxt",
			Usage: "Look up the DNS TXT records",
			Flags: myFlags,
			Action: func (c *cli.Context) error  {
				txts, err := net.LookupTXT(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(txts); i++ {
					fmt.Println(txts[i])
				}
				
				return nil
			},
		},
		{
			Name: "portscan",
			Usage: "Look up the open and closed ports",
			Flags: myFlags,
			Action: func (c *cli.Context) {
				port.ScanPorts(c.String("protocol"), c.String("host"))	
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
