package main

import (
	"fmt"
	"os"
	"strings"
	
)

func main() {
	
	//Pegando os argumentos, sem o programa
	argsWithoutProg := os.Args[1:]

	domain := argsWithoutProg[0]


	output := subfinder(domain)
	output += assetfinder(domain)

	s := strings.Split(output,"\n")

	s = sort(s)
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