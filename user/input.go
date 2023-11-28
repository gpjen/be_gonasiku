package user

type UserRegister struct {
	Name     string `json:"name" `
	Email    string `json:"email" `
	Password string `json:"password" `
	Phone    string `json:"phone" `
	Addres   string `json:"addres" `
	Gender   string `json:"gender" `
}

type UserUpdate struct {
	Name   string `json:"name" `
	Email  string `json:"email" `
	Phone  string `json:"phone" `
	Addres string `json:"addres" `
	Gender string `json:"gender" `
}
