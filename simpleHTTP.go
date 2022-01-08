package main

import (
	"fmt"
	"log"
	"net"
	"os"
 	"net/http"
)

func main() {
	log.Print("simplehttp: Enter main()")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// printing request headers/params
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Message demo class = %s\n", os.Getenv("message"))
}

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}
