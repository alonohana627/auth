package routes

import (
	"auth/database"
	"auth/database/models"
	pb "auth/pb"
	"auth/token"
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s *AuthRoutes) Login(ctx context.Context, user *pb.UserAuthInformation) (*pb.ReturnStatus, error) {
	var password string
	var hashedPassword string
	var jwtToken string
	var err error

	if user == nil {
		return &pb.ReturnStatus{
				Response: "There is no Auth Information!",
				Success:  false,
			}, &ErrorWithRoute{
				ErrorDescription: "There is no Auth Info!",
			}
	}

	// Check if user exists
	userModel := models.Users{}

	result := database.DBInstance.Where("email = ?", user.Email).First(&userModel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			goto emailOrPassNotFound
		}
		return &pb.ReturnStatus{
			Response: result.Error.Error(),
			Success:  false,
		}, nil
	} else if result.RowsAffected == 0 {
		goto emailOrPassNotFound
	}

	password = user.Password
	hashedPassword = userModel.HashedPassword

	if !bcryptCheckPasswordHash(password, hashedPassword) {
		goto emailOrPassNotFound
	}

	jwtToken, err = token.GenerateJWT(userModel.Email, userModel.Nickname)
	if err != nil {
		return &pb.ReturnStatus{
			Response: "Error generate JWT token",
			Success:  false,
		}, err
	}

	return &pb.ReturnStatus{
		Response: "Welcome " + userModel.Nickname,
		Success:  true,
		Token: &pb.TokenDetails{
			AccessToken: jwtToken,
		},
	}, nil

emailOrPassNotFound:
	return &pb.ReturnStatus{
		Response: "Email does not exists or Password is not correct!",
		Success:  false,
	}, nil
}
