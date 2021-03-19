package model

type OrderHistory struct {
	Version                 int64   `json:"_version_"`
	AccountNumber           string  `json:"accountNumber"`
	Attention               string  `json:"attention"`
	BrandName               string  `json:"brandName"`
	ContactEmailAddress     string  `json:"contactEmailAddress"`
	ContactFaxNumber        string  `json:"contactFaxNumber"`
	ContactPhoneNumber      string  `json:"contactPhoneNumber"`
	Currency                string  `json:"currency"`
	CustomerUsername        string  `json:"customerUsername"`
	DeletedFlag             string  `json:"deletedFlag"`
	DeliveryMethod          string  `json:"deliveryMethod"`
	DocTyp                  string  `json:"docTyp"`
	ExtendedItemPrice       float64 `json:"extendedItemPrice"`
	FreightTerms            string  `json:"freightTerms"`
	FuelSurchargeFreight    int64   `json:"fuelSurchargeFreight"`
	GcomOrderNumber         string  `json:"gcomOrderNumber"`
	GsaSchedDesc            string  `json:"gsaSchedDesc"`
	GsaSchedule             string  `json:"gsaSchedule"`
	HigherLevelLineItem     string  `json:"higherLevelLineItem"`
	ID                      string  `json:"id"`
	ItemFreight             float64 `json:"itemFreight"`
	ItemNumber              string  `json:"itemNumber"`
	ItemShortDesc           string  `json:"itemShortDesc"`
	ItemTax                 float64 `json:"itemTax"`
	LastModified            string  `json:"lastModified"`
	LineFreightTerms        string  `json:"lineFreightTerms"`
	MfgPartNum              string  `json:"mfgPartNum"`
	OrderChannel            string  `json:"orderChannel"`
	OrderCreatedDateTime    string  `json:"orderCreatedDateTime"`
	OrderFreight            float64 `json:"orderFreight"`
	OrderHeaderFlag         string  `json:"orderHeaderFlag"`
	OrderSubtotal           float64 `json:"orderSubtotal"`
	OrderTax                float64 `json:"orderTax"`
	OrderTotal              float64 `json:"orderTotal"`
	OrderedBy               string  `json:"orderedBy"`
	OrigSlsOrdlineItm       string  `json:"origSlsOrdlineItm"`
	PaymentMethod           string  `json:"paymentMethod"`
	PurchaseOrderLineNumber string  `json:"purchaseOrderLineNumber"`
	PurchaseOrderNumber     string  `json:"purchaseOrderNumber"`
	PurchaseOrderType       string  `json:"purchaseOrderType"`
	Quantity                int64   `json:"quantity"`
	QuoteExpDt              string  `json:"quoteExpDt"`
	RequisitionerName       string  `json:"requisitionerName"`
	SOrg                    string  `json:"sOrg"`
	SalesOffice             string  `json:"salesOffice"`
	ContactID            	string  `json:"contactId"`
	OrderLineNumber      	string  `json:"orderLineNumber"`
	OrderNumber          	string  `json:"orderNumber"`
	SystemTimestamp      	string  `json:"systemTimestamp"`
	ShipAddress1            string  `json:"shipAddress1"`
	ShipCity                string  `json:"shipCity"`
	ShipCompanyName1        string  `json:"shipCompanyName1"`
	ShipCountry             string  `json:"shipCountry"`
	ShipPostalCode          string  `json:"shipPostalCode"`
	ShipRegion              string  `json:"shipRegion"`
	UnitPrice               float64 `json:"unitPrice"`
	Unspsc                  string  `json:"unspsc"`
	ZDPPItemCond            string  `json:"zDPPItemCond"`
}

type UserAccountIds struct {
	AccountNumber string `json:"accountNumber"`
	ContactId  	  string `json:"contactId"`
}
