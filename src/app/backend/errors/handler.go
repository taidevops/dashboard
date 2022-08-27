package errors

import "net/http"

func HandleHTTPError(err error) int {
	if err == nil {
		return http.StatusInternalServerError
	}
	if err.Error() == MsgTokenExpiredError || err.Error() == MsgLoginUnauthorizedError || err.Error() == MsgEncryptionKeyChanged {
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}
