package geminiai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type GeminiClient struct{}

type GeminiResponse struct {
	RecommendedSpecialty string  `json:"recommended_specialty"`
	PreferredTimeStart   *string `json:"preferred_time_start"`
	PreferredTimeEnd     *string `json:"preferred_time_end"`
	Summary              string  `json:"summary"`
}

type geminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type geminiAPIResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`

	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

func NewGeminiClient() *GeminiClient {
	return &GeminiClient{}
}

func (g *GeminiClient) AnalyzeSymptoms(symptoms string) (*GeminiResponse, error) {
	prompt := fmt.Sprintf(`
		You are a medical triage assistant.
		
		Analyze the symptoms below.
		
		Choose ONLY ONE specialty from this list:
		
		- Dokter Umum
		- Dokter Gigi
		- Dokter Anak
		- Dokter Jantung
		- Dokter Penyakit Dalam
		- Dokter Saraf
		- Dokter Kulit dan Kelamin
		- Dokter Mata
		- Dokter THT
		- Dokter Kandungan
		- Dokter Bedah
		- Dokter Paru
		
		Extract:
		
		1. recommended_specialty
		2. preferred_time_start (if mentioned)
		3. preferred_time_end (if mentioned)
		4. summary
		
		Return ONLY valid JSON.
		
		{
		  "recommended_specialty": "",
		  "preferred_time_start": "",
		  "preferred_time_end": "",
		  "summary": ""
		}
		
		Symptoms:
		%s
		`, symptoms)

	reqBody := geminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{
						Text: prompt,
					},
				},
			},
		},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf(
		"%s/v1beta/models/%s:generateContent?key=%s",
		os.Getenv("GEMINI_BASE_URL"),
		os.Getenv("GEMINI_MODEL"),
		os.Getenv("GEMINI_API_KEY"),
	)

	resp, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geminiResp geminiAPIResponse

	err = json.Unmarshal(
		responseBody,
		&geminiResp,
	)

	if err != nil {
		return nil, err
	}

	if geminiResp.Error.Message != "" {
		return nil, fmt.Errorf(
			"gemini error: %s",
			geminiResp.Error.Message,
		)
	}

	if len(geminiResp.Candidates) == 0 {
		return nil, fmt.Errorf("gemini returned empty candidates")
	}

	rawJSON := geminiResp.Candidates[0].
		Content.
		Parts[0].
		Text

	rawJSON = strings.TrimSpace(rawJSON)
	rawJSON = strings.TrimPrefix(rawJSON, "```json")
	rawJSON = strings.TrimPrefix(rawJSON, "```")
	rawJSON = strings.TrimSuffix(rawJSON, "```")
	rawJSON = strings.TrimSpace(rawJSON)

	var result GeminiResponse

	err = json.Unmarshal(
		[]byte(rawJSON),
		&result,
	)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
