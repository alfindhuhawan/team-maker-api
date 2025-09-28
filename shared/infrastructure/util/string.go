package util

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"team-maker-api/shared/model/vo"
	"time"

	"github.com/leekchan/accounting"
)

func CreateURLGCS(gcsBaseURL string, url string) string {
	completedUrl := "-"

	if url != "" {
		completedUrl = gcsBaseURL + url
	}

	return completedUrl
}

func CreateETAString(firstDay int, additionalTime int) string {
	etaString := "-"

	timeNow := time.Now().AddDate(0, 0, firstDay).In(LocationJakarta()).Format(time.RFC822)

	var timeNext string
	if additionalTime > 0 {
		dateNext := time.Now().In(LocationJakarta()).AddDate(0, 0, additionalTime)
		timeNext = dateNext.Format(time.RFC822)
	}

	if timeNow == timeNext {
		etaString = "Tiba " + timeNow[0:6]
	} else {
		if timeNow != "" && timeNext == "" { // TIME NOW ONLY EXIST
			etaString = "Tiba " + timeNow[0:6]
		} else if timeNow == "" && timeNext != "" { // TIME NEXT ONLY EXIST
			etaString = "Tiba " + timeNext[0:6]
		} else if timeNow != "" || timeNext != "" { // BOTH EXIST
			if timeNow[3:6] == timeNext[3:6] {
				etaString = "Tiba " + timeNow[0:2] + " - " + timeNext[0:6]
			} else {
				etaString = "Tiba " + timeNow[0:6] + " - " + timeNext[0:6]
			}
		}
	}

	return etaString
}

func FormatMoneyInString(money vo.Money) string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 0, Thousand: ".", Decimal: ","}

	stringMoney := ac.FormatMoney(money)

	return stringMoney
}

func StringContains(slices []string, comparizon string) bool {
	for _, a := range slices {
		if a == comparizon {
			return true
		}
	}

	return false
}

func ConvertByteToString(inBytes []byte) string {
	myString := string(inBytes[:])

	return myString
}

func ChangeEnterToSpace(s string) string {
	re := regexp.MustCompile(`\r?\n`)
	returnString := re.ReplaceAllString(s, " ")

	return returnString
}

func CreateETAStringNewLogic(limit int, daysAvailable map[string][]string, daysUnavailable map[string]int) (string, error) {
	timeSlotArrayService := GetTimeSlotArrayService()

	etaString := "-"

	// CHECK LENGTH UNAVAILABLE DAYS
	lenUnavailableDays := len(daysUnavailable)

	// FOR SORTING
	var keys []string
	for key := range daysAvailable {
		keys = append(keys, key)
	}

	// SORT BY KEYS
	sort.Strings(keys)

	// GET BY LIMIT
	var firstDate string
	var firstDaysServices string
	var lastDate string
	var lastDaysServices string
	totalData := 0      // TOTAL DATA (KEY AND SERVICE) TO COMPARE WITH LIMIT
	keyLen := len(keys) // LEN KEY TO GET LAST DATA ON ARRAY
	isLastKey := false  // LAST KEY TO GET LAST DATA ON ARRAY
	totalKey := 0       // TOTAL KEY

	for _, key := range keys {
		// CHECK LAST KEY
		if totalKey == keyLen-1 {
			isLastKey = true
		}

		lenService := len(daysAvailable[key]) // LEN SERVICE TO GET LAST DATA ON ARRAY
		totalService := 0                     // TOTAL SERVICE TO GET LAST DATA ON ARRAY

		lastDateGet := false // FOR BREAK AFTER GETTING LAST DATE

		for _, s := range daysAvailable[key] {
			// GET FIRST DATE
			if totalData == 0 {
				firstDate = key

				firstDaysServices = s
			}

			// GET LAST DATE
			if totalData == limit-1 || (isLastKey && lenService-1 == totalService) { // WE MUST CHECK TOTAL DATA COMPARE TO LIMIT OR ITS THE LAST DATA ON ARRAY
				lastDate = key

				lastDaysServices = s

				lastDateGet = true
				break
			}

			totalData++
			totalService++
		}

		totalKey++

		if lastDateGet {
			break
		}
	}

	if firstDate != "" {
		additionalTimeFirstService := timeSlotArrayService[firstDaysServices]
		additionalTimeLastService := timeSlotArrayService[lastDaysServices]

		// CHECK IF UNAVAILABLE DAYS ON FIRST DATE
		layoutTime := "2006-01-02" // LAYOUT

		// CONVERT FIRST DATE AND LAST DATE
		// FIRST DATE
		convertFirstDate, err := time.Parse(layoutTime, firstDate)
		if err != nil {
			return "", err
		}

		// LAST DATE
		convertLastDate, err := time.Parse(layoutTime, lastDate)
		if err != nil {
			return "", err
		}

		firstDate = convertFirstDate.AddDate(0, 0, additionalTimeFirstService).In(LocationJakarta()).Format(layoutTime)
		lastDate = convertLastDate.AddDate(0, 0, additionalTimeLastService).In(LocationJakarta()).Format(layoutTime)

		for i := 0; i < lenUnavailableDays; i++ {
			// FIRST DATE CHECK ITS UNAVAILABLE OR NOT
			if _, exist := daysUnavailable[firstDate]; exist {
				convertTemporaryFD, err := time.Parse(layoutTime, firstDate)
				if err != nil {
					return "", err
				}

				firstDate = convertTemporaryFD.AddDate(0, 0, 1).In(LocationJakarta()).Format(layoutTime)
			}

			// LAST DATE CHECK ITS UNAVAILABLE OR NOT
			if _, exist := daysUnavailable[lastDate]; exist {
				convertTemporaryLD, err := time.Parse(layoutTime, lastDate)
				if err != nil {
					return "", err
				}

				lastDate = convertTemporaryLD.AddDate(0, 0, 1).In(LocationJakarta()).Format(layoutTime)
			}
		}

		if lastDate == "" {
			lastDate = firstDate
		}

		// CONVERT FIRST DATE AND LAST DATE TO FORMAT ETA
		// FIRST DATE
		convertFirstDate, err = time.Parse(layoutTime, firstDate)
		if err != nil {
			return "", err
		}

		// LAST DATE
		convertLastDate, err = time.Parse(layoutTime, lastDate)
		if err != nil {
			return "", err
		}

		timeFirst := convertFirstDate.In(LocationJakarta()).Format(time.RFC822)
		timeLast := convertLastDate.In(LocationJakarta()).Format(time.RFC822)

		if timeFirst == timeLast {
			etaString = "Tiba " + timeFirst[0:6]
		} else {
			if timeFirst != "" && timeLast == "" { // TIME NOW ONLY EXIST
				etaString = "Tiba " + timeFirst[0:6]
			} else if timeFirst == "" && timeLast != "" { // TIME NEXT ONLY EXIST
				etaString = "Tiba " + timeLast[0:6]
			} else if timeFirst != "" || timeLast != "" { // BOTH EXIST
				if timeFirst[3:6] == timeLast[3:6] {
					etaString = "Tiba " + timeFirst[0:2] + " - " + timeLast[0:6]
				} else {
					etaString = "Tiba " + timeFirst[0:6] + " - " + timeLast[0:6]
				}
			}
		}
	}

	return etaString, nil
}

func DebuggingStruct(input any) (string, error) {
	dataInByte, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	myString := string(dataInByte[:])

	return myString, nil
}

func DebuggingStructWithoutError(input any) string {
	var myString string
	dataInByte, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
	} else {
		myString = string(dataInByte[:])
	}

	return myString
}

func ConvertEstimatedTime(estimatedTime int64) string {
	stringEstimatedTime := strconv.Itoa(int(estimatedTime))
	var result string
	if len(stringEstimatedTime) > 3 {
		result = fmt.Sprintf("%s%s%s", stringEstimatedTime[:2], ":", stringEstimatedTime[2:])
	} else {
		result = fmt.Sprintf("0%s%s%s", stringEstimatedTime[:1], ":", stringEstimatedTime[1:])
	}

	return result
}

func DebuggingResponseFromExternal(input []byte) (string, error) {
	var anyData interface{}
	err := json.Unmarshal(input, &anyData)
	if err != nil {
		fmt.Println(err.Error())
	}

	myString, err := DebuggingStruct(anyData)
	if err != nil {
		fmt.Println(err.Error())
	}

	return myString, nil
}
