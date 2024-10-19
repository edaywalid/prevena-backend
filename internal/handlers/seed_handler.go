package handlers

import (
	"net/http"

	"github.com/edaywalid/pinktober-hackathon-backend/internal/seed"
)

type SeedHandler struct {
	seed *seed.Seeder
}

func NewSeedHandler(
	seed *seed.Seeder,
) *SeedHandler {
	return &SeedHandler{
		seed: seed,
	}
}

func (sh *SeedHandler) Seed(w http.ResponseWriter, r *http.Request) {
	err := sh.seed.SeedAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Seed completed"))
}
