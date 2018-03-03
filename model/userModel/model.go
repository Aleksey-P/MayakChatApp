package userModel

type User struct {
	ID int `json:id`
	Username string `json:username`
	Password int
}

func (u *User) Update(fieldsToUpdate interface{}) {
}

