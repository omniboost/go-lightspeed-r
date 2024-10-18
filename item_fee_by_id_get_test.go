package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestItemFeeByIdGet(t *testing.T) {
	req := client.NewItemFeeByIdGet()
	req.PathParams().ItemFeeID = 2
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
