package logchunk

import "slices"

// LogStreamLineFilters contains the filters for streaming a log.
// If none of the fields are set, the whole log will be streamed.
type LogStreamLineFilters struct {
	// IncludedSeverities specifies the severities that should be included.
	IncludedSeverities []string

	// IncludedCategories specifies the categories that should be included.
	// If set, the ExcludedCategories is ignored.
	IncludedCategories []string

	// ExcludedCategories specifies the categories that should be excluded.
	// Will be ignored if IncludedCategories is set.
	ExcludedCategories []string
}

// HasFilters returns true if at least one of the filters is specified.
func (f LogStreamLineFilters) HasFilters() bool {
	return len(f.IncludedSeverities) > 0 ||
		len(f.IncludedCategories) > 0 ||
		len(f.ExcludedCategories) > 0
}

func (f LogStreamLineFilters) shouldStreamChunk(chunkMetadata logChunkMetadata) bool { //nolint:gocyclo // It is broken up into a few section, this is fine
	if !f.HasFilters() {
		return true
	}

	// Build slices of present severities and categories
	severities := make([]string, 0)
	for sev := range chunkMetadata.SeveritiesCount {
		severities = append(severities, sev)
	}

	categories := make([]string, 0)
	for cat := range chunkMetadata.CategoriesCount {
		categories = append(categories, cat)
	}

	// Check included categories
	if len(f.IncludedCategories) > 0 {
		// If we filter by included categories, return false if we don't have any of them
		hasCategory := false

		for _, category := range f.IncludedCategories {
			if slices.Contains(categories, category) {
				hasCategory = true

				break
			}
		}

		// If we don't have at least one matching category, return false
		if !hasCategory {
			return false
		}

		// So we have at least one matching category, check if we have matching severities too
		return f.chunkHasMatchingSeverities(severities)
	}

	// If we don't have included categories, but we have excluded categories, check if there are categories other than these, return true if so
	if len(f.ExcludedCategories) > 0 {
		// If we filter by included categories, return false if we don't have any of them
		hasOtherCategory := false

		for _, category := range categories {
			if !slices.Contains(f.ExcludedCategories, category) {
				hasOtherCategory = true

				break
			}
		}

		// If we only have ignored categories, return false
		if !hasOtherCategory {
			return false
		}

		// So we have at least one non-excluded category, check if we have matching severities too
		return f.chunkHasMatchingSeverities(severities)
	}

	// If we don't have any category based filtering, return severity matching result
	return f.chunkHasMatchingSeverities(severities)
}

func (f LogStreamLineFilters) chunkHasMatchingSeverities(chunkSeverities []string) bool {
	if len(f.IncludedSeverities) == 0 {
		return true
	}

	// See if there are chunkSeverities present
	for _, severity := range f.IncludedSeverities {
		if slices.Contains(chunkSeverities, severity) {
			return true
		}
	}

	// We have no match
	return false
}
