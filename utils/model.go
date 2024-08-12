package utils

type (
	EPResponse struct {
		ProcessTime string      `json:"processTime"`
		Response    interface{} `json:"response"`
	}
)