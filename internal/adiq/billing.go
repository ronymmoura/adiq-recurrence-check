package adiq

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ronymmoura/adiq-recurrence-check/internal/util"
)

type GetBillingResponse struct {
	Items []Billing `json:"items"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
	Total int       `json:"total"`
}

type Billing struct {
	Status            string             `json:"status"`
	ExpireAt          util.TimeWithoutTZ `json:"expireAt"`
	Amount            float32            `json:"amount"`
	Tid               string             `json:"tid"`
	AuthorizationCode string             `json:"authorizationCode"`
	Installment       int                `json:"installment"`
	Id                string             `json:"id"`
	CreatedBy         string             `json:"createdBy"`
	CreatedDate       util.TimeWithoutTZ `json:"createdDate"`
	ModifiedBy        string             `json:"modifiedBy"`
	ModifiedDate      util.TimeWithoutTZ `json:"modifiedDate"`
	Subscription      Subscription       `json:"subscription"`
}

type Subscription struct {
	VaultId     string             `json:"vaultId"`
	OrderNumber string             `json:"orderNumber"`
	Status      string             `json:"status"`
	Id          string             `json:"id"`
	CreatedDate util.TimeWithoutTZ `json:"createdDate"`
	CreatedBy   string             `json:"createdBy"`
	Plan        Plan               `json:"plan"`
}

type Plan struct {
	MerchantId    string             `json:"merchantId"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Amount        float32            `json:"amount"`
	PlanType      string             `json:"planType"`
	TrialDays     int                `json:"int"`
	PaymentMethod string             `json:"paymentMethod"`
	Interval      int                `json:"interval"`
	Installments  int                `json:"installments"`
	Status        string             `json:"status"`
	Attempts      int                `json:"attempts"`
	Id            string             `json:"id"`
	CreatedDate   util.TimeWithoutTZ `json:"createdDate"`
}

const billingUrl = "https://recorrencia.adiq.io/v1/recurrence/billing"

func GetBilling(accessToken string) (billings GetBillingResponse, err error) {
	req, err := http.NewRequest("GET", billingUrl, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)

	err = json.Unmarshal(resBody, &billings)
	if err != nil {
		return
	}

	return
}
