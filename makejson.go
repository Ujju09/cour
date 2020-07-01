package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var userName string
	var userAddress string
	// fmt.Println("Enter first name:")     /*The following is applicable for one text inputs only*/
	// fmt.Scanln(&userName)
	// fmt.Println("Enter address:")
	// fmt.Scanln(&userAddress)

	fmt.Println("Enter name:") //This is for the name part
	scannedName := bufio.NewScanner(os.Stdin)
	scannedName.Scan()
	userName = scannedName.Text()

	fmt.Println("Enter address:") //This is for the address part
	scannedAddress := bufio.NewScanner(os.Stdin)
	scannedAddress.Scan()
	userAddress = scannedAddress.Text()

	userMap := map[string]string{
		"Address": userAddress,
		"Name":    userName,
	}
	// fmt.Println(userMap)
	userMapToJSON, err := json.Marshal(userMap)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(userMapToJSON)

}
