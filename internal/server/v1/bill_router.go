package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rojasuap/go-backend-api_facturacion/pkg/bill"
	"github.com/rojasuap/go-backend-api_facturacion/pkg/response"
)

// PromotionRouter is the router of the promotions.
type BillRouter struct {
	Repository bill.Repository
}

// CreateHandler Create a new bills.
func (pr *BillRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var p bill.Bill
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Create(ctx, &p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), p.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"bills": p})
}

// RootHandler - Returns all the available APIs
// @Summary This API can be used as health check for this application.
// @Description Tells if the chi-swagger APIs are working or not.
// @Tags chi-swagger
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router / [get]
func (pr *BillRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bills, err := pr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"bills": bills})
}

// GetOneHandler response one bills by id.
func (pr *BillRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	p, err := pr.Repository.GetOne(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"bill": p})
}

// GetAllHandler response all the bills.
func (pr *BillRouter) GetAllDaysHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bills, err := pr.Repository.GetAllDays(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"bills": bills})
}

// GetAllHandler response all the bills.
func (pr *BillRouter) GetAllPaymentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	bills, err := pr.Repository.GetAllPayments(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"bills": bills})
}

// UpdateHandler update a stored bill by id.
func (pr *BillRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var p bill.Bill
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = pr.Repository.Update(ctx, uint(id), p)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, nil)
}

// DeleteHandler Remove a bill by ID.
func (pr *BillRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = pr.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{})
}

// GetByUserHandler response promotions by user id.
/**func (pr *PostRouter) GetByUserHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := chi.URLParam(r, "userId")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	promotions, err := pr.Repository.GetByUser(ctx, uint(userID))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"promotions": promotions})
}*/

// Routes returns bill router with each endpoint.
func (pr *BillRouter) Routes() http.Handler {
	r := chi.NewRouter()

	//r.Use(middleware.Authorizator)

	//r.Get("/user/{userId}", pr.GetByUserHandler)

	r.Get("/", pr.GetAllHandler)

	r.Post("/", pr.CreateHandler)

	r.Get("/{id}", pr.GetOneHandler)

	r.Get("/dias", pr.GetAllDaysHandler)

	r.Get("/totales", pr.GetAllPaymentsHandler)

	r.Put("/{id}", pr.UpdateHandler)

	r.Delete("/{id}", pr.DeleteHandler)

	return r
}
