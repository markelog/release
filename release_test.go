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

			It("should get type", func() {
				Expect(Type()).To(Equal("mac"))
			})

			It("should get name", func() {
				Expect(Name()).To(Equal("El Capitan"))
			})

			It("should get name", func() {
				Expect(Version()).To(Equal("10.11.6"))
			})

			It("should get all", func() {
				typa, name, version := All()

				Expect(typa).To(Equal("mac"))
				Expect(name).To(Equal("El Capitan"))
				Expect(version).To(Equal("10.11.6"))
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

			It("should get all", func() {
				typa, name, version := All()

				Expect(typa).To(Equal("mac"))
				Expect(name).To(Equal("Leopard"))
				Expect(version).To(Equal("10.5.0"))
			})
		})
	})

	Describe("linux", func() {
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
				GOOS = "linux"
				monkey.Patch(LSBRelease, func() string {
					return lsbUbuntu
				})
			})

			It("should get type", func() {
				Expect(Type()).To(Equal("ubuntu"))
			})

			It("should get name", func() {
				Expect(Name()).To(Equal("trusty"))
			})

			It("should get version", func() {
				Expect(Version()).To(Equal("14.04"))
			})

			It("should get all", func() {
				typa, name, version := All()

				Expect(typa).To(Equal("ubuntu"))
				Expect(name).To(Equal("trusty"))
				Expect(version).To(Equal("14.04"))
			})
		})

		Describe("unknown", func() {
			BeforeEach(func() {
				GOOS = "linux"
				monkey.Patch(LSBRelease, func() string {
					return ""
				})
			})

			It("should get type", func() {
				Expect(Type()).To(Equal(""))
			})

			It("should get name", func() {
				Expect(Name()).To(Equal(""))
			})

			It("should get version", func() {
				Expect(Version()).To(Equal(""))
			})

			It("should get all", func() {
				typa, name, version := All()

				Expect(typa).To(Equal(""))
				Expect(name).To(Equal(""))
				Expect(version).To(Equal(""))
			})
		})
	})

	Describe("windows", func() {
		BeforeEach(func() {
			GOOS = "windows"
		})

		It("should get type", func() {
			Expect(Type()).To(Equal(""))
		})

		It("should get name", func() {
			Expect(Name()).To(Equal(""))
		})

		It("should get version", func() {
			Expect(Version()).To(Equal(""))
		})

		It("should get all", func() {
			typa, name, version := All()

			Expect(typa).To(Equal(""))
			Expect(name).To(Equal(""))
			Expect(version).To(Equal(""))
		})
	})
})
