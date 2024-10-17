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
	Next     string `json:"Next"`
	Previous string `json:"Previous"`
	Count    string `json:"Count,omitempty"`
}

type AccountResp struct {
	Attributes Attributes `json:"@attributes"`
	Account    Account    `json:"Account"`
}

type Account struct {
	AccountID string `json:"AccountID"`
	Name      string `json:"Name"`
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
	TaxCategoryID      string    `json:"TaxCategoryID,omitempty"`
	IsTaxInclusive     string    `json:"IsTaxInclusive,omitempty"`
	Tax1Name           string    `json:"Tax1Name,omitempty"`
	Tax2Name           string    `json:"Tax2Name,omitempty"`
	Tax1Rate           string    `json:"Tax1Rate,omitempty"`
	Tax2Rate           string    `json:"Tax2Rate,omitempty"`
	TimeStamp          time.Time `json:"TimeStamp,omitempty"`
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
	TaxClassID string    `json:"TaxClassID,omitempty"`
	Name       string    `json:"Name,omitempty"`
	TimeStamp  time.Time `json:"TimeStamp,omitempty"`
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
	TaxCategoryClassID string    `json:"TaxCategoryClassID,omitempty"`
	Tax1Rate           string    `json:"Tax1Rate,omitempty"`
	Tax2Rate           string    `json:"Tax2Rate,omitempty"`
	TimeStamp          time.Time `json:"TimeStamp,omitempty"`
	TaxCategoryID      string    `json:"TaxCategoryID,omitempty"`
	TaxClassID         string    `json:"TaxClassID,omitempty"`
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
	Amount    string `json:"Amount,omitempty"`
	UseTypeID string `json:"UseTypeID,omitempty"`
	UseType   string `json:"UseType,omitempty"`
}
type Prices struct {
	ItemPrice []ItemPrice `json:"ItemPrice,omitempty"`
}
type Item struct {
	ItemID           string    `json:"ItemID,omitempty"`
	SystemSku        string    `json:"SystemSku,omitempty"`
	DefaultCost      string    `json:"DefaultCost,omitempty"`
	AvgCost          string    `json:"AvgCost,omitempty"`
	Discountable     string    `json:"Discountable,omitempty"`
	Tax              string    `json:"Tax,omitempty"`
	Archived         string    `json:"Archived,omitempty"`
	ItemType         string    `json:"ItemType,omitempty"`
	Serialized       string    `json:"Serialized,omitempty"`
	Description      string    `json:"Description,omitempty"`
	ModelYear        string    `json:"ModelYear,omitempty"`
	Upc              string    `json:"Upc,omitempty"`
	Ean              string    `json:"Ean,omitempty"`
	CustomSku        string    `json:"CustomSku,omitempty"`
	ManufacturerSku  string    `json:"ManufacturerSku,omitempty"`
	CreateTime       time.Time `json:"CreateTime,omitempty"`
	TimeStamp        time.Time `json:"TimeStamp,omitempty"`
	CategoryID       string    `json:"CategoryID,omitempty"`
	TaxClassID       string    `json:"TaxClassID,omitempty"`
	DepartmentID     string    `json:"DepartmentID,omitempty"`
	ItemMatrixID     string    `json:"ItemMatrixID,omitempty"`
	ItemAttributesID string    `json:"ItemAttributesID,omitempty"`
	ManufacturerID   string    `json:"ManufacturerID,omitempty"`
	NoteID           string    `json:"NoteID,omitempty"`
	SeasonID         string    `json:"SeasonID,omitempty"`
	DefaultVendorID  string    `json:"DefaultVendorID,omitempty"`
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
	ID        string `json:"Id,omitempty"`
	Name      string `json:"Name,omitempty"`
	Taxable   string `json:"Taxable,omitempty"`
	Rate      string `json:"Rate,omitempty"`
	Amount    string `json:"Amount,omitempty"`
	Taxname   string `json:"Taxname,omitempty"`
	Subtotal  string `json:"Subtotal,omitempty"`
	Rate2     string `json:"Rate2,omitempty"`
	Amount2   string `json:"Amount2,omitempty"`
	Taxname2  string `json:"Taxname2,omitempty"`
	Subtotal2 string `json:"Subtotal2,omitempty"`
}
type TaxClassTotals struct {
	Tax Tax `json:"Tax,omitempty"`
}
type ContactPhone struct {
	Number  string `json:"Number,omitempty"`
	UseType string `json:"UseType,omitempty"`
}
type Phones struct {
	ContactPhone ContactPhone `json:"ContactPhone,omitempty"`
}
type ContactEmail struct {
	Address string `json:"Address,omitempty"`
	UseType string `json:"UseType,omitempty"`
}

type ContactAddress struct {
	Address1    string `json:"Address1,omitempty"`
	City        string `json:"City,omitempty"`
	State       string `json:"State,omitempty"`
	StateCode   string `json:"StateCode,omitempty"`
	Zip         string `json:"Zip,omitempty"`
	Country     string `json:"Country,omitempty"`
	CountryCode string `json:"CountryCode,omitempty"`
}

type Emails struct {
	ContactEmail ContactEmail `json:"ContactEmail,omitempty"`
}
type Contact struct {
	ContactID string      `json:"ContactID,omitempty"`
	Custom    string      `json:"Custom,omitempty"`
	NoEmail   string      `json:"NoEmail,omitempty"`
	NoPhone   string      `json:"NoPhone,omitempty"`
	NoMail    string      `json:"NoMail,omitempty"`
	TimeStamp time.Time   `json:"TimeStamp,omitempty"`
	Addresses interface{} `json:"Addresses,omitempty"`
	Phones    interface{} `json:"Phones,omitempty"`
	Emails    Emails      `json:"Emails,omitempty"`
	Websites  string      `json:"Websites,omitempty"`
}

type Customer struct {
	CustomerID                string    `json:"CustomerID,omitempty"`
	FirstName                 string    `json:"FirstName,omitempty"`
	LastName                  string    `json:"LastName,omitempty"`
	Dob                       string    `json:"Dob,omitempty"`
	Archived                  string    `json:"Archived,omitempty"`
	Title                     string    `json:"Title,omitempty"`
	Company                   string    `json:"Company,omitempty"`
	CompanyRegistrationNumber string    `json:"CompanyRegistrationNumber,omitempty"`
	VatNumber                 string    `json:"VatNumber,omitempty"`
	CreateTime                time.Time `json:"CreateTime,omitempty"`
	TimeStamp                 time.Time `json:"TimeStamp,omitempty"`
	ContactID                 string    `json:"ContactID,omitempty"`
	CreditAccountID           string    `json:"CreditAccountID,omitempty"`
	CustomerTypeID            string    `json:"CustomerTypeID,omitempty"`
	DiscountID                string    `json:"CiscountID,omitempty"`
	EmployeeID                string    `json:"EmployeeID,omitempty"`
	NoteID                    string    `json:"NoteID,omitempty"`
	TaxCategoryID             string    `json:"TaxCategoryID,omitempty"`
	MeasurementID             string    `json:"MeasurementID,omitempty"`
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
	RegisterID      string      `json:"registerITD,omitempty"`
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
		CustomerID                string    `json:"customerID"`
		FirstName                 string    `json:"firstName"`
		LastName                  string    `json:"lastName"`
		Archived                  string    `json:"archived"`
		Title                     string    `json:"title"`
		Company                   string    `json:"company"`
		CompanyRegistrationNumber string    `json:"companyRegistrationNumber"`
		VatNumber                 string    `json:"vatNumber"`
		CreateTime                time.Time `json:"createTime"`
		TimeStamp                 time.Time `json:"timeStamp"`
		ContactID                 string    `json:"contactID"`
		CreditAccountID           string    `json:"creditAccountID"`
		CustomerTypeID            string    `json:"customerTypeID"`
		DiscountID                string    `json:"discountID"`
		EmployeeID                string    `json:"employeeID"`
		NoteID                    string    `json:"noteID"`
		TaxCategoryID             string    `json:"taxCategoryID"`
		MeasurementID             string    `json:"measurementID"`
		Contact                   struct {
			ContactID string    `json:"contactID"`
			Custom    string    `json:"custom"`
			NoEmail   string    `json:"noEmail"`
			NoPhone   string    `json:"noPhone"`
			NoMail    string    `json:"noMail"`
			TimeStamp time.Time `json:"timeStamp"`
			Addresses string    `json:"Addresses"`
			Phones    string    `json:"Phones"`
			Emails    struct {
				ContactEmail struct {
					Address string `json:"address"`
					UseType string `json:"useType"`
				} `json:"ContactEmail"`
			} `json:"Emails"`
			Websites string `json:"Websites"`
		} `json:"Contact"`
	} `json:"Customer"`
	SalePayments struct {
		SalePayment SalePayments `json:"salePayment,omitempty"`
	} `json:"salePayments,omitempty"`
}
