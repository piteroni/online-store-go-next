package admin_login_handler

import (
	"fmt"
	"piteroni/online-store-api/models"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type AdminAuthenticator struct {
	db *gorm.DB
}

func NewAdminAuthenticator(db *gorm.DB) *AdminAuthenticator {
	return &AdminAuthenticator{
		db: db,
	}
}

func (a *AdminAuthenticator) Authenticate(loginID string, password string) (*models.Stuff, error) {
	user := models.Stuff{}

	condition := &models.Stuff{LoginID: loginID}

	err := a.db.First(&user, condition).Error
	if err != nil {
		if xerrors.Is(err, gorm.ErrRecordNotFound) {
			message := fmt.Sprintf("there is no user matching the specified loginID: %s", loginID)

			return nil, &AuthenticationError{
				message: message,
			}
		} else {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		message := fmt.Sprintf("password is wrong. loginID = %s, backtrace = %+v", loginID, err)

		return nil, &AuthenticationError{
			message: message,
		}
	}

	return &user, nil
}

type AuthenticationError struct {
	message string
}

func (e *AuthenticationError) Error() string {
	return e.message
}
