package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/samitpal/run-ssh/api"
	"github.com/samitpal/run-ssh/plugins/ssh"
)

var (
	addr    string
	apiAddr string
)

func init() {
	flag.StringVar(&ssh.PrivateKey, "i", "", "Path to private key to use")
	flag.StringVar(&ssh.User, "l", os.Getenv("LOGNAME"), "Optional login name")
	flag.StringVar(&addr, "listen-addr", "0.0.0.0:8081", "Http UI address")
	flag.StringVar(&apiAddr, "api-addr", "localhost:8080", "API listen address")
	flag.Parse()

}
func main() {

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(apiAddr, router))
}
