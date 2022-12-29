package main

import "fmt"

type Server struct {
	Name     string
	Protokol string
}

func (s Server) setServerName(name string) {
	s.Name = name
}

func NewProtokol(s Server, protokol string) {
	//sta je zapravo Server ovdje

	s.Protokol = protokol
	fmt.Println(s)
}

func newTask2(ime string, nekaFunkcija func(Server)) {
	// sta je sad mogu sa funkciom nekaFunkcija
}

func main() {
	newServer := Server{}
	fmt.Println("first request ", newServer)

	newServer.Name = "ad"

	NewProtokol(newServer, "novi protokol") //sto ovo ne radi

	fmt.Println(newServer) // ne prikazuje mi novi protokol

	//Sto ce mi zapravo funkcija ko argument
	newTask2("Janko", func(s Server) {
		s.setServerName("Janko") //kako mogu dodati ovo na nepostojeci "objekat"
	})
}
