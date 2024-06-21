package View

import (
	"Project2024kelasV/Controller"
	"fmt"
)

func InsertMember() {
	var id int
	var nama, alamat string
	var point float32
	fmt.Print("masukkan id member : ")
	fmt.Scan(&id)
	fmt.Print("masukkan nama member : ")
	fmt.Scan(&nama)
	fmt.Print("masukkan alamat member : ")
	fmt.Scan(&alamat)
	fmt.Print("masukkan point member : ")
	fmt.Scan(&point)
	cek := Controller.InsertMember(id, nama, alamat, point)
	if cek == true {
		fmt.Println("data berhasil di input")
	} else {
		fmt.Println("data gagal diinput")
	}
}

func ReadAllMember() {
	members := Controller.ReadAllMember()
	if members != nil {
		for members.Next != nil {
			fmt.Println("Nama Member : ", members.Next.Data.Nama)
			fmt.Println("Id Member : ", members.Next.Data.Id)
			fmt.Println("Alamat Member : ", members.Next.Data.Alamat)
			members = members.Next
		}
	}
}
