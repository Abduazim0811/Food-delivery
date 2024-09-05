package user

type User struct {
	ID        int `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegisterReq struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserRes struct {
	ID int `json:"id"`
}

type UserReq struct {
	Email string `json:"email"`
	Code  int32  `json:"code"`
}

type RegisterRes struct {
	Message string `json:"message"`
}

type LoginReq struct {
	Email string `json:"email"`
	Password  string `json:"password"`
}

type LoginRes struct {
	Token string `json:"token"`
}
