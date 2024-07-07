package handlers

import (
	"net/http"

	"github.com/baldurstod/go-shop/src/server/server/errors"
	"github.com/baldurstod/go-shop/src/server/server/model"
	"github.com/baldurstod/go-shop/src/server/utils"
)

type OrdersHandler struct {
}

func (oh *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/orders/create":
		CreateOrder(w, r)
	default:
		utils.JSONError(w, errors.NotFoundError{})
	}
}

func CreateOrder(w http.ResponseWriter, r *http.Request) error {
	//fmt.Print("CreateOrder")
	//panic(0)

	order := model.Order{}
	success := map[string]interface{}{
		"result": order.Init(),
	}

	utils.JSONSuccess(w, success)
	return nil
}
