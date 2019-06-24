package models

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Pass  string `json:"pass" validate:"required"`
	//how much the user will pay
	Bill float64 `json:"bill"`
}

type UpdatableUser struct {
	Filter SecureUser `json: "filter"`
	//how much the user will pay
	Update User `json: "update"`
}

type SecureUser struct {
	Email string `json:"email" validate:"required,email"`
	Pass  string `json:"pass" validate:"required"`
}
