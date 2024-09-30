// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/stock/v1/atomic.proto

package v1

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

// Validate checks the field values on UpdateAtomicRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAtomicRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAtomicRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAtomicRequestMultiError, or nil if none found.
func (m *UpdateAtomicRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAtomicRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateAtomicRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStock() <= 0 {
		err := UpdateAtomicRequestValidationError{
			field:  "Stock",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateAtomicRequestMultiError(errors)
	}

	return nil
}

// UpdateAtomicRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateAtomicRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateAtomicRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAtomicRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAtomicRequestMultiError) AllErrors() []error { return m }

// UpdateAtomicRequestValidationError is the validation error returned by
// UpdateAtomicRequest.Validate if the designated constraints aren't met.
type UpdateAtomicRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAtomicRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAtomicRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAtomicRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAtomicRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAtomicRequestValidationError) ErrorName() string {
	return "UpdateAtomicRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAtomicRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAtomicRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAtomicRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAtomicRequestValidationError{}

// Validate checks the field values on UpdateAtomicRequestReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateAtomicRequestReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateAtomicRequestReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateAtomicRequestReplyMultiError, or nil if none found.
func (m *UpdateAtomicRequestReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateAtomicRequestReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateAtomicRequestReplyMultiError(errors)
	}

	return nil
}

// UpdateAtomicRequestReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateAtomicRequestReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateAtomicRequestReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateAtomicRequestReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateAtomicRequestReplyMultiError) AllErrors() []error { return m }

// UpdateAtomicRequestReplyValidationError is the validation error returned by
// UpdateAtomicRequestReply.Validate if the designated constraints aren't met.
type UpdateAtomicRequestReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateAtomicRequestReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateAtomicRequestReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateAtomicRequestReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateAtomicRequestReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateAtomicRequestReplyValidationError) ErrorName() string {
	return "UpdateAtomicRequestReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateAtomicRequestReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateAtomicRequestReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateAtomicRequestReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateAtomicRequestReplyValidationError{}

// Validate checks the field values on QueryAtomicRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QueryAtomicRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryAtomicRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryAtomicRequestMultiError, or nil if none found.
func (m *QueryAtomicRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryAtomicRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := QueryAtomicRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QueryAtomicRequestMultiError(errors)
	}

	return nil
}

// QueryAtomicRequestMultiError is an error wrapping multiple validation errors
// returned by QueryAtomicRequest.ValidateAll() if the designated constraints
// aren't met.
type QueryAtomicRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryAtomicRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryAtomicRequestMultiError) AllErrors() []error { return m }

// QueryAtomicRequestValidationError is the validation error returned by
// QueryAtomicRequest.Validate if the designated constraints aren't met.
type QueryAtomicRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryAtomicRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryAtomicRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryAtomicRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryAtomicRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryAtomicRequestValidationError) ErrorName() string {
	return "QueryAtomicRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QueryAtomicRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryAtomicRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryAtomicRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryAtomicRequestValidationError{}

// Validate checks the field values on QueryAtomicReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QueryAtomicReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryAtomicReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QueryAtomicReplyMultiError, or nil if none found.
func (m *QueryAtomicReply) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryAtomicReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Stock

	if len(errors) > 0 {
		return QueryAtomicReplyMultiError(errors)
	}

	return nil
}

// QueryAtomicReplyMultiError is an error wrapping multiple validation errors
// returned by QueryAtomicReply.ValidateAll() if the designated constraints
// aren't met.
type QueryAtomicReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryAtomicReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryAtomicReplyMultiError) AllErrors() []error { return m }

// QueryAtomicReplyValidationError is the validation error returned by
// QueryAtomicReply.Validate if the designated constraints aren't met.
type QueryAtomicReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryAtomicReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryAtomicReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryAtomicReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryAtomicReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryAtomicReplyValidationError) ErrorName() string { return "QueryAtomicReplyValidationError" }

// Error satisfies the builtin error interface
func (e QueryAtomicReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryAtomicReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryAtomicReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryAtomicReplyValidationError{}