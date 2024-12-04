package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterStudent(t *testing.T) {
	GetConfig()
	ConnectDB()
	testCases := []struct {
		name     string
		data     interface{}
		wantCode int
	}{
		{
			name:     "Invalid JSON",
			data:     `{"name": "John"`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Valid",
			data: map[string]interface{}{
				"name":            "John Doe",
				"gender":          "Male",
				"dob":             "01-08-2001",
				"stateOfDomicile": "Karnataka",
				"homeDistrict":    "Bangalore",
				"fatherName":      "Robert Doe",
				"boardName":       "ICSE",
				"yearOfPassing":   "2020",
				"rollNumber":      "ICSE123456",
				"address":         "123 Main Street, Oakwood",
				"houseNoVillage":  "Flat 101, Maple Residency",
				"state":           "Karnataka",
				"district":        "Bangalore Urban",
				"city":            "Bangalore",
				"pinCode":         560001,
			},
			wantCode: http.StatusOK,
		},
		{
			name: "Missing field",
			data: map[string]interface{}{
				"name":            "",
				"gender":          "Male",
				"dob":             "21-08-2001",
				"stateOfDomicile": "Karnataka",
				"homeDistrict":    "Bangalore",
				"fatherName":      "Robert Doe",
				"boardName":       "ICSE",
				"yearOfPassing":   "2020",
				"rollNumber":      "ICSE123456",
				"address":         "123 Main Street, Oakwood",
				"houseNoVillage":  "Flat 101, Maple Residency",
				"state":           "Karnataka",
				"district":        "Bangalore Urban",
				"city":            "Bangalore",
				"pinCode":         560001,
			},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// converting data to JSON
			inputJson, err := json.Marshal(tt.data)
			if err != nil {
				t.Errorf("failed to marshal input: %v", err)
			}

			// Creating HTTP POST request
			r := httptest.NewRequest("POST", "/student/create", bytes.NewBuffer(inputJson))
			w := httptest.NewRecorder()

			RegisterStudent(w, r)
			if w.Code != tt.wantCode {
				t.Errorf("\ngot status %v \nwant status %v", w.Code, tt.wantCode)
			}
		})
	}
}

func TestGetStudent(t *testing.T) {
	GetConfig()
	ConnectDB()
	testCases := []struct {
		name     string
		data     interface{}
		wantCode int
	}{
		{
			name:     "Invalid JSON",
			data:     `{"name": "John"`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Missing _id",
			data: map[string]interface{}{
				"name": "dororo",
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "Invalid ObjectID",
			data: map[string]interface{}{
				"_id": "dororo",
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "_id not present in DB",
			data: map[string]interface{}{
				"_id": "67504e69865b091240c08cb7",
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "valid",
			data: map[string]interface{}{
				"_id": "67504e69865b091240c08cb6",
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// converting data to JSON
			inputJson, err := json.Marshal(tt.data)
			if err != nil {
				t.Errorf("failed to marshal input: %v", err)
			}

			// Creating HTTP POST request
			r := httptest.NewRequest("POST", "/student/get", bytes.NewBuffer(inputJson))
			w := httptest.NewRecorder()

			GetStudent(w, r)
			if w.Code != tt.wantCode {
				t.Errorf("\ngot status %v \nwant status %v", w.Code, tt.wantCode)
			}
		})
	}
}
