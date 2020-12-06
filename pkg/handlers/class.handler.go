package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/quangkhaidam93/go_train_api/pkg/models"
	"gorm.io/gorm"
)

type ClassHandler struct {
	gormDb *gorm.DB
}

func ClassHandlerInit(gormDb *gorm.DB) (c *ClassHandler) {
	c = &ClassHandler{gormDb: gormDb}
	return
}

func (c *ClassHandler) GetClassHandler(w http.ResponseWriter, r *http.Request) {
	var classes []models.Class
	result := c.gormDb.Find(&classes)
	if result.Error != nil {
		fmt.Println("Loi")
	} else {
		fmt.Println("Den day", classes)
		json.NewEncoder(w).Encode(classes)
	}
}

func (c *ClassHandler) CreateClassHandler(w http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Printf("Error reading body data: %v", error)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	var class models.Class
	if err := json.Unmarshal(body, &class); err != nil {
		http.Error(w, "can't parse body", http.StatusBadRequest)
		return
	}
	result := c.gormDb.Create(class)
	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Println("Loi")
	} else {
		w.Write([]byte("Created"))
	}
}
