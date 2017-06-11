package main

import (
	"fmt"
	"encoding/json"
)

type Usuario struct{
	Nome string
	SobreNome string
}

func main(){
	usuario := Usuario{"vinicius","dias das silva"}
	usuAux,_ := json.Marshal(usuario)
	fmt.Println(string(usuAux))
}
