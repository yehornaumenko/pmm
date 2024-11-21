// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: management/v1/external.proto

package managementv1

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

// Validate checks the field values on AddExternalServiceParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddExternalServiceParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddExternalServiceParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddExternalServiceParamsMultiError, or nil if none found.
func (m *AddExternalServiceParams) ValidateAll() error {
	return m.validate(true)
}

func (m *AddExternalServiceParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RunsOnNodeId

	// no validation rules for NodeName

	if all {
		switch v := interface{}(m.GetAddNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddExternalServiceParamsValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddExternalServiceParamsValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddExternalServiceParamsValidationError{
				field:  "AddNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Address

	if utf8.RuneCountInString(m.GetServiceName()) < 1 {
		err := AddExternalServiceParamsValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Scheme

	// no validation rules for MetricsPath

	if val := m.GetListenPort(); val <= 0 || val >= 65536 {
		err := AddExternalServiceParamsValidationError{
			field:  "ListenPort",
			reason: "value must be inside range (0, 65536)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NodeId

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	// no validation rules for CustomLabels

	// no validation rules for Group

	// no validation rules for MetricsMode

	// no validation rules for SkipConnectionCheck

	if len(errors) > 0 {
		return AddExternalServiceParamsMultiError(errors)
	}

	return nil
}

// AddExternalServiceParamsMultiError is an error wrapping multiple validation
// errors returned by AddExternalServiceParams.ValidateAll() if the designated
// constraints aren't met.
type AddExternalServiceParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddExternalServiceParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddExternalServiceParamsMultiError) AllErrors() []error { return m }

// AddExternalServiceParamsValidationError is the validation error returned by
// AddExternalServiceParams.Validate if the designated constraints aren't met.
type AddExternalServiceParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddExternalServiceParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddExternalServiceParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddExternalServiceParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddExternalServiceParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddExternalServiceParamsValidationError) ErrorName() string {
	return "AddExternalServiceParamsValidationError"
}

// Error satisfies the builtin error interface
func (e AddExternalServiceParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddExternalServiceParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddExternalServiceParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddExternalServiceParamsValidationError{}

// Validate checks the field values on ExternalServiceResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExternalServiceResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExternalServiceResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExternalServiceResultMultiError, or nil if none found.
func (m *ExternalServiceResult) ValidateAll() error {
	return m.validate(true)
}

func (m *ExternalServiceResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExternalServiceResultValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExternalServiceResultValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExternalServiceResultValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExternalExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExternalServiceResultValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExternalServiceResultValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExternalExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExternalServiceResultValidationError{
				field:  "ExternalExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ExternalServiceResultMultiError(errors)
	}

	return nil
}

// ExternalServiceResultMultiError is an error wrapping multiple validation
// errors returned by ExternalServiceResult.ValidateAll() if the designated
// constraints aren't met.
type ExternalServiceResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExternalServiceResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExternalServiceResultMultiError) AllErrors() []error { return m }

// ExternalServiceResultValidationError is the validation error returned by
// ExternalServiceResult.Validate if the designated constraints aren't met.
type ExternalServiceResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExternalServiceResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExternalServiceResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExternalServiceResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExternalServiceResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExternalServiceResultValidationError) ErrorName() string {
	return "ExternalServiceResultValidationError"
}

// Error satisfies the builtin error interface
func (e ExternalServiceResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExternalServiceResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExternalServiceResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExternalServiceResultValidationError{}