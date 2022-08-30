package models

import (
	"gorm.io/gorm"
)

//User ...
type User struct {
	gorm.Model
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	Phone     string `db:"phone" json:"phone"`
	Address   string `db:"address" json:"address"`
	Country   string `db:"country" json:"country"`
	Zipcode   string `db:"zipcode" json:"zipcode"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}

// func (h User) Signup(userPayload forms.UserSignup) (*User, error) {
// 	db := db.GetDB()
// 	id := uuid.NewV4()
// 	user := User{
// 		ID:        id.String(),
// 		Name:      userPayload.Name,
// 		BirthDay:  userPayload.BirthDay,
// 		Gender:    userPayload.Gender,
// 		PhotoURL:  userPayload.PhotoURL,
// 		Time:      time.Now().UnixNano(),
// 		Active:    true,
// 		UpdatedAt: time.Now().UnixNano(),
// 	}
// 	item, err := dynamodbattribute.MarshalMap(user)
// 	if err != nil {
// 		errors.New("error when try to convert user data to dynamodbattribute")
// 		return nil, err
// 	}
// 	params := &dynamodb.PutItemInput{
// 		Item:      item,
// 		TableName: aws.String("TableUsers"),
// 	}
// 	if _, err := db.PutItem(params); err != nil {
// 		log.Println(err)
// 		return nil, errors.New("error when try to save data to database")
// 	}
// 	return &user, nil
// }

func (h User) GetByID(id string) (*User, error) {
	// db := db.GetDB()
	var user *User
	return user, nil
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
