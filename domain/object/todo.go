package object

type Todo struct {
	Id    int    `json:"id" gorm:"primary_key" uri:"id"`
	Title string `json:"title"`
}
