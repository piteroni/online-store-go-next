package admin_login_handler_test

import (
	common "piteroni/online-store-api"
	"piteroni/online-store-api/database"
	"piteroni/online-store-api/models"
	"piteroni/online-store-api/restapi/handlers/admin_login_handler"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestAuthenticationOK(t *testing.T) {
	db, err := common.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		err := common.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	a := admin_login_handler.NewAdminAuthenticator(db)

	t.Run("login ID and password of registered user", func(t *testing.T) {
		hashed := generateFromPassword(t, "passw0rd")

		user := models.Stuff{
			Model: gorm.Model{
				ID: 1,
			},
			LoginID:  "100",
			Name:     "authed user",
			Password: hashed,
		}

		err = db.Create(&user).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		expected := &models.Stuff{
			Model: gorm.Model{
				ID: 1,
			},
			LoginID:  "100",
			Name:     "authed user",
			Password: hashed,
		}

		u, err := a.Authenticate("100", "passw0rd")

		assert.Nil(t, err)
		assert.Equal(t, expected, u)
	})
}

func TestAuthenticationError(t *testing.T) {
	db, err := common.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		err := common.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}

	a := admin_login_handler.NewAdminAuthenticator(db)

	t.Run("login ID of not registered", func(t *testing.T) {
		hashed := generateFromPassword(t, "passw0rd")

		user := models.Stuff{
			Model: gorm.Model{
				ID: 1,
			},
			LoginID:  "101",
			Name:     "authed user",
			Password: hashed,
		}

		err := db.Create(&user).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		u, err := a.Authenticate("102", "passw0rd")

		assert.Nil(t, u)
		assert.IsType(t, err, &admin_login_handler.AuthenticationError{})
	})

	t.Run("incorrect password of registered user", func(t *testing.T) {
		hashed := generateFromPassword(t, "passw0rd")

		user := models.Stuff{
			Model: gorm.Model{
				ID: 1,
			},
			LoginID:  "102",
			Name:     "authed user",
			Password: hashed,
		}

		err := db.Create(&user).Error
		if err != nil {
			t.Fatal(err)
		}

		defer cleanup()

		u, err := a.Authenticate("102", "incorrect password")

		assert.Nil(t, u)
		assert.IsType(t, err, &admin_login_handler.AuthenticationError{})
	})
}

func generateFromPassword(t *testing.T, password string) string {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}

	return string(p)
}
