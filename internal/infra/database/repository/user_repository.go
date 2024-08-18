package database

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/reangeline/wpa_user_saas/internal/domain/entity"
)

type UserRepository struct {
	svc *dynamodb.Client
}

func NewUserRepository(svc *dynamodb.Client) *UserRepository {
	return &UserRepository{
		svc,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {

	usa, err := attributevalue.MarshalMap(user)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      usa,
		TableName: aws.String("usersTable"),
	}

	_, err = ur.svc.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var userModel *entity.User
	fmt.Println("phone repository 1", phone)

	key := map[string]types.AttributeValue{
		"phone": &types.AttributeValueMemberS{Value: phone},
	}

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String("usersTable"),
		Key:       key,
	}

	fmt.Println("phone repository 2", phone)

	result, err := ur.svc.GetItem(context.TODO(), getItemInput)
	if err != nil {
		log.Fatalf("Failed to get item, %v", err)
	}

	fmt.Println("phone repository 3", result)

	if err != nil {
		return nil, fmt.Errorf("got error calling GetItem: %s", err)
	}

	err = attributevalue.UnmarshalMap(result.Item, &userModel)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Record, %v", err)
	}

	fmt.Println("phone repository 3", phone)

	return userModel, nil
}
