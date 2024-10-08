// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/eshop_gw/v1/eshop_gw.proto

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

// Validate checks the field values on PageParam with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageParam) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageParam with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageParamMultiError, or nil
// if none found.
func (m *PageParam) ValidateAll() error {
	return m.validate(true)
}

func (m *PageParam) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	if m.GetLimit() <= 0 {
		err := PageParamValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PageParamMultiError(errors)
	}

	return nil
}

// PageParamMultiError is an error wrapping multiple validation errors returned
// by PageParam.ValidateAll() if the designated constraints aren't met.
type PageParamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageParamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageParamMultiError) AllErrors() []error { return m }

// PageParamValidationError is the validation error returned by
// PageParam.Validate if the designated constraints aren't met.
type PageParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageParamValidationError) ErrorName() string { return "PageParamValidationError" }

// Error satisfies the builtin error interface
func (e PageParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageParamValidationError{}
