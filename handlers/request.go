package handler

type GenerateOTPRequest struct {
	Phone string `json:"phone"`
}

type ValidateOTPRequest struct {
	Phone string `json:"phone"`
	Otp string `json:"otp"`
}
