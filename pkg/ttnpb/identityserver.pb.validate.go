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

// ValidateFields checks the field values on AuthInfoResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthInfoResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = AuthInfoResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "universal_rights":

			if v, ok := interface{}(m.GetUniversalRights()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return AuthInfoResponseValidationError{
						field:  "universal_rights",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "access_method":
			if len(subs) == 0 {
				subs = []string{
					"api_key", "oauth_access_token",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "api_key":

					if v, ok := interface{}(m.GetAPIKey()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return AuthInfoResponseValidationError{
								field:  "api_key",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "oauth_access_token":

					if v, ok := interface{}(m.GetOAuthAccessToken()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return AuthInfoResponseValidationError{
								field:  "oauth_access_token",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return AuthInfoResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// AuthInfoResponseValidationError is the validation error returned by
// AuthInfoResponse.ValidateFields if the designated constraints aren't met.
type AuthInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthInfoResponseValidationError) ErrorName() string { return "AuthInfoResponseValidationError" }

// Error satisfies the builtin error interface
func (e AuthInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthInfoResponseValidationError{}

// ValidateFields checks the field values on AuthInfoResponse_APIKeyAccess with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *AuthInfoResponse_APIKeyAccess) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = AuthInfoResponse_APIKeyAccessFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "api_key":

			if v, ok := interface{}(&m.APIKey).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return AuthInfoResponse_APIKeyAccessValidationError{
						field:  "api_key",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "entity_ids":

			if v, ok := interface{}(&m.EntityIDs).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return AuthInfoResponse_APIKeyAccessValidationError{
						field:  "entity_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return AuthInfoResponse_APIKeyAccessValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// AuthInfoResponse_APIKeyAccessValidationError is the validation error
// returned by AuthInfoResponse_APIKeyAccess.ValidateFields if the designated
// constraints aren't met.
type AuthInfoResponse_APIKeyAccessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthInfoResponse_APIKeyAccessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthInfoResponse_APIKeyAccessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthInfoResponse_APIKeyAccessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthInfoResponse_APIKeyAccessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthInfoResponse_APIKeyAccessValidationError) ErrorName() string {
	return "AuthInfoResponse_APIKeyAccessValidationError"
}

// Error satisfies the builtin error interface
func (e AuthInfoResponse_APIKeyAccessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthInfoResponse_APIKeyAccess.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthInfoResponse_APIKeyAccessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthInfoResponse_APIKeyAccessValidationError{}
