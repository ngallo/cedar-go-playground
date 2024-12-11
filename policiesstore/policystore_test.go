package policies

import (
	"fmt"

	"testing"

	cedar "github.com/cedar-policy/cedar-go"
)

func TestPolicyStore01(t *testing.T) {
	store, err := NewPolicyStore("01")
	if err != nil {
		fmt.Println(err)
		return
	}
	req := cedar.Request{
		Principal: cedar.NewEntityUID("User", "Alice"),
		Action:    cedar.NewEntityUID("Action", "view"),
		Resource:  cedar.NewEntityUID("Photo", "VacationPhoto94.jpg"),
		Context:   cedar.NewRecord(cedar.RecordMap{
			"demoRequest": cedar.True,
        }),
	}
	_, diagnostic := store.PoliciesSet.IsAuthorized(store.Entities, req)
	if len(diagnostic.Errors) > 0 {
		t.Errorf("Error: %v", diagnostic.Errors)
	}
}