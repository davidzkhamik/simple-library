package main

import "fmt"

func main(){ 
	user1 := Reader{
		ID: 1,
		FirstName: "artema",
		LastName: "nigerksenov",
		IsActive: true,
	}

	fmt.Println(user1)
	user1.DisplayReader()
	user1.Dearctivate

} 