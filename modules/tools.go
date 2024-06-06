package modules

import "time"

func UntoTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")

}

func GetUnix() int64 {
	return time.Now().Unix()

}

func GetDay() string {
	return time.Now().Format("20060102")
}
