package database

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/reangeline/wpa_user_saas/internal/domain/entity"
)

type UserRepository struct {
	svc *dynamodb.DynamoDB
}

func NewUserRepository(svc *dynamodb.DynamoDB) *UserRepository {
	return &UserRepository{
		svc,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {

	usa, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return err
	}

	fmt.Println("teste", usa)

	input := &dynamodb.PutItemInput{
		Item:      usa,
		TableName: aws.String("usersTable"),
	}

	_, err = ur.svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var userModel *entity.User
	result, err := ur.svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("usersTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("got error calling GetItem: %s", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &userModel)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Record, %v", err)
	}

	return userModel, nil

}

func (ur *UserRepository) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var userModel *entity.User
	result, err := ur.svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("usersTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"phone": {
				S: aws.String(phone),
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("got error calling GetItem: %s", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &userModel)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Record, %v", err)
	}

	return userModel, nil
}
