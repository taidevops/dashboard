package errors

import (
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var _ error = &errors.StatusError{}

// NewUnauthoried returns an error indicating the client is not authoried to perform the requested
// action.
func NewUnauthorized(reason string) *errors.StatusError {
	return errors.NewUnauthorized(reason)
}

// NewTokenExpired returns a statusError
// which is an error intended for consumption by a REST API server; it can also be
// reconstructed by clients from a REST response. Public to allow easy type switches
func NewTokenExpired(reason string) *errors.StatusError {
	return &errors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusUnauthorized,
			Reason:  metav1.StatusReasonExpired,
			Message: reason,
		},
	}
}

// NewBadRequest creates an error that indicates that the request is invalid and can not be processed.
func NewBadRequest(reason string) *errors.StatusError {
	return errors.NewBadRequest(reason)
}

// NewInvalid return a statusError
// which is an error intended for consumption by a REST API server; it can also be
// reconstructed by clients from a REST response. Public to allow easy type switches.
func NewInvalid(reason string) *errors.StatusError {
	return &errors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusInternalServerError,
			Reason:  metav1.StatusReasonInvalid,
			Message: reason,
		},
	}
}

func NewNotFound(reason string) *errors.StatusError {
	return &errors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusNotFound,
			Reason:  metav1.StatusReasonNotFound,
			Message: reason,
		},
	}
}

func NewInternal(reason string) *errors.StatusError {
	return &errors.StatusError{ErrStatus: metav1.Status{
		Status: metav1.StatusFailure,
		Code:   http.StatusInternalServerError,
		Reason: metav1.StatusReasonInternalError,
		Details: &metav1.StatusDetails{
			Causes: []metav1.StatusCause{{Message: reason}},
		},
		Message: fmt.Sprintf("Internal error occurred: %s", reason),
	}}
}

func NewUnexpectedObject(obj runtime.Object) *errors.StatusError {
	return &errors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusInternalServerError,
			Reason:  metav1.StatusReasonInternalError,
			Message: errors.FromObject(obj).Error(),
		},
	}
}

func NewGenericResponse(code int, serverMessage string) *errors.StatusError {
	reason := metav1.StatusReasonUnknown
	message := fmt.Sprintf("the server responded with the status code %d but did not return more information", code)
	switch code {
	case http.StatusNotFound:
		reason = metav1.StatusReasonNotFound
		message = "the server could not find the requested resource"
	default:
		if code >= 500 {
			reason = metav1.StatusReasonInternalError
			message = fmt.Sprintf("an error on the server (%q) has prevented the request from succeeding", serverMessage)
		}
	}

	return &errors.StatusError{ErrStatus: metav1.Status{
		Status:  metav1.StatusFailure,
		Code:    int32(code),
		Reason:  reason,
		Message: message,
	}}
}

func IsTokenExpired(err error) bool {
	statusErr, ok := err.(*errors.StatusError)
	if !ok {
		return false
	}

	return statusErr.ErrStatus.Message == MsgTokenExpiredError
}

// IsAlreadyExists determines if the err is an error which indicates that a specified resource already exists.
func IsAlreadyExists(err error) bool {
	return errors.IsAlreadyExists(err)
}

// IsUnauthorized determines if err is an error which indicates that the request is unauthorized and
// requires authentication by the user.
func IsUnauthorized(err error) bool {
	return errors.IsUnauthorized(err)
}
