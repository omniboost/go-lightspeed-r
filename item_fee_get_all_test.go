package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestItemFeeGetAll(t *testing.T) {
	req := client.NewItemFeeGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
