package adiq

import (
	"encoding/json"
	"io"
	"net/http"
)

const planUrl = "https://recorrencia.adiq.io/v1/recurrence/plans/"

func GetPlan(accessToken string, id string) (*Plan, error) {
	req, err := http.NewRequest("GET", planUrl+id, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	var plan *Plan
	err = json.Unmarshal(resBody, &plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}
