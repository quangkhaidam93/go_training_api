package models

// Student ...
type Student struct {
	ID       uint    `json:"id" gorm:"primary_key"`
	FullName string  `json:"fullName"`
	PassWord string  `json:"passWord"`
	Email    string  `json:"email"`
	Classes  []Class `json:"class,omitempty"`
}
