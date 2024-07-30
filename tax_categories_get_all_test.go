package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTaxCategoriesGetAll(t *testing.T) {
	req := client.NewTaxCategoriesGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
