package main

import "fmt"

func main() {

	// 1. Create a map
	stores := map[string]int{
		"Rajat": 10,
		"John":  20,
		"Smith": 30,
	}

	fmt.Println(stores)

	// 2. Add a new key
	stores["Max"] = 100

	// 3. Update existing key
	stores["Rajat"] = 50

	// 4. Delete a key
	delete(stores, "Rajat")

	fmt.Println(stores)

	// 5. Check if key exists
	item, exists := stores["Rajat"]

	if exists {
		fmt.Println("Found:", item)
	} else {
		fmt.Println("Key not found")
	}

	// 6. Create empty map using make
	dict := make(map[string]int)

	dict["Rajat"] = 100
	dict["John"] = 200

	fmt.Println(dict)

	// 7. Iterate: key + value
	servers := map[string]string{
		"web":   "10.0.0.1",
		"db":    "10.0.0.2",
		"cache": "10.0.0.3",
	}

	for key, value := range servers {
		fmt.Println(key, value)
	}

	// 8. Iterate: only keys
	for key := range servers {
		fmt.Println(key)
	}

	// 9. Iterate: only values
	for _, value := range servers {
		fmt.Println(value)
	}
}
