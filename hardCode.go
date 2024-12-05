package main

const (
	serviceName              = "CRUDExamApplication"
	Notok                    = "Not-Ok"
	Ok                       = "OK"
	MsgDataSavedSuccessfully = "data saved successfully"
	MsgValidationSuccessfull = "validation successfull"
	MsgDataDeleted           = "data deleted successfully"
	MsgDBConnected           = "connected to MONGO :)"
	MsgDataFound             = "data retrieved successfully"
	ErrDBUriMissing          = "mongo URI not found"
	ErrDBConnection          = "mongo Connection Failed"
	ErrPing                  = "Ping Failed"
	ErrFieldEmpty            = "field cannot be empty"
	ErrInvalidValue          = "invalid value provided"
	ErrDateInvalid           = "invalid Date format (DD-MM-YYYY)"
	ErrDateParadox           = "date paradox"
	ErrMissingField          = "missing field"
	ErrGenderValue           = "Invalid value of gender"
	ErrPincode               = "value out of range should of 6 digits only"
	ErrIdNotPassed           = "_id not passed"
	ErrInvalidId             = "invalid _id passed"
	ErrNotObjId              = "_id could not be converted from Hex"
	ErrDataNotFound          = "data not found"
	ErrJsonNotMarshal        = "json could not be decoded"
	ErrJsonNotUnmarshal      = "data Unmarshal error"
	ErrDataNotSaved          = "data could not be saved in Database"
	ErrNotStruct             = "Not a struct"
)

var MsgCode = map[string]string{
	MsgDataSavedSuccessfully: "INFO001",
	MsgValidationSuccessfull: "INFO002",
	MsgDataFound:             "INFO003",
	MsgDataDeleted:           "INFO004",
	MsgDBConnected:           "INFO005",
	ErrInvalidValue:          "ERR001",
	ErrJsonNotMarshal:        "ERR002",
	ErrDataNotSaved:          "ERR003",
	ErrMissingField:          "ERR004",
	ErrPincode:               "ERR005",
	ErrNotStruct:             "ERR006",
	ErrGenderValue:           "ERR007",
	ErrIdNotPassed:           "ERR008",
	ErrInvalidId:             "ERR009",
	ErrNotObjId:              "ERR010",
	ErrDataNotFound:          "ERR011",
	ErrDateInvalid:           "ERR012",
	ErrDateParadox:           "ERR013",
	ErrDBConnection:          "ERR013",
	ErrPing:                  "ERR014",
	ErrDBUriMissing:          "ERR015",
	ErrJsonNotUnmarshal:      "ERR016",
}
