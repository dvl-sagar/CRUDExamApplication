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
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[ErrJsonNotMarshal],
			Status:      Notok,
			Msg:         ErrJsonNotMarshal}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	err = RegisterValidation(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[err.Error()],
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
			MessageCode: MsgCode[err.Error()],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode[MsgDataSavedSuccessfully],
		Status:      Ok,
		Msg:         MsgDataSavedSuccessfully,
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
			MessageCode: MsgCode[ErrJsonNotMarshal],
			Status:      Notok,
			Msg:         ErrJsonNotMarshal}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	id, err := fetchID(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[err.Error()],
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
			MessageCode: MsgCode[ErrNotObjId],
			Status:      Notok,
			Msg:         ErrNotObjId}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	std, err := getStudentByID(objID)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[ErrDataNotFound],
			Status:      Notok,
			Msg:         ErrDataNotFound}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode[MsgDataFound],
		Status:      Ok,
		Msg:         MsgDataFound,
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
			MessageCode: MsgCode[ErrJsonNotMarshal],
			Status:      Notok,
			Msg:         ErrJsonNotMarshal}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	if student.Id == "" {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[ErrIdNotPassed],
			Status:      Notok,
			Msg:         ErrIdNotPassed}
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
			MessageCode: MsgCode[ErrNotObjId],
			Status:      Notok,
			Msg:         ErrNotObjId}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	result, err := updateStudentByID(objID, student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[err.Error()],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode[MsgDataSavedSuccessfully],
		Status:      Ok,
		Msg:         MsgDataSavedSuccessfully,
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
			MessageCode: MsgCode[ErrJsonNotMarshal],
			Status:      Notok,
			Msg:         ErrJsonNotMarshal}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	id, err := fetchID(student)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[err.Error()],
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
			MessageCode: MsgCode[ErrNotObjId],
			Status:      Notok,
			Msg:         ErrNotObjId}
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}

	result, err := deleteStudentByID(objID)
	if err != nil {
		res := Response{
			ServiceName: serviceName,
			MessageCode: MsgCode[ErrDataNotFound],
			Status:      Notok,
			Msg:         err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		r, _ := json.Marshal(res)
		w.Write(r)
		return
	}
	res := Response{
		ServiceName: serviceName,
		MessageCode: MsgCode[MsgDataDeleted],
		Status:      Ok,
		Msg:         MsgDataDeleted,
		Data:        result}
	w.WriteHeader(http.StatusOK)
	rd, _ := json.Marshal(res)
	w.Write(rd)
}
