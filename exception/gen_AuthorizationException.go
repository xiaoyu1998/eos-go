// Code generated by gotemplate. DO NOT EDIT.

package exception

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/eosspark/eos-go/log"
)

// template type Exception(PARENT,CODE,WHAT)

var AuthorizationExceptionName = reflect.TypeOf(AuthorizationException{}).Name()

type AuthorizationException struct {
	_AuthorizationException
	Elog log.Messages
}

func NewAuthorizationException(parent _AuthorizationException, message log.Message) *AuthorizationException {
	return &AuthorizationException{parent, log.Messages{message}}
}

func (e AuthorizationException) Code() int64 {
	return 3090000
}

func (e AuthorizationException) Name() string {
	return AuthorizationExceptionName
}

func (e AuthorizationException) What() string {
	return "Authorization exception"
}

func (e *AuthorizationException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e AuthorizationException) GetLog() log.Messages {
	return e.Elog
}

func (e AuthorizationException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e AuthorizationException) DetailMessage() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(int(e.Code())))
	buffer.WriteString(" ")
	buffer.WriteString(e.Name())
	buffer.WriteString(": ")
	buffer.WriteString(e.What())
	buffer.WriteString("\n")
	for _, l := range e.Elog {
		buffer.WriteString("[")
		buffer.WriteString(l.GetMessage())
		buffer.WriteString("]")
		buffer.WriteString("\n")
		buffer.WriteString(l.GetContext().String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (e AuthorizationException) String() string {
	return e.DetailMessage()
}

func (e AuthorizationException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3090000,
		Name: AuthorizationExceptionName,
		What: "Authorization exception",
	}

	return json.Marshal(except)
}

func (e AuthorizationException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*AuthorizationException):
		callback(&e)
		return true
	case func(AuthorizationException):
		callback(e)
		return true
	default:
		return false
	}
}