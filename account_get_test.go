package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountGet(t *testing.T) {
	req := client.NewAccountGet()
	// req.QueryParams().IDs = []string{"97"}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// func TestAccountGetAll(t *testing.T) {
// 	req := client.NewAccountGet()
// 	resp, err := req.All()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	b, _ := json.MarshalIndent(resp, "", "  ")
// 	fmt.Println(string(b))
// }
