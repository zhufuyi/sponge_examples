// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/order/v1/order.proto

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

// Validate checks the field values on SubmitRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubmitRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubmitRequestMultiError, or
// nil if none found.
func (m *SubmitRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserID() <= 0 {
		err := SubmitRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductID() <= 0 {
		err := SubmitRequestValidationError{
			field:  "ProductID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductCount() <= 0 {
		err := SubmitRequestValidationError{
			field:  "ProductCount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAmount() <= 0 {
		err := SubmitRequestValidationError{
			field:  "Amount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CouponID

	if len(errors) > 0 {
		return SubmitRequestMultiError(errors)
	}

	return nil
}

// SubmitRequestMultiError is an error wrapping multiple validation errors
// returned by SubmitRequest.ValidateAll() if the designated constraints
// aren't met.
type SubmitRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitRequestMultiError) AllErrors() []error { return m }

// SubmitRequestValidationError is the validation error returned by
// SubmitRequest.Validate if the designated constraints aren't met.
type SubmitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitRequestValidationError) ErrorName() string { return "SubmitRequestValidationError" }

// Error satisfies the builtin error interface
func (e SubmitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitRequestValidationError{}

// Validate checks the field values on SubmitReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SubmitReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SubmitReplyMultiError, or
// nil if none found.
func (m *SubmitReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderID

	if len(errors) > 0 {
		return SubmitReplyMultiError(errors)
	}

	return nil
}

// SubmitReplyMultiError is an error wrapping multiple validation errors
// returned by SubmitReply.ValidateAll() if the designated constraints aren't met.
type SubmitReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitReplyMultiError) AllErrors() []error { return m }

// SubmitReplyValidationError is the validation error returned by
// SubmitReply.Validate if the designated constraints aren't met.
type SubmitReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitReplyValidationError) ErrorName() string { return "SubmitReplyValidationError" }

// Error satisfies the builtin error interface
func (e SubmitReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitReplyValidationError{}

// Validate checks the field values on CreateOrderReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderReplyMultiError, or nil if none found.
func (m *CreateOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateOrderReplyMultiError(errors)
	}

	return nil
}

// CreateOrderReplyMultiError is an error wrapping multiple validation errors
// returned by CreateOrderReply.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderReplyMultiError) AllErrors() []error { return m }

// CreateOrderReplyValidationError is the validation error returned by
// CreateOrderReply.Validate if the designated constraints aren't met.
type CreateOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderReplyValidationError) ErrorName() string { return "CreateOrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderReplyValidationError{}

// Validate checks the field values on CreateOrderRevertRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRevertRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRevertRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRevertRequestMultiError, or nil if none found.
func (m *CreateOrderRevertRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRevertRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetOrderID()) < 16 {
		err := CreateOrderRevertRequestValidationError{
			field:  "OrderID",
			reason: "value length must be at least 16 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserID() <= 0 {
		err := CreateOrderRevertRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductID() <= 0 {
		err := CreateOrderRevertRequestValidationError{
			field:  "ProductID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductCount() <= 0 {
		err := CreateOrderRevertRequestValidationError{
			field:  "ProductCount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAmount() <= 0 {
		err := CreateOrderRevertRequestValidationError{
			field:  "Amount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CouponID

	if len(errors) > 0 {
		return CreateOrderRevertRequestMultiError(errors)
	}

	return nil
}

// CreateOrderRevertRequestMultiError is an error wrapping multiple validation
// errors returned by CreateOrderRevertRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateOrderRevertRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRevertRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRevertRequestMultiError) AllErrors() []error { return m }

// CreateOrderRevertRequestValidationError is the validation error returned by
// CreateOrderRevertRequest.Validate if the designated constraints aren't met.
type CreateOrderRevertRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRevertRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRevertRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRevertRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRevertRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRevertRequestValidationError) ErrorName() string {
	return "CreateOrderRevertRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRevertRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRevertRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRevertRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRevertRequestValidationError{}

// Validate checks the field values on CreateOrderRevertReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRevertReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRevertReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRevertReplyMultiError, or nil if none found.
func (m *CreateOrderRevertReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRevertReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateOrderRevertReplyMultiError(errors)
	}

	return nil
}

// CreateOrderRevertReplyMultiError is an error wrapping multiple validation
// errors returned by CreateOrderRevertReply.ValidateAll() if the designated
// constraints aren't met.
type CreateOrderRevertReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRevertReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRevertReplyMultiError) AllErrors() []error { return m }

// CreateOrderRevertReplyValidationError is the validation error returned by
// CreateOrderRevertReply.Validate if the designated constraints aren't met.
type CreateOrderRevertReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRevertReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRevertReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRevertReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRevertReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRevertReplyValidationError) ErrorName() string {
	return "CreateOrderRevertReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRevertReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRevertReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRevertReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRevertReplyValidationError{}

// Validate checks the field values on CouponUseRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CouponUseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CouponUseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CouponUseRequestMultiError, or nil if none found.
func (m *CouponUseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CouponUseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCouponID() <= 0 {
		err := CouponUseRequestValidationError{
			field:  "CouponID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CouponUseRequestMultiError(errors)
	}

	return nil
}

// CouponUseRequestMultiError is an error wrapping multiple validation errors
// returned by CouponUseRequest.ValidateAll() if the designated constraints
// aren't met.
type CouponUseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CouponUseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CouponUseRequestMultiError) AllErrors() []error { return m }

// CouponUseRequestValidationError is the validation error returned by
// CouponUseRequest.Validate if the designated constraints aren't met.
type CouponUseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CouponUseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CouponUseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CouponUseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CouponUseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CouponUseRequestValidationError) ErrorName() string { return "CouponUseRequestValidationError" }

// Error satisfies the builtin error interface
func (e CouponUseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCouponUseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CouponUseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CouponUseRequestValidationError{}

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

// Validate checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrderRequestMultiError, or nil if none found.
func (m *CreateOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetOrderID()) < 16 {
		err := CreateOrderRequestValidationError{
			field:  "OrderID",
			reason: "value length must be at least 16 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUserID() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductID() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "ProductID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetProductCount() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "ProductCount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAmount() <= 0 {
		err := CreateOrderRequestValidationError{
			field:  "Amount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for CouponID

	if len(errors) > 0 {
		return CreateOrderRequestMultiError(errors)
	}

	return nil
}

// CreateOrderRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOrderRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrderRequestMultiError) AllErrors() []error { return m }

// CreateOrderRequestValidationError is the validation error returned by
// CreateOrderRequest.Validate if the designated constraints aren't met.
type CreateOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrderRequestValidationError) ErrorName() string {
	return "CreateOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrderRequestValidationError{}

// Validate checks the field values on CreatePayRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreatePayRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePayRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePayRequestMultiError, or nil if none found.
func (m *CreatePayRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePayRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserID() <= 0 {
		err := CreatePayRequestValidationError{
			field:  "UserID",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetOrderID()) < 16 {
		err := CreatePayRequestValidationError{
			field:  "OrderID",
			reason: "value length must be at least 16 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetAmount() <= 0 {
		err := CreatePayRequestValidationError{
			field:  "Amount",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreatePayRequestMultiError(errors)
	}

	return nil
}

// CreatePayRequestMultiError is an error wrapping multiple validation errors
// returned by CreatePayRequest.ValidateAll() if the designated constraints
// aren't met.
type CreatePayRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePayRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePayRequestMultiError) AllErrors() []error { return m }

// CreatePayRequestValidationError is the validation error returned by
// CreatePayRequest.Validate if the designated constraints aren't met.
type CreatePayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePayRequestValidationError) ErrorName() string { return "CreatePayRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreatePayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePayRequestValidationError{}
