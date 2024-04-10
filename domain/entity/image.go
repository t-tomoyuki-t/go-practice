package entity

type Image struct {
	Id int `json:"id" gorm:"primary_key" uri:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}
