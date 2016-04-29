package main

import (
	"log"
	"ms2mysql/lib"
	"time"

	"github.com/robfig/cron"
)

func main() {
	//handle()
	//return
	c := cron.New()
	spec := "0 */30 8-10,17-21 * * ?"
	c.AddFunc(spec, handle)
	c.Start()

	select {}
}

func handle() {
	lastUpadateTime := ms2mysql.GetMaxTime()
	log.Println("更新时间：", time.Now().Format("2006-01-02 15:04:05"))
	log.Println("上次获取到数据库时间点：", lastUpadateTime)
	s := ms2mysql.GetData(lastUpadateTime)
	if len(s) > 0 {
		ms2mysql.InsertTo(s)
	}
}
