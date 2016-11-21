package log

// Handler interface for log handlers
type Handler interface {
	Run() chan<- *Entry
}
