package model

type GDoc struct {
	Account struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		Website           string      `json:"website"`
		LogoURL           interface{} `json:"logo_url"`
		ContractRut       interface{} `json:"contract_rut"`
		ContractName      string      `json:"contract_name"`
		ContractType      interface{} `json:"contract_type"`
		ContractAddress   string      `json:"contract_address"`
		ContractCity      string      `json:"contract_city"`
		ContractContact   interface{} `json:"contract_contact"`
		BankRut           interface{} `json:"bank_rut"`
		BankName          interface{} `json:"bank_name"`
		BankName2         int         `json:"bank_name2"`
		BankAccountHolder interface{} `json:"bank_account_holder"`
		BankAccountNumber interface{} `json:"bank_account_number"`
		PaymentMethod     string      `json:"payment_method"`
		Created           interface{} `json:"created"`
		Modified          interface{} `json:"modified"`
		MembershipInfo    interface{} `json:"membership_info"`
		SfAccountID       string      `json:"sf_account_id"`
		CdCompanyID       int         `json:"cd_company_id"`
		ContactName       interface{} `json:"contact_name"`
		ContactEmail      string      `json:"contact_email"`
	} `json:"account"`
	Category struct {
		ID                 int    `json:"id"`
		ParentID           int    `json:"parent_id"`
		Name               string `json:"name"`
		SpanishName        string `json:"spanish_name"`
		DealbankCategoryID int    `json:"dealbank_category_id"`
		GUID               string `json:"guid"`
	} `json:"category"`
	Country struct {
		ID       int         `json:"id"`
		Name     string      `json:"name"`
		Slug     string      `json:"slug"`
		Code     string      `json:"code"`
		URL      string      `json:"url"`
		Tld      interface{} `json:"tld"`
		Currency interface{} `json:"currency"`
	} `json:"country"`
	Images []struct {
		ID            int         `json:"id"`
		URL           string      `json:"url"`
		MigratedURL   interface{} `json:"migrated_url"`
		OpportunityID int         `json:"opportunity_id"`
		Type          string      `json:"type"`
	} `json:"images"`
	Opportunity struct {
		ID                            int         `json:"id"`
		CdDealID                      int         `json:"cd_deal_id"`
		CountryID                     int         `json:"country_id"`
		CountryDescription            interface{} `json:"country_description"`
		PartnerCountry                interface{} `json:"partner_country"`
		PartnerOriginalPrice          interface{} `json:"partner_original_price"`
		PartnerSpecialPrice           interface{} `json:"partner_special_price"`
		Title                         string      `json:"title"`
		ShortTitle                    string      `json:"short_title"`
		CouponTitle                   string      `json:"coupon_title"`
		NlTitle                       string      `json:"nl_title"`
		Description                   interface{} `json:"description"`
		Details                       string      `json:"details"`
		CompanyID                     int         `json:"company_id"`
		CompanyDescription            interface{} `json:"company_description"`
		Picture                       interface{} `json:"picture"`
		Created                       string      `json:"created"`
		Modified                      string      `json:"modified"`
		ExpirationAlert               bool        `json:"expiration_alert"`
		Commission                    float64     `json:"commission"`
		CommissionTax                 float64     `json:"commission_tax"`
		Highlights                    string      `json:"highlights"`
		SfOpportunityID               string      `json:"sf_opportunity_id"`
		Mulligan                      bool        `json:"mulligan"`
		MulliganDays                  int         `json:"mulligan_days"`
		MulliganTime                  int         `json:"mulligan_time"`
		InstantSending                bool        `json:"instant_sending"`
		SendingCoupons                bool        `json:"sending_coupons"`
		SyncAdditionalInfo            bool        `json:"sync_additional_info"`
		PrimaryDealServiceID          int         `json:"primary_deal_service_id"`
		PrimaryDealServiceDescription interface{} `json:"primary_deal_service_description"`
		Status                        string      `json:"status"`
		NotesToEditorial              string      `json:"notes_to_editorial"`
		ImageNotes                    string      `json:"image_notes"`
		SalesPoints                   string      `json:"sales_points"`
		CategoryID                    interface{} `json:"category_id"`
		Slug                          interface{} `json:"slug"`
		LogisticsNotes                interface{} `json:"logistics_notes"`
		FulfillmentMethod             interface{} `json:"fulfillment_method"`
		RecordType                    string      `json:"record_type"`
		GetawaysTags                  interface{} `json:"getaways_tags"`
		BusinessUnit                  string      `json:"business_unit"`
		WfID                          int         `json:"wf_id"`
		HasTraits                     bool        `json:"has_traits"`
		ClonedFrom                    string      `json:"cloned_from"`
	} `json:"opportunity"`
	OptionList []struct {
		ID                     int         `json:"id"`
		Value                  float64     `json:"value"`
		Discount               int         `json:"discount"`
		FinalPrice             float64     `json:"final_price"`
		CouponTitle            string      `json:"coupon_title"`
		Slug                   interface{} `json:"slug"`
		SoldQty                int         `json:"sold_qty"`
		MaxCoupons             int         `json:"max_coupons"`
		SfMultidealID          string      `json:"sf_multideal_id"`
		MaxCouponsPerBuy       int         `json:"max_coupons_per_buy"`
		InitCoupons            int         `json:"init_coupons"`
		OpportunityID          int         `json:"opportunity_id"`
		OpportunityDescription interface{} `json:"opportunity_description"`
		ProductSKU             interface{} `json:"product_SKU"`
		DisplayOrder           int         `json:"display_order"`
		ProductWeight          float64     `json:"product_weight"`
		CdDealID               interface{} `json:"cd_deal_id"`
		OptionTraits           interface{} `json:"option_traits"`
		VoucherTitle           string      `json:"voucher_title"`
	} `json:"optionList"`
	Traits struct {
		Name         string `json:"name"`
		TraitSummary []struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Position int    `json:"position"`
			Values   []struct {
				Name     string `json:"name"`
				Position int    `json:"position"`
				Options  []int  `json:"options"`
			} `json:"values"`
		} `json:"traitSummary"`
	} `json:"traits"`
}