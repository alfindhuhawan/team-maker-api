package vo

import "regexp"

const uppercaseAlphaNumeric = "^[A-Z0-9]*$"

var regexUppercaseAlphaNumeric = regexp.MustCompile(uppercaseAlphaNumeric)

const numeric = "^[0-9]*$"

var regexNumeric = regexp.MustCompile(numeric)
