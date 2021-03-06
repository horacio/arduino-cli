/*
 * This file is part of arduino-cli.
 *
 * Copyright 2018 ARDUINO SA (http://www.arduino.cc/)
 *
 * This software is released under the GNU General Public License version 3,
 * which covers the main part of arduino-cli.
 * The terms of this license can be found at:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 * You can be released from the requirements of the above licenses by purchasing
 * a commercial license. Buying such a license is mandatory if you want to modify or
 * otherwise use the software for commercial activities involving the Arduino
 * software without disclosing the source code of your own applications. To purchase
 * a commercial license, send an email to license@arduino.cc.
 */

package cores

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlavorCompatibility(t *testing.T) {
	type os struct {
		Os   string
		Arch string
	}
	windowsi386 := &os{"windows", "i386"}
	windowsx8664 := &os{"windows", "amd64"}
	linuxi386 := &os{"linux", "i386"}
	linuxamd64 := &os{"linux", "amd64"}
	linuxarm := &os{"linux", "arm"}
	linuxarmbe := &os{"linux", "armbe"}
	darwini386 := &os{"darwin", "i386"}
	darwinamd646 := &os{"darwin", "amd64"}
	freebsdi386 := &os{"freebsd", "i386"}
	freebsdamd64 := &os{"freebsd", "amd64"}
	oses := []*os{
		windowsi386,
		windowsx8664,
		linuxi386,
		linuxamd64,
		linuxarm,
		linuxarmbe,
		darwini386,
		darwinamd646,
		freebsdi386,
		freebsdamd64,
	}

	type test struct {
		Flavour   *Flavour
		Positives []*os
	}
	tests := []*test{
		&test{&Flavour{OS: "i686-mingw32"}, []*os{windowsi386, windowsx8664}},
		&test{&Flavour{OS: "i386-apple-darwin11"}, []*os{darwini386, darwinamd646}},
		&test{&Flavour{OS: "x86_64-apple-darwin"}, []*os{darwinamd646}},

		// Raspberry PI, BBB or other ARM based host
		// PI: "arm-linux-gnueabihf"
		// Raspbian on PI2: "arm-linux-gnueabihf"
		// Ubuntu Mate on PI2: "arm-linux-gnueabihf"
		// Debian 7.9 on BBB: "arm-linux-gnueabihf"
		// Raspbian on PI Zero: "arm-linux-gnueabihf"
		&test{&Flavour{OS: "arm-linux-gnueabihf"}, []*os{linuxarm, linuxarmbe}},
		// Arch-linux on PI2: "armv7l-unknown-linux-gnueabihf"
		&test{&Flavour{OS: "armv7l-unknown-linux-gnueabihf"}, []*os{linuxarm, linuxarmbe}},

		&test{&Flavour{OS: "i686-linux-gnu"}, []*os{linuxi386}},
		&test{&Flavour{OS: "i686-pc-linux-gnu"}, []*os{linuxi386}},
		&test{&Flavour{OS: "x86_64-linux-gnu"}, []*os{linuxamd64}},
		&test{&Flavour{OS: "x86_64-pc-linux-gnu"}, []*os{linuxamd64}},
	}

	check := func(test *test, os *os) {
		for _, positiveOs := range test.Positives {
			if positiveOs == os {
				require.True(t, test.Flavour.isCompatibleWith(os.Os, os.Arch), "'%s' tag compatible with '%s,%s' pair", test.Flavour.OS, os.Os, os.Arch)
				return
			}
		}
		require.False(t, test.Flavour.isCompatibleWith(os.Os, os.Arch), "'%s' tag compatible with '%s,%s' pair", test.Flavour.OS, os.Os, os.Arch)
	}

	for _, test := range tests {
		for _, os := range oses {
			check(test, os)
		}
	}
}
