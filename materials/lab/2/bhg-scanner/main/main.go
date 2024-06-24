package main

import "bhg-scanner/scanner"

func main(){
	var portsArray []int 
	portsArray = []int{34,87,490,616,65,839,902,200,891,409,10}
	scanner.PortScanner(portsArray)
}
