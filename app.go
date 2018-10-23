package main 

import (
		"encoding/json"
		"gopkg.in/mgo.v2/bson"
		"net/http"
		"github.com/gorilla/mux"
		."DcmStatusReceiver/models"
		."DcmStatusReceiver/dao"
		."DcmStatusReceiver/config"

)

var config = Config{}
var dao = DeviceDAO{}

func DcmStatus(w http.ResponseWriter, r *http.Request) {
	
	defer r.Body.Close()
	var device Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
	device.ID = bson.NewObjectId()
	
	if err := dao.Insert(device); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, device)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main(){
	r := mux.NewRouter();
	r.HandleFunc("/dcm/status",DcmStatus).Methods("POST")
	http.ListenAndServe(":3000", r);
}