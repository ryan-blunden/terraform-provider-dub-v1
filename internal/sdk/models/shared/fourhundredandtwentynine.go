// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// FourHundredAndTwentyNineCode - A short code indicating the error code returned.
type FourHundredAndTwentyNineCode string

const (
	FourHundredAndTwentyNineCodeRateLimitExceeded FourHundredAndTwentyNineCode = "rate_limit_exceeded"
)

func (e FourHundredAndTwentyNineCode) ToPointer() *FourHundredAndTwentyNineCode {
	return &e
}
func (e *FourHundredAndTwentyNineCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rate_limit_exceeded":
		*e = FourHundredAndTwentyNineCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FourHundredAndTwentyNineCode: %v", v)
	}
}

type FourHundredAndTwentyNineError struct {
	// A short code indicating the error code returned.
	Code FourHundredAndTwentyNineCode `json:"code"`
	// A human readable explanation of what went wrong.
	Message string `json:"message"`
	// A link to our documentation with more details about this error code
	DocURL *string `json:"doc_url,omitempty"`
}

func (o *FourHundredAndTwentyNineError) GetCode() FourHundredAndTwentyNineCode {
	if o == nil {
		return FourHundredAndTwentyNineCode("")
	}
	return o.Code
}

func (o *FourHundredAndTwentyNineError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *FourHundredAndTwentyNineError) GetDocURL() *string {
	if o == nil {
		return nil
	}
	return o.DocURL
}

// FourHundredAndTwentyNine - The user has sent too many requests in a given amount of time ("rate limiting")
type FourHundredAndTwentyNine struct {
	Error FourHundredAndTwentyNineError `json:"error"`
}

func (o *FourHundredAndTwentyNine) GetError() FourHundredAndTwentyNineError {
	if o == nil {
		return FourHundredAndTwentyNineError{}
	}
	return o.Error
}
