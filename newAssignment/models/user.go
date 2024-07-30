package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Status   string `json:"status" gorm:"default:'inactive'"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
type LoginResponse struct {
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// SignupResponse represents the signup response
type SignupResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

// UserResponse represents the user information in the response
// func (u *User) Save() error {
// 	hashedPassword, err := utils.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = hashedPassword

// 	result := db.DB.Create(&u)
// 	return result.Error
// }
