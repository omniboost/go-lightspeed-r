package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAuthRequest(t *testing.T) {
	req := client.NewAuthRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
