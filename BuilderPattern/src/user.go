package main

type Kind int

const (
	Free Kind = iota
	MonthlySubscription
	YearlySubscription
)

func (k Kind) String() string {
	switch k {
	case MonthlySubscription:
		return "monthly"
	case YearlySubscription:
		return "yearly"
	default:
		return "free"
	}
}

type User struct {
	Kind   Kind
	Name   string
	Age    int
	Gender string
}
