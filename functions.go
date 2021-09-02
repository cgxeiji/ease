package ease

import "math"

var (
	// Linear does a linear interpolation.
	Linear = &ease{
		fp: func(t float64) float64 {
			return t
		},
		fd: func(t float64) float64 {
			return 1
		},
	}
	// InCubic does a cubic interpolation.
	InCubic = &ease{
		fp: func(t float64) float64 {
			return t * t * t
		},
		fd: func(t float64) float64 {
			return 3 * t * t
		},
	}
	// OutCubic does a cubic interpolation.
	OutCubic = &ease{
		fp: func(t float64) float64 {
			t--
			return t*t*t + 1
		},
		fd: func(t float64) float64 {
			t--
			return 3 * t * t
		},
	}
	// InOutCubic does a cubic interpolation.
	InOutCubic = &ease{
		fp: func(t float64) float64 {
			t *= 2
			if t < 1 {
				return t * t * t / 2
			}
			t -= 2
			return (t*t*t + 2) / 2
		},
		fd: func(t float64) float64 {
			t *= 2
			if t < 1 {
				return 3 * t * t / 2
			}
			t -= 2
			return 3 * t * t / 2
		},
	}
	// InSine does a sine interpolation.
	InSine = &ease{
		fp: func(t float64) float64 {
			return -math.Cos(t*math.Pi/2) + 1
		},
	}
	// OutSine does a sine interpolation.
	OutSine = &ease{
		fp: func(t float64) float64 {
			return math.Sin(t * math.Pi / 2)
		},
	}
	// InOutSine does a sine interpolation.
	InOutSine = &ease{
		fp: func(t float64) float64 {
			return -(math.Cos(math.Pi*t) - 1) / 2
		},
		fd: func(t float64) float64 {
			return math.Pi * math.Sin(math.Pi*t) / 2
		},
	}
	// InBounce does a bouncing interpolation.
	InBounce = &ease{
		fp: func(t float64) float64 {
			return 1 - OutBounce.fp(1-t)
		},
	}
	// OutBounce does a bouncing interpolation.
	OutBounce = &ease{
		fp: func(t float64) float64 {
			switch {
			case t < 0.3636:
				return 7.5625 * t * t
			case t < 0.7272:
				return 9.075*(t-0.5455)*(t-0.5455) + 0.7
			case t < 0.9:
				return 12.0665*(t-0.8136)*(t-0.8136) + 0.91
			}
			return 10.8*(t-0.95)*(t-0.95) + 0.973
		},
	}
	// InOutBounce does a bouncing interpolation.
	InOutBounce = &ease{
		fp: func(t float64) float64 {
			if t < 0.5 {
				return InBounce.fp(2*t) / 2
			}
			return OutBounce.fp(2*t-1)/2 + 0.5
		},
	}
)
