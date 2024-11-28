package bodies

// SearchResults contains the search results.
type SearchResults struct {
	HasMore bool           `json:"hasMore"`
	Results []SearchResult `json:"results"`
}

// SearchResult is a single entry for SearchResults.
type SearchResult struct {
	Type        string `json:"type"`
	Match       string `json:"match"`
	Description string `json:"description"`
	Details     string `json:"details"`
}
