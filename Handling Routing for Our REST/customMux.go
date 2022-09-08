package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// // CustomServeMux is a struct which can be a multiplexer
// type CustomServeMux struct {
// }

// // This is the function handler to be overridden
// func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path == "/" {
// 		giveRandom(w, r)
// 		return
// 	}
// 	if r.URL.Path == "/hello" {
// 		hello(w, r)
// 		return
// 	}
// 	if r.URL.Path == "/post" {
// 		post(w, r)
// 		return
// 	}
// 	http.NotFound(w, r)
// 	return
// }
// func giveRandom(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
// }
// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "hello: %f", rand.Float64())
// }
// func post(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintln( w,"post naumber :",rand.Uint32())
// }


// func main() {
// 	// Any struct that has serveHTTP function can be a multiplexer
// 	mux := &CustomServeMux{}
// 	http.ListenAndServe(":8000", mux)
// }


func main() {
newMux := http.NewServeMux()
newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, rand.Float64())
})
newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, rand.Intn(100))
})
http.ListenAndServe(":8000", newMux)
}
