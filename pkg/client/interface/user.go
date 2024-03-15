package interfaces

import "grpc-api-gateway/pkg/utils/models"

type UserClient interface {
	UsersSignUp(user models.UserSignUp) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
}
