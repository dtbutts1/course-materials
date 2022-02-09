package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){

    got := PortScanner(10, 150) // Currently function returns only number of open ports
    want := 141 // default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns 
	          //consider what would happen if you parameterize the portscanner address and ports to scan

    if got != want {
        t.Errorf("got %d, wanted %d", got, want)
    }
}

func TestTotalPortsScanned(t *testing.T){
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()

    got := PortScanner(20, 100) // Currently function returns only number of open ports
    want := 81 // default value; consider what would happen if you parameterize the portscanner ports to scan

    if got != want {
        t.Errorf("got %d, wanted %d\n", got, want)
    }

	//ADDED NEW TESTS HERE
	got2 := PortScanner(3, 8)
	want2 := 6

	if got2 != want2 {
		t.Errorf("got %d, wanted %d\n", got2, want2)
	}

	got3 := PortScanner(1, 1500)
	want3 := 1500

	if got3 != want3 {
		t.Errorf("got %d, wanted %d\n", got3, want3)
	}

	got4 := PortScanner(35, 47)
	want4 := 13

	if got4 != want4 {
		t.Errorf("got %d, wanted %d\n", got4, want4)
	}
}


