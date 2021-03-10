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

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Buscar IPs de endereço na Web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	//net
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
