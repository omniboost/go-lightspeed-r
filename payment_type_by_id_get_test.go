package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPaymentTypeByIdGet(t *testing.T) {
	req := client.NewPaymentTypeByIdGet()
	req.PathParams().PaymentTypeID = 57
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
