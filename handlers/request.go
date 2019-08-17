package handler

type GenerateOTPRequest struct {
	Phone  string `json:"phone"`
	Entity string `json:"entity"`
}

type ValidateOTPRequest struct {
	Phone string `json:"phone"`
	Otp string `json:"otp"`
}
