package algorithm

// CircularIndex provide index in a curcular way
// Using CircularIndex, any slices can be turned into a circular list
type CircularIndex struct {
	entry      int
	upperBound int
	direction  int
	step       int
	incomming  chan int
	started    bool
	done       bool
}

func (c *CircularIndex) start() {
	c.started = true
	var cur = c.entry
	go func() {
		c.incomming <- cur
		for !c.done {
			switch c.direction {
			case 1:
				cur += c.step
			case -1:
				cur -= c.step
			default:
				panic("direction value can only be 1 or -1")
			}
			if cur > c.upperBound-1 {
				cur -= c.upperBound
			} else if cur < 0 {
				cur += c.upperBound
			}
			c.incomming <- cur
		}
	}()
}

// Next method get the next index number
func (c *CircularIndex) Next() (int, bool) {
	if !c.started {
		c.start()
	}
	return <-c.incomming, !c.done
}

// Stop generating index
func (c *CircularIndex) Stop() {
	c.done = true
}

// Reverse the direction
func (c *CircularIndex) Reverse() {
	c.direction = -c.direction
}

// SetStep change the step, if the provided number
func (c *CircularIndex) SetStep(n int) {
	c.step = n % c.upperBound
}

// NewCircularIndex method create a new CircularIndex object
func NewCircularIndex(entry, upperBound int) *CircularIndex {
	return &CircularIndex{
		entry:      entry,
		upperBound: upperBound,
		direction:  1,
		step:       1,
		incomming:  make(chan int),
		started:    false,
		done:       false,
	}
}
