package errs

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// const for fast initialize
var (
	NotFound     = ErrNotFound{"Record not found"}
	Unknown      = ErrInternalServer{"Server is currently unable to handle the request"}
	Invalid      = ErrInvalidType{"Bad Request"}
	Forbidden    = ErrForbidden{"Permission denied"}
	Unauthorized = ErrUnauthorized{"Unauthorized, access denied please login"}
)

// Response a struct to response error back to requester
type Response struct {
	Error string `json:"error"`
}

// NewError a fucntion that help create an error from msg and code
func NewError(code int, msg string) error {
	return errors.New(strconv.Itoa(code) + ":" + msg)
}

// WriteError a function that help write response error
func WriteError(w http.ResponseWriter, code int, msg string) {
	res := Response{Error: msg}
	resJSON, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resJSON)
}

// ResponseError a function helps response error
func ResponseError(w http.ResponseWriter, err error) {
	msg := strings.Split(err.Error(), ":")
	code, err := strconv.Atoi(msg[0])
	if err != nil {
		WriteError(w, 500, "Server is currently unable to handle the request")
		return
	}
	WriteError(w, code, msg[1])
}

// ErrInvalidType BadRequest Error
type ErrInvalidType struct {
	Err string
}

func (e ErrInvalidType) Error() string {
	return "400:" + e.Err
}

// ErrUnauthorized BadRequest Error
type ErrUnauthorized struct {
	Err string
}

func (e ErrUnauthorized) Error() string {
	return "401:" + e.Err
}

// ErrForbidden Forbidden Error
type ErrForbidden struct {
	Err string
}

func (e ErrForbidden) Error() string {
	return "403:" + e.Err
}

// ErrNotFound NotFound Error
type ErrNotFound struct {
	Err string
}

func (e ErrNotFound) Error() string {
	return "404:" + e.Err
}

// ErrInternalServer Server Error
type ErrInternalServer struct {
	Err string
}

func (e ErrInternalServer) Error() string {
	return "500:" + e.Err
}
