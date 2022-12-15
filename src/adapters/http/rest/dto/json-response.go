package dto

type Response struct {
	Status  bool `json:"status"`
	Message string `json:"mensaje"`
}

type ResponseWithData struct {
	Status  bool      `json:"status"`
	Message string      `json:"mensaje"`
	Data    interface{} `json:"data"`
}
