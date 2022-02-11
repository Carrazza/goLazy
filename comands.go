package main

import (
	"fmt"
	"os/exec"
	
)


func assetfinder(args shellArgs) string {

	//fmt.Printf("HelloWorld")

	cmd, err := exec.Command("assetfinder", args.domain).Output()

	if err != nil {
		fmt.Printf("error %s", err)
	}
	return string(cmd)

}

func amass(args shellArgs) string{

	cmdArgs := []string{"enum", "-noalts", "-passive", "-d", args.domain}
	
	cmd, err := exec.Command("amass", cmdArgs...).Output()

	if( err != nil ){fmt.Println(err)}
	
	return string(cmd)

}

func subfinder(args shellArgs) string{

	cmdArgs := []string{"-silent", "-d",args.domain}
	
	cmd, err := exec.Command("subfinder", cmdArgs...).Output()

	if( err != nil ){fmt.Println(err)}
	
	return string(cmd)
}