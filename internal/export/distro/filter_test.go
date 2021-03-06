//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"testing"

	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
)

const (
	deviceID1 = "DEV1"
	deviceID2 = "DEV2"

	descriptor1 = "Descriptor1"
	descriptor2 = "Descriptor2"
)

func TestFilterDevice(t *testing.T) {
	// Filter only accepting events from device 1
	f := contract.Filter{}
	f.DeviceIDs = append(f.DeviceIDs, "DEV1")

	// Event from device 1
	eventDev1 := contract.Event{
		Device: deviceID1,
	}
	// Event from device 2
	eventDev2 := contract.Event{
		Device: deviceID2,
	}

	filter := newDevIdFilter(f)
	accepted, _ := filter.Filter(nil)
	if accepted {
		t.Fatal("Event should be filtered out")
	}
	accepted, res := filter.Filter(&eventDev1)
	if !accepted {
		t.Fatal("Event should be accepted")
	}
	if res != &eventDev1 {
		t.Fatal("Event should be the same")
	}
	accepted, _ = filter.Filter(&eventDev2)
	if accepted {
		t.Fatal("Event should be filtered")
	}
}

func TestFilterValue(t *testing.T) {
	f1 := contract.Filter{}
	f1.ValueDescriptorIDs = append(f1.ValueDescriptorIDs, descriptor1)

	f12 := contract.Filter{}
	f12.ValueDescriptorIDs = append(f12.ValueDescriptorIDs, descriptor1)
	f12.ValueDescriptorIDs = append(f12.ValueDescriptorIDs, descriptor2)

	// only accepts value descriptor 1
	filter1 := newValueDescFilter(f1)
	// accepts value descriptor 1 and 2
	filter12 := newValueDescFilter(f12)

	// event with a value descriptor 1
	event1 := contract.Event{}
	event1.Readings = append(event1.Readings, contract.Reading{Name: descriptor1})

	// event with a value descriptor 2
	event2 := contract.Event{}
	event2.Readings = append(event2.Readings, contract.Reading{Name: descriptor2})

	// event with a value descriptor 1 and another 2
	event12 := contract.Event{}
	event12.Readings = append(event12.Readings, contract.Reading{Name: descriptor1})
	event12.Readings = append(event12.Readings, contract.Reading{Name: descriptor2})

	accepted, res := filter1.Filter(nil)
	if accepted {
		t.Fatal("Event should be filtered out")
	}

	accepted, res = filter1.Filter(&event1)
	if !accepted {
		t.Fatal("Event should be accepted")
	}
	if len(res.Readings) != 1 {
		t.Fatal("Event should be one reading, there are ", len(res.Readings))
	}

	accepted, res = filter1.Filter(&event12)
	if !accepted {
		t.Fatal("Event should be accepted")
	}
	if len(res.Readings) != 1 {
		t.Fatal("Event should be one reading, there are ", len(res.Readings))
	}

	accepted, res = filter1.Filter(&event2)
	if accepted {
		t.Fatal("Event should be filtered out")
	}

	accepted, res = filter12.Filter(&event1)
	if !accepted {
		t.Fatal("Event should be accepted")
	}
	if len(res.Readings) != 1 {
		t.Fatal("Event should be one reading, there are ", len(res.Readings))
	}

	accepted, res = filter12.Filter(&event12)
	if !accepted {
		t.Fatal("Event should be accepted")
	}
	if len(res.Readings) != 2 {
		t.Fatal("Event should be one reading, there are ", len(res.Readings))
	}

	accepted, res = filter12.Filter(&event2)
	if !accepted {
		t.Fatal("Event should be accepted")
	}
	if len(res.Readings) != 1 {
		t.Fatal("Event should be one reading, there are ", len(res.Readings))
	}
}
