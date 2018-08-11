// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

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
	// TypePayPal DEPRECATED identifies a PayPal account
	TypePayPal Type = "payPal"
	// TypeMerchant DEPRECATED identifies a merchant account
	TypeMerchant Type = "merchantAccount"
	// TypeInvestment DEPRECATED identifies an investment account
	TypeInvestment Type = "investmentAccount"
	// TypeMortgage DEPRECATED identifies a mortgage account
	TypeMortgage Type = "mortgage"
)
