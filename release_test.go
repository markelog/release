package release_test

import (
	"runtime"

	"github.com/bouk/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/markelog/release"
)

var _ = Describe("release", func() {
	AfterEach(func() {
		GOOS = runtime.GOOS
	})

	Describe("mac", func() {
		BeforeEach(func() {
			GOOS = "darwin"
		})

		Describe("newer version", func() {
			BeforeEach(func() {
				monkey.Patch(Uname, func() string {
					return "Darwin test 15.6.0 Darwin Kernel Version 15.6.0: Thu Jun 23 18:25:34 PDT 2016; root:xnu-3248.60.10~1/RELEASE_X86_64 x86_64"
				})
			})

			AfterEach(func() {
				monkey.Unpatch(Uname)
			})

			It("should get name", func() {
				Expect(Name()).To(Equal("El Capitan"))
			})

			It("should get name", func() {
				Expect(Version()).To(Equal("10.11.6"))
			})

			It("should get type", func() {
				Expect(Type()).To(Equal("mac"))
			})
		})

		Describe("older version", func() {
			BeforeEach(func() {
				monkey.Patch(Uname, func() string {
					return "Darwin test 9.0 Darwin Kernel Version 9.0: Thu Jun 23 18:25:34 PDT 2016; root:xnu-3248.60.10~1/RELEASE_X86_64 x86_64"
				})
			})

			AfterEach(func() {
				monkey.Unpatch(Uname)
			})

			It("should get name", func() {
				Expect(Name()).To(Equal("Leopard"))
			})

			It("should get version", func() {
				Expect(Version()).To(Equal("10.5.0"))
			})

			It("should get type", func() {
				Expect(Type()).To(Equal("mac"))
			})
		})
	})

	Describe("unknown", func() {
		BeforeEach(func() {
			GOOS = "linux"
			monkey.Patch(LSBRelease, func() string {
				return ``
			})
		})

		AfterEach(func() {
			monkey.Unpatch(LSBRelease)
		})

		It("should get name", func() {
			Expect(Name()).To(Equal(""))
		})

		It("should get version", func() {
			Expect(Version()).To(Equal(""))
		})

		It("should get type", func() {
			Expect(Type()).To(Equal(""))
		})
	})
})
