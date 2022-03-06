package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type ApiResp struct {
	StatusCode int         `json:"-"`
	OK         bool        `json:"ok"`
	Data       interface{} `json:"data,omitempty"`
	ErrCode    string      `json:"err,omitempty"`
	Message    string      `json:"msg,omitempty"`
}

func NewSuccessResp(data interface{}) ApiResp {
	return ApiResp{
		StatusCode: http.StatusOK,
		OK:         true,
		Data:       data,
	}
}

func NewErrorResp(err error) ApiResp {
	var e *Error
	if !errors.As(err, &e) {
		e = NewErrInternalError(err)
	}
	return ApiResp{
		StatusCode: e.StatusCode,
		OK:         false,
		ErrCode:    e.ErrCode,
		Message:    e.Message,
	}
}

func WriteAPIResp(w http.ResponseWriter, resp ApiResp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	b, _ := json.Marshal(resp)
	w.Write(b)
}

type searchReqBody struct {
	Data string `json:"data"`
}

func (rb searchReqBody) Validate() error {
	if len(rb.Data) == 0 {
		return NewErrBadRequest("missing `data`")
	}
	_, err := rb.GetByte()
	if err != nil {
		return NewErrBadRequest("invalid `data`")
	}
	return nil
}

func (rb searchReqBody) GetByte() ([]byte, error) {
	tokens := strings.Split(rb.Data, ",")
	if len(tokens) != 2 {
		return nil, NewErrBadRequest("unexpected number of tokens")
	}
	data, err := base64.StdEncoding.DecodeString(tokens[1])
	if err != nil {
		return nil, NewErrBadRequest(err.Error())
	}
	return data, nil
}

type dbRecord struct {
	FileName string
	Ima     []byte
}

type similarImage struct {
	FileName        string  `json:"filename"`
	SimilarityScore float64 `json:"similarity_score"`
}
