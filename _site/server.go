/*
Start local server for development: dev_appserver.py ./
*/

package website

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

var contents string

func init() {
    http.HandleFunc("/", root)

    if b, err := ioutil.ReadFile("index.html"); err != nil {
    	log.Fatalf("Error in loading index.html: %v", err)
    } else {
    	contents = string(b)
    }
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, contents)
}