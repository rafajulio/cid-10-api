package controllers

import (
	"cid-10-api/components/token"
	"cid-10-api/configs"
	"cid-10-api/models"
	"cid-10-api/responses"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cidCollection *mongo.Collection = configs.GetCollection(configs.DB, "cid")
var validate = validator.New()
var jwtMaker, jwtMakerError = token.NewJWTMaker(configs.EnvJwtSecret()) 

func CreateCid(c *fiber.Ctx) error {
	if jwtMakerError != nil {return c.Status(http.StatusInternalServerError).JSON(responses.GenericResponse{Status: http.StatusInternalServerError, Message: "Jwt handler error"})}
	bearer := strings.Replace(c.Get("Authorization"), "Bearer ", "", -1)
	_, err := jwtMaker.VerifyToken(bearer)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.GenericResponse{Status: http.StatusUnauthorized, Message: err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cid models.Cid
	defer cancel()

	//validade the request body
	if err := c.BodyParser(&cid); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GenericResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&cid); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.GenericResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	newCid := models.Cid{
		Id:		primitive.NewObjectID(),
		Code:	cid.Code,
		Title:	cid.Title,
	}

	result, err := cidCollection.InsertOne(ctx, newCid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.GenericResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(responses.CidResponse{Status: http.StatusCreated, Message: "success", Response: &fiber.Map{"data": result}})
}

func GetCid(c *fiber.Ctx) error {
	if jwtMakerError != nil {return c.Status(http.StatusInternalServerError).JSON(responses.GenericResponse{Status: http.StatusInternalServerError, Message: "Jwt handler error"})}
	bearer := strings.Replace(c.Get("Authorization"), "Bearer ", "", -1)
	_, err := jwtMaker.VerifyToken(bearer)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.GenericResponse{Status: http.StatusUnauthorized, Message: err.Error()})
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	


	cidCode := strings.ToUpper(c.Query("code"))
	var cid models.Cid
	defer cancel()

	err = cidCollection.FindOne(ctx, bson.M{"code": cidCode}).Decode(&cid)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.GenericResponse{Status: http.StatusNotFound, Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(responses.CidResponse{Status: http.StatusOK, Message: "success", Data: cid})
}

func ValidateCid(c *fiber.Ctx) error {
	if jwtMakerError != nil {return c.Status(http.StatusInternalServerError).JSON(responses.GenericResponse{Status: http.StatusInternalServerError, Message: "Jwt handler error"})}
	bearer := strings.Replace(c.Get("Authorization"), "Bearer ", "", -1)
	_, err := jwtMaker.VerifyToken(bearer)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(responses.GenericResponse{Status: http.StatusUnauthorized, Message: err.Error()})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cid models.Cid
	defer cancel()
	cidCode := strings.ToUpper(c.Query("code"))

	if cidCode == "" {
		return c.Status(http.StatusBadRequest).JSON(responses.GenericResponse{Status: http.StatusBadRequest, Message: "Query param 'code' is mandatory"})
	}
	err = cidCollection.FindOne(ctx, bson.M{"code": cidCode}).Decode(&cid)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.GenericResponse{Status: http.StatusNotFound, Message: "Not a valid CID"})
	}

	return c.Status(http.StatusOK).JSON(responses.GenericResponse{Status: http.StatusOK, Message: "Is a valid CID"})
}