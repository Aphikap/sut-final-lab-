package entity

import (
	"testing"

	. "github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func Test_employee_code_test(t *testing.T) {
	g := NewGomegaWithT(t)

	e := Employees{
		Name:"อภิชาติ",
		Salary:20000.00,
		EmployeeCode:"HR-10245",
	}

	ok, err := ValidateStruct(e)
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by ‘-’ and 4 digits (0-9)"))
}
