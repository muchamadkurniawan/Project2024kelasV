package Controller

import (
	"Project2024kelasV/Model"
	"Project2024kelasV/Node"
)

func InsertMember(id int, nama string, alamat string, point float32) bool {
	if id > 0 && nama != "" {
		Model.InsertMember(id, nama, alamat, point)
		return true
	} else {
		return false
	}
}

func ReadAllMember() *Node.TableMember {
	return Model.ReadAllMember()
}
