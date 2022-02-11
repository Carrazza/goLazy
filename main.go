package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	
)

//Essa struct possui variáveis de execução para os comandos shell
type shellArgs struct {
	domain string
	wgPointer *sync.WaitGroup
	outChannel *chan string
}

func main() {
	
	//Pegando os argumentos
	sysArgs := os.Args[1:]
	
	//Criando os grupos de espera para ajeitar as concorrências
	var wg sync.WaitGroup

	wg.Add(2)

	

	//Criando o canal para adicionar os subdomínios encontrados 

	//domain := sysArgs[0]
	
	shellArgs := shellArgs{ 
		domain: sysArgs[0] , 
		wgPointer: &wg ,
	}

	fmt.Println("domain:", shellArgs.domain)

	//Chamando as funções de enumeração

	out1 := subfinder(&shellArgs)
	out2 := assetfinder(&shellArgs)

	//------------------------------------------------

	wg.Wait()
	output := out1 + out2

	// Sorting a lista para só resultados únicos
	s := strings.Split(output,"\n")
	s = sort(s)

	//------------------------------------------------

	fmt.Println(strings.Join(s, "\n"))
	fmt.Printf("%T\n",s)
	//fmt.Println(output)
	
}

//Sort output
func sort(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}