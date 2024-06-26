package usecases

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/karma-dev-team/karma-docs/internal/user"
	"github.com/karma-dev-team/karma-docs/internal/user/entities"
)

type AuthService struct {
	userService    user.UserRepository
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func (a *AuthService) SignUp(ctx context.Context, username, email, password string) error {
	user, err := entities.NewUser(
		username,
		email,
		password,
	)
	if err != nil {
		return err
	}

	return a.userRepo.CreateUser(ctx, user)
}

func (a *AuthService) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.userRepo.GetUser(ctx, username, password)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(a.signingKey)
}

func (a *AuthService) ParseToken(ctx context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}
