package ease

import "time"

// Ease is an interface that implements the Interpolate function.
type Ease interface {
	// Interpolate interpolates an input (0.0 to 1.0 inclusive) with the
	// corresponding easing function. Use this to get the position.
	Interpolate(float64) float64
}

// Derivative is an interface that implements the derivative function of an ease
// function.
type Derivative interface {
	// Derivative irterpolates an input (0.0 to 1.0 inclusive) with the
	// derivative of the corresponding easing function. Use this to get the
	// velocity.
	Derivative(float64) float64
}

type ease struct {
	fp func(float64) float64
	fd func(float64) float64
}

// Interpolate interpolates an input (0.0 to 1.0 inclusive) with the
// corresponding easing function. Use this to get the position.
func (e *ease) Interpolate(t float64) float64 {
	return e.fp(t)
}

// Derivative interpolates an input (0.0 to 1.0 inclusive) with the
// derivative of the corresponding easing function. Use this to get the
// velocity.
func (e *ease) Derivative(t float64) float64 {
	return e.fd(t)
}

// Tween automatically interpolates a value with a give ease function lasting a
// given duration. A write function with a func(float64) signature must be
// provided to write the value to the desired variable.
//
// Tween returns a stop function that can be used to stop the interpolation.
//
// The update rate of the interpolation is 10ms.
func Tween(write func(float64), from, to float64, duration time.Duration, ease Ease) (stop func()) {
	const tick = 10 * time.Millisecond
	t := 0.0

	delta := to - from
	step := float64(tick) / float64(duration)

	stopper := make(chan struct{})

	go func() {
		ticker := time.NewTicker(tick)
		defer ticker.Stop()

		write(from)
		for {
			select {
			case <-stopper:
				return
			case <-ticker.C:
				t += step
				if t > 1.0 {
					write(to)
					return
				}
				write(delta*ease.Interpolate(t) + from)
			}
		}
	}()

	return func() {
		select {
		case <-stopper:
		default:
			close(stopper)
		}
	}
}
