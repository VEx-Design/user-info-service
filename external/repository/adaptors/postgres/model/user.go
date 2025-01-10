package gorm_model

type User struct {
	ID      string `gorm:"primaryKey"`
	Email   string `gorm:"unique"`
	Name    string
	Picture string
}
