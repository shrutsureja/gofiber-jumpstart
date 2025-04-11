package custom_errors

var CommonErrors = struct {
	BodyParsingFailed     string
	BodyValidationFailed  string
	QueryParsingFailed    string
	QueryValidationFailed string
	UserUnauthorized      string
	UserForbiddenAccess   string
	DatabaseError         string
	BodySizeExceeded      string
	InternalServerError   string
}{
	BodyParsingFailed:     "Failed to parse the request body. Please check the data format and types.",
	BodyValidationFailed:  "Request body validation failed. Ensure all required fields are correctly filled.",
	QueryParsingFailed:    "Failed to parse the query parameters. Please check the data format.",
	QueryValidationFailed: "Query parameter validation failed. Ensure all required parameters are correctly provided.",
	UserUnauthorized:      "User is not authorized. Please log in to continue.",
	UserForbiddenAccess:   "User does not have permission to access this resource.",
	DatabaseError:         "An error occurred while accessing the database. Please try again later.",
	BodySizeExceeded:      "Request body size exceeded the limit.",
	InternalServerError:   "An unexpected error occurred. Please try again later.",
}

var CommonErrorsCodes = struct {
	BodyParsingFailed     string
	BodyValidationFailed  string
	QueryParsingFailed    string
	QueryValidationFailed string
	UserUnauthorized      string
	UserForbiddenAccess   string
	DatabaseError         string
	BodySizeExceeded      string
	InternalServerError   string
}{
	BodyParsingFailed:     "body_parsing_failed",
	BodyValidationFailed:  "body_validation_failed",
	QueryParsingFailed:    "query_parsing_failed",
	QueryValidationFailed: "query_validation_failed",
	UserUnauthorized:      "user_unauthorized",
	UserForbiddenAccess:   "user_forbidden_access",
	DatabaseError:         "database_error",
	BodySizeExceeded:      "body_size_exceeded",
	InternalServerError:   "internal_server_error",
}
