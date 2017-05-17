package main

type Message struct {
	ID string
	Code int
	Subcode int
	Lat float64
	Long float64
}

func main() {
	recRawMsg := []byte(`{"name":"channel add",` + `"data":{"name":"Hardware Support"}}`)
var recMessage 
}
