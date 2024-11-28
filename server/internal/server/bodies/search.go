package bodies

type SearchResults struct {
	HasMore bool           `json:"hasMore"`
	Results []SearchResult `json:"results"`
}

type SearchResult struct {
	Type    string `json:"type"`
	Match   string `json:"match"`
	Details string `json:"details"`
}
