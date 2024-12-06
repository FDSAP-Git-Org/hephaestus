package utils

type (
	EPResponse struct {
		ProcessTime string      `json:"processTime"`
		Request     interface{} `json:"request"`
	}

	UserDetails struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
