package store

type Store interface {
	Images() ImageRepository
	Users() UserRepository
	Error(int64, error)
	LogRequest(int64, string)
}
