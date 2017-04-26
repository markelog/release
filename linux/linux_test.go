package linux_test

import (
	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/markelog/release/linux"
)

var _ = Describe("linux", func() {
	const lsbUbuntu = `
No LSB modules are available.
Distributor ID:	Ubuntu
Description:   	Ubuntu 14.04.1 LTS
Release:       	14.04
Codename:      	trusty
`

	AfterEach(func() {
		monkey.Unpatch(LSBRelease)
	})

	Describe("ubuntu", func() {
		BeforeEach(func() {
			monkey.Patch(LSBRelease, func() string {
				return lsbUbuntu
			})
		})

		It("should get type", func() {
			Expect(New().Type()).To(Equal("ubuntu"))
		})

		It("should get name", func() {
			Expect(New().Name()).To(Equal("trusty"))
		})

		It("should get version", func() {
			Expect(New().Version()).To(Equal("14.04"))
		})
	})

	Describe("unknown", func() {
		BeforeEach(func() {
			monkey.Patch(LSBRelease, func() string {
				return ""
			})
		})

		It("should get type", func() {
			Expect(New().Type()).To(Equal(""))
		})

		It("should get name", func() {
			Expect(New().Name()).To(Equal(""))
		})

		It("should get version", func() {
			Expect(New().Version()).To(Equal(""))
		})
	})
})
