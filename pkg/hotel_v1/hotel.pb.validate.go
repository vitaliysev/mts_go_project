// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hotel.proto

package hotel_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on GetInfoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInfoRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetInfoRequestMultiError,
// or nil if none found.
func (m *GetInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetInfoRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetInfoRequestMultiError(errors)
	}

	return nil
}

// GetInfoRequestMultiError is an error wrapping multiple validation errors
// returned by GetInfoRequest.ValidateAll() if the designated constraints
// aren't met.
type GetInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInfoRequestMultiError) AllErrors() []error { return m }

// GetInfoRequestValidationError is the validation error returned by
// GetInfoRequest.Validate if the designated constraints aren't met.
type GetInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoRequestValidationError) ErrorName() string { return "GetInfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoRequestValidationError{}

// Validate checks the field values on GetInfoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetInfoResponseMultiError, or nil if none found.
func (m *GetInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetHotel()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetInfoResponseValidationError{
					field:  "Hotel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetInfoResponseValidationError{
					field:  "Hotel",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHotel()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetInfoResponseValidationError{
				field:  "Hotel",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetInfoResponseMultiError(errors)
	}

	return nil
}

// GetInfoResponseMultiError is an error wrapping multiple validation errors
// returned by GetInfoResponse.ValidateAll() if the designated constraints
// aren't met.
type GetInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInfoResponseMultiError) AllErrors() []error { return m }

// GetInfoResponseValidationError is the validation error returned by
// GetInfoResponse.Validate if the designated constraints aren't met.
type GetInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoResponseValidationError) ErrorName() string { return "GetInfoResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoResponseValidationError{}

// Validate checks the field values on HotelInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HotelInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HotelInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HotelInfoMultiError, or nil
// if none found.
func (m *HotelInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *HotelInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Location

	// no validation rules for Price

	if len(errors) > 0 {
		return HotelInfoMultiError(errors)
	}

	return nil
}

// HotelInfoMultiError is an error wrapping multiple validation errors returned
// by HotelInfo.ValidateAll() if the designated constraints aren't met.
type HotelInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HotelInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HotelInfoMultiError) AllErrors() []error { return m }

// HotelInfoValidationError is the validation error returned by
// HotelInfo.Validate if the designated constraints aren't met.
type HotelInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HotelInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HotelInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HotelInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HotelInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HotelInfoValidationError) ErrorName() string { return "HotelInfoValidationError" }

// Error satisfies the builtin error interface
func (e HotelInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHotelInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HotelInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HotelInfoValidationError{}