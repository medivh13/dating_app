package errors

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : github.com/medivh13/dating_app
 */

import (
	"net/http"
)

var httpCode = map[ErrorCode]int{
	UNKNOWN_ERROR:            http.StatusInternalServerError,
	DATA_INVALID:             http.StatusBadRequest,
	STATUS_PAGE_NOT_FOUND:    http.StatusNotFound,
	INVALID_HEADER_X_API_KEY: http.StatusBadRequest,
	UNAUTHORIZED:             http.StatusUnauthorized,
	FAILED_RETRIEVE_DATA:     http.StatusInternalServerError,
}
