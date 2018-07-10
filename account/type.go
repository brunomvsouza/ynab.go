package account

// Type identifies an account type
type Type string

const (
	// TypeChecking identifies a checking account
	TypeChecking Type = "checking"
	// TypeSavings identifies a savings account
	TypeSavings Type = "savings"
	// TypeCash identifies a cash account
	TypeCash Type = "cash"
	// TypeCreditCard identifies a credit card account
	TypeCreditCard Type = "creditCard"
	// TypeLineOfCredit identifies a line of credit account
	TypeLineOfCredit Type = "lineOfCredit"
	// TypeOtherAsset identifies an other asset account
	TypeOtherAsset Type = "otherAsset"
	// TypeOtherLiability identifies an other liability account
	TypeOtherLiability Type = "otherLiability"
	// TypePayPal identifies a PayPal account
	TypePayPal Type = "payPal"
	// TypeMerchant identifies a merchant account
	TypeMerchant Type = "merchantAccount"
	// TypeInvestment identifies an investment account
	TypeInvestment Type = "investmentAccount"
	// TypeMortgage identifies a mortgage account
	TypeMortgage Type = "mortgage"
)
