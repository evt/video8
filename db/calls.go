package db

import (
	"github.com/evt/video8/model"
)

// SaveCall saves a call in Postgres
func (db *PgDB) SaveCall(call *model.Call) error {
	_, err := db.Model(call).Returning("*").Insert()
	return err
}