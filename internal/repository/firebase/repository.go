package firebase

import (
	"avec_moi_with_us_api/api/request/user"
	"avec_moi_with_us_api/api/utils"
	"context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Repository struct {
	PrivateKey string
}

func NewRepository() *Repository {
	return &Repository{PrivateKey: utils.GetSecretFile()}
}
func (r Repository) Validate(IdToken string) (user.SignUp, error) {
	opt := option.WithCredentialsFile(r.PrivateKey)
	var response user.SignUp
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return response, err
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return response, err
	}

	client, err := authClient.GetUser(context.Background(), IdToken)
	if err != nil {
		return response, err
	}
	return user.SignUp{Mail: client.Email, Name: client.DisplayName, Gender: "other"}, nil
}
