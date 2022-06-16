package givebutter

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var apiKey string

func SetAPIKey(key string) {
	apiKey = key
}

type TransactionScope string

func (scope TransactionScope) String() string { return string(scope) }

const (
	transactionUri                  = "https://api.givebutter.com/v1/transactions"
	Null           TransactionScope = "null"
	Benefiting     TransactionScope = "benefiting"
	Chapters       TransactionScope = "chapters"
	All            TransactionScope = "all"

	timeout = 5 * time.Second
)

func authHeader() string {
	return fmt.Sprintf("Bearer %s", apiKey)
}

func GetTransactions(ctx context.Context, scope TransactionScope) (*GetTransactionsResponse, error) {
	var response = GetTransactionsResponse{}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var butterURL, err = url.Parse(transactionUri)
	if err != nil {
		return &response, err
	}
	var queryParams = butterURL.Query()
	queryParams.Add("scope", scope.String())
	butterURL.RawQuery = queryParams.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", butterURL.String(), nil)
	if err != nil {
		return &response, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", authHeader())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return &response, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return &response, err
	}

	return &response, nil
}
