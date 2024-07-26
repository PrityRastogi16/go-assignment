package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
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
