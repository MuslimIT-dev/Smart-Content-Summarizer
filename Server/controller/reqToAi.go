package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"smartcontentsummarizer/Server/models"
)

func ReqTextToAi(c *gin.Context) {
	var reqBody models.RequestBody
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	url := "https://api-inference.huggingface.co/models/facebook/bart-large-cnn"
	payload := map[string]interface{}{
		"inputs": reqBody.Text,
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY не установлен!")
		return
	}

	payloadBytes, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	req.Header.Set("Authorization", "Bearer "+apiKey) // вставь свой токен
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contact Hugging Face"})
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body)) // для отладки: увидишь точный ответ HF

	// Исправлено: парсим как массив
	var hfResp []struct {
		SummaryText string `json:"summary_text"`
	}

	if err := json.Unmarshal(body, &hfResp); err != nil {
		log.Println("Unmarshal error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse Hugging Face response"})
		return
	}

	if len(hfResp) == 0 {
		c.JSON(http.StatusOK, gin.H{"summary": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": hfResp[0].SummaryText})
}
