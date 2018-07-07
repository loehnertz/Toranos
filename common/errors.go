package common

import "errors"

// Error substrings
const FoundRowsErrorSubstring = "destination arguments in Scan,"

// Custom errors
var UnknownError = errors.New("an unknown error ocurred")
var NotAuthorizedError = errors.New("you are not authorized to access this path")
