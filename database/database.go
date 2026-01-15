package database

import(
	"fmt"
	"context"
	"slices"
	"errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/bson"
)

type itemStats struct{
	ua  int
	ucc int
	uco int
	rcam int
}

type item struct {
	office string
	date Date
	stadistics itemStats
}

const collectionList := []string{"DailyStadistics"}

func Connect() *mongo.Client{

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opt := options.Client().ApplyURI("mongodb://127.0.0.1:27017/").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opt)

	if err != nil {
		fmt.Println(err)	
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Connected")
	}

	return client

}

func InsertItem(client *mongo.Client, it interface{}, collec string) error{
	
	//check if collection exist
	if slices.Contains(collectionList, collect) == false {
		return errors.New("collection not found")
	}

	//check if item is a valid struct
	switch it.(type){
		case item:
			if collec != ""{
				return errors.New("type do not match collection")
			}
		default:
			return errors.New("type not valid")

	}

	collection := client.Database("StadisticSaime").Collection(collec)

	result, err := collection.InsertOne(context.TODO(), in)

	return err

}
