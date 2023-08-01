package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ZenoHwanEth/go-jwt-with-gin-gonic/helpers"
	"github.com/ZenoHwanEth/go-jwt-with-gin-gonic/models"
	"gopkg.in/mgo.v2/bson"

	"github.com/ZenoHwanEth/go-jwt-with-gin-gonic/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string)

func VerifyPassword(password string)

func Signup(password string)

func Login(password string)

func GetUsers()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helpers.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
	}
}
