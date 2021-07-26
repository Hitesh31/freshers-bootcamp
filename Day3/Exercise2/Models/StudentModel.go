package Models
type Student struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	lastName   string `json:"lastName"`
	dateOfBirth   string `json:"date_of_birth"`
	Address string `json:"address"`
	Subject string `json:"Subject"`
	Marks uint `json:"marks"`
}
func (b *Student) TableName() string {
	return "Student"
}