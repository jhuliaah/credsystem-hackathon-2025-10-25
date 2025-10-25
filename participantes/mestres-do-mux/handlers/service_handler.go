package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// FindServiceRequest (from your original code)
// The client sends: {"intent": "some message..."}
type FindServiceRequest struct {
    Intent string `json:"intent"`
}

// AIHandlerResponse - A new response struct to send back to the client
// We return: {"success": true, "response": "The AI's answer..."}
type AIHandlerResponse struct {
    Success  bool   `json:"success"`
    Response string `json:"response,omitempty"` // Holds the AI's answer
    Error    string `json:"error,omitempty"`
}

// --- Structs for calling the OpenRouter AI API ---

// OpenRouterRequest is the payload we send to the AI
type OpenRouterRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
}

// Message is part of the OpenRouterRequest
type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

// OpenRouterResponse is what we expect back from the AI
type OpenRouterResponse struct {
    Choices []struct {
        Message struct {
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
    // We only care about the first choice's message content
}

// Your new Gin handler function
func HandleAIAssistedIntent(c *gin.Context) {
    // 1. Get the API Key from environment variables
    apiKey := os.Getenv("OPENROUTER_API_KEY")
    if apiKey == "" {
        c.JSON(http.StatusInternalServerError, AIHandlerResponse{
            Success: false,
            Error:   "OPENROUTER_API_KEY is not set on the server",
        })
        return
    }

    // 2. Tenta converter o JSON recebido ({"intent": "..."})
    var req FindServiceRequest
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, AIHandlerResponse{
            Success: false,
            Error:   "Invalid JSON body",
        })
        return
    }

    // 3. Construct the request payload for the AI
    // We use the client's "intent" as the prompt "content"
    aiPayload := OpenRouterRequest{
        Model: "mistralai/mistral-7b-instruct",
        Messages: []Message{
            {
                Role:    "user",
                Content: req.Intent, // Here is where the client's message is passed to the AI
            },
        },
    }

    payloadBytes, err := json.Marshal(aiPayload)
    if err != nil {
        c.JSON(http.StatusInternalServerError, AIHandlerResponse{
            Success: false,
            Error:   "Failed to create AI request payload",
        })
        return
    }

    // 4. Create the HTTP POST request to OpenRouter
    httpReq, err := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", bytes.NewBuffer(payloadBytes))
    if err != nil {
        c.JSON(http.StatusInternalServerError, AIHandlerResponse{
            Success: false,
            Error:   "Failed to create HTTP request",
        })
        return
    }

    // 5. Set the required headers (as seen in your cURL command)
    httpReq.Header.Set("Content-Type", "application/json")
    httpReq.Header.Set("Authorization", "Bearer "+apiKey)

    // 6. Send the request
    client := &http.Client{Timeout: 30 * time.Second}
    httpResp, err := client.Do(httpReq)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, AIHandlerResponse{
            Success: false,
            Error:   "Failed to communicate with AI service",
        })
        return
    }
    defer httpResp.Body.Close()

    // 7. Read the AI's response
    body, err := io.ReadAll(httpResp.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, AIHandlerResponse{
            Success: false,
            Error:   "Failed to read AI response body",
        })
        return
    }

    // 8. Check if the AI service returned an error
    if httpResp.StatusCode != http.StatusOK {
        c.JSON(http.StatusInternalServerError, AIHandlerResponse{
            Success: false,
            Error:   fmt.Sprintf("AI service returned an error: %s", string(body)),
        })
        return
    }

    // 9. Unmarshal the AI's JSON response
    var aiResp OpenRouterResponse
    if err := json.Unmarshal(body, &aiResp); err != nil {
        c.JSON(http.StatusInternalServerError, AIHandlerResponse{
            Success: false,
            Error:   "Failed to parse AI response",
        })
        return
    }

    // 10. Extract the text answer
    if len(aiResp.Choices) == 0 {
        c.JSON(http.StatusInternalServerError, AIHandlerResponse{
            Success: false,
            Error:   "AI response contained no answer",
        })
        return
    }
    aiAnswer := aiResp.Choices[0].Message.Content

    // 11. Retorna a resposta da AI para o cliente original
    c.JSON(http.StatusOK, AIHandlerResponse{
        Success:  true,
        Response: aiAnswer,
    })
}
