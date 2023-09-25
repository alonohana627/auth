package routes

import (
	pb "auth/pb"
	"auth/token"
	"context"
)

func (s *AuthRoutes) RefreshToken(ctx context.Context, tokenDetails *pb.TokenDetails) (*pb.ReturnStatus, error) {
	currentToken := tokenDetails.AccessToken
	
	newToken, err := token.RefreshJWT(currentToken)
	if err != nil {
		return &pb.ReturnStatus{
			Success:  false,
			Response: "The Token is Invalid!",
			Token:    nil,
		}, nil
	}

	return &pb.ReturnStatus{
		Success:  true,
		Response: "A new token was generated!",
		Token:    &pb.TokenDetails{AccessToken: newToken},
	}, nil
}
