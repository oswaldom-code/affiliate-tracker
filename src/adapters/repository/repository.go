package repository

//var sessionConfig = &gorm.Session{SkipDefaultTransaction: true, FullSaveAssociations: false}

func (r *repository) TestDb() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}
