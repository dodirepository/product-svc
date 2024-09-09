package handlers

import (
	"net/http"

	render "github.com/dodirepository/common-lib"
	domain "github.com/dodirepository/product-svc/internal/domain/usecases"
	"github.com/gorilla/mux"
)

type ItemHandlersController struct {
	Usecase domain.ItemsInterface
}

func (u ItemHandlersController) ItemList(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	payload := domain.SearchItemList{
		Name: name}
	result, err := u.Usecase.List(r.Context(), payload)
	if err != nil {
		render.Render(err, err.Status, w)
		return
	}

	render.Render(result, http.StatusOK, w)
}
