package mongo

import (
	"context"

	"github.com/SalomanYu/Academkin/src/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collectionVuz 			*mongo.Collection
	collectionSpec 			*mongo.Collection
	ctx 				 =  context.TODO()
	vuzCount int
	dbName = "academkin"
)

func init(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	checkErr(err)
	
	err = client.Ping(ctx, nil)
	checkErr(err)

	collectionVuz = client.Database(dbName).Collection("Vuz")
	collectionSpec = client.Database(dbName).Collection("Specialization")
}

func AddVuz(vuz *models.Vuz) error{
	_, err := collectionVuz.InsertOne(ctx, vuz)
	vuzCount++
	// fmt.Printf("%d. Vuz: %s\n", vuzCount, vuz.VuzId)
	return err
}

func AddSpecialization(spec *models.Specialization) error{
	_, err := collectionSpec.InsertOne(ctx, spec)
	return err
}

func checkErr(err error) {
	if err != nil{
		panic(err)
	}
}