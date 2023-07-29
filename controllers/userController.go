package controllers

import (
	"context"
	"fmt"
	"golang-jwt-project/helpers"
	helper "golang-jwt-project/helpers"
	"golang-jwt-project/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ZenoHwanEth/go-jwt-with-gin-gonic/database"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string)

func VerifyPassword(password string)

func Signup(password string)

func Login(password string)

func GetUsers()

func GetUser()
