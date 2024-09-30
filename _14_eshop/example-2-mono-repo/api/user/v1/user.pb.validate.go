// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/user/v1/user.proto

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

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Password

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Avatar

	// no validation rules for Age

	// no validation rules for Gender

	// no validation rules for Status

	// no validation rules for LoginAt

	// no validation rules for LoginType

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on CreateUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserReplyMultiError, or nil if none found.
func (m *CreateUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateUserReplyMultiError(errors)
	}

	return nil
}

// CreateUserReplyMultiError is an error wrapping multiple validation errors
// returned by CreateUserReply.ValidateAll() if the designated constraints
// aren't met.
type CreateUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserReplyMultiError) AllErrors() []error { return m }

// CreateUserReplyValidationError is the validation error returned by
// CreateUserReply.Validate if the designated constraints aren't met.
type CreateUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserReplyValidationError) ErrorName() string { return "CreateUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserReplyValidationError{}

// Validate checks the field values on DeleteUserByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserByIDRequestMultiError, or nil if none found.
func (m *DeleteUserByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteUserByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteUserByIDRequestMultiError(errors)
	}

	return nil
}

// DeleteUserByIDRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteUserByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteUserByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserByIDRequestMultiError) AllErrors() []error { return m }

// DeleteUserByIDRequestValidationError is the validation error returned by
// DeleteUserByIDRequest.Validate if the designated constraints aren't met.
type DeleteUserByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserByIDRequestValidationError) ErrorName() string {
	return "DeleteUserByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserByIDRequestValidationError{}

// Validate checks the field values on DeleteUserByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteUserByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteUserByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteUserByIDReplyMultiError, or nil if none found.
func (m *DeleteUserByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteUserByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteUserByIDReplyMultiError(errors)
	}

	return nil
}

// DeleteUserByIDReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteUserByIDReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteUserByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteUserByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteUserByIDReplyMultiError) AllErrors() []error { return m }

// DeleteUserByIDReplyValidationError is the validation error returned by
// DeleteUserByIDReply.Validate if the designated constraints aren't met.
type DeleteUserByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUserByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUserByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUserByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUserByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUserByIDReplyValidationError) ErrorName() string {
	return "DeleteUserByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUserByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUserByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUserByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUserByIDReplyValidationError{}

// Validate checks the field values on UpdateUserByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserByIDRequestMultiError, or nil if none found.
func (m *UpdateUserByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Password

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Avatar

	// no validation rules for Age

	// no validation rules for Gender

	// no validation rules for Status

	// no validation rules for LoginAt

	// no validation rules for LoginType

	if len(errors) > 0 {
		return UpdateUserByIDRequestMultiError(errors)
	}

	return nil
}

// UpdateUserByIDRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateUserByIDRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserByIDRequestMultiError) AllErrors() []error { return m }

// UpdateUserByIDRequestValidationError is the validation error returned by
// UpdateUserByIDRequest.Validate if the designated constraints aren't met.
type UpdateUserByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserByIDRequestValidationError) ErrorName() string {
	return "UpdateUserByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserByIDRequestValidationError{}

// Validate checks the field values on UpdateUserByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserByIDReplyMultiError, or nil if none found.
func (m *UpdateUserByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateUserByIDReplyMultiError(errors)
	}

	return nil
}

// UpdateUserByIDReplyMultiError is an error wrapping multiple validation
// errors returned by UpdateUserByIDReply.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserByIDReplyMultiError) AllErrors() []error { return m }

// UpdateUserByIDReplyValidationError is the validation error returned by
// UpdateUserByIDReply.Validate if the designated constraints aren't met.
type UpdateUserByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserByIDReplyValidationError) ErrorName() string {
	return "UpdateUserByIDReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserByIDReplyValidationError{}

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for Name

	// no validation rules for Password

	// no validation rules for Email

	// no validation rules for Phone

	// no validation rules for Avatar

	// no validation rules for Age

	// no validation rules for Gender

	// no validation rules for Status

	// no validation rules for LoginAt

	// no validation rules for LoginType

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on GetUserByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetUserByIDRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserByIDRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserByIDRequestMultiError, or nil if none found.
func (m *GetUserByIDRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserByIDRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetUserByIDRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserByIDRequestMultiError(errors)
	}

	return nil
}

// GetUserByIDRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserByIDRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserByIDRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserByIDRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserByIDRequestMultiError) AllErrors() []error { return m }

// GetUserByIDRequestValidationError is the validation error returned by
// GetUserByIDRequest.Validate if the designated constraints aren't met.
type GetUserByIDRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserByIDRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserByIDRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserByIDRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserByIDRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserByIDRequestValidationError) ErrorName() string {
	return "GetUserByIDRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserByIDRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserByIDRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserByIDRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserByIDRequestValidationError{}

// Validate checks the field values on GetUserByIDReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserByIDReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserByIDReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserByIDReplyMultiError, or nil if none found.
func (m *GetUserByIDReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserByIDReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserByIDReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserByIDReplyValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserByIDReplyValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserByIDReplyMultiError(errors)
	}

	return nil
}

// GetUserByIDReplyMultiError is an error wrapping multiple validation errors
// returned by GetUserByIDReply.ValidateAll() if the designated constraints
// aren't met.
type GetUserByIDReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserByIDReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserByIDReplyMultiError) AllErrors() []error { return m }

// GetUserByIDReplyValidationError is the validation error returned by
// GetUserByIDReply.Validate if the designated constraints aren't met.
type GetUserByIDReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserByIDReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserByIDReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserByIDReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserByIDReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserByIDReplyValidationError) ErrorName() string { return "GetUserByIDReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserByIDReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserByIDReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserByIDReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserByIDReplyValidationError{}

// Validate checks the field values on ListUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserRequestMultiError, or nil if none found.
func (m *ListUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetParams()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListUserRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListUserRequestValidationError{
					field:  "Params",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListUserRequestValidationError{
				field:  "Params",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListUserRequestMultiError(errors)
	}

	return nil
}

// ListUserRequestMultiError is an error wrapping multiple validation errors
// returned by ListUserRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserRequestMultiError) AllErrors() []error { return m }

// ListUserRequestValidationError is the validation error returned by
// ListUserRequest.Validate if the designated constraints aren't met.
type ListUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserRequestValidationError) ErrorName() string { return "ListUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserRequestValidationError{}

// Validate checks the field values on ListUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListUserReplyMultiError, or
// nil if none found.
func (m *ListUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserReplyValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserReplyValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserReplyValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListUserReplyMultiError(errors)
	}

	return nil
}

// ListUserReplyMultiError is an error wrapping multiple validation errors
// returned by ListUserReply.ValidateAll() if the designated constraints
// aren't met.
type ListUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserReplyMultiError) AllErrors() []error { return m }

// ListUserReplyValidationError is the validation error returned by
// ListUserReply.Validate if the designated constraints aren't met.
type ListUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserReplyValidationError) ErrorName() string { return "ListUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserReplyValidationError{}
