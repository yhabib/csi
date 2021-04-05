package clock

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func (c *Clock) recalculateTime() {
	c.hour += c.minute / 60
	c.minute %= 60
	if c.minute < 0 {
		c.minute += 60
		c.hour--
	}

	c.hour %= 24
	if c.hour < 0 {
		c.hour += 24
	}
}

func New(hour, minute int) Clock {
	c := Clock{hour: hour, minute: minute}
	c.recalculateTime()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return New(c.hour, c.minute)
}

func (c Clock) Subtract(minutes int) Clock {
	return c.Add(-1 * minutes)
}
