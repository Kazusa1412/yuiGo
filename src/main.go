package main

func main() {
	var server = NewServer("127.0.0.1",6667)
	server.Start()
}
