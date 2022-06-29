package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	oldFile string
	newFile string
	oldData interface{}
	newData interface{}
)

func main() {

	fmt.Println("****************************************************")
	fmt.Println("            Lets Start The Application     ")
	fmt.Println("****************************************************")
	fmt.Print("Enter Old Json Data File Name : ")
	fmt.Scanf("%s\n", &oldFile)
	fmt.Print("Enter New Json Data File Name : ")
	fmt.Scanf("%s\n", &newFile)
	data, err := os.ReadFile(oldFile)
	if err != nil {
		log.Fatalf("failed to read data from file : %q", oldFile)
	}
	if err = json.Unmarshal(data, &oldData); err != nil {
		log.Fatalln("failed to unmarshal data into json")
	}
	data2, err := os.ReadFile(newFile)
	if err != nil {
		log.Fatalf("failed to read data from file : %q", newFile)
	}
	if err = json.Unmarshal(data2, &newData); err != nil {
		log.Fatalln("failed to unmarshal data into json")
	}
	fmt.Println(oldData)
	fmt.Println(newData)
}
