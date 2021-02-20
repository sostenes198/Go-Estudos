package main

import (
	"Introducao/Enderecos"
	"fmt"
)

func main(){
	tipoEnderecao := Enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEnderecao)
}
