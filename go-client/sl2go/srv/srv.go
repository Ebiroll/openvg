package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func jsondata(w http.ResponseWriter, r *http.Request) {
	//f, err := os.Open("testdata.json")
	//if err != nil {
    //    panic(err)
    //}
	contents,_ := ioutil.ReadFile("testdata.json")
	io.WriteString(w, string(contents))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", jsondata)
	http.ListenAndServe(":8000", mux)
}
