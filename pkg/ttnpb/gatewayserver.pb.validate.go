// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// ValidateFields checks the field values on GatewayUp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GatewayUp) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GatewayUpFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "uplink_messages":

			for idx, item := range m.GetUplinkMessages() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return GatewayUpValidationError{
							field:  fmt.Sprintf("uplink_messages[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		case "gateway_status":

			if v, ok := interface{}(m.GetGatewayStatus()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GatewayUpValidationError{
						field:  "gateway_status",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "tx_acknowledgment":

			if v, ok := interface{}(m.GetTxAcknowledgment()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GatewayUpValidationError{
						field:  "tx_acknowledgment",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GatewayUpValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GatewayUpValidationError is the validation error returned by
// GatewayUp.ValidateFields if the designated constraints aren't met.
type GatewayUpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GatewayUpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GatewayUpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GatewayUpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GatewayUpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GatewayUpValidationError) ErrorName() string { return "GatewayUpValidationError" }

// Error satisfies the builtin error interface
func (e GatewayUpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGatewayUp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GatewayUpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GatewayUpValidationError{}

// ValidateFields checks the field values on GatewayDown with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GatewayDown) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GatewayDownFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "downlink_message":

			if v, ok := interface{}(m.GetDownlinkMessage()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GatewayDownValidationError{
						field:  "downlink_message",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GatewayDownValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GatewayDownValidationError is the validation error returned by
// GatewayDown.ValidateFields if the designated constraints aren't met.
type GatewayDownValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GatewayDownValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GatewayDownValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GatewayDownValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GatewayDownValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GatewayDownValidationError) ErrorName() string { return "GatewayDownValidationError" }

// Error satisfies the builtin error interface
func (e GatewayDownValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGatewayDown.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GatewayDownValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GatewayDownValidationError{}

// ValidateFields checks the field values on ScheduleDownlinkResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ScheduleDownlinkResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ScheduleDownlinkResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "delay":

			if v, ok := interface{}(&m.Delay).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ScheduleDownlinkResponseValidationError{
						field:  "delay",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return ScheduleDownlinkResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ScheduleDownlinkResponseValidationError is the validation error returned by
// ScheduleDownlinkResponse.ValidateFields if the designated constraints
// aren't met.
type ScheduleDownlinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScheduleDownlinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScheduleDownlinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScheduleDownlinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScheduleDownlinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScheduleDownlinkResponseValidationError) ErrorName() string {
	return "ScheduleDownlinkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ScheduleDownlinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScheduleDownlinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScheduleDownlinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScheduleDownlinkResponseValidationError{}
