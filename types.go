package youtube

// Date represents a date in the Content ID API.
type Date struct {
	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

// Rating represents a content rating.
type Rating struct {
	RatingSystem string `json:"ratingSystem,omitempty"`
	RatingValue  string `json:"ratingValue,omitempty"`
}

// Origin describes the origin of a claim or other resource.
type Origin struct {
	Source string `json:"source,omitempty"`
}

// Source identifies the content owner and user that performed an action.
type Source struct {
	ContentOwnerId string `json:"contentOwnerId,omitempty"`
	Type           string `json:"type,omitempty"`
	UserEmail      string `json:"userEmail,omitempty"`
}

// Origination describes who provided a resource (metadata, reference, etc.).
type Origination struct {
	Owner  string `json:"owner,omitempty"`
	Source string `json:"source,omitempty"`
}

// StudioInfo contains information about a claim from the YouTube Studio.
type StudioInfo struct {
	StudioUrl string `json:"studioUrl,omitempty"`
}

// MatchInfo contains match information for a claim.
type MatchInfo struct {
	LongestMatch  *LongestMatch  `json:"longestMatch,omitempty"`
	MatchSegments []*MatchSegment `json:"matchSegments,omitempty"`
	ReferenceId   string          `json:"referenceId,omitempty"`
	TotalMatch    *TotalMatch    `json:"totalMatch,omitempty"`
}

// LongestMatch contains the longest match between reference and video.
type LongestMatch struct {
	DurationSecs     string `json:"durationSecs,omitempty"`
	ReferenceOffset  string `json:"referenceOffset,omitempty"`
	UserVideoOffset  string `json:"userVideoOffset,omitempty"`
}

// TotalMatch contains total match information.
type TotalMatch struct {
	ReferenceDurationSecs string `json:"referenceDurationSecs,omitempty"`
	UserVideoDurationSecs string `json:"userVideoDurationSecs,omitempty"`
}

// MatchSegment describes a matched segment between a reference and video.
type MatchSegment struct {
	Channel          string   `json:"channel,omitempty"`
	ManualSegment    *Segment2 `json:"manual_segment,omitempty"`
	ReferenceSegment *Segment  `json:"reference_segment,omitempty"`
	VideoSegment     *Segment  `json:"video_segment,omitempty"`
}

// Segment represents a time segment with a start and duration (uint64 strings).
type Segment struct {
	Duration string `json:"duration,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Start    string `json:"start,omitempty"`
}

// Segment2 represents a manual segment with float start/duration.
type Segment2 struct {
	Duration float64 `json:"duration,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	Start    float64 `json:"start,omitempty"`
}

// TypeDetails provides details about a claim event type.
type TypeDetails struct {
	AppealExplanation   string `json:"appealExplanation,omitempty"`
	DisputeNotes        string `json:"disputeNotes,omitempty"`
	DisputeReason       string `json:"disputeReason,omitempty"`
	UpdateStatus        string `json:"updateStatus,omitempty"`
}

// IntervalCondition represents a condition with low/high bounds.
type IntervalCondition struct {
	High float64 `json:"high,omitempty"`
	Low  float64 `json:"low,omitempty"`
}

// TerritoryCondition specifies territory-based conditions.
type TerritoryCondition struct {
	Territories []string `json:"territories,omitempty"`
	Type        string   `json:"type,omitempty"`
}

// ExcludedInterval represents an interval excluded from a reference.
type ExcludedInterval struct {
	High float64 `json:"high,omitempty"`
	Low  float64 `json:"low,omitempty"`
}

// StatusReport provides a status report for a package.
type StatusReport struct {
	StatusContent string `json:"statusContent,omitempty"`
	StatusId      string `json:"statusId,omitempty"`
}

// TerritoryOwners describes ownership in specific territories.
type TerritoryOwners struct {
	Owner       string   `json:"owner,omitempty"`
	Publisher   string   `json:"publisher,omitempty"`
	Ratio       float64  `json:"ratio,omitempty"`
	Territories []string `json:"territories,omitempty"`
	Type        string   `json:"type,omitempty"`
}

// TerritoryConflicts describes ownership conflicts in territories.
type TerritoryConflicts struct {
	ConflictingOwnership []*ConflictingOwnership `json:"conflictingOwnership,omitempty"`
	Territory            string                  `json:"territory,omitempty"`
}

// ConflictingOwnership describes conflicting ownership for an asset.
type ConflictingOwnership struct {
	Owner string `json:"owner,omitempty"`
	Ratio float64 `json:"ratio,omitempty"`
}

// OwnershipConflicts contains ownership conflict information for an asset.
type OwnershipConflicts struct {
	General         []*TerritoryConflicts `json:"general,omitempty"`
	Kind            string                `json:"kind,omitempty"`
	Mechanical      []*TerritoryConflicts `json:"mechanical,omitempty"`
	Performance     []*TerritoryConflicts `json:"performance,omitempty"`
	Synchronization []*TerritoryConflicts `json:"synchronization,omitempty"`
}

// NWayRevenueSharing describes N-way revenue sharing for an asset.
type NWayRevenueSharing struct {
	TerritoriesIneligible []*TerritoriesIneligibleForNWayRevenueSharing `json:"territoriesIneligible,omitempty"`
}

// TerritoriesIneligibleForNWayRevenueSharing describes territories
// ineligible for N-way revenue sharing.
type TerritoriesIneligibleForNWayRevenueSharing struct {
	Reason      string   `json:"reason,omitempty"`
	Territories []string `json:"territories,omitempty"`
}

// AssetLicensability describes the licensability of an asset.
type AssetLicensability struct {
	Kind string `json:"kind,omitempty"`
}

// PromotedContent represents promoted content in a campaign.
type PromotedContent struct {
	Link []*CampaignTargetLink `json:"link,omitempty"`
	Type string                `json:"type,omitempty"`
}

// CampaignTargetLink links a campaign to a target.
type CampaignTargetLink struct {
	TargetId   string `json:"targetId,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

// AdBreak represents an ad break in a video.
type AdBreak struct {
	MidrollSeconds float64 `json:"midrollSeconds,omitempty"`
	Position       string  `json:"position,omitempty"`
}

// CountriesRestriction describes ad restrictions by country.
type CountriesRestriction struct {
	AdFormats  []string `json:"adFormats,omitempty"`
	Territories []string `json:"territories,omitempty"`
}

// ValidateError represents a validation error from the validator or package endpoints.
type ValidateError struct {
	ColumnName  string `json:"columnName,omitempty"`
	ColumnNumber int   `json:"columnNumber,omitempty"`
	LineNumber  int    `json:"lineNumber,omitempty"`
	Message     string `json:"message,omitempty"`
	MessageCode int    `json:"messageCode,omitempty"`
	Severity    string `json:"severity,omitempty"`
}

// Empty is the response for delete operations.
type Empty struct{}
