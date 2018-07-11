package category

// Category represents a category
type Category struct {
	ID                      string `json:"id"`
	CategoryGroupID         string `json:"category_group_id"`
	OriginalCategoryGroupID string `json:"original_category_group_id"`
	Name                    string `json:"name"`
	Hidden                  bool   `json:"hidden"`
	Budgeted                int64  `json:"budgeted"`
	Activity                int64  `json:"activity"`
	Balance                 int64  `json:"balance"`
	Deleted                 bool   `json:"deleted"`

	Note *string `json:"note"`
}

// Group represents a category group
type Group struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Hidden  bool   `json:"hidden"`
	Deleted bool   `json:"deleted"`
}
