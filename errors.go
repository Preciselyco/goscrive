package scrive

import (
	"fmt"

	"github.com/pkg/errors"
)

type ErrorType = string

const (
	ErrorTypeServerError                   ErrorType = "server_error"
	ErrorTypeEndpointNotFound              ErrorType = "endpoint_not_found"
	ErrorTypeInvalidAuthorisation          ErrorType = "invalid_authorisation"
	ErrorTypeInsufficientPrivileges        ErrorType = "insufficient_privileges"
	ErrorTypeResourceNotFound              ErrorType = "resource_not_found"
	ErrorTypeDocumentActionForbidden       ErrorType = "document_action_forbidden"
	ErrorTypeRequestParametersMissing      ErrorType = "request_parameters_missing"
	ErrorTypeRequestParametersParseError   ErrorType = "request_parameters_parse_error"
	ErrorTypeRequestParametersInvalid      ErrorType = "request_parameters_invalid"
	ErrorTypeDocumentObjectVersionMismatch ErrorType = "document_object_version_mismatch"
	ErrorTypeSignatoryStateError           ErrorType = "signatory_state_error"
	ErrorTypeLocal                         ErrorType = "local_error"
)

type se struct {
	ErrorType    *ErrorType `json:"error_type"`
	ErrorMessage *string    `json:"error_message"`
	HttpCode     *int       `json:"http_code"`
}

type ScriveError struct {
	ErrorType    ErrorType
	ErrorMessage string
	HttpCode     int
}

func (c *Client) parseResponseError(body []byte) (*ScriveError, error) {
	se := &se{}
	if err := parseJson(body, se); err != nil {
		return nil, err
	}
	if anyNil(se.ErrorType, se.ErrorMessage, se.HttpCode) {
		return nil, fmt.Errorf("Cannot parse: %s as ScriveError", string(body))
	}
	return &ScriveError{
		ErrorType:    *se.ErrorType,
		ErrorMessage: *se.ErrorMessage,
		HttpCode:     *se.HttpCode,
	}, nil
}

func localError(err error) *ScriveError {
	return &ScriveError{
		ErrorType:    ErrorTypeLocal,
		ErrorMessage: err.Error(),
		HttpCode:     -1,
	}
}

func ScriveErrorToError(e *ScriveError) error {
	return errors.Errorf("%d %s %s", e.HttpCode, string(e.ErrorType), e.ErrorMessage)
}
