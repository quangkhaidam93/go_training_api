package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quangkhaidam93/go_train_api/pkg/models"
)

type StudentHanlder struct {
	db *sql.DB
}

func StudentHandlerInit(db *sql.DB) (s *StudentHanlder) {
	s = &StudentHanlder{db: db}
	return
}

func (stundentHandler *StudentHanlder) GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	rows, error := stundentHandler.db.Query("SELECT * FROM students")
	if error != nil {
		fmt.Println("Loi", error)
	} else {
		studentList := make([]models.Student, 0, 1)
		for rows.Next() {
			student := models.Student{}
			rows.Scan(&student.ID, &student.FullName, &student.PassWord, &student.Email)
			studentList = append(studentList, student)
		}
		fmt.Println(studentList)
		json.NewEncoder(w).Encode(studentList)
	}
}

func (studentHandler *StudentHanlder) CreateNewStudentHandler(w http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Printf("Error reading body data: %v", error)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	var student models.Student
	if err := json.Unmarshal(body, &student); err != nil {
		http.Error(w, "can't parse body", http.StatusBadRequest)
		return
	}
	result, e := studentHandler.db.Exec("INSERT INTO students (full_name, pass_word, email) VALUES (?, ?, ?)", student.FullName, student.PassWord, student.Email)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "query error", http.StatusBadRequest)
		return
	}
	fmt.Println("Result", result)
	w.WriteHeader(http.StatusCreated)
}

func (studentHandler *StudentHanlder) UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		fmt.Println("Error reading body data: %v", error)
		http.Error(w, "can't read body", http.StatusBadGateway)
		return
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	var student models.Student
	if err := json.Unmarshal(body, &student); err != nil {
		http.Error(w, "can't parse body", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id := params["id"]
	result, e := studentHandler.db.Exec("UPDATE students SET full_name = ?, email = ?, pass_word = ? WHERE id = ?",
		student.FullName,
		nil,
		student.PassWord,
		id,
	)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "query error", http.StatusBadRequest)
		return
	}
	if rowAffected, sqlError := result.RowsAffected(); sqlError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Khong co id nay"))
	} else if rowAffected == 0 {
		w.Write([]byte("Id khong ton tai"))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (StudentHanlder *StudentHanlder) DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Den day")
	params := mux.Vars(r)
	id := params["id"]
	if result, e := StudentHanlder.db.Exec(("DELETE FROM students WHERE id = ?"), id); e != nil {
		fmt.Println("Loi", e)
	} else if rowAffected, sqlError := result.RowsAffected(); sqlError != nil {
		fmt.Println("Loi query", sqlError)
	} else if rowAffected == 0 {
		fmt.Println("Khong kiem duoc id")
	} else {
		fmt.Println("Ok done", rowAffected)
	}
}
