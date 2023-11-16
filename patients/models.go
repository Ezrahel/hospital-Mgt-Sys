package patients

import "gorm.io/gorm"

type patientsDetails struct {
	gorm.Model
	Firstname    string `json:"firstname" validate:"required, min=4, max=50"`
	Lastname     string `json:"lastname" validate:"required, min=4, max=50"`
	Email        string `json:"email" validate: "email, required"`
	Title        string `json:"title" validate: "required"`
	Password     string `json:"password" validate: "required, min=8"`
	Phone        string `json:"phone" validate: "required"`
	UserType     string `json:"usertype" validate:"required, eq=ADMIN¦eq=USER`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
	Sickness     string `json:"sickness"`
}
