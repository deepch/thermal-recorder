// thermal-recorder - record thermal video footage of warm moving objects
//  Copyright (C) 2018, The Cacophony Project
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package throttle

type ThrottlerConfig struct {
	ApplyThrottling bool    `yaml:"apply-throttling"`
	ThrottleAfter   uint16  `yaml:"throttle-after-secs"`
	SparseAfter     uint16  `yaml:"sparse-after-secs"`
	SparseLength    uint16  `yaml:"sparse-length-secs"`
	RefillRate      float64 `yaml:"refill-rate"`
}

func DefaultThrottlerConfig() ThrottlerConfig {
	return ThrottlerConfig{
		ApplyThrottling: true,
		SparseAfter:     3600,
		SparseLength:    30,
		ThrottleAfter:   600,
		RefillRate:      1.0,
	}
}
