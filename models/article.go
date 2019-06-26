package models

//type Article struct {
//	Model
//
//	TagID int `json:"tag_id" gorm:"index"`
//	Tag   Tag `json:"tag"`
//
//	Title string `json:"title"`
//	Desc string `json:"desc"`
//	Content string `json:"content"`
//	CreatedBy string `json:"created_by"`
//	ModifiedBy string `json:"modified_by"`
//	State int `json:"state"`
//}

type Article struct {
	Model

	TicketId int `json:"ticket_id" gorm:"index"`
	Ticket Ticket `json:"ticket"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}


//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreatedOn", time.Now().Unix())
//
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//
//	return nil
//}