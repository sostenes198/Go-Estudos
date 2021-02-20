package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct{
	Nome string `json:"-"`
	Raca string `json:"raca"`
	Idade string `json:"idade"s`
}

func main(){
	cachorroEmJSON := `{"nome":"Mel","raca":"Poddle","idade":"16"}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil{
		log.Fatal(erro)
	}

	fmt.Println(c)
}
