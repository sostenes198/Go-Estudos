package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App{

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}

	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando do Soso"
	app.Usage = "Buscas Ips e Nomes de Servidores na internet"
	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Busca ips de endereços na internet",
			Flags: flags,
			Action: buscarIps,
		},
		{
			Name: "servidores",
			Usage: "Busca o nome do servidores na internet",
			Flags: flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context){
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil{
		log.Fatal(erro)
	}

	for _, ip := range ips{
		fmt.Println(ip)
		}
}

func buscarServidores(c *cli.Context){
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil{
		log.Fatal(erro)
	}

	for _, servidor := range servidores{
		fmt.Println(servidor.Host)
	}
}