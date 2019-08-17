package handler

type SuccessResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}

type FailureResponse struct {
	Success bool `json:"success"`
	Error string `json:"error"`
}

type RootResponse struct {
	Message string `json:"message"`
	Status bool `json:"status"`
}
