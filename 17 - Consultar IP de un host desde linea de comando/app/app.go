package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar va a retornar la aplicacion lista para ser ejecutada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicación de linea de comandos"
	app.Usage = "Buscas ips y nombres de servidor en internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "localhost",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca ips de dominios de internet",
			Flags:  flags,
			Action: buscarIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca los servidores en los que se aloja un dominio",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
