package errors

// Error - error object to contain the error details
type Error struct {
	TraceId    string                 `json:"traceId"`
	Code       int                    `json:"code"`
	Message    string                 `json:"message"`
	Info       map[string]interface{} `json:"info,omitempty"`
	HttpStatus int                    `json:"-"`
}

// New - constructor of error struct
func New(code int) Error {
	err, _ := ErrorCodeMapping[code]
	return err
}

// WithInfo - add additional info
func (e *Error) WithInfo(key string, value interface{}) *Error {
	if e.Info == nil {
		e.Info = make(map[string]interface{})
	}
	e.Info[key] = value
	return e
}

func (e *Error) Error() string {
	return e.Message
}
