// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/stock/v1/stock.proto

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

// Validate checks the field values on StockDeductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StockDeductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StockDeductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StockDeductRequestMultiError, or nil if none found.
func (m *StockDeductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StockDeductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetProductID() <= 0 {
		err := StockDeductRequestValidationError{
			field:  "ProductID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductCount() <= 0 {
		err := StockDeductRequestValidationError{
			field:  "ProductCount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StockDeductRequestMultiError(errors)
	}

	return nil
}

// StockDeductRequestMultiError is an error wrapping multiple validation errors
// returned by StockDeductRequest.ValidateAll() if the designated constraints
// aren't met.
type StockDeductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockDeductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockDeductRequestMultiError) AllErrors() []error { return m }

// StockDeductRequestValidationError is the validation error returned by
// StockDeductRequest.Validate if the designated constraints aren't met.
type StockDeductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockDeductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockDeductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockDeductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockDeductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockDeductRequestValidationError) ErrorName() string {
	return "StockDeductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StockDeductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStockDeductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockDeductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockDeductRequestValidationError{}

// Validate checks the field values on StockDeductReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StockDeductReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StockDeductReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StockDeductReplyMultiError, or nil if none found.
func (m *StockDeductReply) ValidateAll() error {
	return m.validate(true)
}

func (m *StockDeductReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StockDeductReplyMultiError(errors)
	}

	return nil
}

// StockDeductReplyMultiError is an error wrapping multiple validation errors
// returned by StockDeductReply.ValidateAll() if the designated constraints
// aren't met.
type StockDeductReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockDeductReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockDeductReplyMultiError) AllErrors() []error { return m }

// StockDeductReplyValidationError is the validation error returned by
// StockDeductReply.Validate if the designated constraints aren't met.
type StockDeductReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockDeductReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockDeductReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockDeductReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockDeductReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockDeductReplyValidationError) ErrorName() string { return "StockDeductReplyValidationError" }

// Error satisfies the builtin error interface
func (e StockDeductReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStockDeductReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockDeductReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockDeductReplyValidationError{}

// Validate checks the field values on StockDeductRevertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StockDeductRevertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StockDeductRevertRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StockDeductRevertRequestMultiError, or nil if none found.
func (m *StockDeductRevertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StockDeductRevertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetProductID() <= 0 {
		err := StockDeductRevertRequestValidationError{
			field:  "ProductID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductCount() <= 0 {
		err := StockDeductRevertRequestValidationError{
			field:  "ProductCount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StockDeductRevertRequestMultiError(errors)
	}

	return nil
}

// StockDeductRevertRequestMultiError is an error wrapping multiple validation
// errors returned by StockDeductRevertRequest.ValidateAll() if the designated
// constraints aren't met.
type StockDeductRevertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockDeductRevertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockDeductRevertRequestMultiError) AllErrors() []error { return m }

// StockDeductRevertRequestValidationError is the validation error returned by
// StockDeductRevertRequest.Validate if the designated constraints aren't met.
type StockDeductRevertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockDeductRevertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockDeductRevertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockDeductRevertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockDeductRevertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockDeductRevertRequestValidationError) ErrorName() string {
	return "StockDeductRevertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StockDeductRevertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStockDeductRevertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockDeductRevertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockDeductRevertRequestValidationError{}

// Validate checks the field values on StockDeductRevertReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StockDeductRevertReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StockDeductRevertReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StockDeductRevertReplyMultiError, or nil if none found.
func (m *StockDeductRevertReply) ValidateAll() error {
	return m.validate(true)
}

func (m *StockDeductRevertReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StockDeductRevertReplyMultiError(errors)
	}

	return nil
}

// StockDeductRevertReplyMultiError is an error wrapping multiple validation
// errors returned by StockDeductRevertReply.ValidateAll() if the designated
// constraints aren't met.
type StockDeductRevertReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StockDeductRevertReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StockDeductRevertReplyMultiError) AllErrors() []error { return m }

// StockDeductRevertReplyValidationError is the validation error returned by
// StockDeductRevertReply.Validate if the designated constraints aren't met.
type StockDeductRevertReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StockDeductRevertReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StockDeductRevertReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StockDeductRevertReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StockDeductRevertReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StockDeductRevertReplyValidationError) ErrorName() string {
	return "StockDeductRevertReplyValidationError"
}

// Error satisfies the builtin error interface
func (e StockDeductRevertReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStockDeductRevertReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StockDeductRevertReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StockDeductRevertReplyValidationError{}
