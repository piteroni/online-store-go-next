package admin_login_handler_test

import (
	"net/http/httptest"
	"os"
	common "piteroni/online-store-api"
	"piteroni/online-store-api/database"
	"piteroni/online-store-api/models"
	"piteroni/online-store-api/restapi/auth"
	"piteroni/online-store-api/restapi/handlers/admin_login_handler"
	"piteroni/online-store-api/restapi/operations"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestAdminLoginHandler(t *testing.T) {
	db, err := common.ConnnectToInMemoryDatabase()
	if err != nil {
		t.Fatal(err)
	}

	err = database.Migrate(db)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		err := common.RefreshInMemoryDatabase(db)
		if err != nil {
			t.Fatal(err)
		}
	}()

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

	tests := []struct {
		name       string
		loginID    string
		password   string
		wantStatus int
	}{
		{
			name:       "ok",
			loginID:    "100",
			password:   "passw0rd",
			wantStatus: 200,
		},
		{
			name:       "missing login id",
			loginID:    "",
			password:   "passw0rd",
			wantStatus: 400,
		},
		{
			name:       "missing password",
			loginID:    "100",
			password:   "",
			wantStatus: 400,
		},
		{
			name:       "incorrect params",
			loginID:    "100",
			password:   "password",
			wantStatus: 401,
		},
	}

	h := admin_login_handler.AdminLoginHandler{
		JWTTokenGenerator:  &auth.JWTTokenGenerator{},
		AppLogger:          common.NewLogger(os.Stdout),
		AdminAuthenticator: admin_login_handler.NewAdminAuthenticator(db),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := operations.NewPostAdminLoginParams()
			params.HTTPRequest = httptest.NewRequest("POST", "http://localhost:9000", nil)
			params.Body = operations.PostAdminLoginBody{
				LoginID:  &tt.loginID,
				Password: &tt.password,
			}

			response := h.Handle(params)

			w := httptest.NewRecorder()
			response.WriteResponse(w, runtime.JSONProducer())

			assert.Equal(t, tt.wantStatus, w.Result().StatusCode)
			assert.NoError(t, err)
		})
	}
}
