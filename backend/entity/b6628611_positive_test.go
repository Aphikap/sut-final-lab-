package entity

import (
	"testing"

	. "github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_positive_test(t *testing.T) {
	g := NewGomegaWithT(t)

	e := Employees{
		Name:"อภิชาติ",
		Salary:20000.00,
		EmployeeCode:"HR-1024",
	}

	ok, err := ValidateStruct(e)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	
}
