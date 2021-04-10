package userRepository

import (
	. "github.com/VusalShahbazov/api.freelance4.az/internal/models"
	"github.com/VusalShahbazov/api.freelance4.az/internal/requests"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(input *requests.SignUpRequest) (*User  , error){
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return nil, err
	}
	user := &User{
		LastName:     input.LastName,
		FirstName:    input.FirstName,
		Email:        input.Email,
		Password:     string(password),
		IsFreelancer: false,
	}

	err = DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return  user , nil

}

func CheckEmailIsExists(email string) (bool , error){
	var user User
	err := DB.Where("email = ? " , email).Select("id").First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func GetUserByEmail(email string) (*User  , error) {
	var user User
	err := DB.Where("email = ? " , email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return  &user , nil
}

func ResetPassword(user *User , password string) error {
	hash , err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	//user.Password = string(hash)
	//DB.Save(user)
	err = DB.Model(user).Update("password" , string(hash)).Error
	if err != nil {
		return err
	}

	return nil
}