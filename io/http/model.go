package http

type SignupRequest struct {
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Email string `json:"email"`
}

type SignupResponse struct {
	Name  string `json:"name"`
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
