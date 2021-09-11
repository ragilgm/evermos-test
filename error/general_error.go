package error

// GeneralError :
type GeneralError map[string]string

var ce GeneralError

func init() {
	ce = map[string]string{
		"3P-TDS-0001": "General Error",
		"3P-TDS-0002": "Parsing Error",
		"MS-TDS-0001": "Request Nil",
		"MS-TDS-0002": "{field} is required",
		"MS-TDS-0003": "Failure Parsing",
	}
}

// ErrorMessage :
func ErrorMessage(errorCode string) string {
	return errorCode + "|" + ce[errorCode] + "|"
}
