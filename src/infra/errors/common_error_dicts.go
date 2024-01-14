package errors

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : github.com/medivh13/dating_app
 */

const (
	UNKNOWN_ERROR            ErrorCode = 0
	DATA_INVALID             ErrorCode = 4001
	FAILED_RETRIEVE_DATA     ErrorCode = 4002
	STATUS_PAGE_NOT_FOUND    ErrorCode = 4003
	INVALID_HEADER_X_API_KEY ErrorCode = 4004
	UNAUTHORIZED             ErrorCode = 4005
	FAILED_CREATE_DATA       ErrorCode = 4006
)

var errorCodes = map[ErrorCode]*CommonError{
	UNKNOWN_ERROR: {
		ClientMessage: "Unknown error.",
		SystemMessage: "Unknown error.",
		ErrorCode:     UNKNOWN_ERROR,
	},
	DATA_INVALID: {
		ClientMessage: "Invalid Data Request",
		SystemMessage: "Some of query params has invalid value.",
		ErrorCode:     DATA_INVALID,
	},
	FAILED_RETRIEVE_DATA: {
		ClientMessage: "Failed to retrieve Data.",
		SystemMessage: "Something wrong happened while retrieve Data.",
		ErrorCode:     FAILED_RETRIEVE_DATA,
	},
	STATUS_PAGE_NOT_FOUND: {
		ClientMessage: "Invalid Status Page.",
		SystemMessage: "Status Page Email Address not found.",
		ErrorCode:     STATUS_PAGE_NOT_FOUND,
	},
	INVALID_HEADER_X_API_KEY: {
		ClientMessage: "Invalid Api Key.",
		SystemMessage: "Invalid value of header X-Api-Key.",
		ErrorCode:     INVALID_HEADER_X_API_KEY,
	},
	UNAUTHORIZED: {
		ClientMessage: "Unauthorized",
		SystemMessage: "Unauthorized",
		ErrorCode:     UNAUTHORIZED,
	},
	FAILED_CREATE_DATA: {
		ClientMessage: "Failed to create data.",
		SystemMessage: "Something wrong happened while create data.",
		ErrorCode:     FAILED_CREATE_DATA,
	},
}
