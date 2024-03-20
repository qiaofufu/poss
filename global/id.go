package global

import "OSS/utils"

var IDGenerator map[string]*utils.SnowFlake

const (
	WorkID = 1
)

func init() {
	IDGenerator = make(map[string]*utils.SnowFlake)
	IDGenerator[UserTable] = utils.NewSnowFlake(WorkID)
}
