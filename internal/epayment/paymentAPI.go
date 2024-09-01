package epayment

var (
	ClientId     = "test"
	ClientSecret = "yF587AV9Ms94qN2QShFzVR3vFnWkhjbAK3sG"
)

type PaymentRequest struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Name        string `json:"name"`
	Cryptogram  string `json:"cryptogram"`
	InvoiceId   string `json:"invoiceId"`
	Description string `json:"description"`
	Email       string `json:"email"`
	PostLink    string `json:"postLink"`
	Cardsave    string `json:"cardsave"`
}

type PaymentResponse struct {
	ID          int    `json:"id"`
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	Name        string `json:"name"`
	InvoiceId   string `json:"invoiceId"`
	Description string `json:"description"`
	Email       string `json:"email"`
	PostLink    string `json:"postLink"`
	Cardsave    string `json:"cardsave"`
}

func pay() {

}
