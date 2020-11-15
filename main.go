package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lulumel0n/m3u8-relay/server/model"
	"github.com/lulumel0n/m3u8-relay/server/router"
)

func main() {
	fmt.Println("v1")
	args := os.Args[1:]
	r := router.Router()

	if len(args) < 2 {
		panic("Usage: <port> <endpoint> [<cert> <key>]")
	}

	model.ENDPOINT = args[1]

	if len(args) < 4 {
		fmt.Println("Usage: <endpoint> [<cert> <key>] to start https, now starting without tls instead")
		fmt.Printf("Starting server on the port %s \n", args[0])
		log.Fatal(http.ListenAndServe(":"+args[0], r))
	} else {
		certFile := args[2]
		privkeyFile := args[3]
		fmt.Printf("Starting server on the port %s \n", args[0])
		log.Fatal(http.ListenAndServeTLS(":"+args[0], certFile, privkeyFile, r))
	}
}
