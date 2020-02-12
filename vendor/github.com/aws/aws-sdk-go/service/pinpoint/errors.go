// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// Provides information about an API request or response.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// Provides information about an API request or response.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// Provides information about an API request or response.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeMethodNotAllowedException for service response error code
	// "MethodNotAllowedException".
	//
	// Provides information about an API request or response.
	ErrCodeMethodNotAllowedException = "MethodNotAllowedException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// Provides information about an API request or response.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Provides information about an API request or response.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":          newErrorBadRequestException,
	"ForbiddenException":           newErrorForbiddenException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"MethodNotAllowedException":    newErrorMethodNotAllowedException,
	"NotFoundException":            newErrorNotFoundException,
	"TooManyRequestsException":     newErrorTooManyRequestsException,
}
