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
	CategoryID   StringInt `json:"categoryID"`
	Name         string    `json:"name"`
	NodeDepth    string    `json:"nodeDepth"`
	FullPathName string    `json:"fullPathName"`
	LeftNode     StringInt `json:"leftNode"`
	RightNode    StringInt `json:"rightNode"`
	ParentID     StringInt `json:"parentID"`
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
	PaymentTypeID         StringInt `json:"paymentTypeID"`
	Name                  string    `json:"name"`
	RequireCustomer       string    `json:"requireCustomer"`
	Archived              string    `json:"archived"`
	Code                  string    `json:"code"`
	InternalReserved      string    `json:"internalReserved"`
	Type                  string    `json:"type"`
	RefundAsPaymentTypeID StringInt `json:"refundAsPaymentTypeID"`
	Channel               string    `json:"channel,omitempty"`
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
	TaxCategoryID      StringInt   `json:"taxCategoryID,omitempty"`
	IsTaxInclusive     string      `json:"isTaxInclusive,omitempty"`
	Tax1Name           string      `json:"tax1Name,omitempty"`
	Tax2Name           string      `json:"tax2Name,omitempty"`
	Tax1Rate           StringFloat `json:"tax1Rate,omitempty"`
	Tax2Rate           StringFloat `json:"tax2Rate,omitempty"`
	TimeStamp          time.Time   `json:"timeStamp,omitempty"`
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
	TaxClassID StringInt `json:"taxClassID,omitempty"`
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
	TaxCategoryClassID StringInt   `json:"taxCategoryClassID,omitempty"`
	Tax1Rate           StringFloat `json:"tax1Rate,omitempty"`
	Tax2Rate           StringFloat `json:"tax2Rate,omitempty"`
	TimeStamp          time.Time   `json:"timeStamp,omitempty"`
	TaxCategoryID      StringInt   `json:"taxCategoryID,omitempty"`
	TaxClassID         StringInt   `json:"taxClassID,omitempty"`
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
	Amount    StringFloat `json:"amount,omitempty"`
	UseTypeID StringInt   `json:"useTypeID,omitempty"`
	UseType   string      `json:"useType,omitempty"`
}
type Prices struct {
	ItemPrice []ItemPrice `json:"ItemPrice,omitempty"`
}
type Discount struct {
	DiscountID      StringInt   `json:"discountID,omitempty"`
	Name            string      `json:"name,omitempty"`
	DiscountAmount  StringFloat `json:"discountAmount,omitempty"`
	DiscountPercent StringFloat `json:"discountPercent,omitempty"`
	RequireCustomer string      `json:"requireCustomer,omitempty"`
	Archived        string      `json:"archived,omitempty"`
	SourceID        StringInt   `json:"sourceID,omitempty"`
	CreateTime      time.Time   `json:"createTime,omitempty"`
	TimeStamp       time.Time   `json:"timeStamp,omitempty"`
}

type Item struct {
	ItemID           StringInt   `json:"itemID,omitempty"`
	SystemSku        StringInt   `json:"systemSku,omitempty"`
	DefaultCost      StringFloat `json:"defaultCost,omitempty"`
	AvgCost          StringFloat `json:"avgCost,omitempty"`
	Discountable     string      `json:"discountable,omitempty"`
	Tax              string      `json:"tax,omitempty"`
	Archived         string      `json:"archived,omitempty"`
	ItemType         string      `json:"itemType,omitempty"`
	Serialized       string      `json:"serialized,omitempty"`
	Description      string      `json:"description,omitempty"`
	ModelYear        StringInt   `json:"modelYear,omitempty"`
	Upc              string      `json:"upc,omitempty"`
	Ean              string      `json:"ean,omitempty"`
	CustomSku        string      `json:"customSku,omitempty"`
	ManufacturerSku  string      `json:"manufacturerSku,omitempty"`
	CreateTime       time.Time   `json:"createTime,omitempty"`
	TimeStamp        time.Time   `json:"timeStamp,omitempty"`
	CategoryID       StringInt   `json:"categoryID,omitempty"`
	TaxClassID       StringInt   `json:"taxClassID,omitempty"`
	DepartmentID     StringInt   `json:"departmentID,omitempty"`
	ItemMatrixID     StringInt   `json:"itemMatrixID,omitempty"`
	ItemAttributesID StringInt   `json:"itemAttributesID,omitempty"`
	ManufacturerID   StringInt   `json:"manufacturerID,omitempty"`
	NoteID           StringInt   `json:"noteID,omitempty"`
	SeasonID         StringInt   `json:"seasonID,omitempty"`
	DefaultVendorID  StringInt   `json:"defaultVendorID,omitempty"`
	Prices           Prices      `json:"Prices,omitempty"`
}
type SaleLine struct {
	SaleLineID              StringInt   `json:"saleLineID,omitempty"`
	CreateTime              time.Time   `json:"createTime,omitempty"`
	TimeStamp               time.Time   `json:"timeStamp,omitempty"`
	UnitQuantity            StringInt   `json:"unitQuantity,omitempty"`
	UnitPrice               StringFloat `json:"unitPrice,omitempty"`
	NormalUnitPrice         StringFloat `json:"normalUnitPrice,omitempty"`
	DiscountAmount          StringFloat `json:"discountAmount,omitempty"`
	DiscountPercent         StringFloat `json:"discountPercent,omitempty"`
	AvgCost                 StringFloat `json:"avgCost,omitempty"`
	FifoCost                StringFloat `json:"fifoCost,omitempty"`
	Tax                     string      `json:"tax,omitempty"`
	Tax1Rate                StringFloat `json:"tax1Rate,omitempty"`
	Tax2Rate                StringFloat `json:"tax2Rate,omitempty"`
	IsLayaway               string      `json:"isLayaway,omitempty"`
	IsWorkorder             string      `json:"isWorkorder,omitempty"`
	IsSpecialOrder          string      `json:"isSpecialOrder,omitempty"`
	DisplayableSubtotal     StringFloat `json:"displayableSubtotal,omitempty"`
	DisplayableUnitPrice    StringFloat `json:"displayableUnitPrice,omitempty"`
	LineType                string      `json:"lineType,omitempty"`
	CalcLineDiscount        StringFloat `json:"calcLineDiscount,omitempty"`
	CalcTransactionDiscount StringFloat `json:"calcTransactionDiscount,omitempty"`
	CalcTotal               StringFloat `json:"calcTotal,omitempty"`
	CalcSubtotal            StringFloat `json:"calcSubtotal,omitempty"`
	CalcTax1                StringFloat `json:"calcTax1,omitempty"`
	CalcTax2                StringFloat `json:"calcTax2,omitempty"`
	TaxClassID              StringInt   `json:"taxClassID,omitempty"`
	CustomerID              StringInt   `json:"customerID,omitempty"`
	DiscountID              StringInt   `json:"discountID,omitempty"`
	EmployeeID              StringInt   `json:"employeeID,omitempty"`
	ItemID                  StringInt   `json:"itemID,omitempty"`
	NoteID                  StringInt   `json:"noteID,omitempty"`
	ParentSaleLineID        StringInt   `json:"parentSaleLineID,omitempty"`
	ShopID                  StringInt   `json:"shopID,omitempty"`
	SaleID                  StringInt   `json:"saleID,omitempty"`
	TaxClass                TaxClass    `json:"TaxClass,omitempty"`
	Discount                Discount    `json:"Discount,omitempty"`
	Item                    Item        `json:"Item,omitempty"`
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
	ID        StringInt   `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Taxable   StringFloat `json:"taxable,omitempty"`
	Rate      StringInt   `json:"rate,omitempty"`
	Amount    StringFloat `json:"amount,omitempty"`
	Taxname   string      `json:"taxname,omitempty"`
	Subtotal  StringFloat `json:"subtotal,omitempty"`
	Rate2     StringInt   `json:"rate2,omitempty"`
	Amount2   StringFloat `json:"amount2,omitempty"`
	Taxname2  string      `json:"taxname2,omitempty"`
	Subtotal2 StringFloat `json:"subtotal2,omitempty"`
}

type Taxes []Tax

func (s *Taxes) UnmarshalJSON(data []byte) error {
	// if the json doesn't start with an '[', force it to be an array
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}

	ss := []Tax{}
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}

	*s = ss
	return nil
}

type TaxClassTotals struct {
	Tax Taxes `json:"Tax,omitempty"`
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
	ContactID StringInt   `json:"contactID,omitempty"`
	Custom    string      `json:"custom,omitempty"`
	NoEmail   string      `json:"noEmail,omitempty"`
	NoPhone   string      `json:"noPhone,omitempty"`
	NoMail    string      `json:"noMail,omitempty"`
	TimeStamp time.Time   `json:"timeStamp,omitempty"`
	Addresses interface{} `json:"addresses,omitempty"`
	Phones    interface{} `json:"phones,omitempty"`
	Emails    Emails      `json:"emails,omitempty"`
	Websites  string      `json:"websites,omitempty"`
}

type Customer struct {
	CustomerID                StringInt `json:"customerID,omitempty"`
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
	ContactID                 StringInt `json:"contactID,omitempty"`
	CreditAccountID           StringInt `json:"creditAccountID,omitempty"`
	CustomerTypeID            StringInt `json:"customerTypeID,omitempty"`
	DiscountID                StringInt `json:"ciscountID,omitempty"`
	EmployeeID                StringInt `json:"employeeID,omitempty"`
	NoteID                    StringInt `json:"noteID,omitempty"`
	TaxCategoryID             StringInt `json:"taxCategoryID,omitempty"`
	MeasurementID             StringInt `json:"measurementID,omitempty"`
	Contact                   Contact   `json:"contact,omitempty"`
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
	SalePaymentID   StringInt   `json:"salePaymentID,omitempty"`
	Amount          StringFloat `json:"amount,omitempty"`
	CreateTime      time.Time   `json:"createTime,omitempty"`
	Archived        string      `json:"archived,omitempty"`
	RemoteReference string      `json:"remoteReference,omitempty"`
	TipAmount       StringFloat `json:"tipAmount,omitempty"`
	PaymentID       string      `json:"paymentID,omitempty"`
	SaleID          StringInt   `json:"saleID,omitempty"`
	PaymentTypeID   StringInt   `json:"paymentTypeID,omitempty"`
	CcChargeID      StringInt   `json:"ccChargeID,omitempty"`
	RefPaymentID    StringInt   `json:"refPaymentID,omitempty"`
	RegisterID      StringInt   `json:"registerITD,omitempty"`
	EmployeeID      StringInt   `json:"employeeID,omitempty"`
	CreditAccountID StringInt   `json:"creditAccountID,omitempty"`
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

type ItemFeeResp struct {
	Attributes Attributes `json:"@attributes"`
	ItemFee    ItemFees   `json:"ItemFee"`
}

type ItemFee struct {
	ItemFeeID         StringInt   `json:"itemFeeID,omitempty"`
	Name              string      `json:"name,omitempty"`
	CalculationMethod string      `json:"calculationMethod,omitempty"`
	FeeValue          StringFloat `json:"feeValue,omitempty"`
	Taxable           string      `json:"taxable,omitempty"`
	Discountable      string      `json:"discountable,omitempty"`
	NonRefundable     string      `json:"nonRefundable,omitempty"`
	Archived          string      `json:"archived,omitempty"`
	CreateTime        time.Time   `json:"createTime,omitempty"`
	Timestamp         time.Time   `json:"timestamp,omitempty"`
}

type ItemFees []ItemFee

func (s *ItemFees) UnmarshalJSON(data []byte) error {
	// if the json doesn't start with an '[', force it to be an array
	if data[0] != '[' {
		data = []byte("[" + string(data) + "]")
	}

	ss := []ItemFee{}
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
	DiscountPercent       StringFloat `json:"discountPercent,omitempty"`
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
	Tax1Rate              StringFloat `json:"tax1Rate,omitempty"`
	Tax2Rate              StringFloat `json:"tax2Rate,omitempty"`
	Change                StringFloat `json:"change,omitempty"`
	ReceiptPreference     string      `json:"receiptPreference,omitempty"`
	DisplayableSubtotal   string      `json:"displayableSubtotal,omitempty"`
	TicketNumber          string      `json:"ticketNumber,omitempty"`
	CalcDiscount          StringFloat `json:"calcDiscount,omitempty"`
	CalcTotal             StringFloat `json:"calcTotal,omitempty"`
	CalcSubtotal          StringFloat `json:"calcSubtotal,omitempty"`
	CalcTaxable           StringFloat `json:"calcTaxable,omitempty"`
	CalcNonTaxable        StringFloat `json:"calcNonTaxable,omitempty"`
	CalcAvgCost           StringFloat `json:"calcAvgCost,omitempty"`
	CalcFIFOCost          StringFloat `json:"calcFIFOCost,omitempty"`
	CalcTax1              StringFloat `json:"calcTax1,omitempty"`
	CalcTax2              StringFloat `json:"calcTax2,omitempty"`
	CalcPayments          StringFloat `json:"calcPayments,omitempty"`
	CalcItemFees          string      `json:"calcItemFees,omitempty"`
	Total                 StringFloat `json:"total,omitempty"`
	TotalDue              StringFloat `json:"totalDue,omitempty"`
	DisplayableTotal      StringFloat `json:"displayableTotal,omitempty"`
	Balance               StringFloat `json:"balance,omitempty"`
	CustomerID            StringInt   `json:"customerID,omitempty"`
	DiscountID            StringInt   `json:"discountID,omitempty"`
	EmployeeID            StringInt   `json:"employeeID,omitempty"`
	QuoteID               StringInt   `json:"quoteID,omitempty"`
	RegisterID            StringInt   `json:"registerID,omitempty"`
	ShipToID              StringInt   `json:"shipToID,omitempty"`
	ShopID                StringInt   `json:"shopID,omitempty"`
	TaxCategoryID         StringInt   `json:"taxCategoryID,omitempty"`
	TaxCategory           TaxCategory `json:"TaxCategory,omitempty"`
	TaxTotal              string      `json:"taxTotal,omitempty"`
	SaleLines             struct {
		SaleLine SaleLines `json:"SaleLine,omitempty"`
	} `json:"SaleLines,omitempty"`
	TaxClassTotals TaxClassTotals `json:"TaxClassTotals,omitempty"`
	Customer       Customer       `json:"customer,omitempty"`
	SalePayments   struct {
		SalePayment SalePayments `json:"salePayment,omitempty"`
	} `json:"salePayments,omitempty"`
}
