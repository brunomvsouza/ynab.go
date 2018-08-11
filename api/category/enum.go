// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package category

// Goal represents the goal of a category
type Goal string

// Pointer returns the pointer of a Goal
func (g Goal) Pointer() *Goal {
	return &g
}

const (
	// GoalTargetCategoryBalance Goal targets category balance
	GoalTargetCategoryBalance Goal = "TB"
	// GoalTargetCategoryBalanceByDate Goal targets category balance by date
	GoalTargetCategoryBalanceByDate Goal = "TBD"
	// GoalMonthlyFunding Goal by monthly funding
	GoalMonthlyFunding Goal = "MF"
)
