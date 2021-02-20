package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct{
	Nome string `json:"nome"`
	Raca string `json:"raca"`
	Idade string `json:"idade"s`
}


func main(){
	c := cachorro{"Mel", "Poddle", "16"}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))
}
