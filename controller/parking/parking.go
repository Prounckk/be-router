package controller

import (
	_ "be-router/entity"
	model "be-router/model/parking"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	ErrorCode string = "4622"
)

// GetParkingByName returns information about parking spot rate and schedule
func GetParkingByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]
	parking, err := model.GetParkingByName(number)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return

	}
	response, err := json.Marshal(parking)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error with your request, please try again. code: " + ErrorCode))
		return
	}
	w.Write(response)
}

//GetParkingStatusByName returns information about the parking spot with information about the booking
func GetParkingStatusByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]
	parking, err := model.GetParkingByName(number)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return

	}
	response, err := json.Marshal(parking.Status)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error with your request, please try again. code: " + ErrorCode))
		return
	}
	w.Write(response)
}

func GetParkingTimerByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number := vars["number"]
	p, err := model.GetParkingByName(number)
	if err != nil {
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}

	if p.Status == model.Paid {
		response, err := calculateExpiration(p.EndTime)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error with your request, please try again. code: " + ErrorCode))
			//The app can continue working, but this one is very specific and critical for the app
			log.Printf("calculateExpiration() failed with '%s'\n\n", err)
			return
		}
		w.Write([]byte(response))
		return
	}

	response, err := json.Marshal(p.Status)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("Error with your request, please try again. code: " + ErrorCode))
		return
	}
	w.Write(response)

}

func calculateExpiration(s string) (string, error) {
	// fail first, if the string is empty
	if s == "" {
		return "", errors.New("empty parking.EndTime provided, code: " + ErrorCode)
	}

	end, err := time.Parse(time.Kitchen, s)
	if err != nil {
		return "", err
	}
	now := time.Now()
	durationNow, err := time.ParseDuration(now.Format("15h04m05s"))
	if err != nil {
		return "", err
	}
	durationEnd, err := time.ParseDuration(end.Format("15h04m05s"))
	if err != nil {
		return "", err
	}
	actualDuration := durationEnd - durationNow
	res := fmt.Sprintf("Your's parking spot expired %s ago", actualDuration)

	if actualDuration > 0 {
		res = fmt.Sprintf("You've got %s of parking time left.", actualDuration)
	}

	return res, nil
}
