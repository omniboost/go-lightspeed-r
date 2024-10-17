package lightspeed_r

import (
	"encoding/json"
	"time"
)

const (
	LoadRelationSaleLines                     LoadRelationType = "SaleLines"
	LoadRelationSaleLinesDiscount             LoadRelationType = "SaleLines.Discount"
	LoadRelationSaleLinesTaxClass             LoadRelationType = "SaleLines.TaxClass"
	LoadRelationSaleLinesItem                 LoadRelationType = "SaleLines.Item"
	LoadRelationSalePayments                  LoadRelationType = "SalePayments"
	LoadRelationSalePaymentsPaymentType       LoadRelationType = "SalePayments.PaymentType"
	LoadRelationSalePaymentsCCCHarge          LoadRelationType = "SalePayments.CCCharge"
	LoadRelationSalePaymentsSaleAccounts      LoadRelationType = "SalePayments.SaleAccounts"
	LoadRelationCustomer                      LoadRelationType = "Customer"
	LoadRelationCustomerContact               LoadRelationType = "Customer.Contact"
	LoadRelationDiscount                      LoadRelationType = "Discount"
	LoadRelationTaxCategory                   LoadRelationType = "TaxCategory"
	LoadRelationTaxCategoryTaxCategoryClasses LoadRelationType = "TaxCategory.TaxCategoryClasses"
)

type LoadRelationType string

type Attributes struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Count    string `json:"count,omitempty"`
}

type AccountResp struct {
	Attributes Attributes `json:"@attributes"`
	Account    Account    `json:"Account"`
}

type Account struct {
	AccountID string `json:"accountID"`
	Name      string `json:"name"`
	Link      struct {
		Attributes struct {
			Href string `json:"href"`
		} `json:"@attributes"`
	} `json:"link"`
}

type CategoryResp struct {
	Attributes Attributes `json:"@attributes"`
	Category   Category   `json:"Category"`
}

type CategoriesResp struct {
	Attributes Attributes `json:"@attributes"`
	Category   []Category `json:"Category"`
}

type Category struct {
	CategoryID   string    `json:"categoryID"`
	Name         string    `json:"name"`
	NodeDepth    string    `json:"nodeDepth"`
	FullPathName string    `json:"fullPathName"`
	LeftNode     string    `json:"leftNode"`
	RightNode    string    `json:"rightNode"`
	ParentID     string    `json:"parentID"`
	CreateTime   time.Time `json:"createTime"`
	TimeStamp    time.Time `json:"timeStamp"`
}

type PaymentTypesResp struct {
	Attributes  Attributes    `json:"@attributes"`
	PaymentType []PaymentType `json:"PaymentType"`
}

type PaymentTypeResp struct {
	Attributes  Attributes  `json:"@attributes"`
	PaymentType PaymentType `json:"PaymentType"`
}

type PaymentType struct {
	PaymentTypeID         string `json:"paymentTypeID"`
	Name                  string `json:"name"`
	RequireCustomer       string `json:"requireCustomer"`
	Archived              string `json:"archived"`
	Code                  string `json:"code"`
	InternalReserved      string `json:"internalReserved"`
	Type                  string `json:"type"`
	RefundAsPaymentTypeID string `json:"refundAsPaymentTypeID"`
	Channel               string `json:"channel,omitempty"`
}

type TaxCategoriesResp struct {
	Attributes  Attributes    `json:"@attributes"`
	TaxCategory []TaxCategory `json:"TaxCategory"`
}

type TaxCategoryResp struct {
	Attributes  Attributes  `json:"@attributes"`
	TaxCategory TaxCategory `json:"TaxCategory"`
}

type TaxCategory struct {
	TaxCategoryID      string    `json:"taxCategoryID,omitempty"`
	IsTaxInclusive     string    `json:"isTaxInclusive,omitempty"`
	Tax1Name           string    `json:"tax1Name,omitempty"`
	Tax2Name           string    `json:"tax2Name,omitempty"`
	Tax1Rate           string    `json:"tax1Rate,omitempty"`
	Tax2Rate           string    `json:"tax2Rate,omitempty"`
	TimeStamp          time.Time `json:"timeStamp,omitempty"`
	TaxCategoryClasses struct {
		TaxCategoryClass TaxCategoryClasses `json:"TaxCategoryClass,omitempty"`
	} `json:"TaxCategoryClasses,omitempty"`
}

type TaxClassesResp struct {
	Attributes Attributes `json:"@attributes"`
	TaxClass   []TaxClass `json:"TaxClass"`
}

type TaxClassResp struct {
	Attributes Attributes `json:"@attributes"`
	TaxClass   TaxClass   `json:"TaxClass"`
}

type TaxClass struct {
	TaxClassID string    `json:"taxClassID,omitempty"`
	Name       string    `json:"name,omitempty"`
	TimeStamp  time.Time `json:"timeStamp,omitempty"`
}

type BearerToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type SalesResp struct {
	Attributes Attributes `json:"@attributes"`
	Sale       Sales      `json:"Sale"`
}

type TaxCategoryClass struct {
	TaxCategoryClassID string    `json:"taxCategoryClassID,omitempty"`
	Tax1Rate           string    `json:"tax1Rate,omitempty"`
	Tax2Rate           string    `json:"tax2Rate,omitempty"`
	TimeStamp          time.Time `json:"timeStamp,omitempty"`
	TaxCategoryID      string    `json:"taxCategoryID,omitempty"`
	TaxClassID         string    `json:"taxClassID,omitempty"`
}

type TaxCategoryClasses []TaxCategoryClass

func (s *TaxCategoryClasses) UnmarshalJSON(data []byte) error {
	// if the json doesn't start with an '[', force it to be an array
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}

	ss := []TaxCategoryClass{}
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}

	*s = ss
	return nil
}

type ItemPrice struct {
	Amount    string `json:"amount,omitempty"`
	UseTypeID string `json:"useTypeID,omitempty"`
	UseType   string `json:"useType,omitempty"`
}
type Prices struct {
	ItemPrice []ItemPrice `json:"ItemPrice,omitempty"`
}
type Item struct {
	ItemID           string    `json:"itemID,omitempty"`
	SystemSku        string    `json:"systemSku,omitempty"`
	DefaultCost      string    `json:"defaultCost,omitempty"`
	AvgCost          string    `json:"avgCost,omitempty"`
	Discountable     string    `json:"discountable,omitempty"`
	Tax              string    `json:"tax,omitempty"`
	Archived         string    `json:"archived,omitempty"`
	ItemType         string    `json:"itemType,omitempty"`
	Serialized       string    `json:"serialized,omitempty"`
	Description      string    `json:"description,omitempty"`
	ModelYear        string    `json:"modelYear,omitempty"`
	Upc              string    `json:"upc,omitempty"`
	Ean              string    `json:"ean,omitempty"`
	CustomSku        string    `json:"customSku,omitempty"`
	ManufacturerSku  string    `json:"manufacturerSku,omitempty"`
	CreateTime       time.Time `json:"createTime,omitempty"`
	TimeStamp        time.Time `json:"timeStamp,omitempty"`
	CategoryID       string    `json:"categoryID,omitempty"`
	TaxClassID       string    `json:"taxClassID,omitempty"`
	DepartmentID     string    `json:"departmentID,omitempty"`
	ItemMatrixID     string    `json:"itemMatrixID,omitempty"`
	ItemAttributesID string    `json:"itemAttributesID,omitempty"`
	ManufacturerID   string    `json:"manufacturerID,omitempty"`
	NoteID           string    `json:"noteID,omitempty"`
	SeasonID         string    `json:"seasonID,omitempty"`
	DefaultVendorID  string    `json:"defaultVendorID,omitempty"`
	Prices           Prices    `json:"Prices,omitempty"`
}
type SaleLine struct {
	SaleLineID              string    `json:"saleLineID,omitempty"`
	CreateTime              time.Time `json:"createTime,omitempty"`
	TimeStamp               time.Time `json:"timeStamp,omitempty"`
	UnitQuantity            string    `json:"unitQuantity,omitempty"`
	UnitPrice               string    `json:"unitPrice,omitempty"`
	NormalUnitPrice         string    `json:"normalUnitPrice,omitempty"`
	DiscountAmount          string    `json:"discountAmount,omitempty"`
	DiscountPercent         string    `json:"discountPercent,omitempty"`
	AvgCost                 string    `json:"avgCost,omitempty"`
	FifoCost                string    `json:"fifoCost,omitempty"`
	Tax                     string    `json:"tax,omitempty"`
	Tax1Rate                string    `json:"tax1Rate,omitempty"`
	Tax2Rate                string    `json:"tax2Rate,omitempty"`
	IsLayaway               string    `json:"isLayaway,omitempty"`
	IsWorkorder             string    `json:"isWorkorder,omitempty"`
	IsSpecialOrder          string    `json:"isSpecialOrder,omitempty"`
	DisplayableSubtotal     string    `json:"displayableSubtotal,omitempty"`
	DisplayableUnitPrice    string    `json:"displayableUnitPrice,omitempty"`
	CalcLineDiscount        string    `json:"calcLineDiscount,omitempty"`
	CalcTransactionDiscount string    `json:"calcTransactionDiscount,omitempty"`
	CalcTotal               string    `json:"calcTotal,omitempty"`
	CalcSubtotal            string    `json:"calcSubtotal,omitempty"`
	CalcTax1                string    `json:"calcTax1,omitempty"`
	CalcTax2                string    `json:"calcTax2,omitempty"`
	TaxClassID              string    `json:"taxClassID,omitempty"`
	CustomerID              string    `json:"customerID,omitempty"`
	DiscountID              string    `json:"discountID,omitempty"`
	EmployeeID              string    `json:"employeeID,omitempty"`
	ItemID                  string    `json:"itemID,omitempty"`
	NoteID                  string    `json:"noteID,omitempty"`
	ParentSaleLineID        string    `json:"parentSaleLineID,omitempty"`
	ShopID                  string    `json:"shopID,omitempty"`
	SaleID                  string    `json:"saleID,omitempty"`
	TaxClass                TaxClass  `json:"TaxClass,omitempty"`
	Item                    Item      `json:"Item,omitempty"`
}
type SaleLines []SaleLine

func (s *SaleLines) UnmarshalJSON(data []byte) error {
	// if the json doesn't start with an '[', force it to be an array
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}

	ss := []SaleLine{}
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}

	*s = ss
	return nil
}

type Tax struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Taxable   string `json:"taxable,omitempty"`
	Rate      string `json:"rate,omitempty"`
	Amount    string `json:"amount,omitempty"`
	Taxname   string `json:"taxname,omitempty"`
	Subtotal  string `json:"subtotal,omitempty"`
	Rate2     string `json:"rate2,omitempty"`
	Amount2   string `json:"amount2,omitempty"`
	Taxname2  string `json:"taxname2,omitempty"`
	Subtotal2 string `json:"subtotal2,omitempty"`
}
type TaxClassTotals struct {
	Tax Tax `json:"Tax,omitempty"`
}
type ContactPhone struct {
	Number  string `json:"number,omitempty"`
	UseType string `json:"useType,omitempty"`
}
type Phones struct {
	ContactPhone ContactPhone `json:"ContactPhone,omitempty"`
}
type ContactEmail struct {
	Address string `json:"address,omitempty"`
	UseType string `json:"useType,omitempty"`
}

type ContactAddress struct {
	Address1    string `json:"address1,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	StateCode   string `json:"stateCode,omitempty"`
	Zip         string `json:"zip,omitempty"`
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
}

type Emails struct {
	ContactEmail ContactEmail `json:"ContactEmail,omitempty"`
}
type Contact struct {
	ContactID string      `json:"contactID,omitempty"`
	Custom    string      `json:"custom,omitempty"`
	NoEmail   string      `json:"noEmail,omitempty"`
	NoPhone   string      `json:"noPhone,omitempty"`
	NoMail    string      `json:"noMail,omitempty"`
	TimeStamp time.Time   `json:"timeStamp,omitempty"`
	Addresses interface{} `json:"Addresses,omitempty"`
	Phones    interface{} `json:"Phones,omitempty"`
	Emails    Emails      `json:"Emails,omitempty"`
	Websites  string      `json:"Websites,omitempty"`
}

type Customer struct {
	CustomerID                string    `json:"customerID,omitempty"`
	FirstName                 string    `json:"firstName,omitempty"`
	LastName                  string    `json:"lastName,omitempty"`
	Dob                       string    `json:"dob,omitempty"`
	Archived                  string    `json:"archived,omitempty"`
	Title                     string    `json:"title,omitempty"`
	Company                   string    `json:"company,omitempty"`
	CompanyRegistrationNumber string    `json:"companyRegistrationNumber,omitempty"`
	VatNumber                 string    `json:"vatNumber,omitempty"`
	CreateTime                time.Time `json:"createTime,omitempty"`
	TimeStamp                 time.Time `json:"timeStamp,omitempty"`
	ContactID                 string    `json:"contactID,omitempty"`
	CreditAccountID           string    `json:"creditAccountID,omitempty"`
	CustomerTypeID            string    `json:"customerTypeID,omitempty"`
	DiscountID                string    `json:"discountID,omitempty"`
	EmployeeID                string    `json:"employeeID,omitempty"`
	NoteID                    string    `json:"noteID,omitempty"`
	TaxCategoryID             string    `json:"taxCategoryID,omitempty"`
	MeasurementID             string    `json:"measurementID,omitempty"`
	Contact                   Contact   `json:"Contact,omitempty"`
}

type Sales []Sale

func (s *Sales) UnmarshalJSON(data []byte) error {
	// if the json doesn't start with an '[', force it to be an array
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}

	ss := []Sale{}
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}

	*s = ss
	return nil
}

type SalePayment struct {
	SalePaymentID   string      `json:"salePaymentID,omitempty"`
	Amount          string      `json:"amount,omitempty"`
	CreateTime      time.Time   `json:"createTime,omitempty"`
	Archived        string      `json:"archived,omitempty"`
	RemoteReference string      `json:"remoteReference,omitempty"`
	TipAmount       string      `json:"tipAmount,omitempty"`
	PaymentID       string      `json:"paymentID,omitempty"`
	SaleID          string      `json:"saleID,omitempty"`
	PaymentTypeID   string      `json:"paymentTypeID,omitempty"`
	CcChargeID      string      `json:"ccChargeID,omitempty"`
	RefPaymentID    string      `json:"refPaymentID,omitempty"`
	RegisterID      string      `json:"registerID,omitempty"`
	EmployeeID      string      `json:"employeeID,omitempty"`
	CreditAccountID string      `json:"creditAccountID,omitempty"`
	PaymentType     PaymentType `json:"PaymentType,omitempty"`
}

type SalePayments []SalePayment

func (s *SalePayments) UnmarshalJSON(data []byte) error {
	// if the json doesn't start with an '[', force it to be an array
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}

	ss := []SalePayment{}
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}

	*s = ss
	return nil
}

type Sale struct {
	SaleID                StringInt   `json:"saleID,omitempty"`
	TimeStamp             time.Time   `json:"timeStamp,omitempty"`
	DiscountPercent       string      `json:"discountPercent,omitempty"`
	Completed             string      `json:"completed,omitempty"`
	Archived              string      `json:"archived,omitempty"`
	Voided                string      `json:"voided,omitempty"`
	EnablePromotions      string      `json:"enablePromotions,omitempty"`
	IsTaxInclusive        string      `json:"isTaxInclusive,omitempty"`
	CreateTime            time.Time   `json:"createTime,omitempty"`
	UpdateTime            time.Time   `json:"updateTime,omitempty"`
	CompleteTime          time.Time   `json:"completeTime,omitempty"`
	ReferenceNumber       string      `json:"referenceNumber,omitempty"`
	ReferenceNumberSource string      `json:"referenceNumberSource,omitempty"`
	Tax1Rate              string      `json:"tax1Rate,omitempty"`
	Tax2Rate              string      `json:"tax2Rate,omitempty"`
	Change                string      `json:"change,omitempty"`
	ReceiptPreference     string      `json:"receiptPreference,omitempty"`
	DisplayableSubtotal   string      `json:"displayableSubtotal,omitempty"`
	TicketNumber          string      `json:"ticketNumber,omitempty"`
	CalcDiscount          string      `json:"calcDiscount,omitempty"`
	CalcTotal             string      `json:"calcTotal,omitempty"`
	CalcSubtotal          string      `json:"calcSubtotal,omitempty"`
	CalcTaxable           string      `json:"calcTaxable,omitempty"`
	CalcNonTaxable        string      `json:"calcNonTaxable,omitempty"`
	CalcAvgCost           string      `json:"calcAvgCost,omitempty"`
	CalcFIFOCost          string      `json:"calcFIFOCost,omitempty"`
	CalcTax1              string      `json:"calcTax1,omitempty"`
	CalcTax2              string      `json:"calcTax2,omitempty"`
	CalcPayments          string      `json:"calcPayments,omitempty"`
	CalcItemFees          string      `json:"calcItemFees,omitempty"`
	Total                 string      `json:"total,omitempty"`
	TotalDue              string      `json:"totalDue,omitempty"`
	DisplayableTotal      string      `json:"displayableTotal,omitempty"`
	Balance               string      `json:"balance,omitempty"`
	CustomerID            string      `json:"customerID,omitempty"`
	DiscountID            string      `json:"discountID,omitempty"`
	EmployeeID            string      `json:"employeeID,omitempty"`
	QuoteID               string      `json:"quoteID,omitempty"`
	RegisterID            string      `json:"registerID,omitempty"`
	ShipToID              string      `json:"shipToID,omitempty"`
	ShopID                string      `json:"shopID,omitempty"`
	TaxCategoryID         string      `json:"taxCategoryID,omitempty"`
	TaxCategory           TaxCategory `json:"TaxCategory,omitempty"`
	TaxTotal              string      `json:"taxTotal,omitempty"`
	SaleLines             struct {
		SaleLine SaleLines `json:"SaleLine,omitempty"`
	} `json:"SaleLines,omitempty"`
	TaxClassTotals TaxClassTotals `json:"TaxClassTotals,omitempty"`
	Customer       struct {
		Customer Customer `json:"Customer,omitempty"`
	} `json:"Customer,omitempty"`
	SalePayments struct {
		SalePayment SalePayments `json:"salePayment,omitempty"`
	} `json:"salePayments,omitempty"`
}
