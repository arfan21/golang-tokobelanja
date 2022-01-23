package modeluser

import "time"

type Request struct {
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestTopUp struct {
	ID      uint `json:"id,omitempty"`
	Balance int  `json:"balance"`
}

type Response struct {
	ID        uint       `json:"id"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Fullname  string     `json:"full_name"`
	Email     string     `json:"email"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}
