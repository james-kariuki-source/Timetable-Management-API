package models

type Admin struct {
	Id       uint   `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
