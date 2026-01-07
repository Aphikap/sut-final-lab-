package entity

import (
	"testing"

	. "github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test__negative_salary_test(t *testing.T) {
	g := NewGomegaWithT(t)

	e := Employees{
		Name:         "อภิชาติ",
		Salary:       20000.00,
		EmployeeCode: "HR-1024",
	}

	ok, err := ValidateStruct(e)
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
}
