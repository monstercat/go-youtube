package youtube

import (
	"net/http"
	"net/url"
)

// AssetMatchPolicy represents the match policy assigned to an asset. A match
// policy defines how YouTube should handle user-uploaded videos that match the
// asset. The match policy can reference a saved policy by ID or specify ad hoc
// rules directly.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetMatchPolicy#AssetMatchPolicy
type AssetMatchPolicy struct {
	// Kind identifies the API resource type. The value is "youtubePartner#assetMatchPolicy".
	Kind string `json:"kind,omitempty"`
	// PolicyId is the ID of a saved policy. If this field is set, the rules
	// field should be empty. If the rules field is non-empty, this field is ignored.
	PolicyId string `json:"policyId,omitempty"`
	// Rules is a list of rules that collectively define the policy. If the
	// policyId field is set, the rules in the saved policy take precedence.
	Rules []*PolicyRule `json:"rules,omitempty"`
}

func assetMatchPolicyUrl(assetId string) string {
	return AssetsUrl + "/" + assetId + "/matchPolicy"
}

// ── assetMatchPolicy.get ────────────────────────────────────────────────────

// GetAssetMatchPolicyParams are parameters for assetMatchPolicy.get.
type GetAssetMatchPolicyParams struct {
	// AssetId is the ID of the asset for which to retrieve the match policy.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
}

func (p *GetAssetMatchPolicyParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// GetAssetMatchPolicy retrieves the match policy assigned to the specified
// asset by the content owner associated with the authenticated user. This
// policy determines how YouTube handles user-uploaded videos that match the asset.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetMatchPolicy/get
func GetAssetMatchPolicy(runner RequestRunner, p *GetAssetMatchPolicyParams) (*AssetMatchPolicy, error) {
	res, err := runner.Run(&Request{
		Method: http.MethodGet,
		Url:    assetMatchPolicyUrl(p.AssetId),
		Params: p.Values(),
	})
	if err != nil {
		return nil, err
	}
	var out AssetMatchPolicy
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assetMatchPolicy.patch ──────────────────────────────────────────────────

// PatchAssetMatchPolicyParams are parameters for assetMatchPolicy.patch.
type PatchAssetMatchPolicyParams struct {
	// AssetId is the ID of the asset for which to patch the match policy.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// MatchPolicy is the AssetMatchPolicy resource with updated fields.
	MatchPolicy *AssetMatchPolicy
}

func (p *PatchAssetMatchPolicyParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// PatchAssetMatchPolicy patches the asset's match policy. This method supports
// patch semantics, meaning only the fields included in the request body will be
// updated; all other fields will retain their current values.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetMatchPolicy/patch
func PatchAssetMatchPolicy(runner RequestRunner, p *PatchAssetMatchPolicyParams) (*AssetMatchPolicy, error) {
	body, err := jsonBody(p.MatchPolicy)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPatch,
		Url:    assetMatchPolicyUrl(p.AssetId),
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out AssetMatchPolicy
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// ── assetMatchPolicy.update ─────────────────────────────────────────────────

// UpdateAssetMatchPolicyParams are parameters for assetMatchPolicy.update.
type UpdateAssetMatchPolicyParams struct {
	// AssetId is the ID of the asset for which to update the match policy.
	AssetId string
	// OnBehalfOfContentOwner identifies the content owner that the user is
	// acting on behalf of. This parameter supports users whose accounts are
	// associated with multiple content owners.
	OnBehalfOfContentOwner string
	// MatchPolicy is the complete AssetMatchPolicy resource to replace the
	// existing match policy.
	MatchPolicy *AssetMatchPolicy
}

func (p *UpdateAssetMatchPolicyParams) Values() url.Values {
	v := url.Values{}
	if p.OnBehalfOfContentOwner != "" {
		v.Set("onBehalfOfContentOwner", p.OnBehalfOfContentOwner)
	}
	return v
}

// UpdateAssetMatchPolicy updates the asset's match policy. This method
// replaces the entire match policy resource, so all fields must be provided.
// Use PatchAssetMatchPolicy for partial updates.
//
// see https://developers.google.com/youtube/partner/reference/rest/v1/assetMatchPolicy/update
func UpdateAssetMatchPolicy(runner RequestRunner, p *UpdateAssetMatchPolicyParams) (*AssetMatchPolicy, error) {
	body, err := jsonBody(p.MatchPolicy)
	if err != nil {
		return nil, err
	}
	res, err := runner.Run(&Request{
		Method: http.MethodPut,
		Url:    assetMatchPolicyUrl(p.AssetId),
		Params: p.Values(),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}
	var out AssetMatchPolicy
	if err := DecodeResponse(res, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
