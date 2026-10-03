package main

import "fmt"

// Nested struct
type Address struct {
	City string
	Zip  int
}

// Main struct
type Server struct {
	Name    string
	IP      string
	Port    int
	Active  bool
	Address Address
	Ports   []int
	Labels  map[string]string
}

func main() {

	// 1. Create struct
	server1 := Server{
		Name:   "web",
		IP:     "10.0.0.1",
		Port:   8080,
		Active: true,

		// Nested struct
		Address: Address{
			City: "Bengaluru",
			Zip:  560001,
		},

		// Slice inside struct
		Ports: []int{8080, 8443},

		// Map inside struct
		Labels: map[string]string{
			"env":  "prod",
			"team": "platform",
		},
	}

	// 2. Access fields
	fmt.Println("Name:", server1.Name)
	fmt.Println("IP:", server1.IP)
	fmt.Println("Port:", server1.Port)
	fmt.Println("Active:", server1.Active)

	// 3. Nested struct access
	fmt.Println("City:", server1.Address.City)

	// 4. Update field
	server1.Port = 9090
	fmt.Println("Updated Port:", server1.Port)

	// 5. Struct assignment / copy
	server2 := server1

	fmt.Println("\nServer2:")
	fmt.Println(server2.Name)
	fmt.Println(server2.Port)

	// 6. Changing normal field in server2
	server2.Port = 7070

	fmt.Println("\nAfter changing server2 Port:")
	fmt.Println("Server1 Port:", server1.Port)
	fmt.Println("Server2 Port:", server2.Port)

	// 7. Slice inside struct
	server1.Ports = append(server1.Ports, 9000)

	fmt.Println("\nPorts:", server1.Ports)

	// 8. Map inside struct
	server1.Labels["region"] = "india"

	fmt.Println("Labels:", server1.Labels)

	// 9. Struct with zero values
	var server3 Server

	fmt.Println("\nServer3:")
	fmt.Println("Name:", server3.Name)
	fmt.Println("Port:", server3.Port)
	fmt.Println("Active:", server3.Active)
}
