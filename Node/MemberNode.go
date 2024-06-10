package Node

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
