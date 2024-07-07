package errors

import (
	//"net/http"
	//"go-shop/src/server/utils"
)

type NotFoundError struct{}

func (e NotFoundError) Error() string {
	return "Not found"
}
