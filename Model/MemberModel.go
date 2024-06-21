package Model

import (
	"Project2024kelasV/Database"
	"Project2024kelasV/Node"
)

func InsertMember(id int, nama string, alamat string, point float32) {
	newDataMember := Node.TableMember{
		Data: Node.Member{id, nama, alamat, point},
		Next: nil,
	}
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		Database.DBmember.Next = &newDataMember
	} else {
		for tempLL.Next != nil {
			tempLL = tempLL.Next
		}
		tempLL.Next = &newDataMember
	}
}

func ReadAllMember() *Node.TableMember {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if Database.DBmember.Next == nil {
		return nil
	} else {
		return tempLL
	}
}

func DeleteMember(id int) bool {
	var tempLL *Node.TableMember
	tempLL = &Database.DBmember
	if tempLL.Next == nil {
		return false
	} else {
		for tempLL.Next != nil {
			if tempLL.Next.Data.Id == id {
				tempLL.Next = tempLL.Next.Next
				return true
			}
			tempLL = tempLL.Next
		}
		return false
	}
}

func SearchMember(id int) *Node.TableMember {
	var tempLL *Node.TableMember
	tempLL = Database.DBmember.Next
	cek := false
	if Database.DBmember.Next == nil {
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

func UpdateMember(member Node.Member) bool {
	addr := SearchMember(member.Id)
	if addr == nil {
		return false
	} else {
		addr.Data.Nama = member.Nama
		addr.Data.Alamat = member.Alamat
		return true
	}
}
