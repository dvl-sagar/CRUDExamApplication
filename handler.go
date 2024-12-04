package main

import (
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	// fmt.Println(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Json could not be decoded"],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	valRes, err := RegisterValidation(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[valRes],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	resp, err := RegisterStudentService(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Data could not saved in Database"],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode["Data saved successfully"],
		Status:      Ok,
		Msg:         "Operation Successfull",
		Data:        resp}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Json could not be decoded"],
			Status:      Notok,
			Msg:         "Invalid"}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	id, err := fetchID(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[id],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	objID, err := primitive.ObjectIDFromHex(id)
	log.Println("ObjID : ", objID)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["_id could not be converted from Hex"],
			Status:      Notok,
			Msg:         "could not convert to the ObjectID"}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	std, err := getStudentByID(objID)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Student Data not found"],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode["Data retrieved successfully"],
		Status:      Ok,
		Msg:         "Operation Successfully",
		Data:        std}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Json could not be decoded"],
			Status:      Notok,
			Msg:         "Invalid"}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	if student.Id == "" {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["_id not passed"],
			Status:      Notok,
			Msg:         "Id cannot be empty"}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	objID, err := primitive.ObjectIDFromHex(student.Id)
	log.Println("ObjID : ", objID)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["_id could not be converted from Hex"],
			Status:      Notok,
			Msg:         "could not convert to the ObjectID"}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	result, err := updateStudentByID(objID, student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Data could not be saved in Database"],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode["Data saved successfully"],
		Status:      Ok,
		Msg:         "Operation Successfully",
		Data:        result}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var student map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Json could not be decoded"],
			Status:      Notok,
			Msg:         "Invalid"}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	id, err := fetchID(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[id],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	log.Println("ObjID : ", objID)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["_id could not be converted from Hex"],
			Status:      Notok,
			Msg:         "could not convert to the ObjectID"}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	result, err := deleteStudentByID(objID)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode["Student Data not found"],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode["Data deleted successfully"],
		Status:      Ok,
		Msg:         "Operation Successfully",
		Data:        result}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}
