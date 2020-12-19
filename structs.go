package goSellix

type Order struct {
	Status int `json:"status"`
	Data   struct {
		Order struct {
			ID                        int           `json:"id"`
			Uniqid                    string        `json:"uniqid"`
			Type                      string        `json:"type"`
			Total                     float64       `json:"total"`
			TotalDisplay              float64       `json:"total_display"`
			ExchangeRate              int           `json:"exchange_rate"`
			CryptoExchangeRate        float64       `json:"crypto_exchange_rate"`
			Currency                  string        `json:"currency"`
			ShopID                    int           `json:"shop_id"`
			Name                      string        `json:"name"`
			CustomerEmail             string        `json:"customer_email"`
			PaypalEmailDelivery       float64       `json:"paypal_email_delivery"`
			ProductID                 string        `json:"product_id"`
			ProductTitle              string        `json:"product_title"`
			ProductType               string        `json:"product_type"`
			SubscriptionID            interface{}   `json:"subscription_id"`
			SubscriptionTime          interface{}   `json:"subscription_time"`
			Gateway                   string        `json:"gateway"`
			PaypalEmail               interface{}   `json:"paypal_email"`
			PaypalOrderID             interface{}   `json:"paypal_order_id"`
			PaypalPayerEmail          interface{}   `json:"paypal_payer_email"`
			PaypalFee                 float64       `json:"paypal_fee"`
			LexOrderID                interface{}   `json:"lex_order_id"`
			LexPaymentMethod          interface{}   `json:"lex_payment_method"`
			StripeClientSecret        interface{}   `json:"stripe_client_secret"`
			SkrillEmail               interface{}   `json:"skrill_email"`
			SkrillSid                 interface{}   `json:"skrill_sid"`
			SkrillLink                interface{}   `json:"skrill_link"`
			PerfectmoneyID            interface{}   `json:"perfectmoney_id"`
			CryptoAddress             string        `json:"crypto_address"`
			CryptoAmount              float64       `json:"crypto_amount"`
			CryptoReceived            float64       `json:"crypto_received"`
			CryptoURI                 string        `json:"crypto_uri"`
			CryptoConfirmationsNeeded float64       `json:"crypto_confirmations_needed"`
			CashappQrcode             interface{}   `json:"cashapp_qrcode"`
			CashappNote               interface{}   `json:"cashapp_note"`
			CashappCashtag            interface{}   `json:"cashapp_cashtag"`
			Country                   string        `json:"country"`
			Location                  string        `json:"location"`
			IP                        string        `json:"ip"`
			IsVpnOrProxy              bool          `json:"is_vpn_or_proxy"`
			UserAgent                 string        `json:"user_agent"`
			Quantity                  int           `json:"quantity"`
			CouponID                  interface{}   `json:"coupon_id"`
			CustomFields              []interface{} `json:"custom_fields"`
			DeveloperInvoice          bool          `json:"developer_invoice"`
			DeveloperTitle            interface{}   `json:"developer_title"`
			DeveloperWebhook          interface{}   `json:"developer_webhook"`
			DeveloperReturnURL        interface{}   `json:"developer_return_url"`
			Status                    string        `json:"status"`
			Discount                  float64       `json:"discount"`
			FeePercentage             float64       `json:"fee_percentage"`
			DayValue                  float64       `json:"day_value"`
			Day                       string        `json:"day"`
			Month                     string        `json:"month"`
			Year                      int           `json:"year"`
			CreatedAt                 int           `json:"created_at"`
			UpdatedAt                 int           `json:"updated_at"`
			UpdatedBy                 int           `json:"updated_by"`
			IPInfo                    struct {
				ID             int     `json:"id"`
				Type           string  `json:"type"`
				RequestID      string  `json:"request_id"`
				IP             string  `json:"ip"`
				UserAgent      string  `json:"user_agent"`
				UserLanguage   string  `json:"user_language"`
				FraudScore     int     `json:"fraud_score"`
				CountryCode    string  `json:"country_code"`
				Region         string  `json:"region"`
				City           string  `json:"city"`
				Isp            string  `json:"isp"`
				Asn            float64 `json:"asn"`
				Organization   string  `json:"organization"`
				Latitude       float64 `json:"latitude"`
				Longitude      float64 `json:"longitude"`
				IsCrawler      float64 `json:"is_crawler"`
				Timezone       string  `json:"timezone"`
				Mobile         int     `json:"mobile"`
				Host           string  `json:"host"`
				Proxy          float64 `json:"proxy"`
				Vpn            float64 `json:"vpn"`
				Tor            float64 `json:"tor"`
				ActiveVpn      float64 `json:"active_vpn"`
				ActiveTor      float64 `json:"active_tor"`
				RecentAbuse    float64 `json:"recent_abuse"`
				BotStatus      float64 `json:"bot_status"`
				ConnectionType string  `json:"connection_type"`
				AbuseVelocity  string  `json:"abuse_velocity"`
				CreatedAt      int     `json:"created_at"`
				UpdatedAt      int     `json:"updated_at"`
			} `json:"ip_info"`
			Serials                 []string      `json:"serials"`
			Webhooks                []interface{} `json:"webhooks"`
			CryptoPayout            bool          `json:"crypto_payout"`
			CryptoPayoutTransaction struct {
				ToAddress    string  `json:"to_address"`
				FromAddress  string  `json:"from_address"`
				CryptoAmount float64 `json:"crypto_amount"`
				Hash         string  `json:"hash"`
				CreatedAt    int     `json:"created_at"`
			} `json:"crypto_payout_transaction"`
			PaypalDispute interface{} `json:"paypal_dispute"`
			StatusHistory []struct {
				ID        int    `json:"id"`
				InvoiceID string `json:"invoice_id"`
				Status    string `json:"status"`
				Details   string `json:"details"`
				CreatedAt int    `json:"created_at"`
			} `json:"status_history"`
			CryptoTransactions []struct {
				CryptoAmount  float64 `json:"crypto_amount"`
				Hash          string  `json:"hash"`
				Confirmations float64 `json:"confirmations"`
				CreatedAt     int     `json:"created_at"`
				UpdatedAt     int     `json:"updated_at"`
			} `json:"crypto_transactions"`
			Product struct {
				ID                        int           `json:"id"`
				Uniqid                    string        `json:"uniqid"`
				ShopID                    int           `json:"shop_id"`
				Type                      string        `json:"type"`
				Title                     string        `json:"title"`
				Currency                  string        `json:"currency"`
				Price                     float64       `json:"price"`
				PriceDisplay              float64       `json:"price_display"`
				Description               string        `json:"description"`
				ImageAttachment           interface{}   `json:"image_attachment"`
				FileAttachment            interface{}   `json:"file_attachment"`
				QuantityMin               float64       `json:"quantity_min"`
				QuantityMax               float64       `json:"quantity_max"`
				QuantityWarning           float64       `json:"quantity_warning"`
				Gateways                  []string      `json:"gateways"`
				CustomFields              []interface{} `json:"custom_fields"`
				CryptoConfirmationsNeeded float64       `json:"crypto_confirmations_needed"`
				MaxRiskLevel              float64       `json:"max_risk_level"`
				BlockVpnProxies           bool          `json:"block_vpn_proxies"`
				DeliveryText              string        `json:"delivery_text"`
				ServiceText               string        `json:"service_text"`
				StockDelimiter            string        `json:"stock_delimiter"`
				Stock                     float64       `json:"stock"`
				DynamicWebhook            string        `json:"dynamic_webhook"`
				SortPriority              float64       `json:"sort_priority"`
				Unlisted                  bool          `json:"unlisted"`
				OnHold                    int           `json:"on_hold"`
				TermsOfService            string        `json:"terms_of_service"`
				Warranty                  int           `json:"warranty"`
				WarrantyText              string        `json:"warranty_text"`
				Private                   bool          `json:"private"`
				CreatedAt                 int           `json:"created_at"`
				UpdatedAt                 int           `json:"updated_at"`
				UpdatedBy                 int           `json:"updated_by"`
				Name                      string        `json:"name"`
				Serials                   []string      `json:"serials"`
				Webhooks                  []interface{} `json:"webhooks"`
				Feedback                  struct {
					Total    int `json:"total"`
					Positive int `json:"positive"`
					Neutral  int `json:"neutral"`
					Negative int `json:"negative"`
				} `json:"feedback"`
				Theme               string `json:"theme"`
				ImageAttachmentInfo struct {
					ID           int     `json:"id"`
					Uniqid       string  `json:"uniqid"`
					Storage      string  `json:"storage"`
					Name         string  `json:"name"`
					OriginalName string  `json:"original_name"`
					Extension    string  `json:"extension"`
					ShopID       int     `json:"shop_id"`
					Size         float64 `json:"size"`
					CreatedAt    int     `json:"created_at"`
				} `json:"image_attachment_info"`
			} `json:"product"`
		} `json:"order"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type CouponDelete struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type CouponCreation struct {
	Status int `json:"status"`
	Data   struct {
		Uniqid string `json:"uniqid"`
	} `json:"data"`
	Message string      `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type CouponResp struct {
	Status int `json:"status"`
	Data   struct {
		Coupons []struct {
			ID            int         `json:"id"`
			Uniqid        string      `json:"uniqid"`
			ShopID        int         `json:"shop_id"`
			Type          string      `json:"type"`
			Code          string      `json:"code"`
			UseType       string      `json:"use_type"`
			Discount      float64     `json:"discount"`
			Currency      interface{} `json:"currency"`
			Used          float64     `json:"used"`
			MaxUses       float64     `json:"max_uses"`
			CreatedAt     int         `json:"created_at"`
			UpdatedAt     int         `json:"updated_at"`
			UpdatedBy     int         `json:"updated_by"`
			ProductsBound []string    `json:"products_bound"`
			ProductsCount float64     `json:"products_count"`
		} `json:"coupons"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type BlacklistCreation struct {
	Status int `json:"status"`
	Data   struct {
		Uniqid string `json:"uniqid"`
	} `json:"data"`
	Message string      `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type BlacklistSpec struct {
	Status int `json:"status"`
	Data   struct {
		Blacklist struct {
			ID        int    `json:"id"`
			Uniqid    string `json:"uniqid"`
			Scope     string `json:"scope"`
			ShopID    int    `json:"shop_id"`
			Type      string `json:"type"`
			Data      string `json:"data"`
			Note      string `json:"note"`
			CreatedAt int    `json:"created_at"`
			UpdatedAt int    `json:"updated_at"`
			UpdatedBy int    `json:"updated_by"`
		} `json:"blacklist"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type Blacklist struct {
	Status int `json:"status"`
	Data   struct {
		Blacklists []struct {
			ID        int    `json:"id"`
			Uniqid    string `json:"uniqid"`
			Scope     string `json:"scope"`
			ShopID    int    `json:"shop_id"`
			Type      string `json:"type"`
			Data      string `json:"data"`
			Note      string `json:"note"`
			CreatedAt int    `json:"created_at"`
			UpdatedAt int    `json:"updated_at"`
			UpdatedBy int    `json:"updated_by"`
		} `json:"blacklists"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type Queries struct {
	Status int `json:"status"`
	Data   struct {
		Queries []struct {
			ID            int    `json:"id"`
			Uniqid        string `json:"uniqid"`
			ShopID        int    `json:"shop_id"`
			InvoiceID     string `json:"invoice_id"`
			CustomerEmail string `json:"customer_email"`
			Title         string `json:"title"`
			Status        string `json:"status"`
			DayValue      int    `json:"day_value"`
			Day           string `json:"day"`
			Month         string `json:"month"`
			Year          int    `json:"year"`
			CreatedAt     int    `json:"created_at"`
			UpdatedAt     int    `json:"updated_at"`
			UpdatedBy     int    `json:"updated_by"`
		} `json:"queries"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type UserFeedback struct {
	Status int `json:"status"`
	Data   struct {
		Feedback struct {
			ID        int         `json:"id"`
			Uniqid    string      `json:"uniqid"`
			InvoiceID string      `json:"invoice_id"`
			ProductID string      `json:"product_id"`
			ShopID    int         `json:"shop_id"`
			Message   string      `json:"message"`
			Reply     interface{} `json:"reply"`
			Score     float64     `json:"score"`
			Blocked   float64     `json:"blocked"`
			Appealed  float64     `json:"appealed"`
			CreatedAt int         `json:"created_at"`
			UpdatedAt int         `json:"updated_at"`
			UpdatedBy int         `json:"updated_by"`
			Invoice   struct {
				ID                        int         `json:"id"`
				Uniqid                    string      `json:"uniqid"`
				Type                      string      `json:"type"`
				ShopID                    int         `json:"shop_id"`
				ProductID                 string      `json:"product_id"`
				ProductType               string      `json:"product_type"`
				ProductTitle              string      `json:"product_title"`
				SubscriptionID            interface{} `json:"subscription_id"`
				SubscriptionTime          interface{} `json:"subscription_time"`
				Gateway                   string      `json:"gateway"`
				Quantity                  int         `json:"quantity"`
				Total                     float64     `json:"total"`
				TotalDisplay              float64     `json:"total_display"`
				CouponID                  interface{} `json:"coupon_id"`
				Discount                  float64     `json:"discount"`
				DiscountDisplay           float64     `json:"discount_display"`
				Currency                  string      `json:"currency"`
				ExchangeRate              float64     `json:"exchange_rate"`
				CryptoExchangeRate        float64     `json:"crypto_exchange_rate"`
				CustomerEmail             string      `json:"customer_email"`
				PaypalEmailDelivery       float64     `json:"paypal_email_delivery"`
				PaypalEmail               interface{} `json:"paypal_email"`
				PaypalOrderID             interface{} `json:"paypal_order_id"`
				PaypalAuthorizationID     interface{} `json:"paypal_authorization_id"`
				PaypalCaptureID           interface{} `json:"paypal_capture_id"`
				PaypalPayerEmail          interface{} `json:"paypal_payer_email"`
				PaypalFee                 float64     `json:"paypal_fee"`
				PaypalFeeCurrency         interface{} `json:"paypal_fee_currency"`
				PaypalAPICredentials      string      `json:"paypal_api_credentials"`
				LexOrderID                interface{} `json:"lex_order_id"`
				LexPaymentMethod          interface{} `json:"lex_payment_method"`
				LexSecret                 interface{} `json:"lex_secret"`
				SkrillEmail               interface{} `json:"skrill_email"`
				SkrillSid                 interface{} `json:"skrill_sid"`
				SkrillLink                interface{} `json:"skrill_link"`
				StripeID                  interface{} `json:"stripe_id"`
				StripeClientSecret        interface{} `json:"stripe_client_secret"`
				StripePriceID             interface{} `json:"stripe_price_id"`
				StripeSubscriptionID      interface{} `json:"stripe_subscription_id"`
				PerfectmoneyID            interface{} `json:"perfectmoney_id"`
				CashappQrcode             interface{} `json:"cashapp_qrcode"`
				CashappCashtag            interface{} `json:"cashapp_cashtag"`
				CashappNote               interface{} `json:"cashapp_note"`
				CryptoAddress             string      `json:"crypto_address"`
				CryptoAmount              float64     `json:"crypto_amount"`
				CryptoReceived            float64     `json:"crypto_received"`
				CryptoURI                 string      `json:"crypto_uri"`
				CryptoConfirmationsNeeded float64     `json:"crypto_confirmations_needed"`
				CryptoWalletVersion       string      `json:"crypto_wallet_version"`
				Country                   string      `json:"country"`
				Location                  string      `json:"location"`
				IP                        string      `json:"ip"`
				IsVpnOrProxy              bool        `json:"is_vpn_or_proxy"`
				UserAgent                 string      `json:"user_agent"`
				CustomFields              struct {
				} `json:"custom_fields"`
				DeveloperInvoice   bool        `json:"developer_invoice"`
				DeveloperTitle     interface{} `json:"developer_title"`
				DeveloperWebhook   interface{} `json:"developer_webhook"`
				DeveloperReturnURL interface{} `json:"developer_return_url"`
				FeePercentage      float64     `json:"fee_percentage"`
				ToProcess          float64     `json:"to_process"`
				Status             string      `json:"status"`
				DayValue           int         `json:"day_value"`
				Day                string      `json:"day"`
				Month              string      `json:"month"`
				Year               int         `json:"year"`
				CreatedAt          int         `json:"created_at"`
				UpdatedAt          int         `json:"updated_at"`
				UpdatedBy          int         `json:"updated_by"`
				Name               string      `json:"name"`
			} `json:"invoice"`
			Product struct {
				ID                        int           `json:"id"`
				Uniqid                    string        `json:"uniqid"`
				ShopID                    int           `json:"shop_id"`
				Type                      string        `json:"type"`
				Title                     string        `json:"title"`
				Currency                  string        `json:"currency"`
				Price                     float64       `json:"price"`
				PriceDisplay              float64       `json:"price_display"`
				Description               string        `json:"description"`
				ImageAttachment           interface{}   `json:"image_attachment"`
				FileAttachment            interface{}   `json:"file_attachment"`
				QuantityMin               float64       `json:"quantity_min"`
				QuantityMax               float64       `json:"quantity_max"`
				QuantityWarning           float64       `json:"quantity_warning"`
				Gateways                  []string      `json:"gateways"`
				CustomFields              []interface{} `json:"custom_fields"`
				CryptoConfirmationsNeeded float64       `json:"crypto_confirmations_needed"`
				MaxRiskLevel              float64       `json:"max_risk_level"`
				BlockVpnProxies           bool          `json:"block_vpn_proxies"`
				DeliveryText              string        `json:"delivery_text"`
				ServiceText               string        `json:"service_text"`
				StockDelimiter            string        `json:"stock_delimiter"`
				Stock                     float64       `json:"stock"`
				DynamicWebhook            string        `json:"dynamic_webhook"`
				SortPriority              float64       `json:"sort_priority"`
				Unlisted                  bool          `json:"unlisted"`
				OnHold                    float64       `json:"on_hold"`
				TermsOfService            string        `json:"terms_of_service"`
				Warranty                  float64       `json:"warranty"`
				WarrantyText              string        `json:"warranty_text"`
				Private                   bool          `json:"private"`
				CreatedAt                 int           `json:"created_at"`
				UpdatedAt                 int           `json:"updated_at"`
				UpdatedBy                 int           `json:"updated_by"`
				Name                      string        `json:"name"`
			} `json:"product"`
		} `json:"feedback"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type Feedback struct {
	Status int `json:"status"`
	Data   struct {
		Feedback []struct {
			ID            int         `json:"id"`
			Uniqid        string      `json:"uniqid"`
			InvoiceID     string      `json:"invoice_id"`
			ProductID     string      `json:"product_id"`
			ProductTitle  string      `json:"product_title"`
			ShopID        int         `json:"shop_id"`
			Message       string      `json:"message"`
			Reply         interface{} `json:"reply"`
			Score         float64     `json:"score"`
			Appealed      float64     `json:"appealed"`
			AppealOutcome string      `json:"appeal_outcome"`
			CreatedAt     int         `json:"created_at"`
			UpdatedAt     int         `json:"updated_at"`
			UpdatedBy     int         `json:"updated_by"`
		} `json:"feedback"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type Product struct {
	Status int `json:"status"`
	Data   struct {
		Products []struct {
			ID                        int           `json:"id"`
			Uniqid                    string        `json:"uniqid"`
			ShopID                    int           `json:"shop_id"`
			Type                      string        `json:"type"`
			Title                     string        `json:"title"`
			Currency                  string        `json:"currency"`
			Price                     float64       `json:"price"`
			PriceDisplay              float64       `json:"price_display"`
			Description               string        `json:"description"`
			ImageAttachment           string        `json:"image_attachment"`
			FileAttachment            interface{}   `json:"file_attachment"`
			QuantityMin               float64       `json:"quantity_min"`
			QuantityMax               float64       `json:"quantity_max"`
			QuantityWarning           float64       `json:"quantity_warning"`
			Gateways                  []string      `json:"gateways"`
			CustomFields              []interface{} `json:"custom_fields"`
			CryptoConfirmationsNeeded float64       `json:"crypto_confirmations_needed"`
			MaxRiskLevel              float64       `json:"max_risk_level"`
			BlockVpnProxies           bool          `json:"block_vpn_proxies"`
			DeliveryText              string        `json:"delivery_text"`
			ServiceText               string        `json:"service_text"`
			StockDelimiter            string        `json:"stock_delimiter"`
			Stock                     float64       `json:"stock"`
			DynamicWebhook            string        `json:"dynamic_webhook"`
			SortPriority              float64       `json:"sort_priority"`
			Unlisted                  bool          `json:"unlisted"`
			OnHold                    float64       `json:"on_hold"`
			TermsOfService            string        `json:"terms_of_service"`
			Warranty                  float64       `json:"warranty"`
			WarrantyText              string        `json:"warranty_text"`
			Private                   bool          `json:"private"`
			CreatedAt                 int           `json:"created_at"`
			UpdatedAt                 int           `json:"updated_at"`
			UpdatedBy                 int           `json:"updated_by"`
			ImageName                 string        `json:"image_name"`
			ImageStorage              string        `json:"image_storage"`
		} `json:"products"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}
