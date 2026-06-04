package youtube

// PageInfo - Paging details for lists of resources, including total
// number of items available and number of resources returned in a
// single page.
type PageInfo struct {
	// ResultsPerPage: The number of results included in the API response.
	ResultsPerPage int64 `json:"resultsPerPage"`

	// StartIndex: The index of the first item in the API response.
	StartIndex int64 `json:"startIndex,omitempty"`

	// TotalResults: The total number of results in the result set.
	TotalResults int64 `json:"totalResults"`
}
