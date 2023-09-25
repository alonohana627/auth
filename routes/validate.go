package routes

import (
	pb "auth/pb"
	"auth/token"
	"context"
	"fmt"
)

func (s *AuthRoutes) ValidateToken(ctx context.Context, tokenDetails *pb.TokenDetails) (*pb.ReturnStatus, error) {
	currentToken := tokenDetails.AccessToken

	err := token.ValidateJWT(currentToken)
	if err != nil {
		fmt.Println(err)
		return &pb.ReturnStatus{
			Success:  false,
			Response: "The Token is Invalid!",
		}, nil
	}

	return &pb.ReturnStatus{
		Success:  true,
		Response: "The token is valid!",
	}, nil
}
