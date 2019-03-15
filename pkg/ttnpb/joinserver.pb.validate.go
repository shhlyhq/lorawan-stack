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

// ValidateFields checks the field values on SessionKeyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SessionKeyRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SessionKeyRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "session_key_id":
			// no validation rules for SessionKeyID
		case "dev_eui":
			// no validation rules for DevEUI
		default:
			return SessionKeyRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SessionKeyRequestValidationError is the validation error returned by
// SessionKeyRequest.ValidateFields if the designated constraints aren't met.
type SessionKeyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SessionKeyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SessionKeyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SessionKeyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SessionKeyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SessionKeyRequestValidationError) ErrorName() string {
	return "SessionKeyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SessionKeyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSessionKeyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SessionKeyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SessionKeyRequestValidationError{}

// ValidateFields checks the field values on NwkSKeysResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NwkSKeysResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = NwkSKeysResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "f_nwk_s_int_key":

			if v, ok := interface{}(&m.FNwkSIntKey).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NwkSKeysResponseValidationError{
						field:  "f_nwk_s_int_key",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "s_nwk_s_int_key":

			if v, ok := interface{}(&m.SNwkSIntKey).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NwkSKeysResponseValidationError{
						field:  "s_nwk_s_int_key",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "nwk_s_enc_key":

			if v, ok := interface{}(&m.NwkSEncKey).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return NwkSKeysResponseValidationError{
						field:  "nwk_s_enc_key",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return NwkSKeysResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// NwkSKeysResponseValidationError is the validation error returned by
// NwkSKeysResponse.ValidateFields if the designated constraints aren't met.
type NwkSKeysResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NwkSKeysResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NwkSKeysResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NwkSKeysResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NwkSKeysResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NwkSKeysResponseValidationError) ErrorName() string { return "NwkSKeysResponseValidationError" }

// Error satisfies the builtin error interface
func (e NwkSKeysResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNwkSKeysResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NwkSKeysResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NwkSKeysResponseValidationError{}

// ValidateFields checks the field values on AppSKeyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AppSKeyResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = AppSKeyResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "app_s_key":

			if v, ok := interface{}(&m.AppSKey).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return AppSKeyResponseValidationError{
						field:  "app_s_key",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return AppSKeyResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// AppSKeyResponseValidationError is the validation error returned by
// AppSKeyResponse.ValidateFields if the designated constraints aren't met.
type AppSKeyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppSKeyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppSKeyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppSKeyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppSKeyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppSKeyResponseValidationError) ErrorName() string { return "AppSKeyResponseValidationError" }

// Error satisfies the builtin error interface
func (e AppSKeyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppSKeyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppSKeyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppSKeyResponseValidationError{}

// ValidateFields checks the field values on CryptoServicePayloadRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *CryptoServicePayloadRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = CryptoServicePayloadRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.EndDeviceIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CryptoServicePayloadRequestValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "lorawan_version":

			if _, ok := MACVersion_name[int32(m.GetLoRaWANVersion())]; !ok {
				return CryptoServicePayloadRequestValidationError{
					field:  "lorawan_version",
					reason: "value must be one of the defined enum values",
				}
			}

		case "payload":
			// no validation rules for Payload
		case "provisioner_id":

			if utf8.RuneCountInString(m.GetProvisionerID()) > 36 {
				return CryptoServicePayloadRequestValidationError{
					field:  "provisioner_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_CryptoServicePayloadRequest_ProvisionerID_Pattern.MatchString(m.GetProvisionerID()) {
				return CryptoServicePayloadRequestValidationError{
					field:  "provisioner_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "provisioning_data":

			if v, ok := interface{}(m.GetProvisioningData()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return CryptoServicePayloadRequestValidationError{
						field:  "provisioning_data",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return CryptoServicePayloadRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// CryptoServicePayloadRequestValidationError is the validation error returned
// by CryptoServicePayloadRequest.ValidateFields if the designated constraints
// aren't met.
type CryptoServicePayloadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CryptoServicePayloadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CryptoServicePayloadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CryptoServicePayloadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CryptoServicePayloadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CryptoServicePayloadRequestValidationError) ErrorName() string {
	return "CryptoServicePayloadRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CryptoServicePayloadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCryptoServicePayloadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CryptoServicePayloadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CryptoServicePayloadRequestValidationError{}

var _CryptoServicePayloadRequest_ProvisionerID_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")

// ValidateFields checks the field values on CryptoServicePayloadResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *CryptoServicePayloadResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = CryptoServicePayloadResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "payload":
			// no validation rules for Payload
		default:
			return CryptoServicePayloadResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// CryptoServicePayloadResponseValidationError is the validation error returned
// by CryptoServicePayloadResponse.ValidateFields if the designated
// constraints aren't met.
type CryptoServicePayloadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CryptoServicePayloadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CryptoServicePayloadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CryptoServicePayloadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CryptoServicePayloadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CryptoServicePayloadResponseValidationError) ErrorName() string {
	return "CryptoServicePayloadResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CryptoServicePayloadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCryptoServicePayloadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CryptoServicePayloadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CryptoServicePayloadResponseValidationError{}

// ValidateFields checks the field values on JoinAcceptMICRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *JoinAcceptMICRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = JoinAcceptMICRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "payload_request":

			if v, ok := interface{}(&m.CryptoServicePayloadRequest).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return JoinAcceptMICRequestValidationError{
						field:  "payload_request",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "join_request_type":

			if _, ok := RejoinType_name[int32(m.GetJoinRequestType())]; !ok {
				return JoinAcceptMICRequestValidationError{
					field:  "join_request_type",
					reason: "value must be one of the defined enum values",
				}
			}

		case "dev_nonce":
			// no validation rules for DevNonce
		default:
			return JoinAcceptMICRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// JoinAcceptMICRequestValidationError is the validation error returned by
// JoinAcceptMICRequest.ValidateFields if the designated constraints aren't met.
type JoinAcceptMICRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinAcceptMICRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinAcceptMICRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinAcceptMICRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinAcceptMICRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinAcceptMICRequestValidationError) ErrorName() string {
	return "JoinAcceptMICRequestValidationError"
}

// Error satisfies the builtin error interface
func (e JoinAcceptMICRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinAcceptMICRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinAcceptMICRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinAcceptMICRequestValidationError{}

// ValidateFields checks the field values on DeriveSessionKeysRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeriveSessionKeysRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = DeriveSessionKeysRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.EndDeviceIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DeriveSessionKeysRequestValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "lorawan_version":

			if _, ok := MACVersion_name[int32(m.GetLoRaWANVersion())]; !ok {
				return DeriveSessionKeysRequestValidationError{
					field:  "lorawan_version",
					reason: "value must be one of the defined enum values",
				}
			}

		case "join_nonce":
			// no validation rules for JoinNonce
		case "dev_nonce":
			// no validation rules for DevNonce
		case "net_id":
			// no validation rules for NetID
		case "provisioner_id":

			if utf8.RuneCountInString(m.GetProvisionerID()) > 36 {
				return DeriveSessionKeysRequestValidationError{
					field:  "provisioner_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_DeriveSessionKeysRequest_ProvisionerID_Pattern.MatchString(m.GetProvisionerID()) {
				return DeriveSessionKeysRequestValidationError{
					field:  "provisioner_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "provisioning_data":

			if v, ok := interface{}(m.GetProvisioningData()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return DeriveSessionKeysRequestValidationError{
						field:  "provisioning_data",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return DeriveSessionKeysRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// DeriveSessionKeysRequestValidationError is the validation error returned by
// DeriveSessionKeysRequest.ValidateFields if the designated constraints
// aren't met.
type DeriveSessionKeysRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeriveSessionKeysRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeriveSessionKeysRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeriveSessionKeysRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeriveSessionKeysRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeriveSessionKeysRequestValidationError) ErrorName() string {
	return "DeriveSessionKeysRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeriveSessionKeysRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeriveSessionKeysRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeriveSessionKeysRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeriveSessionKeysRequestValidationError{}

var _DeriveSessionKeysRequest_ProvisionerID_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")

// ValidateFields checks the field values on GetRootKeysRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetRootKeysRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GetRootKeysRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "ids":

			if v, ok := interface{}(&m.EndDeviceIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetRootKeysRequestValidationError{
						field:  "ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "provisioner_id":

			if utf8.RuneCountInString(m.GetProvisionerID()) > 36 {
				return GetRootKeysRequestValidationError{
					field:  "provisioner_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_GetRootKeysRequest_ProvisionerID_Pattern.MatchString(m.GetProvisionerID()) {
				return GetRootKeysRequestValidationError{
					field:  "provisioner_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$\"",
				}
			}

		case "provisioning_data":

			if v, ok := interface{}(m.GetProvisioningData()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GetRootKeysRequestValidationError{
						field:  "provisioning_data",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GetRootKeysRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GetRootKeysRequestValidationError is the validation error returned by
// GetRootKeysRequest.ValidateFields if the designated constraints aren't met.
type GetRootKeysRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRootKeysRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRootKeysRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRootKeysRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRootKeysRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRootKeysRequestValidationError) ErrorName() string {
	return "GetRootKeysRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetRootKeysRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRootKeysRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRootKeysRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRootKeysRequestValidationError{}

var _GetRootKeysRequest_ProvisionerID_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$|^$")

// ValidateFields checks the field values on ProvisionEndDevicesRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ProvisionEndDevicesRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ProvisionEndDevicesRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "application_ids":

			if v, ok := interface{}(&m.ApplicationIdentifiers).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ProvisionEndDevicesRequestValidationError{
						field:  "application_ids",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "provisioner_id":

			if utf8.RuneCountInString(m.GetProvisionerID()) > 36 {
				return ProvisionEndDevicesRequestValidationError{
					field:  "provisioner_id",
					reason: "value length must be at most 36 runes",
				}
			}

			if !_ProvisionEndDevicesRequest_ProvisionerID_Pattern.MatchString(m.GetProvisionerID()) {
				return ProvisionEndDevicesRequestValidationError{
					field:  "provisioner_id",
					reason: "value does not match regex pattern \"^[a-z0-9](?:[-]?[a-z0-9]){2,}$\"",
				}
			}

		case "provisioning_data":
			// no validation rules for ProvisioningData
		case "end_devices":
			if len(subs) == 0 {
				subs = []string{
					"list", "range", "from_data",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "list":

					if v, ok := interface{}(m.GetList()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return ProvisionEndDevicesRequestValidationError{
								field:  "list",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "range":

					if v, ok := interface{}(m.GetRange()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return ProvisionEndDevicesRequestValidationError{
								field:  "range",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "from_data":

					if v, ok := interface{}(m.GetFromData()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return ProvisionEndDevicesRequestValidationError{
								field:  "from_data",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return ProvisionEndDevicesRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ProvisionEndDevicesRequestValidationError is the validation error returned
// by ProvisionEndDevicesRequest.ValidateFields if the designated constraints
// aren't met.
type ProvisionEndDevicesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProvisionEndDevicesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProvisionEndDevicesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProvisionEndDevicesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProvisionEndDevicesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProvisionEndDevicesRequestValidationError) ErrorName() string {
	return "ProvisionEndDevicesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProvisionEndDevicesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProvisionEndDevicesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProvisionEndDevicesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProvisionEndDevicesRequestValidationError{}

var _ProvisionEndDevicesRequest_ProvisionerID_Pattern = regexp.MustCompile("^[a-z0-9](?:[-]?[a-z0-9]){2,}$")

// ValidateFields checks the field values on
// ProvisionEndDevicesRequest_IdentifiersList with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ProvisionEndDevicesRequest_IdentifiersList) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ProvisionEndDevicesRequest_IdentifiersListFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "join_eui":
			// no validation rules for JoinEUI
		case "end_device_ids":

			for idx, item := range m.EndDeviceIDs {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ProvisionEndDevicesRequest_IdentifiersListValidationError{
							field:  fmt.Sprintf("end_device_ids[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return ProvisionEndDevicesRequest_IdentifiersListValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ProvisionEndDevicesRequest_IdentifiersListValidationError is the validation
// error returned by ProvisionEndDevicesRequest_IdentifiersList.ValidateFields
// if the designated constraints aren't met.
type ProvisionEndDevicesRequest_IdentifiersListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProvisionEndDevicesRequest_IdentifiersListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProvisionEndDevicesRequest_IdentifiersListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProvisionEndDevicesRequest_IdentifiersListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProvisionEndDevicesRequest_IdentifiersListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProvisionEndDevicesRequest_IdentifiersListValidationError) ErrorName() string {
	return "ProvisionEndDevicesRequest_IdentifiersListValidationError"
}

// Error satisfies the builtin error interface
func (e ProvisionEndDevicesRequest_IdentifiersListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProvisionEndDevicesRequest_IdentifiersList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProvisionEndDevicesRequest_IdentifiersListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProvisionEndDevicesRequest_IdentifiersListValidationError{}

// ValidateFields checks the field values on
// ProvisionEndDevicesRequest_IdentifiersRange with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ProvisionEndDevicesRequest_IdentifiersRange) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ProvisionEndDevicesRequest_IdentifiersRangeFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "join_eui":
			// no validation rules for JoinEUI
		case "start_dev_eui":
			// no validation rules for StartDevEUI
		default:
			return ProvisionEndDevicesRequest_IdentifiersRangeValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ProvisionEndDevicesRequest_IdentifiersRangeValidationError is the validation
// error returned by
// ProvisionEndDevicesRequest_IdentifiersRange.ValidateFields if the
// designated constraints aren't met.
type ProvisionEndDevicesRequest_IdentifiersRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProvisionEndDevicesRequest_IdentifiersRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProvisionEndDevicesRequest_IdentifiersRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProvisionEndDevicesRequest_IdentifiersRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProvisionEndDevicesRequest_IdentifiersRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProvisionEndDevicesRequest_IdentifiersRangeValidationError) ErrorName() string {
	return "ProvisionEndDevicesRequest_IdentifiersRangeValidationError"
}

// Error satisfies the builtin error interface
func (e ProvisionEndDevicesRequest_IdentifiersRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProvisionEndDevicesRequest_IdentifiersRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProvisionEndDevicesRequest_IdentifiersRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProvisionEndDevicesRequest_IdentifiersRangeValidationError{}

// ValidateFields checks the field values on
// ProvisionEndDevicesRequest_IdentifiersFromData with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ProvisionEndDevicesRequest_IdentifiersFromData) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ProvisionEndDevicesRequest_IdentifiersFromDataFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "join_eui":
			// no validation rules for JoinEUI
		default:
			return ProvisionEndDevicesRequest_IdentifiersFromDataValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ProvisionEndDevicesRequest_IdentifiersFromDataValidationError is the
// validation error returned by
// ProvisionEndDevicesRequest_IdentifiersFromData.ValidateFields if the
// designated constraints aren't met.
type ProvisionEndDevicesRequest_IdentifiersFromDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProvisionEndDevicesRequest_IdentifiersFromDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProvisionEndDevicesRequest_IdentifiersFromDataValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e ProvisionEndDevicesRequest_IdentifiersFromDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProvisionEndDevicesRequest_IdentifiersFromDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProvisionEndDevicesRequest_IdentifiersFromDataValidationError) ErrorName() string {
	return "ProvisionEndDevicesRequest_IdentifiersFromDataValidationError"
}

// Error satisfies the builtin error interface
func (e ProvisionEndDevicesRequest_IdentifiersFromDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProvisionEndDevicesRequest_IdentifiersFromData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProvisionEndDevicesRequest_IdentifiersFromDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProvisionEndDevicesRequest_IdentifiersFromDataValidationError{}
