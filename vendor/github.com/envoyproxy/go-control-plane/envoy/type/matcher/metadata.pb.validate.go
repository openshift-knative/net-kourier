// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/metadata.proto

package matcher

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

// Validate checks the field values on MetadataMatcher with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MetadataMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetadataMatcher with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetadataMatcherMultiError, or nil if none found.
func (m *MetadataMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *MetadataMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetFilter()) < 1 {
		err := MetadataMatcherValidationError{
			field:  "Filter",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetPath()) < 1 {
		err := MetadataMatcherValidationError{
			field:  "Path",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetPath() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MetadataMatcherValidationError{
						field:  fmt.Sprintf("Path[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MetadataMatcherValidationError{
						field:  fmt.Sprintf("Path[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetadataMatcherValidationError{
					field:  fmt.Sprintf("Path[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetValue() == nil {
		err := MetadataMatcherValidationError{
			field:  "Value",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetadataMatcherValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetadataMatcherValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataMatcherValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MetadataMatcherMultiError(errors)
	}

	return nil
}

// MetadataMatcherMultiError is an error wrapping multiple validation errors
// returned by MetadataMatcher.ValidateAll() if the designated constraints
// aren't met.
type MetadataMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetadataMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetadataMatcherMultiError) AllErrors() []error { return m }

// MetadataMatcherValidationError is the validation error returned by
// MetadataMatcher.Validate if the designated constraints aren't met.
type MetadataMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataMatcherValidationError) ErrorName() string { return "MetadataMatcherValidationError" }

// Error satisfies the builtin error interface
func (e MetadataMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadataMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataMatcherValidationError{}

// Validate checks the field values on MetadataMatcher_PathSegment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *MetadataMatcher_PathSegment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MetadataMatcher_PathSegment with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MetadataMatcher_PathSegmentMultiError, or nil if none found.
func (m *MetadataMatcher_PathSegment) ValidateAll() error {
	return m.validate(true)
}

func (m *MetadataMatcher_PathSegment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofSegmentPresent := false
	switch v := m.Segment.(type) {
	case *MetadataMatcher_PathSegment_Key:
		if v == nil {
			err := MetadataMatcher_PathSegmentValidationError{
				field:  "Segment",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofSegmentPresent = true

		if utf8.RuneCountInString(m.GetKey()) < 1 {
			err := MetadataMatcher_PathSegmentValidationError{
				field:  "Key",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofSegmentPresent {
		err := MetadataMatcher_PathSegmentValidationError{
			field:  "Segment",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MetadataMatcher_PathSegmentMultiError(errors)
	}

	return nil
}

// MetadataMatcher_PathSegmentMultiError is an error wrapping multiple
// validation errors returned by MetadataMatcher_PathSegment.ValidateAll() if
// the designated constraints aren't met.
type MetadataMatcher_PathSegmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetadataMatcher_PathSegmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetadataMatcher_PathSegmentMultiError) AllErrors() []error { return m }

// MetadataMatcher_PathSegmentValidationError is the validation error returned
// by MetadataMatcher_PathSegment.Validate if the designated constraints
// aren't met.
type MetadataMatcher_PathSegmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataMatcher_PathSegmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataMatcher_PathSegmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataMatcher_PathSegmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataMatcher_PathSegmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataMatcher_PathSegmentValidationError) ErrorName() string {
	return "MetadataMatcher_PathSegmentValidationError"
}

// Error satisfies the builtin error interface
func (e MetadataMatcher_PathSegmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadataMatcher_PathSegment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataMatcher_PathSegmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataMatcher_PathSegmentValidationError{}
