package user

type CreateDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
