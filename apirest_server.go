/*
Go server apirest simulation  A.Villanueva

Modeles
curl -X GET -H "Content-Type: application/json" -H "APIKey: cd66a4f9-8a9b-4a8c-a02a-ff2d7d1c3e3c" <URL>
curl -X GET -H "Content-Type: application/json" -H "APIKey: ef3a5fb2-8d89-4f07-ae6c-9b2692ef9f5f" <URL>


OK get gateways with APIKey
curl -X GET -H "Content-Type: application/json" -H "APIKey: cd66a4f9-8a9b-4a8c-a02a-ff2d7d1c3e3c" 127.0.0.1:8080/api/v1/gateways


Erreur no APIKey
curl -X GET -H "Content-Type: application/json"  127.0.0.1:8080/api/v1/gateways
curl http://localhost:8080/api/v1/gateways
*/

package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	URL_BASE = "/api/v1"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Data struct {
	TotalCount int       `json:"totalCount"`
	Results    []Gateway `json:"results"`
}

type Gateway struct {
	GatewayID          int    `json:"gatewayId"`
	GatewayLabel       string `json:"gatewayLabel"`
	CMac               string `json:"cMac"`
	LorawanMac         string `json:"lorawanMac"`
	LorawanConnected   bool   `json:"lorawanConnected"`
	LorawanLastPing    string `json:"lorawanLastPing"`
	ClovernetConnected bool   `json:"clovernetConnected"`
	ClovernetLastPing  string `json:"clovernetLastPing"`
	LevelID            int    `json:"levelId"`
}

var gateways = Data{
	TotalCount: 2,
	Results: []Gateway{
		{
			GatewayID:          1,
			GatewayLabel:       "00800000A0008796_FBR",
			CMac:               "55:66:77:88",
			LorawanMac:         "01:02:03:04:05:06",
			LorawanConnected:   true,
			LorawanLastPing:    "2023-06-28T10:00:00Z",
			ClovernetConnected: false,
			ClovernetLastPing:  "",
			LevelID:            3,
		},
		{
			GatewayID:          2,
			GatewayLabel:       "00800000A00093B7-AXIOME",
			CMac:               "11:22:33:44",
			LorawanMac:         "AA:BB:CC:DD:EE:FF",
			LorawanConnected:   false,
			LorawanLastPing:    "",
			ClovernetConnected: true,
			ClovernetLastPing:  "2023-06-28T12:30:00Z",
			LevelID:            2,
		},
	},
}

// getGateways
func getGateways(c *gin.Context) {
	if testHeader(c) {
		c.IndentedJSON(http.StatusOK, gateways)
	} else {
		// Erreur json message
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{"\n Internal Server Error \n"})
	}
}

func testHeader(c *gin.Context) bool {
	//Vérifier la clé API
	apiKey := c.GetHeader("APIKey")
	if apiKey != "cd66a4f9-8a9b-4a8c-a02a-ff2d7d1c3e3c" && apiKey != "ef3a5fb2-8d89-4f07-ae6c-9b2692ef9f5f" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return false
	}

	//Vérifiez le header Content-Type
	contentType := c.GetHeader("Content-Type")
	if contentType != "application/json" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Content-Type"})
		return false
	}

	return true
}

func main() {
	router := gin.Default()
	router.GET(URL_BASE+"/gateways", getGateways)

	router.Run("localhost:8080")
	os.Exit(0)
}
