package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID                  string         `gorm:"primaryKey" json:"id"`
	FirstName           string         `gorm:"column:first_name" json:"firstName"`
	LastName            *string        `gorm:"column:last_name" json:"lastName"`
	Email               string         `gorm:"unique" json:"email"`
	Password            *string        `json:"-"`
	PhotoURL            *string        `gorm:"column:photo_url" json:"photoURL"`
	IsVerified          bool           `gorm:"column:is_verified;default:false" json:"isVerified"`
	IsTermsAccepted     bool           `gorm:"column:is_terms_accepted;default:false" json:"isTermsAccepted"`
	OnBoardingCompleted bool           `gorm:"column:on_boarding_completed;default:false" json:"onBoardingCompleted"`
	GoogleID            *string        `gorm:"column:google_id" json:"googleID"`
	LinkedinID          *string        `gorm:"column:linkedin_id" json:"linkedinID"`
	FacebookID          *string        `gorm:"column:facebook_id" json:"facebookID"`
	GithubID            *string        `gorm:"column:github_id" json:"githubID"`
	BirthDate           *time.Time     `gorm:"column:birth_date" json:"birthDate"`
	Country             *string        `json:"country"`
	State               *string        `json:"state"`
	Skills              pq.StringArray `gorm:"type:text[]" json:"skills"`
	City                *string        `json:"city"`
	Linkedin            *string        `json:"linkedin"`
	Github              *string        `json:"github"`
	Website             *string        `json:"website"`
	YearsOfExperience   *float64       `gorm:"column:years_of_experience" json:"yearsOfExperience"`
	Description         *string        `json:"description"`
	JobTitle            *string        `gorm:"column:job_title" json:"jobTitle"`
	JobCompany          *string        `gorm:"column:job_company" json:"jobCompany"`
	Biography           *string        `json:"biography"`
	IsMentor            bool           `gorm:"column:is_mentor;default:false" json:"isMentor"`
	Status              string         `gorm:"default:'PENDING'" json:"status"`
	Active              bool           `gorm:"default:true" json:"active"`
	CreatedAt           time.Time      `gorm:"default:current_timestamp;column:created_at" json:"createdAt"`
	UpdatedAt           time.Time      `gorm:"default:current_timestamp;column:updated_at" json:"updatedAt"`
	Events              []Event        `gorm:"many2many:EventsOnUsers;" json:"events"`
}

func (User) TableName() string {
	return "users"
}
