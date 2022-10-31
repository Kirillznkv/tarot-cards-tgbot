package store

type Store interface {
	Images() ImageRepository
	Users() UserRepository
}
