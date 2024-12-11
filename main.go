package main

import (
	"fmt"
	"time"

	policiesstore "my-go-project/policiesstore"
	cedar "github.com/cedar-policy/cedar-go"
)



func main() {
	store, err := policiesstore.NewPolicyStore("01")
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
	start := time.Now()
	decision, diagnostic := store.PoliciesSet.IsAuthorized(store.Entities, req)
	elapsed := time.Since(start)
	fmt.Printf("[ExecutionTime] %s\n", elapsed)
	fmt.Printf("[Decision] %s\n", decision)
	fmt.Printf("[Diagnostic] %v\n", diagnostic)
}