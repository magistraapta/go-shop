package dto

// Message represents the message response
// @Description Message response
// @Accept json
// @Produce json
// @Param message body string true "Message"
type OrderMessage struct {
	Message string `json:"message"`
}

// CheckoutRequest represents the checkout request
// @Description Checkout request
// @Accept json
// @Produce json
// @Param payment_method body string true "Payment method"
type CheckoutRequest struct {
	PaymentMethod string `json:"payment_method" binding:"required"`
}
