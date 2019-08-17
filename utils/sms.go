package utils


func SendOTPMessage(phone string, otp string, entity string) error {
	message := "Welcome to " + entity + ". Your verification code is " + otp
	_, sendMessageError := HttpClient.Get(
		"https://apps.mnotify.net/smsapi?key=Fj369KxppKU0QJsngKBGdSNFS&to=" + phone + "&msg=" + message + "&sender_id=Polymorph")

	if sendMessageError != nil {
		return sendMessageError
	}

	return nil
}
