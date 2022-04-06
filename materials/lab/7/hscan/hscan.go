package hscan

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

//==========================================================================\\

var shalookup  = make(map[string]string)
var md5lookup  = make(map[string]string)

var wg sync.WaitGroup

var hash string

func GuessSingle(sourceHash string, filename string) {

	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text()

		// TODO - From the length of the hash you should know which one of these to check ...
		// add a check and logicial structure
		
		if(len(sourceHash) == 32){
			//it is a md5 hash
			hash := fmt.Sprintf("%x", md5.Sum([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (MD5): %s\n", password)
			}
		}else {
			//it is a sha256 hash
			hash = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
			if hash == sourceHash {
				fmt.Printf("[+] Password found (SHA-256): %s\n", password)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func GenHashMaps(filename string) {

	//TODO
	//DONE - itterate through a file (look in the guessSingle function above)
	//rather than check for equality add each hash:passwd entry to a map SHA and MD5 where the key = hash and the value = password
	//TODO at the very least use go subroutines to generate the sha and md5 hashes at the same time
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text() // grab password

		//do for hashmd5
		hashmd5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		//fmt.Printf("newHASH added: %s\n", hashmd5)
		md5lookup[hashmd5] = password

		//now for hashsha
		hashsha := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		//fmt.Printf("newHASH added: %s\n", hashsha)
		shalookup[hashsha] = password

	}
}

func GenHashMapsWithSubRoutines(filename string){
	
	wg.Add(2)
	go GenerateMD5s(filename)
	go GenerateSHAs(filename)
	wg.Wait()
	fmt.Println("finished the withSubroutiens call")
}

func GenerateMD5s(filename string){
	defer wg.Done()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text() // grab password

		//do for hashmd5
		hashmd5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
		md5lookup[hashmd5] = password

	}
	fmt.Println("actually finished md5 hashmap")
}

func GenerateSHAs(filename string){
	defer wg.Done()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		password := scanner.Text() // grab password

		//now for hashsha
		hashsha := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
		shalookup[hashsha] = password

	}
	fmt.Println("actually finished sha hashmap")
}



func GetSHA(hash string) (string, error) {
	password, ok := shalookup[hash]
	if ok {
		fmt.Printf("[+] Password (found with hashmap) (SHA): %s\n", password)
		return password, nil

	} else {
		fmt.Printf("did not find password in hashmap\n")
		return "", errors.New("password does not exist")

	}
}

func GetMD5(hash string) (string, error) {
	password, ok := md5lookup[hash]
	if ok {
		fmt.Printf("[+] Password (found with hashmap) (MD5): %s\n", password)
		return password, nil
	}else{
		fmt.Printf("did not find password in hashmap\n")
		return "", errors.New("password does not exist")
	}
}
