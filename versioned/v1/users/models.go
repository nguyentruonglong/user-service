package users

import (
	"time"
	db "user_service/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//User ...
type UserModel struct {
	gorm.Model
	ID        int       `db:"id, primarykey, autoincrement" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	Name      string    `db:"name" json:"name"`
	Phone     string    `db:"phone" json:"phone"`
	Address   string    `db:"address" json:"address"`
	Country   string    `db:"country" json:"country"`
	Zipcode   string    `db:"zipcode" json:"zipcode"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}

func (u UserModel) SaveOne(data interface{}) error {
	return db.GetDB().Create(data).Error
}

func (u UserModel) HashPassword(data *UserModel) (bool, error) {
	cipherBytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		return false, err
	}
	data.Password = string(cipherBytes)
	return true, nil
}

func (u UserModel) CheckPassword(providedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u UserModel) FindOneUser(condition interface{}) (UserModel, int, error) {
	var model UserModel
	var count_int64 int64
	tx := db.GetDB().Begin()
	tx.Where(condition).First(&model).Count(&count_int64)
	count_int := int(count_int64)
	err := tx.Commit().Error
	return model, count_int, err
}

func (u UserModel) FindManyUser(limit int, offset int) ([]UserModel, int, error) {
	var models []UserModel
	var count_int64 int64

	tx := db.GetDB().Begin()
	tx.Find(&models).Offset(offset).Limit(limit).Count(&count_int64)
	count_int := int(count_int64)
	err := tx.Commit().Error

	return models, count_int, err
}
