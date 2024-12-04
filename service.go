package main

import (
	"context"
	"fmt"

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
		fmt.Println("Document not found")
		return nil, err
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
		fmt.Println("Student Data not found")
		return nil, err
	}
	return resp, nil
}

func deleteStudentByID(objID primitive.ObjectID) (any, error) {

	Applicant := GetCollection("ExamApplication")
	deleted, err := Applicant.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		fmt.Println("Student Data not found")
		return nil, err
	}

	return deleted, nil
}
