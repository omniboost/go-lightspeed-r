package lightspeed_r_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCategoryByIdGet(t *testing.T) {
	req := client.NewCategoryByIdGet()
	req.PathParams().AccountID = 183202
	req.PathParams().CategoryID = 1
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// func TestCategoryByIdGetAll(t *testing.T) {
// 	req := client.NewCategoryByIdGet()
// 	resp, err := req.All()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	b, _ := json.MarshalIndent(resp, "", "  ")
// 	fmt.Println(string(b))
// }
