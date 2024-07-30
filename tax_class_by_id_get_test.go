package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTaxClassByIdGet(t *testing.T) {
	req := client.NewTaxClassByIdGet()
	req.PathParams().TaxClassID = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
