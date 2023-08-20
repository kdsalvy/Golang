package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//Handler - Perform single product operations
func Handler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product := getProduct(productID)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// return a single product
		productJSON, err := json.Marshal(*product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)
	case http.MethodPut:
		//update the product
		var updateProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updateProduct)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updateProduct.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		addOrUpdateProduct(updateProduct)
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodDelete:
		deleteProduct(productID)
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

//BulkHandler - Perform multi product operations
func BulkHandler(e http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(getAllProducts())
		if err != nil {
			log.Fatal(err)
			e.WriteHeader(http.StatusInternalServerError)
			return
		}
		e.Header().Set("Content-Type", "application/json")
		e.Write(productsJSON)
	case http.MethodPost:
		// add a new product to the list
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
			e.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			log.Fatal(err)
			e.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			e.WriteHeader(http.StatusBadRequest)
			return
		}
		addOrUpdateProduct(newProduct)
		e.WriteHeader(http.StatusCreated)
		return
	}
}
