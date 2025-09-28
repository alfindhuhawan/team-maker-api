package util

import (
	"strconv"
	"time"
)

func GetTimeSlotArrayService() map[string]int {
	arrayService := make(map[string]int, 0)

	arrayService["same_day"] = 0
	arrayService["next_day"] = 1
	arrayService["2_days"] = 2
	arrayService["3_days"] = 3
	arrayService["4_days"] = 4
	arrayService["5_days"] = 5
	arrayService["6_days"] = 6
	arrayService["7_days"] = 7

	return arrayService
}

func DefineIsSameDayAndAdditionalTime(service string) int {
	arrayService := GetTimeSlotArrayService()

	additionalTime := 0

	if _, exist := arrayService[service]; exist {
		additionalTime = arrayService[service]
	}

	return additionalTime
}

func DefineFirstDay(isFound bool, firstDay int, service string) (bool, int) {
	arrayService := GetTimeSlotArrayService()

	if !isFound {
		if _, exist := arrayService[service]; exist {
			firstDay = arrayService[service]
			isFound = true
		}
	}

	return isFound, firstDay
}

func MergeMinMaxTimeSlot(minTimeSlot int, maxTimeSlot int) string {
	var preferedTimeSlot string

	// MIN
	strMinTimeSlot := convertTimeSlotToString(minTimeSlot)
	// MAX
	strMaxTimeSlot := convertTimeSlotToString(maxTimeSlot)

	// SET PREFERED TIME SLOT
	if strMinTimeSlot != "" && strMaxTimeSlot != "" {
		preferedTimeSlot = strMinTimeSlot + " - " + strMaxTimeSlot
	}

	return preferedTimeSlot
}

func convertTimeSlotToString(timeslot int) string {
	var returnTimeSlot string
	stringTimeSlot := strconv.Itoa(timeslot)

	count := 0
	for range stringTimeSlot {
		count++
	}
	if count == 3 {
		stringTimeSlot = "0" + stringTimeSlot
	}

	t, _ := time.Parse("1504", stringTimeSlot)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	returnTimeSlot = t.Format("15:04")

	return returnTimeSlot
}
