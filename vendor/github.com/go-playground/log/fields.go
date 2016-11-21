package log

// Field contains a single key + value for structured logging.
type Field struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
