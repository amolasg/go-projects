package main

import (
	"context"
	"testing"

	"github.com/amolasg/go-projects/microservices-grpc/global"
	proto "github.com/amolasg/go-projects/microservices-grpc/proto/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Test_authServer_Login(t *testing.T) {
	global.ConnectToTestDB()
	pw, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	global.DB.Collection("user").InsertOne(context.Background(), global.User{ID: primitive.NewObjectID(), Email: "amol.asg@gmail.com", Username: "amol", Password: string(pw)})

	server := authServer{}
	_, err := server.Login(context.Background(), &proto.LoginRequest{Login: "amol.asg@gmail.com", Password: "password"})

	if err != nil {

		t.Error("1.An error returned: ", err.Error())
	}

	_, err = server.Login(context.Background(), &proto.LoginRequest{Login: "something", Password: "something"})

	if err == nil {
		t.Error("2.Error is nil")
	}
	_, err = server.Login(context.Background(), &proto.LoginRequest{Login: "amol", Password: "password"})

	if err != nil {
		t.Error("3.An error returned: ", err.Error())

	}
}
