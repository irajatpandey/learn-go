package main

import "fmt"

type Server struct {
	Name string
	Port int
}

func (s *Server) changePort() {
	s.Port = 9090
}

func main() {
	x := 100

	var ptr *int
	ptr = &x
	fmt.Println(x)
	fmt.Println(ptr)
	fmt.Println(*ptr)

	x = 200
	fmt.Println(x)

	server := Server{
		Name: "web",
		Port: 8080,
	}

	fmt.Println("Port before method call ", server.Port)
	server.changePort()
	fmt.Println("Port after method call ", server.Port)
}
