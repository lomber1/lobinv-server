package units

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func Get(w http.ResponseWriter, r *http.Request) {
	units := GetAllUnits()

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(units)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Post(w http.ResponseWriter, r *http.Request) {
	var body UnitDTO
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unit, err := CreateUnit(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(unit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Put(w http.ResponseWriter, r *http.Request) {
	var body UnitDTO
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unitID := chi.URLParam(r, "unitID")
	id, err := strconv.Atoi(unitID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	unit, err := UpdateUnit(id, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(unit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	unitID := chi.URLParam(r, "unitID")
	id, err := strconv.Atoi(unitID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = RemoveUnit(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
