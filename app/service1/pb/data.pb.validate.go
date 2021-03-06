// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: data.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _data_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Employee with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Employee) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EmpNo

	if v, ok := interface{}(m.GetBirthDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EmployeeValidationError{
				field:  "BirthDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FirstName

	// no validation rules for LastName

	// no validation rules for Gender

	if v, ok := interface{}(m.GetHireDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EmployeeValidationError{
				field:  "HireDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EmployeeValidationError is the validation error returned by
// Employee.Validate if the designated constraints aren't met.
type EmployeeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmployeeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmployeeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmployeeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmployeeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmployeeValidationError) ErrorName() string { return "EmployeeValidationError" }

// Error satisfies the builtin error interface
func (e EmployeeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmployee.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmployeeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmployeeValidationError{}
