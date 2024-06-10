package main

import (
	"Project2024kelasV/Model"
	"fmt"
)

//}
//
//
//

//

//
//func SearchMember(id int) {
//	var tempLL *TableMember
//	tempLL = DBmember.Next
//	cek := false
//	if DBmember.Next == nil {
//		fmt.Println("DATA MEMBER KOSONG")
//	} else {
//		for tempLL != nil {
//			if id == tempLL.Data.Id {
//				cek = true
//				fmt.Println("------------------------------------")
//				fmt.Println("nomer member :", tempLL.Data.Id)
//				fmt.Println("nama member :", tempLL.Data.Nama)
//				fmt.Println("alamat member :", tempLL.Data.Alamat)
//				fmt.Println("point member :", tempLL.Data.Point)
//			}
//			tempLL = tempLL.Next
//		}
//		if cek != true {
//			fmt.Println("ID tidak ditemukan")
//		}
//	}
//}
//
//func cariMember(id int) *TableMember {
//	var tempLL *TableMember
//	tempLL = DBmember.Next
//	cek := false
//	if DBmember.Next == nil {
//		return nil
//	} else {
//		for tempLL != nil {
//			if id == tempLL.Data.Id {
//				cek = true
//				return tempLL
//			}
//			tempLL = tempLL.Next
//		}
//		if cek != true {
//			return nil
//		}
//	}
//	return nil
//}
//
//func UpdateMember(id int, nama string, alamat string) {
//	member := cariMember(id)
//	if member != nil {
//		member.Data.Nama = nama
//		member.Data.Alamat = alamat
//		fmt.Println("update berhasil")
//	} else {
//		fmt.Println("tidak ada data yang dicari")
//	}
//}
//
//func DeleteMember(id int) {
//	var tempLL *TableMember
//	tempLL = &DBmember
//	if tempLL.Next == nil {
//		//cek data kosong
//		fmt.Println("data kosong")
//	} else {
//		for tempLL.Next != nil {
//			//fmt.Println(tempLL.Next.Data)
//			if tempLL.Next.Data.Id == id {
//				tempLL.Next = tempLL.Next.Next
//				return
//			}
//			tempLL = tempLL.Next
//		}
//	}
//}

func main() {
	Model.InsertMember(1, "Kurniawan", "surabaya", 100)
	Model.InsertMember(2, "Aan", "surabaya", 50)
	members := Model.ReadAllMember()
	//fmt.Println(members)
	//fmt.Println(members.Next)
	if members != nil {
		for members.Next != nil {
			fmt.Println(members.Next.Data)
			members = members.Next
		}
	}
}
