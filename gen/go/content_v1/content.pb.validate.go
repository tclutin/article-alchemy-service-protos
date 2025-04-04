// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: content_v1/content.proto

package content_v1

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

// Validate checks the field values on ExtractContentPreviewRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExtractContentPreviewRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtractContentPreviewRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExtractContentPreviewRequestMultiError, or nil if none found.
func (m *ExtractContentPreviewRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtractContentPreviewRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUrl()); l < 1 || l > 2083 {
		err := ExtractContentPreviewRequestValidationError{
			field:  "Url",
			reason: "value length must be between 1 and 2083 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		err = ExtractContentPreviewRequestValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := ExtractContentPreviewRequestValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetEmail()); l < 5 || l > 255 {
		err := ExtractContentPreviewRequestValidationError{
			field:  "Email",
			reason: "value length must be between 5 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = ExtractContentPreviewRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ExtractContentPreviewRequestMultiError(errors)
	}

	return nil
}

func (m *ExtractContentPreviewRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *ExtractContentPreviewRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// ExtractContentPreviewRequestMultiError is an error wrapping multiple
// validation errors returned by ExtractContentPreviewRequest.ValidateAll() if
// the designated constraints aren't met.
type ExtractContentPreviewRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtractContentPreviewRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtractContentPreviewRequestMultiError) AllErrors() []error { return m }

// ExtractContentPreviewRequestValidationError is the validation error returned
// by ExtractContentPreviewRequest.Validate if the designated constraints
// aren't met.
type ExtractContentPreviewRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtractContentPreviewRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtractContentPreviewRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtractContentPreviewRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtractContentPreviewRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtractContentPreviewRequestValidationError) ErrorName() string {
	return "ExtractContentPreviewRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExtractContentPreviewRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtractContentPreviewRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtractContentPreviewRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtractContentPreviewRequestValidationError{}

// Validate checks the field values on ExtractContentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExtractContentRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtractContentRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExtractContentRequestMultiError, or nil if none found.
func (m *ExtractContentRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtractContentRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUrl()); l < 1 || l > 2083 {
		err := ExtractContentRequestValidationError{
			field:  "Url",
			reason: "value length must be between 1 and 2083 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		err = ExtractContentRequestValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := ExtractContentRequestValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ExtractContentRequestMultiError(errors)
	}

	return nil
}

// ExtractContentRequestMultiError is an error wrapping multiple validation
// errors returned by ExtractContentRequest.ValidateAll() if the designated
// constraints aren't met.
type ExtractContentRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtractContentRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtractContentRequestMultiError) AllErrors() []error { return m }

// ExtractContentRequestValidationError is the validation error returned by
// ExtractContentRequest.Validate if the designated constraints aren't met.
type ExtractContentRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtractContentRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtractContentRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtractContentRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtractContentRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtractContentRequestValidationError) ErrorName() string {
	return "ExtractContentRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExtractContentRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtractContentRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtractContentRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtractContentRequestValidationError{}

// Validate checks the field values on ExtractContentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExtractContentResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtractContentResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExtractContentResponseMultiError, or nil if none found.
func (m *ExtractContentResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtractContentResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ContentId

	if len(errors) > 0 {
		return ExtractContentResponseMultiError(errors)
	}

	return nil
}

// ExtractContentResponseMultiError is an error wrapping multiple validation
// errors returned by ExtractContentResponse.ValidateAll() if the designated
// constraints aren't met.
type ExtractContentResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtractContentResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtractContentResponseMultiError) AllErrors() []error { return m }

// ExtractContentResponseValidationError is the validation error returned by
// ExtractContentResponse.Validate if the designated constraints aren't met.
type ExtractContentResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtractContentResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtractContentResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtractContentResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtractContentResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtractContentResponseValidationError) ErrorName() string {
	return "ExtractContentResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ExtractContentResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtractContentResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtractContentResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtractContentResponseValidationError{}

// Validate checks the field values on GetHistoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetHistoryResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHistoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHistoryResponseMultiError, or nil if none found.
func (m *GetHistoryResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHistoryResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetHistoryResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetHistoryResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetHistoryResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetHistoryResponseMultiError(errors)
	}

	return nil
}

// GetHistoryResponseMultiError is an error wrapping multiple validation errors
// returned by GetHistoryResponse.ValidateAll() if the designated constraints
// aren't met.
type GetHistoryResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHistoryResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHistoryResponseMultiError) AllErrors() []error { return m }

// GetHistoryResponseValidationError is the validation error returned by
// GetHistoryResponse.Validate if the designated constraints aren't met.
type GetHistoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoryResponseValidationError) ErrorName() string {
	return "GetHistoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetHistoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoryResponseValidationError{}

// Validate checks the field values on GetHistoryItem with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetHistoryItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHistoryItem with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetHistoryItemMultiError,
// or nil if none found.
func (m *GetHistoryItem) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHistoryItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ContentId

	// no validation rules for Status

	// no validation rules for Url

	// no validation rules for Data

	// no validation rules for Error

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetHistoryItemValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetHistoryItemValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetHistoryItemValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetHistoryItemMultiError(errors)
	}

	return nil
}

// GetHistoryItemMultiError is an error wrapping multiple validation errors
// returned by GetHistoryItem.ValidateAll() if the designated constraints
// aren't met.
type GetHistoryItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHistoryItemMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHistoryItemMultiError) AllErrors() []error { return m }

// GetHistoryItemValidationError is the validation error returned by
// GetHistoryItem.Validate if the designated constraints aren't met.
type GetHistoryItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoryItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoryItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoryItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoryItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoryItemValidationError) ErrorName() string { return "GetHistoryItemValidationError" }

// Error satisfies the builtin error interface
func (e GetHistoryItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoryItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoryItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoryItemValidationError{}
