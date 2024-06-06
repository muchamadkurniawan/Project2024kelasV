package main

import "fmt"

type Member struct {
	Id     int
	Nama   string
	Alamat string
	Point  float32
}

type TableMember struct {
	Data Member
	Next *TableMember
}

var DBmember TableMember

func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := TableMember{
		Data: Member{id, nama, alamat, point},
		Next: nil,
	}
	var tempLL *TableMember
	tempLL = &DBmember
	if DBmember.Next == nil {
		DBmember.Next = &newDataMember
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataMember
	}
}

func ReadAllMember() {
	var tempLL *TableMember
	tempLL = DBmember.Next
	if DBmember.Next == nil {
		fmt.Println("tidak ada data")
	} else {
		for tempLL != nil {
			fmt.Println("------------------------------------")
			fmt.Println("nomer member :", tempLL.Data.Id)
			fmt.Println("nama member :", tempLL.Data.Nama)
			fmt.Println("alamat member :", tempLL.Data.Alamat)
			fmt.Println("point member :", tempLL.Data.Point)
			tempLL = tempLL.Next
			fmt.Println("------------------------------------")
		}
	}
}

func SearchMember(id int) {
	var tempLL *TableMember
	tempLL = DBmember.Next
	cek := false
	if DBmember.Next == nil {
		fmt.Println("DATA MEMBER KOSONG")
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				fmt.Println("------------------------------------")
				fmt.Println("nomer member :", tempLL.Data.Id)
				fmt.Println("nama member :", tempLL.Data.Nama)
				fmt.Println("alamat member :", tempLL.Data.Alamat)
				fmt.Println("point member :", tempLL.Data.Point)
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			fmt.Println("ID tidak ditemukan")
		}
	}
}

func cariMember(id int) *TableMember {
	var tempLL *TableMember
	tempLL = DBmember.Next
	cek := false
	if DBmember.Next == nil {
		return nil
	} else {
		for tempLL != nil {
			if id == tempLL.Data.Id {
				cek = true
				return tempLL
			}
			tempLL = tempLL.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

func UpdateMember(id int, nama string, alamat string) {
	member := cariMember(id)
	if member != nil {
		member.Data.Nama = nama
		member.Data.Alamat = alamat
		fmt.Println("update berhasil")
	} else {
		fmt.Println("tidak ada data yang dicari")
	}
}

func DeleteMember(id int) {
	var tempLL *TableMember
	tempLL = &DBmember
	if tempLL.Next == nil {
		//cek data kosong
		fmt.Println("data kosong")
	} else {
		for tempLL.Next != nil {
			//fmt.Println(tempLL.Next.Data)
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return
			}
			tempLL = tempLL.Next
		}
	}
}

func main() {
	InsertMember(1, "kurniawan", "surabaya", 100)
	InsertMember(2, "aan", "malang", 10)
	InsertMember(3, "nino", "jakarta", 50)
	InsertMember(4, "adel", "jakarta", 50)
	//
	//ReadAllMember()
	//SearchMember(8)
	//testt update
	//UpdateMember(1, "siti", "malang")
	DeleteMember(3)
	ReadAllMember()
}
