package givebutter

type (
	GetTransactionsResponse struct {
		Data []struct {
			FirstName string  `json:"first_name,omitempty"`
			LastName  string  `json:"last_name,omitempty"`
			Amount    uint64  `json:"amount,omitempty"`
			Fee       float64 `json:"fee,omitempty"`
			Donated   uint64  `json:"donated,omitempty"`
		} `json:"data"`
	}
)
