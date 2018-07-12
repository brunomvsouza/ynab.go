package category

// Category represents a category for a budget
type Category struct {
	ID              string `json:"id"`
	CategoryGroupID string `json:"category_group_id"`
	Name            string `json:"name"`
	Hidden          bool   `json:"hidden"`
	Budgeted        int64  `json:"budgeted"`
	Activity        int64  `json:"activity"`
	Balance         int64  `json:"balance"`
	Deleted         bool   `json:"deleted"`

	Note                    *string `json:"note"`
	OriginalCategoryGroupID *string `json:"original_category_group_id"`
}

// Group represents a category group for a budget
type Group struct {
	ResumedGroup

	Categories []*Category `json:"categories"`
}

// ResumedGroup represents a resumed category group for a budget
type ResumedGroup struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Hidden  bool   `json:"hidden"`
	Deleted bool   `json:"deleted"`
}
