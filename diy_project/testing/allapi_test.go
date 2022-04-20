package testing

import (
	"bytes"
	"github.com/Jeffail/gabs"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

var url = "http://localhost:8080"

func TestCreateUser(t *testing.T) {
	jsonStr := []byte(`{"name":"kartik","email":"kartik@gmail.com","type": "merchant"}}`)
	resp, err := http.Post(url+"/user/create", "application/json", bytes.NewBuffer(jsonStr))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	parsejson, err := gabs.ParseJSON(data)
	assert.Nil(t, err)
	assert.Equal(t, "kartik", parsejson.Path("data.name").Data())
}

func TestUpdateUser(t *testing.T) {
	jsonStr := []byte(`{"name":"kartikA","email":"kartik@gmail.com","type": "merchant"}}`)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url+"/user/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	parsejson, err := gabs.ParseJSON(data)

	assert.Equal(t, "kartikA", parsejson.Path("data.name").Data())
}

func TestUpdateUserNotfound(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url+"/user/1000", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCreateProduct(t *testing.T) {
	jsonStr := []byte(`{"user_id":1,"product_list": [{"name":"orange","category":"fruit","quantity":100}]}`)
	resp, err := http.Post(url+"/add_products_list", "application/json", bytes.NewBuffer(jsonStr))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	parsejson, err := gabs.ParseJSON(data)
	assert.Nil(t, err)
	assert.Equal(t, "Products added successfully !!!", parsejson.Path("data").Data())
}

func TestGetProductsById(t *testing.T) {
	resp, err := http.Get(url + "/product/1")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	parsejson, err := gabs.ParseJSON(data)
	assert.Nil(t, err)
	assert.Equal(t, "orange", parsejson.Path("data.name").Data())
}

func TestGetProductsByIdNotFound(t *testing.T) {
	resp, err := http.Get(url + "/product/1000")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	defer resp.Body.Close()
	assert.Nil(t, err)
}

func TestGetAllProductsFound(t *testing.T) {
	resp, err := http.Get(url + "/all_products")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	assert.Nil(t, err)
}

func TestCreateOrder(t *testing.T) {
	jsonStr := []byte(`{"user_id":1,"cart_products": [{"product_id":1,"product_quantity":2}]}`)
	resp, err := http.Post(url+"/place_order", "application/json", bytes.NewBuffer(jsonStr))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()
	assert.Nil(t, err)
}

func TestGetAllOrder(t *testing.T) {
	resp, err := http.Get(url + "/all_orders")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	assert.Nil(t, err)
}

func TestGetTopProducts(t *testing.T) {
	resp, err := http.Get(url + "/top_products")
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	assert.Nil(t, err)
}
