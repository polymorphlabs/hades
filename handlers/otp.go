package handler

import (
	"encoding/json"
	"hades/repository"
	"net/http"
)

func CreateNewOTP(writer http.ResponseWriter, request *http.Request) {
	var otp GenerateOTPRequest

	decodeError := json.NewDecoder(request.Body).Decode(&otp)

	if decodeError != nil {
		http.Error(writer, decodeError.Error(), 400)
		return
	}

	err := repository.SetOTP(otp.Phone, "12345")

	if err != nil {
		failureResponse := FailureResponse{
			Success:false,
			Error: err.Error(),
		}
		if err := json.NewEncoder(writer).Encode(failureResponse); err != nil {
			panic(err)
		}
	}


	successResponse := SuccessResponse{
		Success:true,
		Message:"OTP generated successfully",
	}

	if err := json.NewEncoder(writer).Encode(successResponse); err != nil {
		panic(err)
	}
	return
}


func ValidateOTP(writer http.ResponseWriter, request *http.Request) {
	var otp ValidateOTPRequest

	decodeError := json.NewDecoder(request.Body).Decode(&otp)

	if decodeError != nil {
		http.Error(writer, decodeError.Error(), 400)
		return
	}

	valid, validationError := repository.ValidateOTP(otp.Phone, otp.Otp)

	if validationError != nil {
		failureResponse := FailureResponse{
			Success:false,
			Error: validationError.Error(),
		}
		if err := json.NewEncoder(writer).Encode(failureResponse); err != nil {
			panic(err)
		}
	}

	if !valid {
		failureResponse := FailureResponse{
			Success:false,
			Error: "Invalid OTP",
		}
		if err := json.NewEncoder(writer).Encode(failureResponse); err != nil {
			panic(err)
		}

		return
	}

	successResponse := SuccessResponse{
		Success:true,
		Message:"OTP validated successfully",
	}

	if err := json.NewEncoder(writer).Encode(successResponse); err != nil {
		panic(err)
	}
	return
}