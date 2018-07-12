package transaction

// Type represents the type of a transaction
type Type string

const (
	// TypeUncategorized identifies an uncategorized transaction
	TypeUncategorized Type = "uncategorized"
	// TypeUnapproved identifies an unapproved transaction
	TypeUnapproved Type = "unapproved"
)

// ClearingStatus represents the clearing status of a transaction
type ClearingStatus string

const (
	// StatusUncleared identifies an uncleared transaction
	StatusUncleared ClearingStatus = "uncleared"
	// StatusCleared identifies a cleared transaction
	StatusCleared ClearingStatus = "cleared"
	// StatusReconciled identifies a reconciled transaction
	StatusReconciled ClearingStatus = "reconciled"
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
