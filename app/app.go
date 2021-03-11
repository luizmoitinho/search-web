package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Generate ... return app line comand to execute
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nome de servidores na Web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Buscar IPs de endereço na Web",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Busca nome dos servidores na Web",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
