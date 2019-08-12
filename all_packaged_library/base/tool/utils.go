/*
 * @Author: howie
 * @Date: 2019-07-02 11:38:19
 * @Last Modified by: holiday
 * @Last Modified time: 2019-07-02 13:19:17
 */
package tool

import (
	"time"
)

func GetTime() int64 {
	return time.Now().Unix()

}

