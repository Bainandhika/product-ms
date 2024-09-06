package views

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DefaultResponse struct {
	Status Status `json:"status"`
	Data   any    `json:"data"`
}
