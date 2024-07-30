package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTaxCategoryByIdGet(t *testing.T) {
	req := client.NewTaxCategoryByIdGet()
	req.PathParams().TaxCategoryID = 15
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
