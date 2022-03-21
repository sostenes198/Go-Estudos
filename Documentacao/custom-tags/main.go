package main

import (
	"fmt"
	"myexamples.com/custom_tags/validators"
)

type User struct {
	Name  string `validate:"string,required=true,min=2,max=10"`
	Name1 string `validate:"string,required=false"`
	Name5 string `validate:"string,required=true"`
	Name2 string `validate:"string,required=true,min=2,max=10"`
	Name3 string `validate:"string,required=true,min=0,max=2"`
	Age   int
}

func main() {
	user := User{Name: "", Name1: "", Name2: "a", Name3: "aaaa"}

	isValid, errs := validators.ValidateStruct(user)
	if !isValid {
		for i, err := range errs {
			fmt.Printf("\t%d. %s\n", i+1, err.Error())
		}
	}
}
