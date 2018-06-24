package commons

import "errors"

// Error substrings
const FoundRowsErrorSubstring = "destination arguments in Scan,"
const RedisNotFoundErrorSubstring = "redis: nil"

// Custom errors
var UnknownError = errors.New("an unknown error ocurred")
