package htsgeterror

import (
	"net/http"
	"strconv"
)

const codeBadRequest = http.StatusBadRequest
const codeInvalidAuthentication = http.StatusUnauthorized
const codePermissionDenied = http.StatusForbidden
const codeNotFound = http.StatusNotFound
const codeInternalServerError = http.StatusInternalServerError

const errorBadRequestUnsupportedFormat = "UnsupportedFormat"
const errorBadRequestInvalidInput = "InvalidInput"
const errorBadRequestInvalidRange = "InvalidRange"
const errorInvalidAuthentication = "InvalidAuthentication"
const errorPermissionDenied = "PermissionDenied"
const errorNotFound = "NotFound"
const errorInternalServerError = "InternalServerError"

const dfltMsgBadRequestUnsupportedFormat = "The requested file format is not supported by the server"
const dfltMsgBadRequestInvalidInput = "The request parameters do not adhere to the specification"
const dfltMsgBadRequestInvalidRange = "The requested range cannot be satisfied"
const dfltMsgInvalidAuthentication = "Authorization provided is invalid"
const dfltMsgPermissionDenied = "Authorization is required to access the resource"
const dfltMsgNotFound = "The resource requested was not found"
const dfltMsgInternalServerError = "Internal server error"

var errorInfoMap = map[string]map[string]string{
	errorBadRequestUnsupportedFormat: {
		"code":    strconv.Itoa(codeBadRequest),
		"dfltMsg": dfltMsgBadRequestUnsupportedFormat,
	},
	errorBadRequestInvalidInput: {
		"code":    strconv.Itoa(codeBadRequest),
		"dfltMsg": dfltMsgBadRequestInvalidInput,
	},
	errorBadRequestInvalidRange: {
		"code":    strconv.Itoa(codeBadRequest),
		"dfltMsg": dfltMsgBadRequestInvalidRange,
	},
	errorInvalidAuthentication: {
		"code":    strconv.Itoa(codeInvalidAuthentication),
		"dfltMsg": dfltMsgInvalidAuthentication,
	},
	errorPermissionDenied: {
		"code":    strconv.Itoa(codePermissionDenied),
		"dfltMsg": dfltMsgPermissionDenied,
	},
	errorNotFound: {
		"code":    strconv.Itoa(codeNotFound),
		"dfltMsg": dfltMsgNotFound,
	},
	errorInternalServerError: {
		"code":    strconv.Itoa(codeInternalServerError),
		"dfltMsg": dfltMsgInternalServerError,
	},
}
