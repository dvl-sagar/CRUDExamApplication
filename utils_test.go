package main

import (
	"errors"
	"testing"
)

func TestRegisterValidation(t *testing.T) {
	testCases := []struct {
		testName string
		student  Student
		want     string
		wantRes  string
	}{
		{
			testName: "Valid",
			student: Student{
				Name:            "John Doe",
				Gender:          "Male",
				DOB:             "2000-01-01",
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
			want:    "",
			wantRes: "Validation successfull",
		},
		{
			testName: "Missing Required Fields",
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
			want:    "field Name is empty",
			wantRes: "Missing field",
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
				BoardName:       "Some Extremely Long Board Name That Might Be Truncated",
				YearOfPassing:   "abcd",
				RollNumber:      "0000",
				Address:         "123",
				HouseNoVillage:  "wazirabad",
				State:           "Unknown",
				District:        "gurugram",
				City:            "gurugram",
				PinCode:         0, // Invalid pin code
			},
			want:    "field PinCode is empty",
			wantRes: "Missing field",
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
				BoardName:       "Some Extremely Long Board Name That Might Be Truncated",
				YearOfPassing:   "abcd",
				RollNumber:      "0000",
				Address:         "123",
				HouseNoVillage:  "wazirabad",
				State:           "Unknown",
				District:        "gurugram",
				City:            "gurugram",
				PinCode:         85, // Invalid pin code
			},
			want:    "PinCode must be of six digits",
			wantRes: "Value out of range",
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
				BoardName:       "Some Extremely Long Board Name That Might Be Truncated",
				YearOfPassing:   "abcd",
				RollNumber:      "0000",
				Address:         "123",
				HouseNoVillage:  "wazirabad",
				State:           "Unknown",
				District:        "gurugram",
				City:            "gurugram",
				PinCode:         858978, // Invalid pin code
			},
			want:    "field Gender must be of either 'male' or 'female' case insensitive",
			wantRes: "Invalid value of gender",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			valRes, err := RegisterValidation(tt.student)
			if err != nil && err.Error() != tt.want {
				t.Errorf("RegisterValidation() error = %v, want %v", err.Error(), tt.want)
			}
			if valRes != tt.wantRes {
				t.Errorf("got message code %s want %s", valRes, tt.want)
			}
		})
	}
}

func TestSetStruct(t *testing.T) {
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
				DOB:             "04-12-2024",
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
			wantStr: "_id not passed",
			wantErr: errors.New("_id is not present"),
		},
		{
			name: "_id not string",
			data: map[string]interface{}{
				"_id": 123456,
			},
			wantStr: "Invalid _id passed",
			wantErr: errors.New("_id is not a string"),
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