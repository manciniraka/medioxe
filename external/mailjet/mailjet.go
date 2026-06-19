package mailjet

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type MailjetConfig struct {
	MailjetBaseURL           string
	MailjetBasicAuthUsername string
	MailjetBasicAuthPassword string
	MailjetSenderEmail       string
	MailjetSenderName        string
}

type MailjetRepository struct {
	mailjetConfig MailjetConfig
}

func NewMailjetRepository(
	cfg MailjetConfig,
) *MailjetRepository {
	return &MailjetRepository{
		mailjetConfig: cfg,
	}
}

type payloadSendEmail struct {
	Messages []Messages `json:"Messages"`
}

type From struct {
	Email string `json:"Email"`
	Name  string `json:"Name"`
}

type To struct {
	Email string `json:"Email"`
	Name  string `json:"Name"`
}

type Messages struct {
	From     From   `json:"From"`
	To       []To   `json:"To"`
	Subject  string `json:"Subject"`
	TextPart string `json:"TextPart"`
	HTMLPart string `json:"HTMLPart"`
}

func (r *MailjetRepository) SendEmail(
	toName string,
	toEmail string,
	subject string,
	message string,
) error {
	url := r.mailjetConfig.MailjetBaseURL + "/v3.1/send"

	toBody := []To{
		{
			Email: toEmail,
			Name:  toName,
		},
	}

	payload := payloadSendEmail{
		Messages: []Messages{
			{
				From: From{
					Email: r.mailjetConfig.MailjetSenderEmail,
					Name:  r.mailjetConfig.MailjetSenderName,
				},
				To:       toBody,
				Subject:  subject,
				TextPart: message,
				HTMLPart: message,
			},
		},
	}

	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		url,
		strings.NewReader(string(payloadByte)),
	)

	if err != nil {
		return err
	}

	auth := base64.StdEncoding.EncodeToString(
		[]byte(
			r.mailjetConfig.MailjetBasicAuthUsername +
				":" +
				r.mailjetConfig.MailjetBasicAuthPassword,
		),
	)

	req.Header.Set(
		"Authorization",
		"Basic "+auth,
	)

	req.Header.Set(
		"Content-type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 ||
		resp.StatusCode > 299 {
		return fmt.Errorf(
			"mailjet returned status %d",
			resp.StatusCode,
		)
	}

	return nil
}
