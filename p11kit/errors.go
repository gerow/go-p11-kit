// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by errors_generate.go; DO NOT EDIT.

//go:generate go run errors_generate.go

package p11kit

import "fmt"

// pkcs11Error represents a PKCS #11 return code.
type pkcs11Error uint64

// Error returns the spec name of the error.
func (e pkcs11Error) Error() string {
	if s, ok := errStrings[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown pkcs11 error: 0x%08x", uint64(e))
}

var errStrings = map[pkcs11Error]string{
	errCancel:                        "CKR_CANCEL",
	errHostMemory:                    "CKR_HOST_MEMORY",
	errSlotIDInvalid:                 "CKR_SLOT_ID_INVALID",
	errGeneralError:                  "CKR_GENERAL_ERROR",
	errFunctionFailed:                "CKR_FUNCTION_FAILED",
	errArgumentsBad:                  "CKR_ARGUMENTS_BAD",
	errNoEvent:                       "CKR_NO_EVENT",
	errNeedToCreateThreads:           "CKR_NEED_TO_CREATE_THREADS",
	errCantLock:                      "CKR_CANT_LOCK",
	errAttributeReadOnly:             "CKR_ATTRIBUTE_READ_ONLY",
	errAttributeSensitive:            "CKR_ATTRIBUTE_SENSITIVE",
	errAttributeTypeInvalid:          "CKR_ATTRIBUTE_TYPE_INVALID",
	errAttributeValueInvalid:         "CKR_ATTRIBUTE_VALUE_INVALID",
	errCopyProhibited:                "CKR_COPY_PROHIBITED",
	errActionProhibited:              "CKR_ACTION_PROHIBITED",
	errDataInvalid:                   "CKR_DATA_INVALID",
	errDataLenRange:                  "CKR_DATA_LEN_RANGE",
	errDeviceError:                   "CKR_DEVICE_ERROR",
	errDeviceMemory:                  "CKR_DEVICE_MEMORY",
	errDeviceRemoved:                 "CKR_DEVICE_REMOVED",
	errEncryptedDataInvalid:          "CKR_ENCRYPTED_DATA_INVALID",
	errEncryptedDataLenRange:         "CKR_ENCRYPTED_DATA_LEN_RANGE",
	errFunctionCanceled:              "CKR_FUNCTION_CANCELED",
	errFunctionNotParallel:           "CKR_FUNCTION_NOT_PARALLEL",
	errFunctionNotSupported:          "CKR_FUNCTION_NOT_SUPPORTED",
	errKeyHandleInvalid:              "CKR_KEY_HANDLE_INVALID",
	errKeySizeRange:                  "CKR_KEY_SIZE_RANGE",
	errKeyTypeInconsistent:           "CKR_KEY_TYPE_INCONSISTENT",
	errKeyNotNeeded:                  "CKR_KEY_NOT_NEEDED",
	errKeyChanged:                    "CKR_KEY_CHANGED",
	errKeyNeeded:                     "CKR_KEY_NEEDED",
	errKeyIndigestible:               "CKR_KEY_INDIGESTIBLE",
	errKeyFunctionNotPermitted:       "CKR_KEY_FUNCTION_NOT_PERMITTED",
	errKeyNotWrappable:               "CKR_KEY_NOT_WRAPPABLE",
	errKeyUnextractable:              "CKR_KEY_UNEXTRACTABLE",
	errMechanismInvalid:              "CKR_MECHANISM_INVALID",
	errMechanismParamInvalid:         "CKR_MECHANISM_PARAM_INVALID",
	errObjectHandleInvalid:           "CKR_OBJECT_HANDLE_INVALID",
	errOperationActive:               "CKR_OPERATION_ACTIVE",
	errOperationNotInitialized:       "CKR_OPERATION_NOT_INITIALIZED",
	errPINIncorrect:                  "CKR_PIN_INCORRECT",
	errPINInvalid:                    "CKR_PIN_INVALID",
	errPINLenRange:                   "CKR_PIN_LEN_RANGE",
	errPINExpired:                    "CKR_PIN_EXPIRED",
	errPINLocked:                     "CKR_PIN_LOCKED",
	errSessionClosed:                 "CKR_SESSION_CLOSED",
	errSessionCount:                  "CKR_SESSION_COUNT",
	errSessionHandleInvalid:          "CKR_SESSION_HANDLE_INVALID",
	errSessionParallelNotSupported:   "CKR_SESSION_PARALLEL_NOT_SUPPORTED",
	errSessionReadOnly:               "CKR_SESSION_READ_ONLY",
	errSessionExists:                 "CKR_SESSION_EXISTS",
	errSessionReadOnlyExists:         "CKR_SESSION_READ_ONLY_EXISTS",
	errSessionReadWriteSoExists:      "CKR_SESSION_READ_WRITE_SO_EXISTS",
	errSignatureInvalid:              "CKR_SIGNATURE_INVALID",
	errSignatureLenRange:             "CKR_SIGNATURE_LEN_RANGE",
	errTemplateIncomplete:            "CKR_TEMPLATE_INCOMPLETE",
	errTemplateInconsistent:          "CKR_TEMPLATE_INCONSISTENT",
	errTokenNotPresent:               "CKR_TOKEN_NOT_PRESENT",
	errTokenNotRecognized:            "CKR_TOKEN_NOT_RECOGNIZED",
	errTokenWriteProtected:           "CKR_TOKEN_WRITE_PROTECTED",
	errUnwrappingKeyHandleInvalid:    "CKR_UNWRAPPING_KEY_HANDLE_INVALID",
	errUnwrappingKeySizeRange:        "CKR_UNWRAPPING_KEY_SIZE_RANGE",
	errUnwrappingKeyTypeInconsistent: "CKR_UNWRAPPING_KEY_TYPE_INCONSISTENT",
	errUserAlreadyLoggedIn:           "CKR_USER_ALREADY_LOGGED_IN",
	errUserNotLoggedIn:               "CKR_USER_NOT_LOGGED_IN",
	errUserPINNotInitialized:         "CKR_USER_PIN_NOT_INITIALIZED",
	errUserTypeInvalid:               "CKR_USER_TYPE_INVALID",
	errUserAnotherAlreadyLoggedIn:    "CKR_USER_ANOTHER_ALREADY_LOGGED_IN",
	errUserTooManyTypes:              "CKR_USER_TOO_MANY_TYPES",
	errWrappedKeyInvalid:             "CKR_WRAPPED_KEY_INVALID",
	errWrappedKeyLenRange:            "CKR_WRAPPED_KEY_LEN_RANGE",
	errWrappingKeyHandleInvalid:      "CKR_WRAPPING_KEY_HANDLE_INVALID",
	errWrappingKeySizeRange:          "CKR_WRAPPING_KEY_SIZE_RANGE",
	errWrappingKeyTypeInconsistent:   "CKR_WRAPPING_KEY_TYPE_INCONSISTENT",
	errRandomSeedNotSupported:        "CKR_RANDOM_SEED_NOT_SUPPORTED",
	errRandomNoRNG:                   "CKR_RANDOM_NO_RNG",
	errDomainParamsInvalid:           "CKR_DOMAIN_PARAMS_INVALID",
	errCurveNotSupported:             "CKR_CURVE_NOT_SUPPORTED",
	errBufferTooSmall:                "CKR_BUFFER_TOO_SMALL",
	errSavedStateInvalid:             "CKR_SAVED_STATE_INVALID",
	errInformationSensitive:          "CKR_INFORMATION_SENSITIVE",
	errStateUnsaveable:               "CKR_STATE_UNSAVEABLE",
	errCryptokiNotInitialized:        "CKR_CRYPTOKI_NOT_INITIALIZED",
	errCryptokiAlreadyInitialized:    "CKR_CRYPTOKI_ALREADY_INITIALIZED",
	errMutexBad:                      "CKR_MUTEX_BAD",
	errMutexNotLocked:                "CKR_MUTEX_NOT_LOCKED",
	errFunctionRejected:              "CKR_FUNCTION_REJECTED",
	errVendorDefined:                 "CKR_VENDOR_DEFINED",
}

// Error codes defined PKCS #11.
//
// http://docs.oasis-open.org/pkcs11/pkcs11-base/v2.40/csd02/pkcs11-base-v2.40-csd02.html#_Toc385435538
const (
	errCancel                        pkcs11Error = 0x00000001
	errHostMemory                    pkcs11Error = 0x00000002
	errSlotIDInvalid                 pkcs11Error = 0x00000003
	errGeneralError                  pkcs11Error = 0x00000005
	errFunctionFailed                pkcs11Error = 0x00000006
	errArgumentsBad                  pkcs11Error = 0x00000007
	errNoEvent                       pkcs11Error = 0x00000008
	errNeedToCreateThreads           pkcs11Error = 0x00000009
	errCantLock                      pkcs11Error = 0x0000000a
	errAttributeReadOnly             pkcs11Error = 0x00000010
	errAttributeSensitive            pkcs11Error = 0x00000011
	errAttributeTypeInvalid          pkcs11Error = 0x00000012
	errAttributeValueInvalid         pkcs11Error = 0x00000013
	errCopyProhibited                pkcs11Error = 0x0000001a
	errActionProhibited              pkcs11Error = 0x0000001b
	errDataInvalid                   pkcs11Error = 0x00000020
	errDataLenRange                  pkcs11Error = 0x00000021
	errDeviceError                   pkcs11Error = 0x00000030
	errDeviceMemory                  pkcs11Error = 0x00000031
	errDeviceRemoved                 pkcs11Error = 0x00000032
	errEncryptedDataInvalid          pkcs11Error = 0x00000040
	errEncryptedDataLenRange         pkcs11Error = 0x00000041
	errFunctionCanceled              pkcs11Error = 0x00000050
	errFunctionNotParallel           pkcs11Error = 0x00000051
	errFunctionNotSupported          pkcs11Error = 0x00000054
	errKeyHandleInvalid              pkcs11Error = 0x00000060
	errKeySizeRange                  pkcs11Error = 0x00000062
	errKeyTypeInconsistent           pkcs11Error = 0x00000063
	errKeyNotNeeded                  pkcs11Error = 0x00000064
	errKeyChanged                    pkcs11Error = 0x00000065
	errKeyNeeded                     pkcs11Error = 0x00000066
	errKeyIndigestible               pkcs11Error = 0x00000067
	errKeyFunctionNotPermitted       pkcs11Error = 0x00000068
	errKeyNotWrappable               pkcs11Error = 0x00000069
	errKeyUnextractable              pkcs11Error = 0x0000006a
	errMechanismInvalid              pkcs11Error = 0x00000070
	errMechanismParamInvalid         pkcs11Error = 0x00000071
	errObjectHandleInvalid           pkcs11Error = 0x00000082
	errOperationActive               pkcs11Error = 0x00000090
	errOperationNotInitialized       pkcs11Error = 0x00000091
	errPINIncorrect                  pkcs11Error = 0x000000a0
	errPINInvalid                    pkcs11Error = 0x000000a1
	errPINLenRange                   pkcs11Error = 0x000000a2
	errPINExpired                    pkcs11Error = 0x000000a3
	errPINLocked                     pkcs11Error = 0x000000a4
	errSessionClosed                 pkcs11Error = 0x000000b0
	errSessionCount                  pkcs11Error = 0x000000b1
	errSessionHandleInvalid          pkcs11Error = 0x000000b3
	errSessionParallelNotSupported   pkcs11Error = 0x000000b4
	errSessionReadOnly               pkcs11Error = 0x000000b5
	errSessionExists                 pkcs11Error = 0x000000b6
	errSessionReadOnlyExists         pkcs11Error = 0x000000b7
	errSessionReadWriteSoExists      pkcs11Error = 0x000000b8
	errSignatureInvalid              pkcs11Error = 0x000000c0
	errSignatureLenRange             pkcs11Error = 0x000000c1
	errTemplateIncomplete            pkcs11Error = 0x000000d0
	errTemplateInconsistent          pkcs11Error = 0x000000d1
	errTokenNotPresent               pkcs11Error = 0x000000e0
	errTokenNotRecognized            pkcs11Error = 0x000000e1
	errTokenWriteProtected           pkcs11Error = 0x000000e2
	errUnwrappingKeyHandleInvalid    pkcs11Error = 0x000000f0
	errUnwrappingKeySizeRange        pkcs11Error = 0x000000f1
	errUnwrappingKeyTypeInconsistent pkcs11Error = 0x000000f2
	errUserAlreadyLoggedIn           pkcs11Error = 0x00000100
	errUserNotLoggedIn               pkcs11Error = 0x00000101
	errUserPINNotInitialized         pkcs11Error = 0x00000102
	errUserTypeInvalid               pkcs11Error = 0x00000103
	errUserAnotherAlreadyLoggedIn    pkcs11Error = 0x00000104
	errUserTooManyTypes              pkcs11Error = 0x00000105
	errWrappedKeyInvalid             pkcs11Error = 0x00000110
	errWrappedKeyLenRange            pkcs11Error = 0x00000112
	errWrappingKeyHandleInvalid      pkcs11Error = 0x00000113
	errWrappingKeySizeRange          pkcs11Error = 0x00000114
	errWrappingKeyTypeInconsistent   pkcs11Error = 0x00000115
	errRandomSeedNotSupported        pkcs11Error = 0x00000120
	errRandomNoRNG                   pkcs11Error = 0x00000121
	errDomainParamsInvalid           pkcs11Error = 0x00000130
	errCurveNotSupported             pkcs11Error = 0x00000140
	errBufferTooSmall                pkcs11Error = 0x00000150
	errSavedStateInvalid             pkcs11Error = 0x00000160
	errInformationSensitive          pkcs11Error = 0x00000170
	errStateUnsaveable               pkcs11Error = 0x00000180
	errCryptokiNotInitialized        pkcs11Error = 0x00000190
	errCryptokiAlreadyInitialized    pkcs11Error = 0x00000191
	errMutexBad                      pkcs11Error = 0x000001a0
	errMutexNotLocked                pkcs11Error = 0x000001a1
	errFunctionRejected              pkcs11Error = 0x00000200
	errVendorDefined                 pkcs11Error = 0x80000000
)
