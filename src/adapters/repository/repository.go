package repository

//var sessionConfig = &gorm.Session{SkipDefaultTransaction: true, FullSaveAssociations: false}

func (s *repository) TestDb() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
