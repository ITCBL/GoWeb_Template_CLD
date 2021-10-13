package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

/**
 * @Description:
 * @param startTime 时间因子：指公司的成立时间，往后能用69年。
 * @param machineID 机器ID int64
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

/**
 * @Description:
 * @return int64
 */
func GenID() int64 {
	return node.Generate().Int64()
}
