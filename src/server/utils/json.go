package utils

import (
	"encoding/json"
	"net/http"
	//"go-shop/src/server/utils"
)

func WriteJSON(w *http.ResponseWriter, datas *map[string]interface{}) {
	(*w).Header().Add("Content-Type", "application/json")
	(*w).Header().Add("Access-Control-Allow-Origin", "TODO")

	if datas != nil {
		j, err := json.Marshal(datas)
		if err == nil {
			(*w).Write(j)
			return
		}
	}
}

func JSONError(w http.ResponseWriter, e error) {
	failure := map[string]interface{}{
		"success": false,
		"error":   e.Error(),
	}
	WriteJSON(&w, &failure)
}

func JSONSuccess(w http.ResponseWriter, data map[string]interface{}) {
	failure := map[string]interface{}{
		"success": true,
		"result":  data,
	}
	WriteJSON(&w, &failure)
}
