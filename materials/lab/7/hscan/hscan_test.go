// Optional Todo

package hscan

import (
	"testing"
)


//got rid of test already in here 

// func TestGuessSingle(t *testing.T) {
// 	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f") // Currently function returns only number of open ports
// 	want := "Nickelback4life"
// 	if got != want {
// 		t.Errorf("got %d, wanted %d", got, want)
// 	}

// }



func TestGenHashMapsWIHTOUTsubroutines(t *testing.T){
	//numPasswords := 303872
	GenHashMaps("/home/cabox/workspace/course-materials/materials/lab/7/main/Top304Thousand-probable-v2.txt")

}

func TestGenHashMapsWITHsubroutines(t *testing.T){
	//numPasswords := 303872
	GenHashMapsWithSubRoutines("/home/cabox/workspace/course-materials/materials/lab/7/main/Top304Thousand-probable-v2.txt")
	
}

//TODO create a test in hscan_test.go so that you can time the performance of your implementation
	//Test and record the time it takes to scan to generate these Maps
	// 1. With and without using go subroutines
	// 2. Compute the time per password (hint the number of passwords for each file is listed on the site...)