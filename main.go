package main

import json_decoder "github.com/goinaction/code/json-decoder"

func main() {

	// you cannot have both servers running at the same time. to test comment one of these two.
	//http_server.GoHttpServer()
	//http_server.GorillaServer()
	//jsondecoder.DecodeStructure()
	json_decoder.DecodeWStringJsonWithoutStructure()

}
