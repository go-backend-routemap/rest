package main
import (
"io"
"net/http"
"log"
)
// hello world, the web server
func MyServer(w http.ResponseWriter, req *http.Request) {
io.WriteString(w, "hello, world!\n")
}
func basicHandler() {
http.HandleFunc("/h", MyServer)
log.Fatal(http.ListenAndServe(":8000", nil))
}