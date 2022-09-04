package users

import (
	"time"
	db "user_service/db"

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

func SaveOne(data interface{}) error {
	return db.GetDB().Create(data).Error
}

func (h UserModel) FindOneUser(condition interface{}) (UserModel, error) {
	var model UserModel
	tx := db.GetDB().Begin()
	tx.Where(condition).First(&model)
	err := tx.Commit().Error
	return model, err
}

func (h UserModel) FindManyUser(limit int, offset int) ([]UserModel, int, error) {
	var models []UserModel
	var count_int64 int64

	tx := db.GetDB().Begin()
	tx.Find(&models).Offset(offset).Limit(limit).Count(&count_int64)
	count_int := int(count_int64)
	err := tx.Commit().Error

	return models, count_int, err
}

// //UserModel ...
// type UserModel struct{}

// var authModel = new(AuthModel)

// //Login ...
// func (m UserModel) Login(form forms.LoginForm) (user User, token Token, err error) {

// 	err = db.GetDB().SelectOne(&user, "SELECT id, email, password, name, updated_at, created_at FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)

// 	if err != nil {
// 		return user, token, err
// 	}

// 	//Compare the password form and database if match
// 	bytePassword := []byte(form.Password)
// 	byteHashedPassword := []byte(user.Password)

// 	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

// 	if err != nil {
// 		return user, token, err
// 	}

// 	//Generate the JWT auth token
// 	tokenDetails, err := authModel.CreateToken(user.ID)
// 	if err != nil {
// 		return user, token, err
// 	}

// 	saveErr := authModel.CreateAuth(user.ID, tokenDetails)
// 	if saveErr == nil {
// 		token.AccessToken = tokenDetails.AccessToken
// 		token.RefreshToken = tokenDetails.RefreshToken
// 	}

// 	return user, token, nil
// }

// //Register ...
// func (m UserModel) Register(form forms.RegisterForm) (user User, err error) {
// 	getDb := db.GetDB()

// 	//Check if the user exists in database
// 	checkUser, err := getDb.SelectInt("SELECT count(id) FROM public.user WHERE email=LOWER($1) LIMIT 1", form.Email)
// 	if err != nil {
// 		return user, errors.New("something went wrong, please try again later")
// 	}

// 	if checkUser > 0 {
// 		return user, errors.New("email already exists")
// 	}

// 	bytePassword := []byte(form.Password)
// 	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
// 	if err != nil {
// 		return user, errors.New("something went wrong, please try again later")
// 	}

// 	//Create the user and return back the user ID
// 	err = getDb.QueryRow("INSERT INTO public.user(email, password, name) VALUES($1, $2, $3) RETURNING id", form.Email, string(hashedPassword), form.Name).Scan(&user.ID)
// 	if err != nil {
// 		return user, errors.New("something went wrong, please try again later")
// 	}

// 	user.Name = form.Name
// 	user.Email = form.Email

// 	return user, err
// }

// //One ...
// func (m UserModel) One(userID int64) (user User, err error) {
// 	err = db.GetDB().SelectOne(&user, "SELECT id, email, name FROM public.user WHERE id=$1 LIMIT 1", userID)
// 	return user, err
// }
