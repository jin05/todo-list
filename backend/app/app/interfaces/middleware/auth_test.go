package middleware

import (
	"encoding/base64"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/tj/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	test_config "todo-list/app/library/test/config"
)

func newAuthMiddleware(ctrl *gomock.Controller) *authMiddleware {
	return &authMiddleware{conf: test_config.NewTestConfig()}
}

func TestNewAuthMiddleware(t *testing.T) {
	t.Run("normal test", func(t *testing.T) {
		conf := test_config.NewTestConfig()
		result, err := NewAuthMiddleware(conf)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func Test_authMiddleware_Handler(t *testing.T) {
	t.Run("method options test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		auth := newAuthMiddleware(ctrl)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodOptions,
			"/todo",
			nil,
		)

		res, req, err := auth.Handler(w, r)
		user := UserForContext(req.Context())
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotNil(t, res)
		assert.NotNil(t, req)
		assert.Nil(t, err)
		assert.Nil(t, user)
	})

	t.Run("empty idToken test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		auth := newAuthMiddleware(ctrl)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			"/todo",
			nil,
		)

		res, req, err := auth.Handler(w, r)
		user := UserForContext(req.Context())
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.NotNil(t, res)
		assert.NotNil(t, req)
		assert.Equal(t, "invalid jws message: invalid byte sequence", err.Error())
		assert.Nil(t, user)
	})

	t.Run("jwt error test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		auth := newAuthMiddleware(ctrl)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			"/todo",
			nil,
		)
		r.Header.Set("Authorization", "idToken")

		res, req, err := auth.Handler(w, r)
		user := UserForContext(req.Context())
		assert.Equal(t, http.StatusForbidden, w.Code)
		assert.NotNil(t, res)
		assert.NotNil(t, req)
		assert.Equal(t, "invalid jws message: invalid compact serialization format: invalid number of segments", err.Error())
		assert.Nil(t, user)
	})

	t.Run("normal test", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		auth := newAuthMiddleware(ctrl)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			"/todo",
			nil,
		)

		now := time.Now()
		head := "eyJraWQiOiJJampLUXNpQUJcLzlGMjVoS3JrREdkT241NEdTMWx2OW44MWVQcVwvbEc4eG89IiwiYWxnIjoiUlMyNTYifQ."
		signature := ".kUobQtVpXueHOYAur4Ihw6o3w6bjdCRV0QFp93elAHoEQTNYI2q4EgmEbWhFeOTllg_CDg4e5VFA3oC_6H-YWGIucn6FGQ-UBAu8YLNny3K0rOcM8fctKl7Ct5uGmavipCajqVBZ1Pc3GWaChRrq2x2SO_iBEspFhu4NkzXbdWVshvVJOpiaIkdnvIiLDGNeCPR90OI-GImraFGrf1LBd2VojS8WKFxS-BRQQf9BiVbzPRbgiHAkplmCPMsRCGdvHF-cibQufKcb6P9Xrvi5JsPZSAYrAiLiaAc54KMmstyBSMGJbnH4BtNIRol0-zC2uLAR0hSM5u4mRBucqU08xw"
		tokenMap := map[string]interface{}{
			"sub":              "auth_id",
			"aud":              "aws_user_pool_client_id",
			"email_verified":   true,
			"token_use":        "id",
			"auth_time":        now.Unix(),
			"iss":              "https://cognito-idp.ap-northeast-1.amazonaws.com/aws_user_pool_id",
			"cognito:username": "user_name",
			"exp":              now.Unix() + 604800,
			"iat":              now.Unix(),
			"email":            "test@test.com",
		}
		tokenJson, _ := json.Marshal(tokenMap)
		idTokenData := base64.StdEncoding.EncodeToString(tokenJson)
		r.Header.Set("Authorization", head+idTokenData+signature)

		res, req, err := auth.Handler(w, r)
		user := UserForContext(req.Context())
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotNil(t, res)
		assert.NotNil(t, req)
		assert.Nil(t, err)
		assert.Equal(t, "auth_id", user.AuthID)
		assert.Equal(t, "user_name", user.Name)
		assert.Equal(t, "test@test.com", user.Email)
	})
}
