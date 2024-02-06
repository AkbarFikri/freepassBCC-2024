package schemas

type UserResponse struct {
	ID    string
	Email string
	Name  string
}

type UserLoginRequest struct {
	Email    string
	Password string
}
