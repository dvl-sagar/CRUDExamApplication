package main

const serviceName = "CRUDExamApplication"

const (
	Notok = "Not-Ok"
	Ok    = "OK"
)

const (
	errFieldEmpty   = "field cannot be empty"
	ErrInvalidValue = "invalid value provided"
	ErrDateInvalid  = "invalid Date format (DD-MM-YYYY)"
	ErrDateParadox  = "date paradox"
)

var MsgCode = map[string]string{
	"Data saved successfully":             "INFO001",
	"Validation successfull":              "INFO002",
	"Data retrieved successfully":         "INFO003",
	"Data deleted successfully":           "INFO004",
	"Invalid Input Received":              "ERR001",
	"Json could not be decoded":           "ERR002",
	"Data could not be saved in Database": "ERR003",
	"Missing field":                       "ERR004",
	"Value out of range":                  "ERR005",
	"Not a struct":                        "ERR006",
	"Invalid value of gender":             "ERR007",
	"_id not passed":                      "ERR008",
	"Invalid _id passed":                  "ERR009",
	"_id could not be converted from Hex": "ERR010",
	"Student Data not found":              "ERR011",
	"Invalid Date format (DD-MM-YYYY)":    "ERR0012",
	"Date paradox":                        "ERR0013",
}
