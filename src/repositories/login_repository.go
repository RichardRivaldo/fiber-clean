package repositories

import (
	"context"
	"fiber-clean/src/configs"
	"fiber-clean/src/models"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func checkAdmin(email string, password string) (*models.Admin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	var result models.Admin

	err := adminCollection.FindOne(ctx, filter).Decode(&result)
	if err == nil {
		if result.Password == password {
			return &result, nil
		}
	}

	return nil, err
}

func checkUser(email string, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"email": email}

	var result models.User

	err := userCollection.FindOne(ctx, filter).Decode(&result)
	if err == nil {
		if result.Password == password {
			return &result, nil
		}
	}
	return nil, err
}

func Login(email string, password string) (string, error) {
	adminData, err := checkAdmin(email, password)
	if err == nil && adminData != nil {
		claims := jwt.MapClaims{
			"id":       adminData.ID.String(),
			"email":    adminData.Email,
			"exp":      time.Now().Add(time.Hour * 24 * 1).Unix(),
			"is_admin": true,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		secret := configs.GetEnv("JWT_AUTH_SECRET")
		t, err := token.SignedString([]byte(secret))

		return t, err
	}

	userData, err := checkUser(email, password)
	if err == nil && userData != nil {
		claims := jwt.MapClaims{
			"id":       userData.ID.String(),
			"email":    userData.Email,
			"exp":      time.Now().Add(time.Hour * 24 * 1).Unix(),
			"is_admin": false,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		secret := configs.GetEnv("JWT_AUTH_SECRET")
		t, err := token.SignedString([]byte(secret))

		return t, err
	}

	return "", err
}
