package db

import "github.com/jinzhu/gorm"

//Db instance
var Db *gorm.DB

//Open db connection
func Open() error {
	var err error
	Db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		return err
	}
	return nil
}

// Close db connection
func Close() error {
	return Db.Close()
}
