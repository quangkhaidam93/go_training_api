package handlers

import (
	"database/sql"
	"net/http"

	"gorm.io/gorm"
)

type (
	Handler struct {
		DB             *sql.DB
		gormDB         *gorm.DB
		StudentHandler StudentInterface
		ClassHandler   ClassInterface
	}

	StudentInterface interface {
		GetStudentHandler(w http.ResponseWriter, r *http.Request)
		CreateNewStudentHandler(w http.ResponseWriter, r *http.Request)
		UpdateStudentHandler(w http.ResponseWriter, r *http.Request)
		DeleteStudentHandler(w http.ResponseWriter, r *http.Request)
	}

	ClassInterface interface {
		GetClassHandler(w http.ResponseWriter, r *http.Request)
		CreateClassHandler(w http.ResponseWriter, r *http.Request)
		// UpdateClassHandler(w http.ResponseWriter, r *http.Request)
		// DeleteClassHandler(w http.ResponseWriter, r *http.Request)
	}
)
