package goSellix

type Order struct {
	Status int `json:"status"`
	Data   struct {
		Order struct {
			ID                        int           `json:"id"`
			Uniqid                    string        `json:"uniqid"`
			Type                      string        `json:"type"`
			Total                     int           `json:"total"`
			TotalDisplay              int           `json:"total_display"`
			ExchangeRate              int           `json:"exchange_rate"`
			CryptoExchangeRate        int           `json:"crypto_exchange_rate"`
			Currency                  string        `json:"currency"`
			ShopID                    int           `json:"shop_id"`
			Name                      string        `json:"name"`
			CustomerEmail             string        `json:"customer_email"`
			PaypalEmailDelivery       int           `json:"paypal_email_delivery"`
			ProductID                 string        `json:"product_id"`
			ProductTitle              string        `json:"product_title"`
			ProductType               string        `json:"product_type"`
			SubscriptionID            interface{}   `json:"subscription_id"`
			SubscriptionTime          interface{}   `json:"subscription_time"`
			Gateway                   string        `json:"gateway"`
			PaypalEmail               string        `json:"paypal_email"`
			PaypalOrderID             string        `json:"paypal_order_id"`
			PaypalPayerEmail          string        `json:"paypal_payer_email"`
			PaypalFee                 float64       `json:"paypal_fee"`
			LexOrderID                interface{}   `json:"lex_order_id"`
			LexPaymentMethod          interface{}   `json:"lex_payment_method"`
			StripeClientSecret        interface{}   `json:"stripe_client_secret"`
			SkrillEmail               interface{}   `json:"skrill_email"`
			SkrillSid                 interface{}   `json:"skrill_sid"`
			SkrillLink                interface{}   `json:"skrill_link"`
			PerfectmoneyID            interface{}   `json:"perfectmoney_id"`
			CryptoAddress             interface{}   `json:"crypto_address"`
			CryptoAmount              int           `json:"crypto_amount"`
			CryptoReceived            int           `json:"crypto_received"`
			CryptoURI                 interface{}   `json:"crypto_uri"`
			CryptoConfirmationsNeeded int           `json:"crypto_confirmations_needed"`
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
			Discount                  int           `json:"discount"`
			FeePercentage             int           `json:"fee_percentage"`
			DayValue                  int           `json:"day_value"`
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
				Asn            int     `json:"asn"`
				Organization   string  `json:"organization"`
				Latitude       float64 `json:"latitude"`
				Longitude      float64 `json:"longitude"`
				IsCrawler      int     `json:"is_crawler"`
				Timezone       string  `json:"timezone"`
				Mobile         int     `json:"mobile"`
				Host           string  `json:"host"`
				Proxy          int     `json:"proxy"`
				Vpn            int     `json:"vpn"`
				Tor            int     `json:"tor"`
				ActiveVpn      int     `json:"active_vpn"`
				ActiveTor      int     `json:"active_tor"`
				RecentAbuse    int     `json:"recent_abuse"`
				BotStatus      int     `json:"bot_status"`
				ConnectionType string  `json:"connection_type"`
				AbuseVelocity  string  `json:"abuse_velocity"`
				CreatedAt      int     `json:"created_at"`
				UpdatedAt      int     `json:"updated_at"`
			} `json:"ip_info"`
			Serials                 []string      `json:"serials"`
			Webhooks                []interface{} `json:"webhooks"`
			CryptoPayout            bool          `json:"crypto_payout"`
			CryptoPayoutTransaction interface{}   `json:"crypto_payout_transaction"`
			PaypalDispute           interface{}   `json:"paypal_dispute"`
			StatusHistory           []struct {
				ID        int    `json:"id"`
				InvoiceID string `json:"invoice_id"`
				Status    string `json:"status"`
				Details   string `json:"details"`
				CreatedAt int    `json:"created_at"`
			} `json:"status_history"`
			CryptoTransactions []interface{} `json:"crypto_transactions"`
			PaypalClientID     string        `json:"paypal_client_id"`
			Product            struct {
				ID                        int           `json:"id"`
				Uniqid                    string        `json:"uniqid"`
				ShopID                    int           `json:"shop_id"`
				Type                      string        `json:"type"`
				Title                     string        `json:"title"`
				Currency                  string        `json:"currency"`
				Price                     int           `json:"price"`
				PriceDisplay              int           `json:"price_display"`
				Description               string        `json:"description"`
				ImageAttachment           string        `json:"image_attachment"`
				FileAttachment            interface{}   `json:"file_attachment"`
				QuantityMin               int           `json:"quantity_min"`
				QuantityMax               int           `json:"quantity_max"`
				QuantityWarning           int           `json:"quantity_warning"`
				Gateways                  []string      `json:"gateways"`
				CustomFields              []interface{} `json:"custom_fields"`
				CryptoConfirmationsNeeded int           `json:"crypto_confirmations_needed"`
				MaxRiskLevel              int           `json:"max_risk_level"`
				BlockVpnProxies           bool          `json:"block_vpn_proxies"`
				DeliveryText              string        `json:"delivery_text"`
				ServiceText               string        `json:"service_text"`
				StockDelimiter            string        `json:"stock_delimiter"`
				Stock                     int           `json:"stock"`
				DynamicWebhook            string        `json:"dynamic_webhook"`
				SortPriority              int           `json:"sort_priority"`
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
					ID           int    `json:"id"`
					Uniqid       string `json:"uniqid"`
					Storage      string `json:"storage"`
					Name         string `json:"name"`
					OriginalName string `json:"original_name"`
					Extension    string `json:"extension"`
					ShopID       int    `json:"shop_id"`
					Size         int    `json:"size"`
					CreatedAt    int    `json:"created_at"`
				} `json:"image_attachment_info"`
			} `json:"product"`
		} `json:"order"`
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
