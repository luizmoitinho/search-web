package app

import "github.com/urfave/cli"

//Generate ... return app line comand to execute
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nome de servidores na Web"

	return app
}
