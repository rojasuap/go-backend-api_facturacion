package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/promotion"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/response"
)

// PromotionRouter is the router of the promotions.
type PromotionRouter struct {
	Repository promotion.Repository
}

// CreateHandler Create a new promotions.
func (pr *PromotionRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var p promotion.Promotion
	
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
	response.JSON(w, r, http.StatusCreated, response.Map{"promotions": p})
}

// GetAllHandler response all the ps.
func (pr *PromotionRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	promotions, err := pr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"promotions": promotions})
}

// GetOneHandler response one promotion by id.
func (pr *PromotionRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
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

	response.JSON(w, r, http.StatusOK, response.Map{"promotion": p})
}

// UpdateHandler update a stored promotion by id.
func (pr *PromotionRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var p promotion.Promotion
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

// DeleteHandler Remove a promotion by ID.
func (pr *PromotionRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
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

// Routes returns promotion router with each endpoint.
func (pr *PromotionRouter) Routes() http.Handler {
	r := chi.NewRouter()

	//r.Use(middleware.Authorizator)

	//r.Get("/user/{userId}", pr.GetByUserHandler)

	r.Get("/", pr.GetAllHandler)

	r.Post("/", pr.CreateHandler)

	r.Get("/{id}", pr.GetOneHandler)

	r.Put("/{id}", pr.UpdateHandler)

	r.Delete("/{id}", pr.DeleteHandler)

	return r
}
