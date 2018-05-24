package g_test

import (
	"github.com/rockamring/goproject/server/go"
	"fmt"
	"time"
)

func Example() {
	d := g.New(10)

	var res int
	d.Go(func() {
		fmt.Println("1+1=?")
		res = 1 +1
	}, func() {
		fmt.Println(res)
	})

	d.Cb(<-d.ChanCb)


	d.Go(func() {
		fmt.Print("my name is ")
	}, func() {
		fmt.Println("luo")
	})

	d.Close()

	// Output:
	// 1+1=?
	// 2
	// my name is luo
}

func ExampleLinearContext() {
	d := g.New(10)

	d.Go(func() {
		time.Sleep(1*time.Second)
		fmt.Println("first")
	}, nil)

	d.Go(func() {
		fmt.Println("second")
	}, nil)

	d.Cb(<-d.ChanCb)
	d.Cb(<-d.ChanCb)

	c := d.NewLinearContext()
	c.Go(func() {
		time.Sleep(1*time.Second)
		fmt.Println("first")
	}, nil)
	c.Go(func() {
		fmt.Println("second")
	}, nil)

	d.Close()

	// Output:
	// second
	// first
	// first
	// second
}