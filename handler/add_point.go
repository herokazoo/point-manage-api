package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/herokazoo/point-manage-api/entity"
	"github.com/herokazoo/point-manage-api/store"
)

type AddPoint struct {
	Store     *store.HistoryStore
	Validator *validator.Validate
}

func (ap *AddPoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		UserPointID int `json:"user_point_id" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	err := ap.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}
	h := &entity.UserPointHistory{
		UserPointID: entity.UserPointID(b.UserPointID),
		Created:     time.Now(),
	}
	id, err := store.Add(h)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID int `json:"id"`
	}{ID: id}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
