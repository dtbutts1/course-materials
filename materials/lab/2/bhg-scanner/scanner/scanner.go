// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: University of Wyoming cybersecurity class 
// 		   Run test in main with: go build -> ./main
//		   Run test in scanner_test.go with: go test

// Name: Dawson Butts


package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)




func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)    
		conn, err := net.DialTimeout("tcp", address, 1 * time.Second) 
		if err != nil { 
			results <- -1*p 	//if closed port, make it negative 
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails

func PortScanner(start, end int) int {  //ADD RANGE TO SCAN OVER WITH INTS

	//TODO 3 : ADD closed ports; currently code only tracks open ports
	// MOVED TO THIS FUNCTION
	var openports []int  // notice the capitalization here. access limited!
	var closedports []int //keeps tracked of closed ports 

	// fmt.Printf("Start: %d\n", start)
	// fmt.Printf("End: %d\n", end)

	ports := make(chan int, 100)   
	results := make(chan int)

	for i := start; i <= end; i++ {
		go worker(ports, results)
	}

	go func() {
		for i := start; i <= end; i++ {
			ports <- i
		}
	}()

	for i := start; i <= end; i++ {
		port := <-results
		if port > 0 { 	//is open port
			openports = append(openports, port)
		}else{
						//save the closed port
			closedports = append(closedports, -1*port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	sort.Ints(closedports)		//sort the closed ports as well


	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("Port %d OPEN\n", port)
	}
	// ALSO PRINT OUT CLOSED PORTS 
	for _, port := range closedports {
		fmt.Printf("Port %d CLOSED\n", port)
	}

	//RETURN ADDED TO THE CLOSED PORTS
	return len(openports) + len(closedports) // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
