package response

import (
	"github.com/mrefawaladam/server/pkg/entity"
)

type SuplierResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewSuplierResponse(Suplier *entity.Suplier) *SuplierResponse {
	return &SuplierResponse{
		ID:   Suplier.ID,
		Name: Suplier.Name,
	}
}
