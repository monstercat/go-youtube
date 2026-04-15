package youtube

import (
	"net/http"
	"net/url"
)

const (
	PoliciesUrl = YoutubePartnerV1 + "/policies"
)

// Policy specifies rules that define a particular usage or match policy that a
// partner can associate with an asset or claim.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies#Policy
type Policy struct {
	// Description is the policy's description.
	Description string `json:"description,omitempty"`
	// Id is a value that YouTube assigns and uses to uniquely identify the policy.
	Id string `json:"id,omitempty"`
	// Kind identifies this as a policy. Value: "youtubePartner#policy"
	Kind string `json:"kind,omitempty"`
	// Name is the policy's name.
	Name string `json:"name,omitempty"`
	// Rules is a list of rules that specify the action that YouTube should take
	// and may optionally specify the conditions under which that action is
	// enforced.
	Rules []*PolicyRule `json:"rules,omitempty"`
	// TimeUpdated is the time the policy was updated.
	TimeUpdated string `json:"timeUpdated,omitempty"`
}

// PolicyRule represents a policy rule, which specifies a set of conditions that
// must be met and the action that YouTube should take when those conditions are
// met. For a rule to be valid, all of the rule's conditions must be satisfied.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies#PolicyRule
type PolicyRule struct {
	// Action is the policy that YouTube should enforce if the rule's conditions
	// are all valid for an asset or for an attempt to view that asset on YouTube.
	Action string `json:"action,omitempty"`
	// Conditions is a set of conditions that must be met for the rule's action
	// (and subactions) to be enforced. For a rule to be valid, all of its
	// conditions must be met.
	Conditions *Conditions `json:"conditions,omitempty"`
	// Subaction is a list of additional actions that YouTube should take if the
	// conditions in the rule are met.
	Subaction []string `json:"subaction,omitempty"`
}

// Conditions represents conditions for a policy rule. YouTube enforces a rights
// policy if any of the rules specified for the policy are valid. For a rule to
// be valid, all of the rule's conditions must be satisfied, and a condition is
// true when all set parts are satisfied. Unset conditions are disregarded (or
// always satisfied) for a rule.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies#Conditions
type Conditions struct {
	// ContentMatchType specifies whether the user- or partner-uploaded content
	// needs to match the audio, video or audiovisual content of a reference file
	// for the rule to apply.
	ContentMatchType []string `json:"contentMatchType,omitempty"`
	// DateInterval indicates that the specified date range must match the user- or
	// partner-uploaded content for the rule to apply.
	DateInterval []*IntervalCondition `json:"dateInterval,omitempty"`
	// MatchDuration specifies an amount of time that the user- or partner-uploaded
	// content needs to match a reference file for the rule to apply.
	MatchDuration []*IntervalCondition `json:"matchDuration,omitempty"`
	// MatchPercent specifies a percentage of the user- or partner-uploaded content
	// that needs to match a reference file for the rule to apply.
	MatchPercent []*IntervalCondition `json:"matchPercent,omitempty"`
	// ReferenceDuration indicates that the reference must be a certain duration
	// for the rule to apply.
	ReferenceDuration []*IntervalCondition `json:"referenceDuration,omitempty"`
	// ReferencePercent indicates that the specified percentage of a reference file
	// must match the user- or partner-uploaded content for the rule to apply.
	ReferencePercent []*IntervalCondition `json:"referencePercent,omitempty"`
	// RequiredTerritories specifies where users are (or are not) allowed to watch
	// (or listen to) an asset. YouTube determines whether the condition is
	// satisfied based on the user's location.
	RequiredTerritories *TerritoryCondition `json:"requiredTerritories,omitempty"`
}

// PolicyList is the response for policies.list.
type PolicyList struct {
	Items []*Policy `json:"items"`
	Kind  string    `json:"kind,omitempty"`
}

// ── policies.get ────────────────────────────────────────────────────────────

// GetPolicyParams are parameters for policies.get.
type GetPolicyParams struct {
	// PolicyId specifies a value that uniquely identifies the policy being
	// retrieved.
	PolicyId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetPolicyParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetPolicy retrieves the specified saved policy.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies/get
func GetPolicy(runner RequestRunner, p *GetPolicyParams) (*Policy, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    PoliciesUrl + "/" + p.PolicyId,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out Policy
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── policies.insert ─────────────────────────────────────────────────────────

// InsertPolicyParams are parameters for policies.insert.
type InsertPolicyParams struct {
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	Policy                 *Policy
}

func (p *InsertPolicyParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// InsertPolicy creates a saved policy.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies/insert
func InsertPolicy(runner RequestRunner, p *InsertPolicyParams) (*Policy, error) {
	body, err := jsonBody(p.Policy)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPost,
		Url:    PoliciesUrl,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Policy
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── policies.list ───────────────────────────────────────────────────────────

// ListPoliciesParams are parameters for policies.list.
type ListPoliciesParams struct {
	// Id specifies a comma-separated list of saved policy IDs to retrieve. Only
	// policies belonging to the currently authenticated content owner will be
	// available.
	Id string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	// Sort specifies how the search results should be sorted.
	Sort string
}

func (p *ListPoliciesParams) Values() url.Values {
	v := url.Values{}
	if p.Id != "" {
		v.Set("id", p.Id)
	}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	if p.Sort != "" {
		v.Set("sort", p.Sort)
	}
	return v
}

// ListPolicies retrieves a list of the content owner's saved policies.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies/list
func ListPolicies(runner RequestRunner, p *ListPoliciesParams) (*PolicyList, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    PoliciesUrl,
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out PolicyList
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── policies.patch ──────────────────────────────────────────────────────────

// PatchPolicyParams are parameters for policies.patch.
type PatchPolicyParams struct {
	// PolicyId specifies a value that uniquely identifies the policy being
	// updated.
	PolicyId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	Policy                 *Policy
}

func (p *PatchPolicyParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// PatchPolicy patches the specified saved policy.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies/patch
func PatchPolicy(runner RequestRunner, p *PatchPolicyParams) (*Policy, error) {
	body, err := jsonBody(p.Policy)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    PoliciesUrl + "/" + p.PolicyId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Policy
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── policies.update ─────────────────────────────────────────────────────────

// UpdatePolicyParams are parameters for policies.update.
type UpdatePolicyParams struct {
	// PolicyId specifies a value that uniquely identifies the policy being
	// updated.
	PolicyId string
	// OnBehalfOfContentOwner identifies the content owner that the user is acting
	// on behalf of. This parameter supports users whose accounts are associated
	// with multiple content owners.
	OnBehalfOfContentOwner string
	Policy                 *Policy
}

func (p *UpdatePolicyParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// UpdatePolicy updates the specified saved policy.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/policies/update
func UpdatePolicy(runner RequestRunner, p *UpdatePolicyParams) (*Policy, error) {
	body, err := jsonBody(p.Policy)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPut,
		Url:    PoliciesUrl + "/" + p.PolicyId,
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out Policy
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
