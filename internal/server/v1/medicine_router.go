package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/medicine"
	"github.com/orlmonteverde/go-postgres-microblog/pkg/response"
)

// PostRouter is the router of the posts.
type MedicineRouter struct {
	Repository medicine.Repository
}

// CreateHandler Create a new post.
func (pr *MedicineRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {

	var p medicine.Medicine
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

	response.JSON(w, r, http.StatusCreated, response.Map{"medicine": p})
}

// GetAllHandler response all the posts.
func (pr *MedicineRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	posts, err := pr.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"medicines": posts})
}

// GetOneHandler response one post by id.
func (pr *MedicineRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {

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

	response.JSON(w, r, http.StatusOK, response.Map{"medicine": p})
}

// UpdateHandler update a stored post by id.
func (pr *MedicineRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var p medicine.Medicine
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

// DeleteHandler Remove a post by ID.
func (pr *MedicineRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {

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

// Routes returns post router with each endpoint.
func (pr *MedicineRouter) Routes() http.Handler {
	r := chi.NewRouter()

	//r.Use(middleware.Authorizator)

	r.Get("/", pr.GetAllHandler)

	r.Post("/", pr.CreateHandler)

	r.Get("/{id}", pr.GetOneHandler)

	r.Put("/{id}", pr.UpdateHandler)

	r.Delete("/{id}", pr.DeleteHandler)

	return r
}
