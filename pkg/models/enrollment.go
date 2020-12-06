package models

// Enrollment ...
type Enrollment struct {
	ID        uint `gorm:"primary_key"`
	StudentID uint
	ClassID   uint
}
