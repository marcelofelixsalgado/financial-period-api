package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const categoryAPIUrl = "http://localhost:8081"
const accessToken = "[ACCESS_TOKEN]"
const categoryDVId = "357ede5b-c662-48d3-9442-745fc9e9ac15"

var categories = []inputCreateCategoryDto{
	{
		Code: "DF",
		Name: "Despesa fixa",
		TransactionType: transactionTypeInput{
			Code: "EXPENSE",
		},
	},
	{
		Code: "DV",
		Name: "Despesa variável",
		TransactionType: transactionTypeInput{
			Code: "EXPENSE",
		},
	},
}

var subCategories = []inputCreateSubCategoryDto{
	{
		Code: "AG",
		Name: "Açouge",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{
		Code: "AL",
		Name: "Alimentação",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "CA",
		Name: "Casa",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "CM",
		Name: "Cabeleleira/Manicure",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "IC",
		Name: "Investimento de carreira",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "CT",
		Name: "Consultório",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "DS",
		Name: "Desconhecido",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "DV",
		Name: "Diversos",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "FA",
		Name: "Farmácia",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "ME",
		Name: "Mercado",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "PD",
		Name: "Padaria",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "PL",
		Name: "Papelaria",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "PF",
		Name: "Perfumaria",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "PR",
		Name: "Presente",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "RO",
		Name: "Roupas",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
	{

		Code: "TR",
		Name: "Transporte",
		Category: categoryInput{
			Id: categoryDVId,
		},
	},
}

type inputCreateCategoryDto struct {
	Code            string               `json:"code"`
	Name            string               `json:"name"`
	TransactionType transactionTypeInput `json:"transaction_type"`
}

type inputCreateSubCategoryDto struct {
	Code     string        `json:"code"`
	Name     string        `json:"name"`
	Category categoryInput `json:"category"`
}

type transactionTypeInput struct {
	Code string `json:"code"`
}

type categoryInput struct {
	Id string `json:"id"`
}

// func main() {
// 	// for _, item := range categories {
// 	// 	createCategory(item)
// 	// }
// 	// for _, item := range subCategories {
// 	// 	createSubCategory(item)
// 	// }
// }

func createCategory(input inputCreateCategoryDto) {

	category, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("%s/v1/categories", categoryAPIUrl)
	response, err := makeUpstreamRequest(http.MethodPost, url, category)
	fmt.Println(response.Status)
	if err != nil {
		fmt.Println(err)
	}
}

func createSubCategory(input inputCreateSubCategoryDto) {

	subCategory, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	}

	url := fmt.Sprintf("%s/v1/subcategories", categoryAPIUrl)
	response, err := makeUpstreamRequest(http.MethodPost, url, subCategory)
	fmt.Println(response.Status)
	if err != nil {
		fmt.Println(err)
	}
}

func makeUpstreamRequest(method, url string, data []byte) (*http.Response, error) {

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
