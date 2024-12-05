package main

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TestRegisterValidation(t *testing.T) {
	testCases := []struct {
		testName string
		student  Student
		want     error
	}{
		{
			testName: "Valid",
			student: Student{
				Name:            "John Doe",
				Gender:          "Male",
				DOB:             "01-02-2001",
				StateOfDomicile: "California",
				HomeDistrict:    "Los Angeles",
				FatherName:      "Doe Senior",
				BoardName:       "State Board",
				YearOfPassing:   "2018",
				RollNumber:      "123456",
				Address:         "123 Elm Street",
				HouseNoVillage:  "12A",
				State:           "California",
				District:        "Los Angeles",
				City:            "Los Angeles",
				PinCode:         900001,
			},
			want: errors.New(MsgValidationSuccessfull),
		},
		{
			testName: "Missing Fields",
			student: Student{
				Name:            "",
				Gender:          "Male",
				DOB:             "2000-01-01",
				StateOfDomicile: "California",
				HomeDistrict:    "",
				FatherName:      "Doe Senior",
				BoardName:       "State Board",
				YearOfPassing:   "",
				RollNumber:      "123456",
				Address:         "123 Elm Street",
				HouseNoVillage:  "12A",
				State:           "California",
				District:        "",
				City:            "Los Angeles",
				PinCode:         0,
			},
			want: errors.New(ErrMissingField),
		},
		{
			testName: "Missing PinCode",
			student: Student{
				Name:            "A",
				Gender:          "male",
				DOB:             "2025-12-31", // Future date
				StateOfDomicile: "haryana",
				HomeDistrict:    "12345",
				FatherName:      "N/A",
				BoardName:       "Some Extremely Long Board Name",
				YearOfPassing:   "abcd",
				RollNumber:      "0000",
				Address:         "123",
				HouseNoVillage:  "wazirabad",
				State:           "Unknown",
				District:        "gurugram",
				City:            "gurugram",
				PinCode:         0, // Invalid pin code
			},
			want: errors.New(ErrMissingField),
		},
		{
			testName: "Invalid PinCode",
			student: Student{
				Name:            "A",
				Gender:          "male",
				DOB:             "2025-12-31", // Future date
				StateOfDomicile: "haryana",
				HomeDistrict:    "12345",
				FatherName:      "N/A",
				BoardName:       "ICSE",
				YearOfPassing:   "abcd",
				RollNumber:      "0000",
				Address:         "123",
				HouseNoVillage:  "wazirabad",
				State:           "Unknown",
				District:        "gurugram",
				City:            "gurugram",
				PinCode:         85, // Invalid pin code
			},
			want: errors.New(ErrPincode),
		},
		{
			testName: "Invalid gender",
			student: Student{
				Name:            "A",
				Gender:          "shemale",
				DOB:             "2025-12-31", // Future date
				StateOfDomicile: "haryana",
				HomeDistrict:    "12345",
				FatherName:      "N/A",
				BoardName:       "HSBE",
				YearOfPassing:   "abcd",
				RollNumber:      "0000",
				Address:         "123",
				HouseNoVillage:  "wazirabad",
				State:           "Unknown",
				District:        "gurugram",
				City:            "gurugram",
				PinCode:         858978, // Invalid pin code
			},
			want: errors.New(ErrGenderValue),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			err := RegisterValidation(tt.student)
			if err != nil {
				if err.Error() != tt.want.Error() {
					t.Errorf("\ngot %v\nwant %v", err, tt.want)
				}
			}
		})
	}
}

func TestSetStruct(t *testing.T) {

	today := time.Now()
	Date := today.Format("02-01-2006")

	testCases := []struct {
		name    string
		data    Student
		want    StudentDB
		wantErr error
	}{
		{
			name: "Valid 1",
			data: Student{
				Name:            "John Doe",
				Gender:          "Male",
				DOB:             "15-01-1990",
				StateOfDomicile: "California",
				HomeDistrict:    "Los Angeles",
				FatherName:      "Michael Doe",
				BoardName:       "Central Board of Education",
				YearOfPassing:   "2007",
				RollNumber:      "A12345678",
				Address:         "123 Main Street, Apt 4B",
				HouseNoVillage:  "123",
				State:           "California",
				District:        "Los Angeles",
				City:            "Los Angeles",
				PinCode:         213254,
			},
			want: StudentDB{
				FirstName:       "John",
				LastName:        "Doe",
				Gender:          "MALE",
				DOB:             "15-01-1990",
				StateOfDomicile: "California",
				HomeDistrict:    "Los Angeles",
				FatherFirstName: "Michael",
				FatherLastName:  "Doe",
				BoardName:       "Central Board of Education",
				YearOfPassing:   "2007",
				RollNumber:      "A12345678",
				Address:         "123 Main Street, Apt 4B",
				HouseNoVillage:  "123",
				State:           "California",
				District:        "Los Angeles",
				City:            "Los Angeles",
				PinCode:         213254,
				Age:             34,
			},
			wantErr: nil,
		},
		{
			name: "Invalid Date Format",
			data: Student{
				Name:            "John Doe",
				Gender:          "Male",
				DOB:             "1998-12-05",
				StateOfDomicile: "California",
				HomeDistrict:    "Los Angeles",
				FatherName:      "Michael Doe",
				BoardName:       "Central Board of Education",
				YearOfPassing:   "2007",
				RollNumber:      "A12345678",
				Address:         "123 Main Street, Apt 4B",
				HouseNoVillage:  "123",
				State:           "California",
				District:        "Los Angeles",
				City:            "Los Angeles",
				PinCode:         213254,
			},
			want:    StudentDB{},
			wantErr: errors.New(ErrDateInvalid),
		},
		{
			name: "Date of today",
			data: Student{
				Name:            "John Doe",
				Gender:          "Male",
				DOB:             Date,
				StateOfDomicile: "California",
				HomeDistrict:    "Los Angeles",
				FatherName:      "Michael Doe",
				BoardName:       "Central Board of Education",
				YearOfPassing:   "2007",
				RollNumber:      "A12345678",
				Address:         "123 Main Street, Apt 4B",
				HouseNoVillage:  "123",
				State:           "California",
				District:        "Los Angeles",
				City:            "Los Angeles",
				PinCode:         213254,
			},
			want:    StudentDB{},
			wantErr: errors.New(ErrDateParadox),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetStruct(tt.data)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("got %v want %v", err, tt.wantErr)
				}
			}
			if got != tt.want {
				t.Errorf("\ngot %+v \nwant %+v", got, tt.want)
			}
		})
	}
}

func TestSplitName(t *testing.T) {
	testCases := []struct {
		name       string
		data       string
		firstName  string
		middleName string
		lastName   string
	}{
		{
			name:       "complete name",
			data:       "shubham kumar sharma",
			firstName:  "shubham",
			middleName: "kumar",
			lastName:   "sharma",
		},
		{
			name:       "First and last name only",
			data:       "shubham sharma",
			firstName:  "shubham",
			middleName: "",
			lastName:   "sharma",
		},
		{
			name:       "First name only",
			data:       "shubham",
			firstName:  "shubham",
			middleName: "",
			lastName:   "",
		},
		{
			name:       "empty",
			data:       "",
			firstName:  "",
			middleName: "",
			lastName:   "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			fName, mName, lName := SplitName(tt.data)
			if fName != tt.firstName || mName != tt.middleName || lName != tt.lastName {
				t.Errorf("\ngot %s %s %s \nwant %s %s %s", fName, mName, lName, tt.firstName, tt.middleName, tt.lastName)
			}
		})
	}
}

func TestFetchID(t *testing.T) {
	testCases := []struct {
		name    string
		data    map[string]interface{}
		wantStr string
		wantErr error
	}{
		{
			name: "valid",
			data: map[string]interface{}{
				"_id": "12345",
			},
			wantStr: "12345",
			wantErr: nil,
		},
		{
			name: "missing key _id",
			data: map[string]interface{}{
				"Name": "12345",
			},
			wantStr: "",
			wantErr: errors.New(ErrIdNotPassed),
		},
		{
			name: "_id not string",
			data: map[string]interface{}{
				"_id": 123456,
			},
			wantStr: "",
			wantErr: errors.New(ErrInvalidId),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchID(tt.data)
			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("\ngot %v\nwant %v", err, tt.wantErr)
				}
			}
			if got != tt.wantStr {
				t.Errorf("\ngot %+v \nwant %+v", got, tt.wantStr)
			}
		})
	}
}

func TestPrepareQuery(t *testing.T) {
	today := time.Now()
	date := today.Format("02-01-2006")
	tests := []struct {
		name      string
		input     Student
		want      bson.D
		expectErr error
	}{
		{
			name: "Complete Input",
			input: Student{
				Name:            "John Doe",
				Gender:          "Male",
				DOB:             "21-08-2001",
				StateOfDomicile: "State",
				HomeDistrict:    "District",
				FatherName:      "Robert Doe",
				BoardName:       "CBSE",
				YearOfPassing:   "2018",
				RollNumber:      "123456",
				Address:         "123 Street",
				HouseNoVillage:  "A1",
				State:           "State",
				District:        "District",
				City:            "City",
				PinCode:         123456,
			},
			want: bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "firstName", Value: "John"},
					{Key: "middleName", Value: ""},
					{Key: "lastName", Value: "Doe"},
					{Key: "gender", Value: "Male"},
					{Key: "dob", Value: "21-08-2001"},
					{Key: "age", Value: 23},
					{Key: "stateOfDomicile", Value: "State"},
					{Key: "homeDistrict", Value: "District"},
					{Key: "fatherFirstName", Value: "Robert"},
					{Key: "fatherMiddleName", Value: ""},
					{Key: "fatherLastName", Value: "Doe"},
					{Key: "boardName", Value: "CBSE"},
					{Key: "yearOfPassing", Value: "2018"},
					{Key: "rollNumber", Value: "123456"},
					{Key: "address", Value: "123 Street"},
					{Key: "houseNoVillage", Value: "A1"},
					{Key: "state", Value: "State"},
					{Key: "district", Value: "District"},
					{Key: "city", Value: "City"},
					{Key: "pinCode", Value: 123456},
				}},
			},
			expectErr: nil,
		},
		{
			name: "Invalid DOB",
			input: Student{
				Name: "John Doe",
				DOB:  "invalid-date",
			},
			want: bson.D{
				{Key: "firstName", Value: "John"},
				{Key: "middleName", Value: ""},
				{Key: "lastName", Value: "Doe"},
			},
			expectErr: errors.New(ErrDateInvalid),
		},
		{
			name: "Date Paradox",
			input: Student{
				Name: "John Doe",
				DOB:  date,
			},
			want: bson.D{
				{Key: "firstName", Value: "John"},
				{Key: "middleName", Value: ""},
				{Key: "lastName", Value: "Doe"},
			},
			expectErr: errors.New(ErrDateParadox),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := prepareQuery(tt.input)

			if err != nil {
				if err.Error() != tt.expectErr.Error() {
					t.Errorf("\ngot error %v\nwant error %v", err, tt.expectErr)
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\ngot %v \nwant %v", got, tt.want)
			}
		})
	}
}
