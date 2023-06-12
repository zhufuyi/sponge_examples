// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/creation/v1/like.proto

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

// Validate checks the field values on LikeInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LikeInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LikeInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LikeInfoMultiError, or nil
// if none found.
func (m *LikeInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *LikeInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for ObjType

	// no validation rules for ObjId

	// no validation rules for Status

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return LikeInfoMultiError(errors)
	}

	return nil
}

// LikeInfoMultiError is an error wrapping multiple validation errors returned
// by LikeInfo.ValidateAll() if the designated constraints aren't met.
type LikeInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LikeInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LikeInfoMultiError) AllErrors() []error { return m }

// LikeInfoValidationError is the validation error returned by
// LikeInfo.Validate if the designated constraints aren't met.
type LikeInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LikeInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LikeInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LikeInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LikeInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LikeInfoValidationError) ErrorName() string { return "LikeInfoValidationError" }

// Error satisfies the builtin error interface
func (e LikeInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLikeInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LikeInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LikeInfoValidationError{}

// Validate checks the field values on CreateLikeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateLikeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLikeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLikeRequestMultiError, or nil if none found.
func (m *CreateLikeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLikeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := CreateLikeRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetObjType(); val < 1 || val > 2 {
		err := CreateLikeRequestValidationError{
			field:  "ObjType",
			reason: "value must be inside range [1, 2]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetObjId() < 1 {
		err := CreateLikeRequestValidationError{
			field:  "ObjId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateLikeRequestMultiError(errors)
	}

	return nil
}

// CreateLikeRequestMultiError is an error wrapping multiple validation errors
// returned by CreateLikeRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateLikeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLikeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLikeRequestMultiError) AllErrors() []error { return m }

// CreateLikeRequestValidationError is the validation error returned by
// CreateLikeRequest.Validate if the designated constraints aren't met.
type CreateLikeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLikeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLikeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLikeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLikeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLikeRequestValidationError) ErrorName() string {
	return "CreateLikeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLikeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLikeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLikeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLikeRequestValidationError{}

// Validate checks the field values on CreateLikeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateLikeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLikeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLikeReplyMultiError, or nil if none found.
func (m *CreateLikeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLikeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateLikeReplyMultiError(errors)
	}

	return nil
}

// CreateLikeReplyMultiError is an error wrapping multiple validation errors
// returned by CreateLikeReply.ValidateAll() if the designated constraints
// aren't met.
type CreateLikeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLikeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLikeReplyMultiError) AllErrors() []error { return m }

// CreateLikeReplyValidationError is the validation error returned by
// CreateLikeReply.Validate if the designated constraints aren't met.
type CreateLikeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLikeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLikeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLikeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLikeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLikeReplyValidationError) ErrorName() string { return "CreateLikeReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateLikeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLikeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLikeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLikeReplyValidationError{}

// Validate checks the field values on DeleteLikeRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteLikeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteLikeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteLikeRequestMultiError, or nil if none found.
func (m *DeleteLikeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteLikeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 1 {
		err := DeleteLikeRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetObjType(); val < 1 || val > 2 {
		err := DeleteLikeRequestValidationError{
			field:  "ObjType",
			reason: "value must be inside range [1, 2]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetObjId() < 1 {
		err := DeleteLikeRequestValidationError{
			field:  "ObjId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteLikeRequestMultiError(errors)
	}

	return nil
}

// DeleteLikeRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteLikeRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteLikeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteLikeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteLikeRequestMultiError) AllErrors() []error { return m }

// DeleteLikeRequestValidationError is the validation error returned by
// DeleteLikeRequest.Validate if the designated constraints aren't met.
type DeleteLikeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteLikeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteLikeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteLikeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteLikeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteLikeRequestValidationError) ErrorName() string {
	return "DeleteLikeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteLikeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteLikeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteLikeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteLikeRequestValidationError{}

// Validate checks the field values on DeleteLikeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteLikeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteLikeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteLikeReplyMultiError, or nil if none found.
func (m *DeleteLikeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteLikeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteLikeReplyMultiError(errors)
	}

	return nil
}

// DeleteLikeReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteLikeReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteLikeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteLikeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteLikeReplyMultiError) AllErrors() []error { return m }

// DeleteLikeReplyValidationError is the validation error returned by
// DeleteLikeReply.Validate if the designated constraints aren't met.
type DeleteLikeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteLikeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteLikeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteLikeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteLikeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteLikeReplyValidationError) ErrorName() string { return "DeleteLikeReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteLikeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteLikeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteLikeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteLikeReplyValidationError{}

// Validate checks the field values on ListPostLikeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPostLikeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPostLikeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPostLikeRequestMultiError, or nil if none found.
func (m *ListPostLikeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPostLikeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPostId() < 1 {
		err := ListPostLikeRequestValidationError{
			field:  "PostId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() < 0 {
		err := ListPostLikeRequestValidationError{
			field:  "Page",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLimit(); val <= 0 || val > 100 {
		err := ListPostLikeRequestValidationError{
			field:  "Limit",
			reason: "value must be inside range (0, 100]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListPostLikeRequestMultiError(errors)
	}

	return nil
}

// ListPostLikeRequestMultiError is an error wrapping multiple validation
// errors returned by ListPostLikeRequest.ValidateAll() if the designated
// constraints aren't met.
type ListPostLikeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPostLikeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPostLikeRequestMultiError) AllErrors() []error { return m }

// ListPostLikeRequestValidationError is the validation error returned by
// ListPostLikeRequest.Validate if the designated constraints aren't met.
type ListPostLikeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostLikeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostLikeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostLikeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostLikeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostLikeRequestValidationError) ErrorName() string {
	return "ListPostLikeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPostLikeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostLikeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostLikeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostLikeRequestValidationError{}

// Validate checks the field values on ListPostLikeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListPostLikeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPostLikeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPostLikeReplyMultiError, or nil if none found.
func (m *ListPostLikeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPostLikeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLikes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListPostLikeReplyValidationError{
						field:  fmt.Sprintf("Likes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListPostLikeReplyValidationError{
						field:  fmt.Sprintf("Likes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListPostLikeReplyValidationError{
					field:  fmt.Sprintf("Likes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListPostLikeReplyMultiError(errors)
	}

	return nil
}

// ListPostLikeReplyMultiError is an error wrapping multiple validation errors
// returned by ListPostLikeReply.ValidateAll() if the designated constraints
// aren't met.
type ListPostLikeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPostLikeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPostLikeReplyMultiError) AllErrors() []error { return m }

// ListPostLikeReplyValidationError is the validation error returned by
// ListPostLikeReply.Validate if the designated constraints aren't met.
type ListPostLikeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPostLikeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPostLikeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPostLikeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPostLikeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPostLikeReplyValidationError) ErrorName() string {
	return "ListPostLikeReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListPostLikeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPostLikeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPostLikeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPostLikeReplyValidationError{}

// Validate checks the field values on ListCommentLikeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCommentLikeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCommentLikeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCommentLikeRequestMultiError, or nil if none found.
func (m *ListCommentLikeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCommentLikeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCommentId() < 1 {
		err := ListCommentLikeRequestValidationError{
			field:  "CommentId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() < 0 {
		err := ListCommentLikeRequestValidationError{
			field:  "Page",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLimit(); val <= 0 || val > 100 {
		err := ListCommentLikeRequestValidationError{
			field:  "Limit",
			reason: "value must be inside range (0, 100]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListCommentLikeRequestMultiError(errors)
	}

	return nil
}

// ListCommentLikeRequestMultiError is an error wrapping multiple validation
// errors returned by ListCommentLikeRequest.ValidateAll() if the designated
// constraints aren't met.
type ListCommentLikeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCommentLikeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCommentLikeRequestMultiError) AllErrors() []error { return m }

// ListCommentLikeRequestValidationError is the validation error returned by
// ListCommentLikeRequest.Validate if the designated constraints aren't met.
type ListCommentLikeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCommentLikeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCommentLikeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCommentLikeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCommentLikeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCommentLikeRequestValidationError) ErrorName() string {
	return "ListCommentLikeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCommentLikeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCommentLikeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCommentLikeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCommentLikeRequestValidationError{}

// Validate checks the field values on ListCommentLikeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCommentLikeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCommentLikeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCommentLikeReplyMultiError, or nil if none found.
func (m *ListCommentLikeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCommentLikeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetLikes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCommentLikeReplyValidationError{
						field:  fmt.Sprintf("Likes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCommentLikeReplyValidationError{
						field:  fmt.Sprintf("Likes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCommentLikeReplyValidationError{
					field:  fmt.Sprintf("Likes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListCommentLikeReplyMultiError(errors)
	}

	return nil
}

// ListCommentLikeReplyMultiError is an error wrapping multiple validation
// errors returned by ListCommentLikeReply.ValidateAll() if the designated
// constraints aren't met.
type ListCommentLikeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCommentLikeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCommentLikeReplyMultiError) AllErrors() []error { return m }

// ListCommentLikeReplyValidationError is the validation error returned by
// ListCommentLikeReply.Validate if the designated constraints aren't met.
type ListCommentLikeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCommentLikeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCommentLikeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCommentLikeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCommentLikeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCommentLikeReplyValidationError) ErrorName() string {
	return "ListCommentLikeReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListCommentLikeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCommentLikeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCommentLikeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCommentLikeReplyValidationError{}
