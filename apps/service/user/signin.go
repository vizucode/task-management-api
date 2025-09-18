package user

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"task-management-api/apps/domain"
	"task-management-api/helpers/constants/httpstd"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vizucode/gokit/logger"
	"github.com/vizucode/gokit/utils/env"
	"github.com/vizucode/gokit/utils/errorkit"
	"golang.org/x/crypto/bcrypt"
)

func (uc *user) SignIn(ctx context.Context, payload domain.User) (resp domain.SignInResponse, err error) {
	// Validate input
	err = uc.validator.StructCtx(ctx, payload)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	// Get user by email from database
	user, err := uc.db.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		err = errorkit.NewErrorStd(http.StatusUnauthorized, "", "Invalid email or password")
		logger.Log.Error(ctx, err)
		return resp, err
	}

	// Generate JWT token
	accessToken, err := uc.generateAccessToken(ctx, user.ID.String())
	if err != nil {
		logger.Log.Error(ctx, err)
		return resp, errorkit.NewErrorStd(http.StatusInternalServerError, "", httpstd.InternalServerError)
	}

	fmt.Println(accessToken)

	// Prepare response
	resp = domain.SignInResponse{
		AccessToken: accessToken,
		Username:    user.Username,
	}

	return resp, nil
}

// generateAccessToken creates a JWT access token for the user
func (uc *user) generateAccessToken(ctx context.Context, userID string) (string, error) {

	// Get secret key from environment
	secretKey := env.GetString("ACCESS_SECRET_KEY")
	fmt.Println("secret", secretKey)
	if secretKey == "" {
		err := errorkit.NewErrorStd(http.StatusInternalServerError, "", "JWT secret key not configured")
		logger.Log.Error(ctx, err)
		return "", err
	}

	// Set token expiration time (24 hours from now)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create JWT claims
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(env.GetString("ACCESS_SECRET_KEY")))
	if err != nil {
		logger.Log.Error(ctx, err)
		return "", err
	}

	return tokenString, nil
}
