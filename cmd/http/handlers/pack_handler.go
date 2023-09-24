package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"repartners/internal/pack/repository"
	"repartners/internal/pack/services"
	"strconv"
)

type PackHandler struct {
	Repo repository.Repository
}

func (h *PackHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}
	value := r.FormValue("pack_size")

	i, err := strconv.Atoi(value)

	if err != nil {
		http.Error(w, "wrong input for pack size: "+value, http.StatusBadRequest)
		return
	}

	packs, err := h.Repo.GetAll()

	if err != nil {
		http.Error(w, "internal error processing the request", http.StatusInternalServerError)
		log.Println("error getting the pack data:" + err.Error())
		return
	}

	packSizes := services.Calculate(i, packs)

	err = json.NewEncoder(w).Encode(packSizes)

	if err != nil {
		http.Error(w, "wrong input for pack size: "+value, http.StatusInternalServerError)
		log.Println("error generating response:" + err.Error())
		return
	}
}
