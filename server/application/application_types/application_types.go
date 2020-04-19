package application_types

import "errors"

// ---------------------------------------------------------------------------------------------------------------------
//		t o k e n s
// ---------------------------------------------------------------------------------------------------------------------

const AppToken = "iuhdiu87w23ruh897dfyc2w3r"

// ---------------------------------------------------------------------------------------------------------------------
//		e r r o r s
// ---------------------------------------------------------------------------------------------------------------------

var (
	PanicSystemError           = errors.New("panic_system_error")
	MismatchConfigurationError = errors.New("mismatch_configuration")
)
