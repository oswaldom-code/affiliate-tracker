package ports

type Repository interface {
	// system_services ports
	TestDb() error
}
