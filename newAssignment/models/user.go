package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func (u *User) Save() error {
// 	hashedPassword, err := utils.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = hashedPassword

// 	result := db.DB.Create(&u)
// 	return result.Error
// }
