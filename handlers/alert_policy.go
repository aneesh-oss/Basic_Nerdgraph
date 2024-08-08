package handlers

import (
	"encoding/json"
	"net/http"
	"testGo/services"

	"github.com/gorilla/mux"
)

type AlertPolicyRequest struct {
	Name string `json:"name"`
}

func CreateAlertPolicy(w http.ResponseWriter, r *http.Request) {
	var req AlertPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	policyID, err := services.CreateAlertPolicy(req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"id": policyID})
}

func UpdateAlertPolicy(w http.ResponseWriter, r *http.Request) {
	var req AlertPolicyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	err = services.UpdateAlertPolicy(id, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteAlertPolicy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := services.DeleteAlertPolicy(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func FetchAlertPolicy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	policyDetails, err := services.FetchAlertPolicy(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(policyDetails)
}

// package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"testGo/services"

// 	"github.com/gorilla/mux"
// )

// type AlertPolicyRequest struct {
// 	Name string `json:"name"`
// }

// func CreateAlertPolicy(w http.ResponseWriter, r *http.Request) {
// 	var req AlertPolicyRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	policyID, err := services.CreateAlertPolicy(req.Name)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(map[string]string{"id": policyID})
// }

// func UpdateAlertPolicy(w http.ResponseWriter, r *http.Request) {
// 	var req AlertPolicyRequest
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		http.Error(w, "Invalid request payload", http.StatusBadRequest)
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	err = services.UpdateAlertPolicy(id, req.Name)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func DeleteAlertPolicy(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	err := services.DeleteAlertPolicy(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

// func FetchAlertPolicy(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	policyDetails, err := services.FetchAlertPolicy(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write(policyDetails)
// }
