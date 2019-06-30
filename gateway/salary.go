package gateway

import "github.com/sirupsen/logrus"

// SalaryGateway is an interface to salary related resources
type SalaryGateway interface {
	FetchSalary(string) int
}

// SalaryInMemoryGateway is an in memory gateway for salary resources
type SalaryInMemoryGateway struct {
	salaries map[string]int
}

// NewSalaryInMemoryGateway creates an in memory salary gateway
func NewSalaryInMemoryGateway() *SalaryInMemoryGateway {
	salaries := map[string]int{
		"alice":   100,
		"bob":     200,
		"charlie": 100,
		"betty":   200,
		"david":   100,
	}
	return &SalaryInMemoryGateway{salaries: salaries}
}

// FetchSalary fetches the salary for the user matching the id
func (gw *SalaryInMemoryGateway) FetchSalary(id string) int {
	salary := gw.salaries[id]
	logrus.WithFields(logrus.Fields{
		"id":     id,
		"salary": salary,
	}).Info("Fetched salary")
	return salary
}
