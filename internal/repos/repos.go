package repos

import "go.mongodb.org/mongo-driver/mongo"

type MongoRepo struct {
	mgClient *mongo.Client
}

func NewMongoDBRepo(mgClient *mongo.Client) IRepo {
	return &MongoRepo{
		mgClient: mgClient,
	}
}

func (m *MongoRepo) SendService() ISendServiceRepo {
	return NewSendServiceNoSQLRepo(m.mgClient)
}
