package services

import (
	"context"
	"encoding/json"
	"fmt"

	"audio-series-app/backend/internal/config"
	"audio-series-app/backend/internal/models"

	"github.com/google/uuid"
)

type PaymentService struct {
	config   *config.Config
	supabase *SupabaseService
}

func NewPaymentService(cfg *config.Config, supabase *SupabaseService) *PaymentService {
	return &PaymentService{
		config:   cfg,
		supabase: supabase,
	}
}

// Coin bundles for different regions
var coinBundles = map[string][]models.CoinBundle{
	"INR": {
		{ID: uuid.New(), Name: "50 Coins", Coins: 50, Price: 5000, Currency: "INR", IsActive: true},    // ₹50
		{ID: uuid.New(), Name: "120 Coins", Coins: 120, Price: 9900, Currency: "INR", IsActive: true},  // ₹99
		{ID: uuid.New(), Name: "250 Coins", Coins: 250, Price: 19900, Currency: "INR", IsActive: true}, // ₹199
		{ID: uuid.New(), Name: "500 Coins", Coins: 500, Price: 39900, Currency: "INR", IsActive: true}, // ₹399
	},
	"NGN": {
		{ID: uuid.New(), Name: "50 Coins", Coins: 50, Price: 5000, Currency: "NGN", IsActive: true},    // ₦50
		{ID: uuid.New(), Name: "120 Coins", Coins: 120, Price: 9900, Currency: "NGN", IsActive: true},  // ₦99
		{ID: uuid.New(), Name: "250 Coins", Coins: 250, Price: 19900, Currency: "NGN", IsActive: true}, // ₦199
		{ID: uuid.New(), Name: "500 Coins", Coins: 500, Price: 39900, Currency: "NGN", IsActive: true}, // ₦399
	},
}

func (s *PaymentService) GetCoinBundles(currency string) []models.CoinBundle {
	if bundles, exists := coinBundles[currency]; exists {
		return bundles
	}
	return coinBundles["INR"] // Default to INR
}

func (s *PaymentService) InitiatePayment(ctx context.Context, userIDStr, bundleID, currency string) (*models.PaymentResponse, error) {
	// Parse user ID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	// Find the bundle
	var selectedBundle *models.CoinBundle
	bundles := s.GetCoinBundles(currency)
	for _, bundle := range bundles {
		if bundle.ID.String() == bundleID {
			selectedBundle = &bundle
			break
		}
	}

	if selectedBundle == nil {
		return nil, fmt.Errorf("invalid bundle ID")
	}

	// Create payment record
	payment := &models.Payment{
		UserID:   userID,
		Amount:   selectedBundle.Price,
		Currency: currency,
		Coins:    selectedBundle.Coins,
		Status:   "pending",
	}

	err = s.supabase.CreatePayment(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment record: %w", err)
	}

	// Initialize payment based on currency
	var paymentResponse *models.PaymentResponse
	switch currency {
	case "INR":
		paymentResponse, err = s.initiateRazorpayPayment(payment, selectedBundle)
	case "NGN":
		paymentResponse, err = s.initiatePaystackPayment(payment, selectedBundle)
	default:
		return nil, fmt.Errorf("unsupported currency: %s", currency)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to initiate payment: %w", err)
	}

	return paymentResponse, nil
}

func (s *PaymentService) initiateRazorpayPayment(payment *models.Payment, bundle *models.CoinBundle) (*models.PaymentResponse, error) {
	// In a real implementation, you would use the Razorpay SDK
	// For now, we'll create a mock response
	gatewayRef := "rzp_" + uuid.New().String()[:16]

	// Update payment with gateway reference
	payment.Gateway = "razorpay"
	payment.GatewayRef = gatewayRef

	paymentData := map[string]interface{}{
		"key_id":      s.config.RazorpayKeyID,
		"amount":      bundle.Price,
		"currency":    bundle.Currency,
		"order_id":    gatewayRef,
		"description": fmt.Sprintf("Purchase of %d coins", bundle.Coins),
		"prefill": map[string]string{
			"email": "user@example.com",
		},
	}

	paymentDataJSON, _ := json.Marshal(paymentData)

	err := s.supabase.UpdatePayment(context.Background(), payment.ID, "pending", string(paymentDataJSON))
	if err != nil {
		return nil, err
	}

	return &models.PaymentResponse{
		PaymentID:   payment.ID.String(),
		GatewayRef:  gatewayRef,
		Amount:      bundle.Price,
		Currency:    bundle.Currency,
		Gateway:     "razorpay",
		RedirectURL: fmt.Sprintf("https://checkout.razorpay.com/v1/checkout.html?%s", gatewayRef),
	}, nil
}

func (s *PaymentService) initiatePaystackPayment(payment *models.Payment, bundle *models.CoinBundle) (*models.PaymentResponse, error) {
	// In a real implementation, you would use the Paystack SDK
	// For now, we'll create a mock response
	gatewayRef := "ps_" + uuid.New().String()[:16]

	// Update payment with gateway reference
	payment.Gateway = "paystack"
	payment.GatewayRef = gatewayRef

	paymentData := map[string]interface{}{
		"public_key":   s.config.PaystackPublicKey,
		"amount":       bundle.Price,
		"currency":     bundle.Currency,
		"reference":    gatewayRef,
		"email":        "user@example.com",
		"callback_url": "https://yourapp.com/payment/callback",
		"metadata": map[string]interface{}{
			"payment_id": payment.ID.String(),
			"coins":      bundle.Coins,
		},
	}

	paymentDataJSON, _ := json.Marshal(paymentData)

	err := s.supabase.UpdatePayment(context.Background(), payment.ID, "pending", string(paymentDataJSON))
	if err != nil {
		return nil, err
	}

	return &models.PaymentResponse{
		PaymentID:   payment.ID.String(),
		GatewayRef:  gatewayRef,
		Amount:      bundle.Price,
		Currency:    bundle.Currency,
		Gateway:     "paystack",
		RedirectURL: fmt.Sprintf("https://checkout.paystack.com/%s", gatewayRef),
	}, nil
}

func (s *PaymentService) HandlePaymentCallback(ctx context.Context, gateway string, paymentData map[string]interface{}) error {
	// Extract payment reference
	var paymentRef string
	switch gateway {
	case "razorpay":
		if ref, ok := paymentData["razorpay_payment_id"].(string); ok {
			paymentRef = ref
		}
	case "paystack":
		if ref, ok := paymentData["reference"].(string); ok {
			paymentRef = ref
		}
	default:
		return fmt.Errorf("unsupported gateway: %s", gateway)
	}

	if paymentRef == "" {
		return fmt.Errorf("invalid payment reference")
	}

	// In a real implementation, you would:
	// 1. Verify the payment with the gateway
	// 2. Update the payment status
	// 3. Add coins to user balance
	// 4. Send confirmation email

	// For now, we'll simulate a successful payment
	// You would need to implement proper verification logic

	return nil
}

func (s *PaymentService) VerifyPayment(ctx context.Context, paymentID uuid.UUID) error {
	// In a real implementation, you would verify the payment with the gateway
	// For now, we'll assume the payment is successful
	return nil
}
