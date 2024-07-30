package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"

	lightspeed_r "github.com/omniboost/go-lightspeed-r"
)

func TestSalesGet(t *testing.T) {
	req := client.NewSalesGet()
	req.QueryParams().LoadRelations = lightspeed_r.LoadRelations{
		string(lightspeed_r.LoadRelationCustomer),
		string(lightspeed_r.LoadRelationSaleLines),
		string(lightspeed_r.LoadRelationSaleLinesDiscount),
		string(lightspeed_r.LoadRelationSaleLinesTaxClass),
		string(lightspeed_r.LoadRelationSaleLinesItem),
		string(lightspeed_r.LoadRelationSalePayments),
		string(lightspeed_r.LoadRelationSalePaymentsPaymentType),
		string(lightspeed_r.LoadRelationSalePaymentsCCCHarge),
		string(lightspeed_r.LoadRelationSalePaymentsSaleAccounts),
		string(lightspeed_r.LoadRelationCustomerContact),
		string(lightspeed_r.LoadRelationDiscount),
		string(lightspeed_r.LoadRelationTaxCategory),
		string(lightspeed_r.LoadRelationTaxCategoryTaxCategoryClasses),
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
