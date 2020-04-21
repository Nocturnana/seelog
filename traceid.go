package seelog

import (
	"github.com/petermattis/goid"
	"strconv"
	"sync"
	"time"
)

var TraceIdMap map[int64]string
var mutex sync.Mutex

func init() {
	TraceIdMap = make(map[int64]string)
}

func getDefaultTraceId(id int64) string {
	now := time.Now().UnixNano()
	return strconv.FormatInt(now/1e6, 10) + strconv.FormatInt(id, 10)

}

func InitTraceId() {
	defer mutex.Unlock()
	mutex.Lock()
	id := goid.Get()
	TraceIdMap[id] = getDefaultTraceId(id)
}

func InitTraceIdWithId(traceid string) {
	defer mutex.Unlock()
	mutex.Lock()
	id := goid.Get()
	TraceIdMap[id] = traceid
}

func GetTraceId() string {
	defer mutex.Unlock()
	mutex.Lock()
	id := goid.Get()
	traceId, ok := TraceIdMap[id]
	if ok {
		return traceId
	} else {
		traceId := getDefaultTraceId(id)
		TraceIdMap[id] = traceId
		return traceId
	}

}

func DeleteTraceId() {
	defer mutex.Unlock()
	mutex.Lock()
	id := goid.Get()
	delete(TraceIdMap, id)
}
