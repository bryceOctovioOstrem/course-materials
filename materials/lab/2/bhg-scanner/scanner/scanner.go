// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int  // notice the capitalization here. access limited!
var closedports []int // array to hold closed ports

func worker(thePort int, results chan int) {
	//for Port := range thePort {
		
		address := fmt.Sprintf("scanme.nmap.org:%d", thePort)    
		conn, err := net.DialTimeout("tcp", address, 20000 * time.Millisecond) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		if err != nil { 
			results <- 0
			//continue
		}
		conn.Close()
		results <- thePort
	//}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object 
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
func PortScanner(thePorts []int) (int, int){  

	ports := make(chan int, 100)   // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE
	results := make(chan int)

	for  thePort:= range thePorts {
		go worker(thePort, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}else { // runs if detecting a closed port
			closedports = append(closedports, port) // appends port number closed port 
		}
	}

	//close(ports)
	//close(results)
	sort.Ints(openports)
	sort.Ints(closedports) // sorts closed port slice for printing
	//TODO 5 : Enhance the output for easier consumption, include closed ports
	for ports := range closedports { // iterates though closed ports
		fmt.Printf("%d closed\n", ports) //prints: closed *port number*
	}
	for ports := range openports { // iterates though closed ports
		fmt.Printf("%d open\n", ports) //prints: closed *port number*
	}
	
	/*for _, ports := range ports {
		if ports != 0{ // runs if open port
			fmt.Printf("%d open\n", ports) //prints: open *port number*
		} else{ // runs if closed port
			fmt.Printf("%d closed\n", ports) //prints: closed *port number*
		}
		
	}*/
	close(ports)
	close(results)
	return len(closedports) , len(openports) // TODO 6 : Return total number of ports scanned (number open, number closed); 
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
