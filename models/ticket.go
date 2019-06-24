package models

type Ticket struct {
	Model

	OrderNo     string `json:"order_no"`
	OrgId       uint   `json:"org_id"`
	orgType     uint8  `json:"org_type"`
	SourceType  uint8  `json:"source_type"`
	KsuId       uint8  `json:"ksu_id"`
	KsuName     string `json:"ksu_name"`
	TeacherId   int    `json:"teacher_id"`
	TeacherName string `json:"teacher_name"`
	StudentId   int    `json:"student_id"`
	OkayId      int    `json:"okay_id"`
	StudentName string `json:"student_name"`
	Status      uint8  `json:"status"`
	IsTasked    int8   `json:"is_tasked"`
}

func GetTickets(pageNum int, pageSize int, maps interface{}) (tickets []Ticket) {
	//db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tickets)
	return
}

func GetTotalTickets(maps interface{}) (count int) {
	db.Model(&Ticket{}).Where(maps).Count(&count)
	return
}
