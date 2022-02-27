package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

const (
	CommandIp = "ip"
	CommandServidores = "servidores"
)

const (
	FlagHost = "host"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidores na internet"

	app.Commands = []cli.Command{
		{
			Name:  CommandIp,
			Usage: "Busca IPS de endereços na internet",
			Flags: flags(),
			Action: buscarIps,
		},
		{
			Name: CommandServidores,
			Usage: "Buscar nome dos servidores",
			Flags: flags(),
			Action: buscarServidores,
		},
	}

	return app
}

func flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  FlagHost,
			Value: "devbook.com.br",
		},
	}
}

func buscarIps(c *cli.Context) {
	host := c.String(FlagHost)

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatalln(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context){
	host := c.String(FlagHost)

	servidores, error := net.LookupNS(host)
	if error != nil {
		log.Fatalln(error)
	}

	for _, servidor := range servidores{
		fmt.Println(servidor.Host)
	}
}
