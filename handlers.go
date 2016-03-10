package orders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func parentKey(c context.Context, kind string) *datastore.Key {
	return datastore.NewKey(c, kind, "parent_order", 0, nil)
}

func HandleError(w http.ResponseWriter, e error) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)

		panic(e)
	}
}

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
}

type OrderObjects struct {
	Objects    Orders     `json:"objects"`
	Pagination Pagination `json:"pagination"`
}

func OrderIndex(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("OrderTest").Ancestor(parentKey(c, "OrderTest"))
	var orders Orders

	_, err := q.GetAll(c, &orders)
	HandleError(w, err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// TODO: Hardcoded please change this
	pagination := Pagination{
		Limit:  len(orders),
		Offset: 0,
		Count:  len(orders),
	}
	log.Infof(c, "%v", pagination)
	order_objects := OrderObjects{
		Objects:    orders,
		Pagination: pagination,
	}

	err = json.NewEncoder(w).Encode(order_objects)
	HandleError(w, err)
}

func LoadFixtures(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var orders Orders
	dat, err := ioutil.ReadFile("fixtures/orders.json")
	HandleError(w, err)

	err = json.Unmarshal(dat, &orders)
	HandleError(w, err)

	for _, order := range orders {
		log.Infof(c, "Order: %v", order)

		key := datastore.NewIncompleteKey(c, "OrderTest", parentKey(c, "OrderTest"))
		complete_key, err := datastore.Put(c, key, &order)
		HandleError(w, err)
		log.Infof(c, "Order Stored: %v", complete_key.IntID())

		fmt.Fprintf(w, "It is all Good man")
	}
}
