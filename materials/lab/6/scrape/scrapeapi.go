package scrape

// scrapeapi.go HAS TEN TODOS - TODO_5-TODO_14 and an OPTIONAL "ADVANCED" ASK

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "os"
    "path/filepath"
    "strconv"
	"github.com/gorilla/mux"
	"regexp"
)

var dontAdd bool = false	//whether it is duplicate
var numOfHits int = 0	//number of files
var rootDir string = "/home/cabox"


//==========================================================================\\

// Helper function walk function, modfied from Chap 7 BHG to enable passing in of
// additional parameter http responsewriter; also appends items to global Files and 
// if responsewriter is passed, outputs to http 

func walkFn(w http.ResponseWriter) filepath.WalkFunc {
    return func(path string, f os.FileInfo, err error) error {
        w.Header().Set("Content-Type", "application/json")

        for _, r := range regexes {
            if r.MatchString(path) {
                var tfile FileInfo
                dir, filename := filepath.Split(path)
                tfile.Filename = string(filename)
                tfile.Location = string(dir)

                //TODO_5: As it currently stands the same file can be added to the array more than once 
                //TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
				
				for _, currentFile := range Files {
					if(tfile.Filename == currentFile.Filename && tfile.Location == currentFile.Location){
						//it is a duplicate, flag it
						dontAdd = true
						numOfHits++	//still counts as hit since it is in there
					}
				}
				//only add if not duplicate
				if(!dontAdd){
                	Files = append(Files, tfile)
					numOfHits++
				}
				dontAdd = false //reset 

                if w != nil && len(Files)>0 {

                    //TODO_6: The current key value is the LEN of Files (this terrible); 
                    //TODO_6: Create some variable to track how many files have been added
                    w.Write([]byte(`"`+(strconv.FormatInt(int64(numOfHits), 10))+`":  `))
                    json.NewEncoder(w).Encode(tfile)
                    w.Write([]byte(`,`))

                } 
                
				if(LOG_LEVEL > 1){
                	log.Printf("[+] HIT: %s\n", path)
				}
            }

        }
        return nil
    }

}

//TODO_7: One of the options for the API is a query command
//TODO_7: Create a walkFn2 function based on the walkFn function, 
//TODO_7: Instead of using the regexes array, define a single regex 
//TODO_7: Hint look at the logic in scrape.go to see how to do that; 
//TODO_7: You won't have to itterate through the regexes for loop in this func!

func walkFn2(w http.ResponseWriter, query string) filepath.WalkFunc {
    return func(path string, f os.FileInfo, err error) error {
		w.Header().Set("Content-Type", "application/json")

		//grab r for regex
		r := regexp.MustCompile(`(?i)`+query)
		
		if r.MatchString(path) {
			var tfile FileInfo
			dir, filename := filepath.Split(path)
			tfile.Filename = string(filename)
			tfile.Location = string(dir)

			//TODO_5: As it currently stands the same file can be added to the array more than once 
			//TODO_5: Prevent this from happening by checking if the file AND location already exist as a single record
			for _, currentFile := range Files {
				if(tfile.Filename == currentFile.Filename && tfile.Location == currentFile.Filename){
					//it is a duplicate, flag it
					dontAdd = true
					numOfHits++	//still counts as hit since it is in there
				}
			}
			//only add if not duplicate
			if(!dontAdd){
				Files = append(Files, tfile)
				numOfHits++
			}
			dontAdd = false //reset 

			if w != nil && len(Files)>0 {

				//TODO_6: The current key value is the LEN of Files (this terrible); 
				//TODO_6: Create some variable to track how many files have been added
				w.Write([]byte(`"`+(strconv.FormatInt(int64(numOfHits), 10))+`":  `))
				json.NewEncoder(w).Encode(tfile)
				w.Write([]byte(`,`))

			} 
			
			if(LOG_LEVEL > 1){
				log.Printf("[+] HIT: %s\n", path)
			}
		}

        return nil

    }
}

//==========================================================================\\

func APISTATUS(w http.ResponseWriter, r *http.Request) {

	if(LOG_LEVEL > 0){
		log.Printf("Entering %s end point", r.URL.Path)
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{ "status" : "API is up and running ",`))
    var regexstrings []string
    
    for _, regex := range regexes{
        regexstrings = append(regexstrings, regex.String())
    }

    w.Write([]byte(` "regexs" :`))
    json.NewEncoder(w).Encode(regexstrings)
    w.Write([]byte(`}`))
	log.Println(regexes)

}


func MainPage(w http.ResponseWriter, r *http.Request) {
	if(LOG_LEVEL > 0 ){
		log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "text/html")

	w.WriteHeader(http.StatusOK)
    //TODO_8 - Write out something better than this that describes what this api does

	fmt.Fprintf(w, "<html><body><H1>Welcome to my awesome File page</H1><p>/indexer will add hits to Files</p><p>/search will pull specific regular expressions</p><p>/api-status will show regexes and api status</p><p>/clear will empty preset regexes</p><p>/reset will add presets back</p><p>/addsearch/{REGEX} will add a new regex to the searches</p></body>")
}


func FindFile(w http.ResponseWriter, r *http.Request) {
	if(LOG_LEVEL > 0){
		log.Printf("Entering %s end point", r.URL.Path)
	}
    q, ok := r.URL.Query()["q"]

    w.WriteHeader(http.StatusOK)
    if ok && len(q[0]) > 0 {

		if(LOG_LEVEL > 0){
        	log.Printf("Entering search with query=%s",q[0])
		}

        // ADVANCED: Create a function in scrape.go that returns a list of file locations; call and use the result here
        // e.g., func finder(query string) []string { ... }

		var found bool = false 	
        for _, File := range Files {
		    if File.Filename == q[0] {
                json.NewEncoder(w).Encode(File.Location)
                found = true
		    }
        }
        //TODO_9: Handle when no matches exist; print a useful json response to the user; hint you might need a "FOUND variable" to check here ...
		if(!found){
			fmt.Fprintf(w, "<html><body><H1>No matches exist</H1></body>")
		}

    } else {
        // didn't pass in a search term, show all that you've found
        w.Write([]byte(`"files":`))    
        json.NewEncoder(w).Encode(Files)
    }
}

func IndexFiles(w http.ResponseWriter, r *http.Request) {
	if(LOG_LEVEL > 0){
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")

    location, locOK := r.URL.Query()["location"]
	regex, regexOK := r.URL.Query()["regex"]

    
    //TODO_10: Currently there is a huge risk with this code ... namely, we can search from the root /
    //TODO_10: Assume the location passed starts at /home/ (or in Windows pick some "safe?" location)
    //TODO_10: something like ...  rootDir string := "???"
    //TODO_10: create another variable and append location[0] to rootDir (where appropriate) to patch this hole
	newLocation := rootDir + location[0]

    if locOK && len(location[0]) > 0 {
        w.WriteHeader(http.StatusOK)

    } else {
        w.WriteHeader(http.StatusFailedDependency)
        w.Write([]byte(`{ "parameters" : {"required": "location",`))    
        w.Write([]byte(`"optional": "regex"},`))    
        w.Write([]byte(`"examples" : { "required": "/indexer?location=/xyz",`))
        w.Write([]byte(`"optional": "/indexer?location=/xyz&regex=(i?).md"}}`))
        return 
    }

    //wrapper to make "nice json"
    w.Write([]byte(`{ `))
    
    // TODO_11: Currently the code DOES NOT do anything with an optionally passed regex parameter
    // Define the logic required here to call the new function walkFn2(w,regex[0])
    // Hint, you need to grab the regex parameter (see how it's done for location above...) 
    
    // if regexOK
    //   call filepath.Walk(location[0], walkFn2(w, `(i?)`+regex[0]))
    // else run code to locate files matching stored regular expression
	numOfHits = 0 	//reset this here, to restart for each indexing

	//go into walkFn2 if necessary
	if regexOK {
		filepath.Walk(newLocation, walkFn2(w, regex[0]))
	}else{
		filepath.Walk(newLocation, walkFn(w))
	}

    //wrapper to make "nice json"
    w.Write([]byte(` "status": "completed"} `))

}


//TODO_12 create endpoint that calls resetRegEx AND *** clears the current Files found; ***
//TODO_12 Make sure to connect the name of your function back to the reset endpoint main.go!
func ResetArray(w http.ResponseWriter, r *http.Request){
	if(LOG_LEVEL > 0){
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")

	resetRegEx()
	Files = nil
}

//TODO_13 create endpoint that calls clearRegEx ; 
//TODO_13 Make sure to connect the name of your function back to the clear endpoint main.go!
func ClearRegEx(w http.ResponseWriter, r *http.Request){
	if(LOG_LEVEL > 0){
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")

	clearRegEx()
}

//TODO_14 create endpoint that calls addRegEx ; 
//TODO_14 Make sure to connect the name of your function back to the addsearch endpoint in main.go!
// consider using the mux feature
// params := mux.Vars(r)
// params["regex"] should contain your string that you pass to addRegEx
// If you try to pass in (?i) on the command line you'll likely encounter issues
// Suggestion : prepend (?i) to the search query in this endpoint
func AddRegEx(w http.ResponseWriter, r *http.Request){
	if(LOG_LEVEL > 0){
    	log.Printf("Entering %s end point", r.URL.Path)
	}
    w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	addRegEx(`(?i)`+params["regex"])


}
