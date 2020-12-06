package handlers

import (
	"database/sql"

	"gorm.io/gorm"
)

func HandlerInit(db *sql.DB, gormDb *gorm.DB) (h *Handler) {
	h = &Handler{
		DB:     db,
		gormDB: gormDb,
	}

	h.StudentHandler = StudentHandlerInit(db)
	h.ClassHandler = ClassHandlerInit(gormDb)
	return
}
