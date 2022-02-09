package main

import (
	"bhg-scanner/scanner"
	"fmt"
)

func main(){
	var numberReturned int = scanner.PortScanner(20, 50)
	
	if numberReturned == 31 {
		fmt.Printf("Test passed\n")
	}else{
		fmt.Printf("Test FAILED: actual %d\n", numberReturned)
	}
}