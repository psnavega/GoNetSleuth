package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI application"
	app.Usage = "Search ips and server names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "globo.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPS on internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "Find the name of server",
			Flags:  flags,
			Action: SearchServer,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func SearchServer(c *cli.Context) {
	host := c.String("host")

	servers, error := net.LookupNS(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}
