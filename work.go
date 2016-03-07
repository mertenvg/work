package work

// Make a new Work queue
func Make(ts ...Task) (w Work) {
	w.Add(ts...)
	return
}

// Task represents a callable task to be executed
type Task func() error

// Work slice of consecutive tasks that each may or may not have a single point of failure
type Work []Task

// Add a Task to the work queue
func (w *Work) Add(ts ...Task) {
	*w = append(*w, ts...)
}

// Do the work
func (w Work) Do() (err error) {
	for _, t := range w {
		err = t()
		if err != nil {
			return
		}
	}
	return
}
