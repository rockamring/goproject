package timer_test

import (
	"github.com/rockamring/goproject/server/timer"
	"fmt"
	"time"
)

func ExampleTimer() {
	d := timer.NewDispatcher(10)

	d.AfterFunc(1, func() {
		fmt.Println("my name is luo")
	})

	t := d.AfterFunc(1, func() {
		fmt.Println("hehe")
	})
	t.Stop()

	(<-d.ChanTimer).Cb()

	// Output:
	// my name is luo
}

func ExampleCronExpr() {
	cronExpr, err := timer.NewCronExpr("0 22 * * *")
	if err != nil {
		return
	}

	fmt.Println(cronExpr.Next(time.Date(
		2000, 1, 1,
		20, 10, 5,
		0, time.UTC,
	)))

	// Output:
	// 2000-01-01 22:00:00 +0000 UTC
}

func ExampleCron() {
	d := timer.NewDispatcher(10)

	cronExpr, err := timer.NewCronExpr("* * * * * *")
	if err != nil {
		return
	}

	var c *timer.Cron
	c = d.CronFunc(cronExpr, func() {
		fmt.Println("my name is luo")
		c.Stop()
	})

	(<-d.ChanTimer).Cb()

	// Output:
	// my name is luo
}