package model

import "fmt"
import "time"

type Order struct {
	created time.Time
}
type orderJSON map[string]interface{}

func (this *Order) Init() (error) {
	this.created = time.Now()
	fmt.Printf("Go launched at %s\n", this.created.Local())

	return nil
}

func (order *Order) toJSON() (*orderJSON) {
	/*this.created = time.Now()
	fmt.Printf("Go launched at %s\n", this.created.Local())

	return nil*/
	json := make(orderJSON)

	return &json

}


/*
	failure := map[string]interface{}{
		"success": true,
		"result": data,
	}*/
