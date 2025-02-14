package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance >= 0.0 && balance < 1000.0 {
        return float32(0.5)
    } else if balance >= 1000.0 && balance < 5000.0 {
        return float32(1.621)
    } else if balance >= 5000.0 {
        return float32(2.475)
    }
    return float32(3.213)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)) / 100.0
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    currentBalance := balance
    years := 0  
    for currentBalance < targetBalance {
        currentBalance = AnnualBalanceUpdate(currentBalance)
        years++
    }
    return years
}

