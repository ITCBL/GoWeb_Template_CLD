package main

import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// Init
/**
 * @Description:
 * @param startTime 时间因子：指公司的成立时间，往后能用69年。
 * @param machineID 机器ID
 * @return err
 */
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// GenID
/**
 * @Description:
 * @return int64
 */
func GenID() int64 {
	return node.Generate().Int64()
}
func main() {
	if err := Init("2020-07-01", 1); err != nil {
		fmt.Printf("init failed, err:%v\n", err)
		return
	}
	id := GenID()
	fmt.Println("雪花ID：", id)
}
