package main

import (
	"testing"

	cv "github.com/smartystreets/goconvey/convey"
)

func TestWebServerReturnsOnlyWhenItIsReachable(t *testing.T) {
	addr := "localhost:3000"
	cv.Convey("Given a call to StartWebServer", t, func() {
		cv.Convey("the webserver should be already listening on the chosen port when the function returns", func() {
			webserv := NewWebServer(addr)
			webserv.Start()
			defer webserv.Stop()
			cv.So(PortIsBound(addr), cv.ShouldEqual, true)
		})
	})
}

func TestWebServerShutsdownWhenRequested(t *testing.T) {
	cv.Convey("Given a call to StartWebServer", t, func() {
		cv.Convey("the webserver be up after returning, and should terminate when requested", func() {
			addr := "localhost:3000"
			webserv := NewWebServer(addr)
			webserv.Start()
			cv.So(PortIsBound(addr), cv.ShouldEqual, true)
			webserv.Stop()
			cv.So(PortIsBound(addr), cv.ShouldEqual, false)

			// and again right away
			webserv = NewWebServer(addr)
			webserv.Start()
			cv.So(PortIsBound(addr), cv.ShouldEqual, true)
			webserv.Stop()
			cv.So(PortIsBound(addr), cv.ShouldEqual, false)
		})
	})
}
