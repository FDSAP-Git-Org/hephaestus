package utils

type (
	EPResponse struct {
		ProcessTime string      `json:"processTime"`
		Request     interface{} `json:"request"`
	}
)
