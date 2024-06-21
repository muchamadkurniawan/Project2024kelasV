package main

import (
	"Project2024kelasVnew/Model"
	"Project2024kelasVnew/Node"
	"fmt"
)

func main() {
	Model.InsertMember(1, "Kurniawan", "surabaya", 100)
	Model.InsertMember(2, "Aan", "surabaya", 50)
	Model.InsertMember(3, "Nino", "Malang", 150)

	//tes Model update
	member1 := Node.Member{
		Id:     2,
		Nama:   "Adel",
		Alamat: "Malang",
		Point:  0,
	}
	tes := Model.UpdateMember(member1)
	fmt.Println("return update member :", tes)
	// model ReadAll Testing
	members := Model.ReadAllMember()
	if members != nil {
		for members.Next != nil {
			fmt.Println(members.Next.Data)
			members = members.Next
		}
	}
}
