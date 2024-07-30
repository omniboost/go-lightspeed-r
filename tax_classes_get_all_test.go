package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTaxClassesGetAll(t *testing.T) {
	req := client.NewTaxClassesGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
