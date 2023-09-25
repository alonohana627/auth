package routes

import (
	"auth/database"
	"auth/database/models"
	pb "auth/pb"
	"context"
	"fmt"
)

// TODO: Email Verification

func (s *AuthRoutes) Register(ctx context.Context, user *pb.User) (*pb.ReturnStatus, error) {
	fmt.Println("Register invoked")

	authInfo := user.GetAuthInfo()
	if authInfo == nil {
		return &pb.ReturnStatus{
				Response: "There is no Auth Information!",
				Success:  false,
			}, &ErrorWithRoute{
				ErrorDescription: "There is no Auth Info!",
			}
	}

	hashedPassword, err := bcryptHashPassword(authInfo.Password)
	if err != nil {
		return &pb.ReturnStatus{
			Response: "Hash Password Failed!!",
			Success:  false,
		}, err
	}

	newUser := &models.Users{
		Email:          authInfo.Email,
		HashedPassword: hashedPassword, // TODO: change
		Nickname:       user.Nickname,
	}

	result := database.DBInstance.Create(newUser)
	if result.Error != nil {
		return &pb.ReturnStatus{
			Success:  false,
			Response: result.Error.Error(),
		}, nil
	}

	return &pb.ReturnStatus{
		Success:  true,
		Response: "Created the user successfully!",
	}, nil
}
