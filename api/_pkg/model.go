package pkg

type  User struct{
	Username  string `json:"username"`
}

func NewModel() User {
	model := User{
		Username: "coba",
	}
	return model
}