package main

import "fmt"

func main() {

	fmt.Println("Welcome!")
	//declare & assign slice
	ids := []int{23, 22, 30, 66, 45, 66}

	for i, id := range ids {
		fmt.Printf("Index: %d - ID: %d\n", i, id)
	}

	//slice without index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//sum of ids in slices
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum of ids: %d\n", sum)

	//Range with map
	emails := map[string]string{"bob": "bob@gmail.com", "sharon": "sharon@gmail.com"}
	for ky, vl := range emails {
		fmt.Printf("Key: %s - Value: %s\n", ky, vl)
	}

	//get only key
	for ky := range emails {
		fmt.Printf("Value: %s\n", ky)
	}
}
