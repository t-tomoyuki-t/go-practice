package entity

type User struct {
	Id    int    `json:"id" gorm:"primary_key" uri:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	password string
}

func (user *User) GetPassword() string {
	return user.password
}
