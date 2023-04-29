/*
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type SearchResponse struct {
	Documents []Item `json:"documents"`
}

const (
	DbUrl = "https://us-west-2.aws.data.mongodb-api.com"
)

var ApiKey string

func init() {
	// Get the file path of the current function
	_, filePath, _, ok := runtime.Caller(0)
	if !ok {
		panic("Could not get file path")
	}

	// Get the absolute path of the directory containing the file
	dirPath := filepath.Dir(filePath)
	secretsFilePath := filepath.Join(dirPath, "secrets.txt") //"."s
	var err error
	ApiKey, err = readAPIKeyFromFile(secretsFilePath)
	if err != nil {
		panic(err)
	}
	// Use the API key to initialize the database connection
}

func readAPIKeyFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	apiKey := scanner.Text()

	return apiKey, nil
}

func insertItem(newItem Item) error {
	url := DbUrl + "/app/data-ibscs/endpoint/data/v1/action/insertOne"

	payload := map[string]interface{}{
		"collection": "product",
		"database":   "rd-poc-db",
		"dataSource": "RundooCluster0",
		"document": map[string]interface{}{
			"name":     newItem.Name,
			"category": newItem.Category,
			"sku":      newItem.Sku,
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Control-Request-Headers", "*")
	req.Header.Set("api-key", ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to insert item. Status code: %d", resp.StatusCode)
	}

	return nil
}

func searchItems(query string) ([]Item, error) {
	url := "https://us-west-2.aws.data.mongodb-api.com/app/data-ibscs/endpoint/data/v1/action/aggregate"

	searchQuery := fmt.Sprintf(`{
        "collection":"product",
        "database":"rd-poc-db",
        "dataSource":"RundooCluster0",
        "pipeline": [
          {
            "$search": {
                "index": "product-search",
                "regex": {
                  "allowAnalyzedField": true,
                  "query": "%s",										
                   "path": {
                      "wildcard": "*"
                  }
                }
              }
          }
        ]
    }`, query)

	//"query": "(.*)%s",

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(searchQuery)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Access-Control-Request-Headers", "*")
	req.Header.Set("api-key", ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response body
	var searchResponse SearchResponse
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return nil, err
	}

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(body))

	return searchResponse.Documents, nil

}
