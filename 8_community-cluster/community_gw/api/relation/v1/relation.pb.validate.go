// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/relation/v1/relation.proto

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

// Validate checks the field values on FollowRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FollowRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FollowRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FollowRequestMultiError, or
// nil if none found.
func (m *FollowRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FollowRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := FollowRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetFollowedUid() < 1 {
		err := FollowRequestValidationError{
			field:  "FollowedUid",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return FollowRequestMultiError(errors)
	}

	return nil
}

// FollowRequestMultiError is an error wrapping multiple validation errors
// returned by FollowRequest.ValidateAll() if the designated constraints
// aren't met.
type FollowRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FollowRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FollowRequestMultiError) AllErrors() []error { return m }

// FollowRequestValidationError is the validation error returned by
// FollowRequest.Validate if the designated constraints aren't met.
type FollowRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FollowRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FollowRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FollowRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FollowRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FollowRequestValidationError) ErrorName() string { return "FollowRequestValidationError" }

// Error satisfies the builtin error interface
func (e FollowRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFollowRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FollowRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FollowRequestValidationError{}

// Validate checks the field values on FollowReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FollowReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FollowReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FollowReplyMultiError, or
// nil if none found.
func (m *FollowReply) ValidateAll() error {
	return m.validate(true)
}

func (m *FollowReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FollowReplyMultiError(errors)
	}

	return nil
}

// FollowReplyMultiError is an error wrapping multiple validation errors
// returned by FollowReply.ValidateAll() if the designated constraints aren't met.
type FollowReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FollowReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FollowReplyMultiError) AllErrors() []error { return m }

// FollowReplyValidationError is the validation error returned by
// FollowReply.Validate if the designated constraints aren't met.
type FollowReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FollowReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FollowReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FollowReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FollowReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FollowReplyValidationError) ErrorName() string { return "FollowReplyValidationError" }

// Error satisfies the builtin error interface
func (e FollowReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFollowReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FollowReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FollowReplyValidationError{}

// Validate checks the field values on UnfollowRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UnfollowRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnfollowRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnfollowRequestMultiError, or nil if none found.
func (m *UnfollowRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UnfollowRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := UnfollowRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetFollowedUid() < 1 {
		err := UnfollowRequestValidationError{
			field:  "FollowedUid",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UnfollowRequestMultiError(errors)
	}

	return nil
}

// UnfollowRequestMultiError is an error wrapping multiple validation errors
// returned by UnfollowRequest.ValidateAll() if the designated constraints
// aren't met.
type UnfollowRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnfollowRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnfollowRequestMultiError) AllErrors() []error { return m }

// UnfollowRequestValidationError is the validation error returned by
// UnfollowRequest.Validate if the designated constraints aren't met.
type UnfollowRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnfollowRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnfollowRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnfollowRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnfollowRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnfollowRequestValidationError) ErrorName() string { return "UnfollowRequestValidationError" }

// Error satisfies the builtin error interface
func (e UnfollowRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnfollowRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnfollowRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnfollowRequestValidationError{}

// Validate checks the field values on UnfollowReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UnfollowReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnfollowReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UnfollowReplyMultiError, or
// nil if none found.
func (m *UnfollowReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UnfollowReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UnfollowReplyMultiError(errors)
	}

	return nil
}

// UnfollowReplyMultiError is an error wrapping multiple validation errors
// returned by UnfollowReply.ValidateAll() if the designated constraints
// aren't met.
type UnfollowReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnfollowReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnfollowReplyMultiError) AllErrors() []error { return m }

// UnfollowReplyValidationError is the validation error returned by
// UnfollowReply.Validate if the designated constraints aren't met.
type UnfollowReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnfollowReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnfollowReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnfollowReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnfollowReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnfollowReplyValidationError) ErrorName() string { return "UnfollowReplyValidationError" }

// Error satisfies the builtin error interface
func (e UnfollowReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnfollowReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnfollowReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnfollowReplyValidationError{}

// Validate checks the field values on ListFollowingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListFollowingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListFollowingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListFollowingRequestMultiError, or nil if none found.
func (m *ListFollowingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListFollowingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := ListFollowingRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() < 0 {
		err := ListFollowingRequestValidationError{
			field:  "Page",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() < 1 {
		err := ListFollowingRequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListFollowingRequestMultiError(errors)
	}

	return nil
}

// ListFollowingRequestMultiError is an error wrapping multiple validation
// errors returned by ListFollowingRequest.ValidateAll() if the designated
// constraints aren't met.
type ListFollowingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListFollowingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListFollowingRequestMultiError) AllErrors() []error { return m }

// ListFollowingRequestValidationError is the validation error returned by
// ListFollowingRequest.Validate if the designated constraints aren't met.
type ListFollowingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFollowingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFollowingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFollowingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFollowingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFollowingRequestValidationError) ErrorName() string {
	return "ListFollowingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListFollowingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFollowingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFollowingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFollowingRequestValidationError{}

// Validate checks the field values on ListFollowingReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListFollowingReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListFollowingReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListFollowingReplyMultiError, or nil if none found.
func (m *ListFollowingReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListFollowingReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return ListFollowingReplyMultiError(errors)
	}

	return nil
}

// ListFollowingReplyMultiError is an error wrapping multiple validation errors
// returned by ListFollowingReply.ValidateAll() if the designated constraints
// aren't met.
type ListFollowingReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListFollowingReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListFollowingReplyMultiError) AllErrors() []error { return m }

// ListFollowingReplyValidationError is the validation error returned by
// ListFollowingReply.Validate if the designated constraints aren't met.
type ListFollowingReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFollowingReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFollowingReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFollowingReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFollowingReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFollowingReplyValidationError) ErrorName() string {
	return "ListFollowingReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListFollowingReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFollowingReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFollowingReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFollowingReplyValidationError{}

// Validate checks the field values on ListFollowerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListFollowerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListFollowerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListFollowerRequestMultiError, or nil if none found.
func (m *ListFollowerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListFollowerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := ListFollowerRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() < 0 {
		err := ListFollowerRequestValidationError{
			field:  "Page",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() < 1 {
		err := ListFollowerRequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListFollowerRequestMultiError(errors)
	}

	return nil
}

// ListFollowerRequestMultiError is an error wrapping multiple validation
// errors returned by ListFollowerRequest.ValidateAll() if the designated
// constraints aren't met.
type ListFollowerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListFollowerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListFollowerRequestMultiError) AllErrors() []error { return m }

// ListFollowerRequestValidationError is the validation error returned by
// ListFollowerRequest.Validate if the designated constraints aren't met.
type ListFollowerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFollowerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFollowerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFollowerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFollowerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFollowerRequestValidationError) ErrorName() string {
	return "ListFollowerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListFollowerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFollowerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFollowerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFollowerRequestValidationError{}

// Validate checks the field values on ListFollowerReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListFollowerReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListFollowerReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListFollowerReplyMultiError, or nil if none found.
func (m *ListFollowerReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListFollowerReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return ListFollowerReplyMultiError(errors)
	}

	return nil
}

// ListFollowerReplyMultiError is an error wrapping multiple validation errors
// returned by ListFollowerReply.ValidateAll() if the designated constraints
// aren't met.
type ListFollowerReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListFollowerReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListFollowerReplyMultiError) AllErrors() []error { return m }

// ListFollowerReplyValidationError is the validation error returned by
// ListFollowerReply.Validate if the designated constraints aren't met.
type ListFollowerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFollowerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFollowerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFollowerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFollowerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFollowerReplyValidationError) ErrorName() string {
	return "ListFollowerReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListFollowerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFollowerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFollowerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFollowerReplyValidationError{}

// Validate checks the field values on BatchGetRelationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BatchGetRelationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BatchGetRelationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BatchGetRelationRequestMultiError, or nil if none found.
func (m *BatchGetRelationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BatchGetRelationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := BatchGetRelationRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetUids()) < 1 {
		err := BatchGetRelationRequestValidationError{
			field:  "Uids",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return BatchGetRelationRequestMultiError(errors)
	}

	return nil
}

// BatchGetRelationRequestMultiError is an error wrapping multiple validation
// errors returned by BatchGetRelationRequest.ValidateAll() if the designated
// constraints aren't met.
type BatchGetRelationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BatchGetRelationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BatchGetRelationRequestMultiError) AllErrors() []error { return m }

// BatchGetRelationRequestValidationError is the validation error returned by
// BatchGetRelationRequest.Validate if the designated constraints aren't met.
type BatchGetRelationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetRelationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetRelationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetRelationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetRelationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetRelationRequestValidationError) ErrorName() string {
	return "BatchGetRelationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BatchGetRelationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetRelationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetRelationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetRelationRequestValidationError{}

// Validate checks the field values on BatchGetRelationReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BatchGetRelationReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BatchGetRelationReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BatchGetRelationReplyMultiError, or nil if none found.
func (m *BatchGetRelationReply) ValidateAll() error {
	return m.validate(true)
}

func (m *BatchGetRelationReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Result

	if len(errors) > 0 {
		return BatchGetRelationReplyMultiError(errors)
	}

	return nil
}

// BatchGetRelationReplyMultiError is an error wrapping multiple validation
// errors returned by BatchGetRelationReply.ValidateAll() if the designated
// constraints aren't met.
type BatchGetRelationReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BatchGetRelationReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BatchGetRelationReplyMultiError) AllErrors() []error { return m }

// BatchGetRelationReplyValidationError is the validation error returned by
// BatchGetRelationReply.Validate if the designated constraints aren't met.
type BatchGetRelationReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetRelationReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetRelationReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetRelationReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetRelationReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetRelationReplyValidationError) ErrorName() string {
	return "BatchGetRelationReplyValidationError"
}

// Error satisfies the builtin error interface
func (e BatchGetRelationReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetRelationReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetRelationReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetRelationReplyValidationError{}