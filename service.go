package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// creating new user
func RegisterStudentService(student Student) (any, error) {
	response := make(map[string]interface{})
	studentDB, err := SetStruct(student)
	if err != nil {
		return nil, err
	}
	Applicant := GetCollection("ExamApplication")
	result, err := Applicant.InsertOne(context.Background(), studentDB)
	if err != nil {
		return nil, err
	}
	response["_id"] = result.InsertedID
	return response, nil
}

func getStudentByID(objID primitive.ObjectID) (any, error) {
	var oneStudent StudentDB
	Applicant := GetCollection("ExamApplication")
	err := Applicant.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&oneStudent)
	if err != nil {
		return nil, errors.New(ErrJsonNotUnmarshal)
	}
	if oneStudent.Id.String() == "" {
		return nil, errors.New(ErrDataNotFound)
	}
	return &oneStudent, nil
}

func updateStudentByID(objID primitive.ObjectID, std Student) (any, error) {

	var resp StudentDB
	finalQ, err := prepareQuery(std)
	if err != nil {
		return nil, err
	}
	Applicant := GetCollection("ExamApplication")
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = Applicant.FindOneAndUpdate(context.Background(), bson.M{"_id": objID}, finalQ, opt).Decode(&resp)
	if err != nil {
		return nil, errors.New(ErrJsonNotMarshal)
	}
	if resp.Id.String() == "" {
		return nil, errors.New(ErrDataNotFound)
	}
	return resp, nil
}

func deleteStudentByID(objID primitive.ObjectID) (any, error) {
	var oneStudent StudentDB
	Applicant := GetCollection("ExamApplication")
	err := Applicant.FindOneAndDelete(context.Background(), bson.M{"_id": objID}).Decode(&oneStudent)
	if err != nil {
		return nil, errors.New(ErrJsonNotMarshal)
	}
	if oneStudent.Id.String() == "" {
		return nil, errors.New(ErrDataNotFound)
	}
	return &oneStudent, nil
}
