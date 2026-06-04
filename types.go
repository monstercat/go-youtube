package youtube

// Date represents a date in the Content ID API.
type Date struct {
	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

// Rating represents a content rating.
type Rating struct {
	// Rating is the rating that the asset received.
	Rating string `json:"rating,omitempty"`
	// RatingSystem is the rating system associated with the rating.
	RatingSystem string `json:"ratingSystem,omitempty"`
	// RatingValue is retained for backwards compatibility. The spec
	// field is `rating`; populate the Rating field above for new code.
	RatingValue string `json:"ratingValue,omitempty"`
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

// StudioInfo contains URLs linking back to claim-related pages in Studio.
type StudioInfo struct {
	// ClaimUrl links to the claim page in Studio. Note: this page loads
	// differently depending on whether the claim has "action required"
	// issue or not.
	ClaimUrl string `json:"claimUrl,omitempty"`
	// IssueUrl, when the claim has an "action required" issue (guaranteed
	// to be at most 1), links to the issue page in Studio.
	IssueUrl string `json:"issueUrl,omitempty"`
	// StudioUrl is retained for backwards compatibility. The spec does
	// not define this field; use ClaimUrl / IssueUrl / VideoUrl instead.
	StudioUrl string `json:"studioUrl,omitempty"`
	// VideoUrl links to the claimed video page in Studio.
	VideoUrl string `json:"videoUrl,omitempty"`
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

// Segment2 represents the manual_segment payload on a MatchSegment.
//
// Start and Finish are deliberately kept as raw strings rather than a
// parsed time/duration type: the YouTube Partner v1 discovery document
// contradicts itself about their format, so any decoded shape we
// committed to would be wrong half the time. The schema-level
// description says
//
//	"...start and finish time formatted as a \"hh:mm:ss.mmm\" string."
//
// while the per-field descriptions on the very same schema say
//
//	"...measured in milliseconds from the beginning."
//
// Google has shipped responses in both formats in the wild, and they
// can swap which one they send at any time without changing the
// schema. Sampling the production response to "decide" the type isn't
// safe — the contract is just "string." Callers that need a typed
// value must sniff (a `:` means hh:mm:ss.mmm, otherwise treat as a
// uint64 millisecond count) and tolerate either form.
type Segment2 struct {
	// Finish is the finish time of the segment. See the type-level
	// doc above for why this is a string with no committed format.
	Finish string `json:"finish,omitempty"`
	// Kind is the type of the API resource. For segment resources, the
	// value is `youtubePartner#segment`.
	Kind string `json:"kind,omitempty"`
	// Start is the start time of the segment. See the type-level doc
	// above for why this is a string with no committed format.
	Start string `json:"start,omitempty"`
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

// ExcludedInterval defines a time window within the reference that will be
// ignored during the match process.
type ExcludedInterval struct {
	// High is the end (inclusive) time in seconds of the time window. The
	// value can be any value greater than `low`. If `high` is greater
	// than the length of the reference, the interval between `low` and
	// the end of the reference will be excluded. Every interval must
	// specify a value for this field.
	High float64 `json:"high,omitempty"`
	// Low is the start (inclusive) time in seconds of the time window.
	// The value can be any value between `0` and `high`. Every interval
	// must specify a value for this field.
	Low float64 `json:"low,omitempty"`
	// Origin is the source of the request to exclude the interval from
	// Content ID matching.
	Origin string `json:"origin,omitempty"`
	// TimeCreated is the date and time that the exclusion was created.
	// The value is specified in RFC 3339 (`YYYY-MM-DDThh:mm:ss.000Z`) format.
	TimeCreated string `json:"timeCreated,omitempty"`
}

// StatusReport provides a status report for a package.
type StatusReport struct {
	// StatusContent is the content of the status report.
	StatusContent string `json:"statusContent,omitempty"`
	// StatusFileName is the file name of the status report.
	StatusFileName string `json:"statusFileName,omitempty"`
	// StatusId is retained for backwards compatibility. The spec does
	// not define this field; use StatusFileName instead.
	StatusId string `json:"statusId,omitempty"`
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

// NWayRevenueSharing carries information about an asset's n-way revshare.
type NWayRevenueSharing struct {
	// EligibleTerritories is the list of territories in which the asset is
	// eligible for n-way revenue sharing. Each country is represented by
	// its two-letter ISO country code (ISO 3166-1 alpha-2).
	EligibleTerritories []string `json:"eligibleTerritories,omitempty"`
	// IneligibleTerritories carries information about territories in which
	// the asset is ineligible for n-way revenue sharing.
	IneligibleTerritories []*TerritoriesIneligibleForNWayRevenueSharing `json:"ineligibleTerritories,omitempty"`
	// Status is the status of n-way revenue sharing.
	Status string `json:"status,omitempty"`
	// TerritoriesIneligible is retained for backwards compatibility. The
	// spec field is `ineligibleTerritories`; populate IneligibleTerritories
	// above for new code.
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

// AdBreak contains information about a time when YouTube can show an
// in-stream advertisement during video playback.
type AdBreak struct {
	// MidrollSeconds is the time of the ad break specified as the number
	// of seconds after the start of the video when the break occurs.
	MidrollSeconds int32 `json:"midrollSeconds,omitempty"`
	// Position is the point at which the break occurs during the video
	// playback.
	Position string `json:"position,omitempty"`
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
