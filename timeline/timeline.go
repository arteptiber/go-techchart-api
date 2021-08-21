package timeline

import (
	"math/big"
	"time"
)

type Timeline struct {
	Store []*TimelineSlot
}

type TimelineSlot struct {
	Time  time.Time // todo hide
	Value big.Int
}

func AppendToTimeline(timelineInstance Timeline, slot *TimelineSlot) Timeline {
	storeLen := len(timelineInstance.Store)
	if storeLen == 0 {
		timelineInstance.Store = []*TimelineSlot{slot}
		return timelineInstance
	}
	if storeLen == 1 {
		availableSlot := timelineInstance.Store[0]
		if slot.Time.After(availableSlot.Time) {
			return Timeline{Store: []*TimelineSlot{availableSlot, slot}}
		}

		return Timeline{Store: []*TimelineSlot{slot, availableSlot}}
	}

	low := 0
	high := storeLen - 1
	mid := storeLen / 2
	for {
		if mid == 0 {
			return appendToTimelineStore(timelineInstance, slot, 0)
		}

		lowStoreTime, highStoreTime := timelineInstance.Store[mid-1].Time, timelineInstance.Store[mid].Time

		if inTimeSpan(lowStoreTime, highStoreTime, slot.Time) {
			return appendToTimelineStore(timelineInstance, slot, mid)
		}

		if slot.Time.After(lowStoreTime) {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = (low + high) / 2
	}

}

func appendToTimelineStore(timelineInstance Timeline, slot *TimelineSlot, index int) Timeline {
	if index == 0 {
		timelineInstance.Store = append([]*TimelineSlot{slot}, timelineInstance.Store...)
		return timelineInstance
	}

	storeSliceCopy := make([]*TimelineSlot, index)
	copy(storeSliceCopy, timelineInstance.Store[0:index])

	store := append(storeSliceCopy, slot)
	store = append(store, timelineInstance.Store[index:len(timelineInstance.Store)]...)
	timelineInstance.Store = store
	return timelineInstance
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func (t *Timeline) Len() int {
	return len(t.Store)
}
