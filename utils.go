package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetConfig() {
	mongoUri := "mongodb+srv://admin:admin123@cluster0.gfsyk.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	if mongoUri == "" {
		log.Fatal("Database URI not found")
	}
	Config.MongoUri = mongoUri
	fmt.Println(Config)
}

var Client *mongo.Client

func ConnectDB() {
	// connecting to MONGODB server
	var err error
	Client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(Config.MongoUri))
	if err != nil {
		log.Fatal("Mongo Connection Failed", err)
	}

	err = Client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Ping failed ", err)
	}

	fmt.Println("Connected to MongoDB!")
}

func GetCollection(name string) *mongo.Collection {
	if Client == nil {
		ConnectDB()
	}
	return Client.Database("CRUDExamApplication").Collection(name)
}

func AgeCalc(dobstr string) (int, error) {
	layout := "02-01-2006"

	date, err := time.Parse(layout, dobstr)
	if err != nil {
		return 0, errors.New(ErrDateInvalid)

	}

	today := time.Now().Truncate(24 * time.Hour)
	if date.After(today) || date.Equal(today) {
		return 0, errors.New(ErrDateParadox)
	}

	currentTime := time.Now()

	yearsDif := currentTime.Year() - date.Year()
	return yearsDif, nil
}

func SplitName(fullname string) (string, string, string) {
	firstname := ""
	middleName := ""
	lastName := ""

	if fullname == "" {
		return firstname, middleName, lastName
	}

	sName := strings.Split(fullname, " ")
	if len(sName) >= 3 {
		firstname = sName[0]
		middleName = sName[1]
		lastName = sName[2]
	} else if len(sName) == 2 {
		firstname = sName[0]
		lastName = sName[1]
	} else {
		firstname = sName[0]
	}

	return firstname, middleName, lastName
}

func RegisterValidation(s Student) (string, error) {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		// Skip the Id field
		if fieldName == "Id" {
			continue
		}

		// Check for empty fields
		if field.Kind() == reflect.String {
			if strings.TrimSpace(field.String()) == "" {
				return "Missing field", fmt.Errorf("field %s is empty", fieldName)
			}
			if fieldName == "Gender" {
				if !strings.EqualFold(field.String(), "male") && !strings.EqualFold(field.String(), "female") {
					return "Invalid value of gender", fmt.Errorf("field %s must be of either 'male' or 'female' case insensitive", fieldName)
				}
			}

		} else if field.Kind() == reflect.Int {
			if field.Int() == 0 {
				return "Missing field", fmt.Errorf("field %s is empty", fieldName)
			} else if field.Int() > 999999 || field.Int() < 100000 {
				return "Value out of range", fmt.Errorf("%s must be of six digits", fieldName)
			}
		}
	}
	return "Validation successfull", nil
}

func SetStruct(student Student) (StudentDB, error) {
	age, err := AgeCalc(student.DOB)
	if err != nil {
		return StudentDB{}, err
	}
	firstName, middleName, lastName := SplitName(student.Name)
	fFirstName, fMiddleName, fLastName := SplitName(student.FatherName)

	return StudentDB{
		FirstName:        firstName,
		MiddleName:       middleName,
		LastName:         lastName,
		Gender:           strings.ToUpper(student.Gender),
		DOB:              student.DOB,
		StateOfDomicile:  student.StateOfDomicile,
		HomeDistrict:     student.HomeDistrict,
		FatherFirstName:  fFirstName,
		FatherMiddleName: fMiddleName,
		FatherLastName:   fLastName,
		BoardName:        student.BoardName,
		YearOfPassing:    student.YearOfPassing,
		RollNumber:       student.RollNumber,
		Address:          student.Address,
		HouseNoVillage:   student.HouseNoVillage,
		State:            student.State,
		District:         student.District,
		City:             student.City,
		PinCode:          student.PinCode,
		Age:              age}, nil
}

func fetchID(student map[string]interface{}) (string, error) {
	idMap, ok := student["_id"]
	if !ok {
		return "_id not passed", errors.New("_id is not present")
	}
	id, ok := idMap.(string)
	if !ok {
		return "Invalid _id passed", errors.New("_id is not a string")
	}
	return id, nil
}

func prepareQuery(std Student) (bson.D, error) {
	query := bson.D{}
	flag := false
	if std.Name != "" {
		fName, mName, lName := SplitName(std.Name)
		query = append(query, primitive.E{Key: "firstName", Value: fName})
		query = append(query, primitive.E{Key: "middleName", Value: mName})
		query = append(query, primitive.E{Key: "lastName", Value: lName})
		flag = true
	}
	if std.Gender != "" {
		query = append(query, primitive.E{Key: "address", Value: std.Address})
		flag = true
	}
	if std.DOB != "" {
		age, err := AgeCalc(std.DOB)
		if err != nil {
			return query, err
		}
		query = append(query, primitive.E{Key: "dob", Value: std.DOB})
		query = append(query, primitive.E{Key: "age", Value: age})
		flag = true
	}

	if std.StateOfDomicile != "" {
		query = append(query, primitive.E{Key: "stateOfDomicile", Value: std.StateOfDomicile})
		flag = true
	}

	if std.HomeDistrict != "" {
		query = append(query, primitive.E{Key: "homeDistrict", Value: std.HomeDistrict})
		flag = true
	}

	if std.FatherName != "" {
		fName, mName, lName := SplitName(std.FatherName)
		query = append(query, primitive.E{Key: "fatherFirstName", Value: fName})
		query = append(query, primitive.E{Key: "fatherMiddleName", Value: mName})
		query = append(query, primitive.E{Key: "fatherLastName", Value: lName})
		flag = true
	}

	if std.BoardName != "" {
		query = append(query, primitive.E{Key: "boardName", Value: std.BoardName})
		flag = true
	}

	if std.YearOfPassing != "" {
		query = append(query, primitive.E{Key: "yearOfPassing", Value: std.YearOfPassing})
		flag = true
	}

	if std.RollNumber != "" {
		query = append(query, primitive.E{Key: "rollNumber", Value: std.RollNumber})
		flag = true
	}

	if std.Address != "" {
		query = append(query, primitive.E{Key: "address", Value: std.Address})
		flag = true
	}

	if std.HouseNoVillage != "" {
		query = append(query, primitive.E{Key: "houseNoVillage", Value: std.HouseNoVillage})
		flag = true
	}

	if std.State != "" {
		query = append(query, primitive.E{Key: "state", Value: std.State})
		flag = true
	}

	if std.District != "" {
		query = append(query, primitive.E{Key: "district", Value: std.District})
		flag = true
	}

	if std.City != "" {
		query = append(query, primitive.E{Key: "city", Value: std.City})
		flag = true
	}

	if std.PinCode != 0 {
		query = append(query, primitive.E{Key: "pinCode", Value: std.PinCode})
		flag = true
	}

	if flag {
		query = append(query, primitive.E{Key: "updated_at", Value: time.Now()})
	}
	update := bson.D{primitive.E{Key: "$set", Value: query}}
	return update, nil
}
