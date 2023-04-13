package repos

import (
	"booking-service/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type SendServiceNoSQLRepo struct {
	mongoClient *mongo.Client
}

func NewSendServiceNoSQLRepo(mongoClient *mongo.Client) ISendServiceRepo {
	return &SendServiceNoSQLRepo{
		mongoClient: mongoClient,
	}
}

func (s SendServiceNoSQLRepo) GetAvaiableEmpl() (*models.Employee, error) {
	employeeCollection := s.mongoClient.Database("send-service-db").Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	empl := &models.Employee{}
	err := employeeCollection.FindOne(ctx, bson.M{"status": models.StatusAvailable}).Decode(empl)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("get available employee fail with err: %v", err.Error())
		return nil, err
	}

	return empl, nil
}
