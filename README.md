# work
Create a work queue to run tasks in sequence and stop when an error is encountered

### Install

`go get github.com/mertenvg/work`

### Usage

```go
work := work.Make()

// task A
work.Add(func() (err error) {
	fmt.Println("A")
	return
})

// task B
work.Add(func() (err error) {
	fmt.Println("B")
	return
})

// task C
work.Add(func() (err error) {
	fmt.Println("C")
	err = errors.New("I don't know why, but Fail!")
	return
})

// task D
work.Add(func() (err error) {
	fmt.Println("D")
	return
})

err := work.Do()
if err != nil {
	fmt.Println(err.Error())
}
```
