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

var OverflowExceptionName = reflect.TypeOf(OverflowException{}).Name()

type OverflowException struct {
	Exception
	Elog log.Messages
}

func NewOverflowException(parent Exception, message log.Message) *OverflowException {
	return &OverflowException{parent, log.Messages{message}}
}

func (e OverflowException) Code() int64 {
	return OverflowCode
}

func (e OverflowException) Name() string {
	return OverflowExceptionName
}

func (e OverflowException) What() string {
	return "Integer Overflow"
}

func (e *OverflowException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e OverflowException) GetLog() log.Messages {
	return e.Elog
}

func (e OverflowException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e OverflowException) DetailMessage() string {
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
		buffer.WriteString("] ")
		buffer.WriteString(l.GetContext().String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (e OverflowException) String() string {
	return e.DetailMessage()
}

func (e OverflowException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: OverflowCode,
		Name: OverflowExceptionName,
		What: "Integer Overflow",
	}

	return json.Marshal(except)
}

func (e OverflowException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*OverflowException):
		callback(&e)
		return true
	case func(OverflowException):
		callback(e)
		return true
	default:
		return false
	}
}
