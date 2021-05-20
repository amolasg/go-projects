package main

import (
	"context"
	"errors"
	"log"
	"net"
	"time"

	"github.com/amolasg/go-projects/microservices-grpc/global"
	proto "github.com/amolasg/go-projects/microservices-grpc/proto/pb"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type authServer struct {
	proto.UnimplementedAuthServiceServer
}

// Login
func (authServer) Login(_ context.Context, in *proto.LoginRequest) (*proto.AuthResponse, error) {
	login, password := in.GetLogin(), in.GetPassword()

	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()

	var user global.User

	global.DB.Collection("user").FindOne(ctx, bson.M{"$or": []bson.M{bson.M{"username": login}, bson.M{"email": login}}}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, errors.New("Wrong login credientials provided")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &proto.AuthResponse{}, errors.New("Wrong login credientials provided")
	}

	return &proto.AuthResponse{Token: user.GetToken()}, nil

}

// Signup
func (authServer) Signup(_ context.Context, in *proto.SignupRequest) (*proto.AuthResponse, error) {

	username, email, password := in.GetUsername(), in.GetEmail(), in.GetPassword()

	if len(username) < 4 || len(username) > 20 || len(email) < 7 || len(email) > 20 || len(password) < 8 || len(password) > 128 {
		return &proto.AuthResponse{}, errors.New("Validation failed")
	}

	return &proto.AuthResponse{}, nil
}

// UsernameUsed
func (authServer) UsernameUsed(_ context.Context, in *proto.UsernameUsedRequest) (*proto.UsedResponse, error) {
	username := in.GetUsername()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()

	var result global.User

	global.DB.Collection("user").FindOne(ctx, bson.M{"username": username}).Decode(&result)

	return &proto.UsedResponse{
		Used: result != global.NilUser,
	}, nil
}

// EmailUsed

func (authServer) EmailUsed(_ context.Context, in *proto.EmailUsedRequest) (*proto.UsedResponse, error) {
	return &proto.UsedResponse{}, nil
}

func main() {
	server := grpc.NewServer()

	proto.RegisterAuthServiceServer(server, authServer{})

	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatal("Error while listing", err.Error())
	}

	server.Serve(listener)

}
