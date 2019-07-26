package main

import "github.com/goinaction/code/http-server"

func main() {

	// you cannot have both servers running at the same time. to test comment one of these two.
	http_server.GoHttpServer()
	http_server.GorillaServer()
}
