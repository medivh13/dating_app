package errors

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : github.com/medivh13/dating_app
 */

import (
	http "net/http"

	constants "github.com/medivh13/dating_app/src/infra/constants"
)

type HttpError struct {
	CommonError
	HttpStatusNumber int    `json:"-"`
	HttpStatusName   string `json:"type"`
}

func (err HttpError) Error() string {
	return err.ClientMessage
}

func (err CommonError) GetHttpStatus() int {
	if httpCode[err.ErrorCode] == 0 {
		return http.StatusInternalServerError
	}

	return httpCode[err.ErrorCode]
}

func (err CommonError) ToHttpError() HttpError {
	httpStatusNumber := err.GetHttpStatus()

	return HttpError{
		CommonError:      err,
		HttpStatusNumber: httpStatusNumber,
		HttpStatusName:   constants.GetHttpStatusText(httpStatusNumber),
	}
}
