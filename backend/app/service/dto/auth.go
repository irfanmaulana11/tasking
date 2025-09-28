package dto

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	DisplayName string  `json:"display_name"`
	Username    string  `json:"user_name"`
	Role        string  `json:"role"`
	Expired     int64   `json:"expired"`
	Token       *string `json:"token"`
}

type Register struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
