package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Create will return the cli application ready to run
func Create() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplication CLI"
	app.Usage = "Search IPs and Server Names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs on the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Search Server Names on the internet",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
