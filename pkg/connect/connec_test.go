package connect

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// go test ./... 跑所有的test
func TestGet(t *testing.T) {
	convey.Convey("基础用例", t, func() {
		var (
			url    = "https://www.liwenzhou.com/posts/Go/unit-test-5/"
			expect = true
		)
		got := Get(url)
		convey.So(got, convey.ShouldEqual, expect) // 断言
	})

	convey.Convey("ping 不通的用例", t, func() {
		var (
			url = "posts/Go/unit-test-5/"
			//expect = false
		)
		got := Get(url)
		//convey.So(got, convey.ShouldEqual, expect) // 断言
		convey.ShouldBeTrue(got)
	})
}
