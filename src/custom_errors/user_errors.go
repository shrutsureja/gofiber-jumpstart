package custom_errors

var UserErrors = struct {
	ErrorInGeneratingToken string
	UserWithEmailExists    string
}{
	ErrorInGeneratingToken: "An error occurred while generating the token. Please try again later.",
	UserWithEmailExists:    "A user with the provided email already exists. Please try again with a different email.",
}

var UserErrorsCodes = struct {
	ErrorInGeneratingToken string
	UserWithEmailExists    string
}{
	ErrorInGeneratingToken: "error_in_generating_token",
	UserWithEmailExists:    "user_with_email_exists",
}
