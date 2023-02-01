/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */
package index

import (
	"fmt"
	"testing"

	"github.com/pydio/cells/common/utils/mtree"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSQL(t *testing.T) {
	Convey("Test getting a mpath equals", t, func() {
		mpath := mtree.NewMPath(100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 10000, 123456789, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000)

		str, args := getMPathEquals([]byte(mpath.String()))
		So(str, ShouldEqual, `mpath2 LIKE ? and mpath1 LIKE ?`)
		So(len(args), ShouldEqual, 2)
		So(string(args[0].([]byte)), ShouldEqual, "89.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000")
		So(string(args[1].([]byte)), ShouldEqual, "100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.10000.1234567")
	})

	Convey("Test getting a mpath like", t, func() {
		mpath := mtree.NewMPath(100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000)

		str, args := getMPathLike([]byte(mpath.String()))

		So(str, ShouldEqual, `mpath2 LIKE ? and mpath1 LIKE ?`)

		So(len(args), ShouldEqual, 2)

		So(string(args[0].([]byte)), ShouldEqual, "100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.%")
		So(string(args[1].([]byte)), ShouldEqual, "100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.")
	})

	Convey("Test getting mpathes in ", t, func() {
		str, args := getMPathesIn("1", "1.1", "1.1.2", mtree.NewMPath(100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000, 100000000000).String())

		So(str, ShouldEqual, `(mpath1 LIKE ?) or (mpath1 LIKE ?) or (mpath1 LIKE ?) or (mpath2 LIKE ? and mpath1 LIKE ?)`)

		So(len(args), ShouldEqual, 5)

		So(string(args[0].([]byte)), ShouldEqual, "1")
		So(string(args[1].([]byte)), ShouldEqual, "1.1")
		So(string(args[2].([]byte)), ShouldEqual, "1.1.2")
		So(string(args[3].([]byte)), ShouldEqual, "100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000")
		So(string(args[4].([]byte)), ShouldEqual, "100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.100000000000.")
	})

}

func TestMpath(t *testing.T) {

	str, args := getMPathLike([]byte("1"))
	fmt.Println(str, args)
}