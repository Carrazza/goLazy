package main

import (
	"fmt"
	"os/exec"
	"os"
)

func main() {
	
	//Pegando os argumentos, sem o programa
	argsWithoutProg := os.Args[1:]

	domain := argsWithoutProg[0]


	output := subfinder(domain)
	output += assetfinder(domain)


	fmt.Println(output)
	//fmt.Println(output)
	
}


//Todos os comandos usados para enumerar 


func assetfinder(domain string) string {

	//fmt.Printf("HelloWorld")

	cmd, err := exec.Command("assetfinder", domain).Output()

	if err != nil {
		fmt.Printf("error %s", err)
	}
	return string(cmd)

}

func amass(domain string) string{

	cmdArgs := []string{"enum", "-noalts", "-passive", "-d", domain}
	
	cmd, err := exec.Command("amass", cmdArgs...).Output()

	if( err != nil ){fmt.Println(err)}
	
	return string(cmd)

}

func subfinder(domain string) string{

	cmdArgs := []string{"-silent", "-d",domain}
	
	cmd, err := exec.Command("subfinder", cmdArgs...).Output()

	if( err != nil ){fmt.Println(err)}
	
	return string(cmd)
}

//--------------------------------------