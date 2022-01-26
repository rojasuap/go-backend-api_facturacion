package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rojasuap/go-backend-api_facturacion/internal/data"
)

// New returns the API V1 Handler with configuration.
func New() http.Handler {
	r := chi.NewRouter()

	ur := &UserRouter{
		Repository: &data.UserRepository{
			Data: data.New(),
		},
	}

	r.Mount("/users", ur.Routes())

	pr := &MedicineRouter{
		Repository: &data.MedicineRepository{
			Data: data.New(),
		},
	}

	r.Mount("/medicines", pr.Routes())

	px := &PromotionRouter{
		Repository: &data.PromotionRepository{
			Data: data.New(),
		},
	}

	r.Mount("/promotions", px.Routes())

	br := &BillRouter{
		Repository: &data.BillRepository{
			Data: data.New(),
		},
	}

	r.Mount("/bills", br.Routes())

	return r
}
