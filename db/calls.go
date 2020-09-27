package db

import (
	"github.com/evt/video12/model"
)

// SaveCall saves a call in Postgres
func (db *PgDB) SaveCall(call *model.Call) error {
	_, err := db.Model(call).Returning("*").Insert()
	return err
}
