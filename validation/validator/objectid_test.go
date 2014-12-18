package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestObjectIDValidator(t *testing.T) {

	Convey("When creating a ObjectID validator", t, func() {
		v := ObjectID{}
		Convey("give a ObjectID paramater", func() {
			So(v.Validate("507f191e810c19729de860ea"), ShouldBeTrue)
		})
		Convey("give a NOT ObjectID paramater (short)", func() {
			So(v.Validate("507f191e810c1972"), ShouldBeFalse)
		})
		Convey("give a NOT ObjectID paramater (long)", func() {
			So(v.Validate("507f191e810c19729de860ev"), ShouldBeFalse)
		})
		Convey("give a 'empty' paramater", func() {
			So(v.Validate(""), ShouldBeTrue)
		})
	})

}
