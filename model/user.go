package model

import "time"

// User represents the user model
// swagger:model
type User struct {
	ID          int        `json:"id" orm:"primary"`
	Code        string     `json:"code" orm:"size(16)"`
	Status      string     `json:"status" orm:"size(64)"`
	FirstName   string     `json:"first_name" orm:"size(32)"`
	LastName    string     `json:"last_name" orm:"size(32)"`
	FullName    string     `json:"full_name" orm:"size(100)"`
	Gender      string     `json:"gender" sql:"type:ENUM('M', 'F')"`
	PhoneNumber string     `json:"phone_number,omitempty" orm:"size(20)"`
	Email       string     `json:"email" orm:"size(255)"`
	BirthDate   *time.Time `json:"birth_date" orm:"type(date);default(null)"`

	RegistrationDate *time.Time `json:"registration_date" orm:"default(null)"`

	NationalID           string     `json:"national_id" orm:"size(12)"`
	NationalIDIssueDate  *time.Time `json:"national_id_issue_date" orm:"type(date);default(null)"`
	NationalIDIssueBy    string     `json:"national_id_issue_by" orm:"size(255)"`
	NationalIDExpiryDate *time.Time `json:"national_id_expiry_date" orm:"type(date);default(null)"`
	NationalCardType     string     `json:"national_card_type" orm:"size(32)"`

	Base
}
