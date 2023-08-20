package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var productMap = struct {
	sync.RWMutex
	m map[int]Product
}{m: make(map[int]Product)}

func init() {
	productMap.m, _ = loadProductsMap()
}

func loadProductsMap() (map[int]Product, error) {
	fileName := "products.json"
	// check if the file exists
	_, err := os.Stat("products.json")
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	// read the file content
	file, _ := ioutil.ReadFile(fileName)
	productList := make([]Product, 0)
	err = json.Unmarshal([]byte(file), &productList)
	if err != nil {
		log.Fatal(err)
	}
	prodMap := make(map[int]Product)
	for i := 0; i < len(productList); i++ {
		prodMap[productList[i].ProductID] = productList[i]
	}
	return prodMap, nil
}

func getProduct(productID int) *Product {
	productMap.RLock()
	defer productMap.RUnlock()
	if product, ok := productMap.m[productID]; ok {
		return &product
	}
	return nil
}

func getAllProducts() []Product {
	var prodList []Product
	productMap.RLock()
	for _, product := range productMap.m {
		prodList = append(prodList, product)
	}
	productMap.RUnlock()
	return prodList
}

func getNextProductID() int {
	productIDs := getProductIDs()
	return productIDs[len(productIDs)-1] + 1
}

func getProductIDs() []int {
	var prodIDs []int
	for _, product := range productMap.m {
		prodIDs = append(prodIDs, product.ProductID)
	}
	return prodIDs
}

func addOrUpdateProduct(product Product) (int, error) {
	// if the productID is set, update, otherwise add
	addOrUpdateID := -1
	if product.ProductID > 0 {
		oldProduct := getProduct(product.ProductID)
		// if exists replace it, otherwise return error
		if oldProduct == nil {
			return 0, fmt.Errorf("product id [%d] doesn't exist", product.ProductID)
		}
		addOrUpdateID = product.ProductID
	} else {
		addOrUpdateID = getNextProductID()
		product.ProductID = addOrUpdateID
	}
	productMap.Lock()
	productMap.m[addOrUpdateID] = product
	productMap.Unlock()
	return addOrUpdateID, nil
}

func deleteProduct(productID int) {
	productMap.Lock()
	delete(productMap.m, productID)
	productMap.Unlock()
}
