package osx_test

import (
	"io/ioutil"

	"bou.ke/monkey"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/markelog/release/osx"
)

var _ = Describe("osx", func() {
	Describe("El Captain", func() {
		license := `{\rtf1\ansi\ansicpg1252\cocoartf1504
\cocoascreenfonts1{\fonttbl\f0\fnil\fcharset0 HelveticaNeue;}
{\colortbl;\red255\green255\blue255;}
{\*\expandedcolortbl;\csgray\c100000;}
\deftab720
\pard\pardeftab720\partightenfactor0

\f0\b\fs20 \cf0 ENGLISH\
\
APPLE INC.\
SOFTWARE LICENSE AGREEMENT FOR OS X EL CAPITAN\
For use on Apple-branded Systems\
\
`

		sysInfo := `
Software:

    System Software Overview:

      System Version: macOS 10.11.6 (15G1421)
      Kernel Version: Darwin 15.6.0
      Boot Volume: Macintosh HD
      Boot Mode: Normal
      Computer Name: dusty
      User Name: Oleg Gaidarenko (markelog)
      Secure Virtual Memory: Enabled
      System Integrity Protection: Enabled
      Time since boot: 4:11
`

		BeforeEach(func() {
			monkey.Patch(Execute, func(name, arg string) string {
				return sysInfo
			})

			monkey.Patch(ioutil.ReadFile, func(filename string) ([]byte, error) {
				return []byte(license), nil
			})
		})

		AfterEach(func() {
			monkey.Unpatch(Execute)
			monkey.Unpatch(ioutil.ReadFile)
		})

		It("should get type", func() {
			Expect(New().Type()).To(Equal("osx"))
		})

		It("should not freak out if there is no file for the name", func() {
			monkey.Unpatch(ioutil.ReadFile)

			monkey.Patch(ioutil.ReadFile, func(filename string) ([]byte, error) {
				return []byte("test"), nil
			})

			Expect(New().Name()).To(Equal(""))
		})

		It("should get name", func() {
			Expect(New().Name()).To(Equal("OS X EL CAPITAN"))
		})

		It("should get name", func() {
			Expect(New().Version()).To(Equal("10.11.6"))
		})

		It("should get name", func() {
			Expect(New().Version()).To(Equal("10.11.6"))
		})

		It("should not freak out if there is no correct output for version", func() {
			monkey.Unpatch(Execute)

			monkey.Patch(Execute, func(name, arg string) string {
				return "test"
			})

			Expect(New().Version()).To(Equal(""))
		})
	})

	Describe("Sierra", func() {
		sysInfo := `
Software:

    System Software Overview:

      System Version: macOS 10.12.4 (16E195)
      Kernel Version: Darwin 16.5.0
      Boot Volume: Macintosh HD
      Boot Mode: Normal
      Computer Name: dusty
      User Name: Oleg Gaidarenko (markelog)
      Secure Virtual Memory: Enabled
      System Integrity Protection: Enabled
      Time since boot: 4:11
`
		license := `{\rtf1\ansi\ansicpg1252\cocoartf1504
\cocoascreenfonts1{\fonttbl\f0\fnil\fcharset0 HelveticaNeue;}
{\colortbl;\red255\green255\blue255;}
{\*\expandedcolortbl;\csgray\c100000;}
\deftab720
\pard\pardeftab720\partightenfactor0

\f0\b\fs20 \cf0 ENGLISH\
\
APPLE INC.\
SOFTWARE LICENSE AGREEMENT FOR macOS Sierra\
For use on Apple-branded Systems\
\
`

		BeforeEach(func() {
			monkey.Patch(ioutil.ReadFile, func(filename string) ([]byte, error) {
				return []byte(license), nil
			})

			monkey.Patch(Execute, func(name, arg string) string {
				return sysInfo
			})
		})

		AfterEach(func() {
			monkey.Unpatch(ioutil.ReadFile)
			monkey.Unpatch(Execute)
		})

		It("should get type", func() {
			Expect(New().Type()).To(Equal("osx"))
		})

		It("should get name", func() {
			Expect(New().Name()).To(Equal("macOS Sierra"))
		})

		It("should get name", func() {
			Expect(New().Version()).To(Equal("10.12.4"))
		})
	})

	Describe("High Sierra", func() {
		sysInfo := `Software:

    System Software Overview:

      System Version: macOS 10.13 (17A405)
      Kernel Version: Darwin 17.0.0
      Boot Volume: Macintosh HD
      Boot Mode: Normal
      Computer Name: demi
      User Name: Oleg Gaidarenko (markelog)
      Secure Virtual Memory: Enabled
      System Integrity Protection: Enabled
      Time since boot: 55 minutes`

		license := `{\rtf1\ansi\ansicpg1252\cocoartf1504
\cocoascreenfonts1{\fonttbl\f0\fnil\fcharset0 HelveticaNeue;}
{\colortbl;\red255\green255\blue255;}
{\*\expandedcolortbl;\csgray\c100000;}
\deftab720
\pard\pardeftab720\partightenfactor0

\f0\b\fs20 \cf0 ENGLISH\
\
APPLE INC.\
SOFTWARE LICENSE AGREEMENT FOR macOS High Sierra\
For use on Apple-branded Systems\`

		BeforeEach(func() {
			monkey.Patch(ioutil.ReadFile, func(filename string) ([]byte, error) {
				return []byte(license), nil
			})

			monkey.Patch(Execute, func(name, arg string) string {
				return sysInfo
			})
		})

		AfterEach(func() {
			monkey.Unpatch(ioutil.ReadFile)
			monkey.Unpatch(Execute)
		})

		It("should get type", func() {
			Expect(New().Type()).To(Equal("osx"))
		})

		It("should get name", func() {
			Expect(New().Name()).To(Equal("macOS High Sierra"))
		})

		It("should get name", func() {
			Expect(New().Version()).To(Equal("10.13"))
		})
	})
})
