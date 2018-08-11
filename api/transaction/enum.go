// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package transaction

// Status represents the type of a transaction
type Status string

// Pointer returns the pointer of a Status
func (s Status) Pointer() *Status {
	return &s
}

const (
	// StatusUncategorized identifies an uncategorized transaction
	StatusUncategorized Status = "uncategorized"
	// StatusUnapproved identifies an unapproved transaction
	StatusUnapproved Status = "unapproved"
)

// ClearingStatus represents the clearing status of a transaction
type ClearingStatus string

const (
	// ClearingStatusUncleared identifies an uncleared transaction
	ClearingStatusUncleared ClearingStatus = "uncleared"
	// ClearingStatusCleared identifies a cleared transaction
	ClearingStatusCleared ClearingStatus = "cleared"
	// ClearingStatusReconciled identifies a reconciled transaction
	ClearingStatusReconciled ClearingStatus = "reconciled"
)

// FlagColor represents the flag color of a transaction
type FlagColor string

const (
	// FlagColorRed identifies a transaction flagged red
	FlagColorRed FlagColor = "red"
	// FlagColorOrange identifies a transaction flagged orange
	FlagColorOrange FlagColor = "orange"
	// FlagColorYellow identifies a transaction flagged yellow
	FlagColorYellow FlagColor = "yellow"
	// FlagColorGreen identifies a transaction flagged green
	FlagColorGreen FlagColor = "green"
	// FlagColorBlue identifies a transaction flagged blue
	FlagColorBlue FlagColor = "blue"
	// FlagColorPurple identifies a transaction flagged purple
	FlagColorPurple FlagColor = "purple"
)

// ScheduledFrequency represents the frequency of a scheduled transaction
// or sub-transaction
type ScheduledFrequency string

const (
	// FrequencyNever identifies a transaction that will never repeat
	FrequencyNever ScheduledFrequency = "never"
	// FrequencyDaily identifies a transaction that will repeat daily
	FrequencyDaily ScheduledFrequency = "daily"
	// FrequencyWeekly identifies a transaction that will repeat weekly
	FrequencyWeekly ScheduledFrequency = "weekly"
	// FrequencyEveryOtherWeek identifies a transaction that will repeat
	// every other week
	FrequencyEveryOtherWeek ScheduledFrequency = "everyOtherWeek"
	// FrequencyTwiceAMonth identifies a transaction that will repeat
	// twice a month
	FrequencyTwiceAMonth ScheduledFrequency = "twiceAMonth"
	// FrequencyEveryFourWeeks identifies a transaction that will repeat
	// every four weeks
	FrequencyEveryFourWeeks ScheduledFrequency = "every4Weeks"
	// FrequencyMonthly identifies a transaction that will repeat monthly
	FrequencyMonthly ScheduledFrequency = "monthly"
	// FrequencyEveryOtherMonth identifies a transaction that will repeat
	// every other month
	FrequencyEveryOtherMonth ScheduledFrequency = "everyOtherMonth"
	// FrequencyEveryThreeMonths identifies a transaction that will repeat
	// every three months
	FrequencyEveryThreeMonths ScheduledFrequency = "every3Months"
	// FrequencyEveryFourMonths identifies a transaction that will repeat
	// every four months
	FrequencyEveryFourMonths ScheduledFrequency = "every4Months"
	// FrequencyTwiceAYear identifies a transaction that will repeat
	// twice a year
	FrequencyTwiceAYear ScheduledFrequency = "twiceAYear"
	// FrequencyYearly identifies a transaction that will repeat yearly
	FrequencyYearly ScheduledFrequency = "yearly"
	// FrequencyEveryOtherYear identifies a transaction that will repeat
	// every other year
	FrequencyEveryOtherYear ScheduledFrequency = "everyOtherYear"
)

// Type represents the type of a hybrid transaction
type Type string

const (
	// TypeTransaction identifies a hybrid transaction as transaction
	TypeTransaction Type = "transaction"
	// TypeSubTransaction identifies a hybrid transaction as sub-transaction
	TypeSubTransaction Type = "subtransaction"
)
