package views

type Status struct {
	Code    int
	Message string
}

type DefaultResponse struct {
	Status Status
	Data   any
}
