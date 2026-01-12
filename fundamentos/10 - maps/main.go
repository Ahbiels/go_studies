package main

import (
	"fmt"
)

var name string = "Angelo"

func main(){
	user := map[string]string{
		"name": name,
		"surname": "Gabriel",
	}

	fmt.Println(user["name"])
	fmt.Println(user)

	user2 := map[string]map[string]string{
		"full_name": {
			"first_name": name,
			"last_name": "Gabriel",
		},
		"class": {
			"name": "Engenharia da computação",
			"Period": "4th semester",
		},
	}

	fmt.Printf("\n\n%v",user2)
	fmt.Printf("\n%v",user2["full_name"]["first_name"])
	
	delete(user2, "class")
	
	fmt.Printf("\n%v",user2)
	user2["work"] = map[string]string{
		"Company": "CompanyXYZ",
		"area": "Data",
	}
	fmt.Printf("\n%v\n",user2)

	

}