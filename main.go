package main

var call *request

func main() {
	call = &request{}

	//Flags()

	SendRequest()

}

var hdr map[string][]string
