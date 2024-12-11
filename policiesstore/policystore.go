package policies

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	cedar "github.com/cedar-policy/cedar-go"
)

// PolicyStore is a set of named policies against which a request can be authorized.
type PolicyStore struct {
	PolicyStoreID string
	PoliciesSet *cedar.PolicySet
	Entities cedar.EntityMap
}

// NewPolicyStore creates a new, empty PolicyStore
func NewPolicyStore(storeID string) (*PolicyStore, error) {
	policyStore := &PolicyStore{}
	policyStore.PolicyStoreID = storeID

	// Build the Entities
	entitiesJSONPath := fmt.Sprintf("policiesstore/cedar/%s/entities.json", storeID)
	entitiesJSON, err := os.ReadFile(entitiesJSONPath)
	if err != nil {
		return nil, err
	}
	var entities cedar.EntityMap
	if err := json.Unmarshal(entitiesJSON, &entities); err != nil {
		return nil, err
	}
	policyStore.Entities = entities

	// Build the Policy Set
	policyStore.PoliciesSet = cedar.NewPolicySet()
	cedarPath := fmt.Sprintf("policiesstore/cedar/%s/store.cedar", storeID)
	cedarCode, err := os.ReadFile(cedarPath)
	if err != nil {
		return nil, err
	}
	policyRegex := regexp.MustCompile(`permit\s*\([^;]+;`)
	codePolicies := policyRegex.FindAllString(string(cedarCode), -1)
	for i, codePolicy := range codePolicies {
		policyID := fmt.Sprintf("policy%d", i+1)
		var policy cedar.Policy
		if err := policy.UnmarshalCedar([]byte(codePolicy)); err != nil {
			return nil, err
		}
		policyStore.PoliciesSet.Add(cedar.PolicyID(policyID), &policy)
	}

	return policyStore, nil
}
