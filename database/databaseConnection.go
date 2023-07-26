package database

import (
	"fmt"
	"log"
	"time"
	"os"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo" )

func DBinstance() *mongo.Client{
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	MongoDb:= os.Getenv("MONGODB_URL")


}