package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentDB struct {
	Id               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName        string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	MiddleName       string             `json:"middleName,omitempty" bson:"middleName,omitempty"`
	LastName         string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Gender           string             `json:"gender,omitempty" bson:"gender,omitempty"`
	DOB              string             `json:"dob,omitempty" bson:"dob,omitempty"`
	StateOfDomicile  string             `json:"stateOfDomicile,omitempty" bson:"stateOfDomicile,omitempty"`
	HomeDistrict     string             `json:"homeDistrict,omitempty" bson:"homeDistrict,omitempty"`
	FatherFirstName  string             `json:"fatherFirstName,omitempty" bson:"fatherFirstName,omitempty"`
	FatherMiddleName string             `json:"fatherMiddleName,omitempty" bson:"fatherMiddleName,omitempty"`
	FatherLastName   string             `json:"fatherLastName,omitempty" bson:"fatherLastName,omitempty"`
	BoardName        string             `json:"boardName,omitempty" bson:"boardName,omitempty"`
	YearOfPassing    string             `json:"yearOfPassing,omitempty" bson:"yearOfPassing,omitempty"`
	RollNumber       string             `json:"rollNumber,omitempty" bson:"rollNumber,omitempty"`
	Address          string             `json:"address,omitempty" bson:"address,omitempty"`
	HouseNoVillage   string             `json:"houseNoVillage,omitempty" bson:"houseNoVillage,omitempty"`
	State            string             `json:"state,omitempty" bson:"state,omitempty"`
	District         string             `json:"district,omitempty" bson:"district,omitempty"`
	City             string             `json:"city,omitempty" bson:"city,omitempty"`
	PinCode          int                `json:"pinCode,omitempty" bson:"pinCode,omitempty"`
	Age              int                `json:"age,omitempty" bson:"age,omitempty"`
}

type Student struct {
	Id              string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name            string `json:"name,omitempty" bson:"name,omitempty"`
	Gender          string `json:"gender,omitempty" bson:"gender,omitempty"`
	DOB             string `json:"dob,omitempty" bson:"dob,omitempty"`
	StateOfDomicile string `json:"stateOfDomicile,omitempty" bson:"stateOfDomicile,omitempty"`
	HomeDistrict    string `json:"homeDistrict,omitempty" bson:"homeDistrict,omitempty"`
	FatherName      string `json:"fatherName,omitempty" bson:"fatherName,omitempty"`
	BoardName       string `json:"boardName,omitempty" bson:"boardName,omitempty"`
	YearOfPassing   string `json:"yearOfPassing,omitempty" bson:"yearOfPassing,omitempty"`
	RollNumber      string `json:"rollNumber,omitempty" bson:"rollNumber,omitempty"`
	Address         string `json:"address,omitempty" bson:"address,omitempty"`
	HouseNoVillage  string `json:"houseNoVillage,omitempty" bson:"houseNoVillage,omitempty"`
	State           string `json:"state,omitempty" bson:"state,omitempty"`
	District        string `json:"district,omitempty" bson:"district,omitempty"`
	City            string `json:"city,omitempty" bson:"city,omitempty"`
	PinCode         int    `json:"pinCode,omitempty" bson:"pinCode,omitempty"`
}

type Response struct {
	ServiceName string      `json:"serviceName,omitempty"`
	MessageCode string      `json:"messageCode,omitempty"`
	Status      string      `json:"status,omitempty"`
	Msg         string      `json:"msg,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
