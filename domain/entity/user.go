package entity

type User struct {
	Id       int    `json:"id" gorm:"primary_key" uri:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string
}

// TODO: Passwordをprivateにしたい
// Gormはpublicでないといけない 

func NewUser(id int, name string, email string, password string) *User {
	return &User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}
}

func (user *User) GetPassword() string {
	return user.Password
}
