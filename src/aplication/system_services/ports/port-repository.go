package ports

type Store interface {
	TestDb() error
}
