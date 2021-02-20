package core

import "log"

func TratarErro(erro error){
	if erro != nil{
		log.Fatal(erro)
	}
}