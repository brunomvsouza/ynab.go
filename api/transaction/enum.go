package transaction

// Status represents the status of a transaction
type Status string

const (
	// StatusUncleared identifies an uncleared transaction
	StatusUncleared Status = "uncleared"
	// StatusCleared identifies a cleared transaction
	StatusCleared Status = "cleared"
	// StatusReconciled identifies a reconciled transaction
	StatusReconciled Status = "reconciled"
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

// Frequency represents the frequency of a scheduled transaction
// or sub-transaction
type Frequency string

const (
	// FrequencyNever identifies a transaction that will never repeat
	FrequencyNever Frequency = "never"
	// FrequencyDaily identifies a transaction that will repeat daily
	FrequencyDaily Frequency = "daily"
	// FrequencyWeekly identifies a transaction that will repeat weekly
	FrequencyWeekly Frequency = "weekly"
	// FrequencyEveryOtherWeek identifies a transaction that will repeat
	// every other week
	FrequencyEveryOtherWeek Frequency = "everyOtherWeek"
	// FrequencyTwiceAMonth identifies a transaction that will repeat
	// twice a month
	FrequencyTwiceAMonth Frequency = "twiceAMonth"
	// FrequencyEveryFourWeeks identifies a transaction that will repeat
	// every four weeks
	FrequencyEveryFourWeeks Frequency = "every4Weeks"
	// FrequencyMonthly identifies a transaction that will repeat monthly
	FrequencyMonthly Frequency = "monthly"
	// FrequencyEveryOtherMonth identifies a transaction that will repeat
	// every other month
	FrequencyEveryOtherMonth Frequency = "everyOtherMonth"
	// FrequencyEveryThreeMonths identifies a transaction that will repeat
	// every three months
	FrequencyEveryThreeMonths Frequency = "every3Months"
	// FrequencyEveryFourMonths identifies a transaction that will repeat
	// every four months
	FrequencyEveryFourMonths Frequency = "every4Months"
	// FrequencyTwiceAYear identifies a transaction that will repeat
	// twice a year
	FrequencyTwiceAYear Frequency = "twiceAYear"
	// FrequencyYearly identifies a transaction that will repeat yearly
	FrequencyYearly Frequency = "yearly"
	// FrequencyEveryOtherYear identifies a transaction that will repeat
	// every other year
	FrequencyEveryOtherYear Frequency = "everyOtherYear"
)
