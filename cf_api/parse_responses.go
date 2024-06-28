package cfapi

import (
	"encoding/json"
	"fmt"
)

// Define structs to match the JSON response structure
type FindDBIDResponse struct {
	Result []struct {
		UUID      string `json:"uuid"`
		Name      string `json:"name"`
		Version   string `json:"version"`
		CreatedAt string `json:"created_at"`
	} `json:"result"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

// Function to find the UUID of a D1 database by its name
func findD1DatabaseUUID(body []byte, dbName string) (string, error) {
	var response FindDBIDResponse

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	// Iterate over the results to find the database by name
	for _, db := range response.Result {
		if db.Name == dbName {
			return db.UUID, nil
		}
	}

	// Return an error if the database is not found
	return "", fmt.Errorf("database with name %s not found", dbName)
}
