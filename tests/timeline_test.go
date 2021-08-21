package timeline_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/arteptiber/techchart/timeline"
)

func TestAppendToTimelineBasic(t *testing.T) {
	slot1 := timeline.TimelineSlot{Time: time.Unix(10, 0), Value: *big.NewInt(100)}
	slot2 := timeline.TimelineSlot{Time: time.Unix(30, 0), Value: *big.NewInt(300)}
	timelineInstance := timeline.Timeline{Store: []*timeline.TimelineSlot{&slot1, &slot2}}

	slotToAppend := timeline.TimelineSlot{Time: time.Unix(20, 0), Value: *big.NewInt(200)}
	expectedTimeline := timeline.Timeline{Store: []*timeline.TimelineSlot{&slot1, &slotToAppend, &slot2}}

	newTimeline := timeline.AppendToTimeline(timelineInstance, &slotToAppend)

	if !compareTimelines(&newTimeline, &expectedTimeline) {
		t.Error("not equal")
	}
}

func TestAppendToTimelineAdvanced(t *testing.T) {
	slot1 := timeline.TimelineSlot{Time: time.Unix(10, 0), Value: *big.NewInt(100)}
	slot2 := timeline.TimelineSlot{Time: time.Unix(20, 0), Value: *big.NewInt(200)}
	slot3 := timeline.TimelineSlot{Time: time.Unix(30, 0), Value: *big.NewInt(300)}
	slot4 := timeline.TimelineSlot{Time: time.Unix(50, 0), Value: *big.NewInt(500)}
	timelineInstance := timeline.Timeline{Store: []*timeline.TimelineSlot{&slot1, &slot2, &slot3, &slot4}}

	slotToAppend := timeline.TimelineSlot{Time: time.Unix(40, 0), Value: *big.NewInt(400)}
	expectedTimeline := timeline.Timeline{Store: []*timeline.TimelineSlot{&slot1, &slot2, &slot3, &slotToAppend, &slot4}}

	newTimeline := timeline.AppendToTimeline(timelineInstance, &slotToAppend)

	if !compareTimelines(&newTimeline, &expectedTimeline) {
		t.Error("not equal")
	}
}

func TestAppendToEmptyTimeline(t *testing.T) {
	timelineInstance := timeline.Timeline{Store: []*timeline.TimelineSlot{}}
	slotToAppend := timeline.TimelineSlot{Time: time.Unix(10, 0), Value: *big.NewInt(100)}

	expectedTimeline := timeline.Timeline{Store: []*timeline.TimelineSlot{&slotToAppend}}
	newTimeline := timeline.AppendToTimeline(timelineInstance, &slotToAppend)

	if !compareTimelines(&newTimeline, &expectedTimeline) {
		t.Error("not equal")
	}
}

func TestAppendSlotWithSmallerValueToToTimelineWithSingleSlot(t *testing.T) {
	slot := timeline.TimelineSlot{Time: time.Unix(20, 0), Value: *big.NewInt(200)}
	timelineInstance := timeline.Timeline{Store: []*timeline.TimelineSlot{&slot}}

	slotToAppend := timeline.TimelineSlot{Time: time.Unix(10, 0), Value: *big.NewInt(100)}
	expectedTimeline := timeline.Timeline{Store: []*timeline.TimelineSlot{&slotToAppend, &slot}}
	newTimeline := timeline.AppendToTimeline(timelineInstance, &slotToAppend)

	if !compareTimelines(&newTimeline, &expectedTimeline) {
		t.Error("not equal")
	}
}

func TestAppendSlotWithLargerValueToToTimelineWithSingleSlot(t *testing.T) {
	slot := timeline.TimelineSlot{Time: time.Unix(10, 0), Value: *big.NewInt(100)}
	timelineInstance := timeline.Timeline{Store: []*timeline.TimelineSlot{&slot}}

	slotToAppend := timeline.TimelineSlot{Time: time.Unix(20, 0), Value: *big.NewInt(200)}
	expectedTimeline := timeline.Timeline{Store: []*timeline.TimelineSlot{&slot, &slotToAppend}}
	newTimeline := timeline.AppendToTimeline(timelineInstance, &slotToAppend)

	if !compareTimelines(&newTimeline, &expectedTimeline) {
		t.Error("not equal")
	}
}

func compareTimelines(newTimeline, expectedTimeline *timeline.Timeline) bool {
	if newTimeline.Len() != expectedTimeline.Len() {
		return false
	}

	for i := 0; i < newTimeline.Len(); i++ {
		if newTimeline.Store[i].Value.Cmp(&expectedTimeline.Store[i].Value) != 0 {
			return false
		}
	}

	return true
}
