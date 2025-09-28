package util

import (
	"time"
)

func GetDelayInSecond(hour float32) int {
	// receive hour and convert to second
	return int(hour * 3600)
}

func DelayCalculation(hours float32, publicHolidayMap map[string]string) int {
	processTime := int(hours * 3600)
	startWorkingHour := 9
	endWorkingHour := 17
	workingHour := (endWorkingHour - startWorkingHour) * 3600 // in second
	dayInSecond := 86400

	now := time.Now().In(LocationJakarta())
	yyyy, mm, dd := now.Date()
	tomorrow := time.Date(yyyy, mm, dd+1, 0, 0, 0, 0, now.Location())
	openTime := time.Date(yyyy, mm, dd, startWorkingHour, 0, 0, 0, now.Location()) // waktu buka toko di set pada jam 09.00
	closeTime := time.Date(yyyy, mm, dd, endWorkingHour, 0, 0, 0, now.Location())  // waktu tutup toko di set pada jam 17.00

	// fmt.Println("processTime >>>>>>>>>>>>>>>>>>>", processTime)
	// fmt.Println("openTime >>>>>>>>>>>>>>>>>>>>>>", openTime)
	// fmt.Println("closeTime >>>>>>>>>>>>>>>>>>>>>>", closeTime)

	remainingSLA := 0 // sisa process time/SLA
	addtionalProcessTime := 0

	// Jika termasuk hari ini libur
	if _, ok := publicHolidayMap[now.Format("2006-01-02")]; ok {
		// fmt.Println("hari ini libur >>>>>>>>>>>>>>>>>>>>>>")
		remainingSLA = processTime
	} else {
		if now.Sub(openTime).Seconds() < 0 {
			// fmt.Println("sebelum buka >>>>>>>>>>>>>>>>>>>>>>")
			todaysCut := closeTime.Sub(openTime).Seconds() // waktu yang di gunakan untuk memotong process time didapat dari (sekarang - 17pm)
			// fmt.Println("pengurang >>>>>>>>>>>>>>>>>>>>>>", todaysCut)
			remainingSLA = processTime - int(todaysCut)
			// waktu yang akan ditambahkan jika submit sebelum jam buka diambil dari now hingga jam buka
			addtionalProcessTime = int(openTime.Sub(now).Seconds())
			processTime += addtionalProcessTime
		} else if now.Sub(closeTime).Seconds() > 0 {
			// fmt.Println("setelah tutup >>>>>>>>>>>>>>>>>>>>>>")
			remainingSLA = processTime
		} else {
			// fmt.Println("saat buka >>>>>>>>>>>>>>>>>>>>>>")
			todaysCut := closeTime.Sub(now).Seconds() // waktu yang di gunakan untuk memotong process time didapat dari (sekarang - 17pm)
			// fmt.Println("pengurang >>>>>>>>>>>>>>>>>>>>>>", todaysCut)
			remainingSLA = processTime - int(todaysCut)
		}
	}

	// fmt.Println("addtionalProcessTime >>>>>>>>>>>>>>>>>>>>>>", addtionalProcessTime)
	// fmt.Println("remainingSLA >>>>>>>>>>>>>>>>>>>>>>", remainingSLA)

	totalTime := 0
	// jika tersisa waktu proses di hari berikutnya
	if remainingSLA > 0 {
		// hitung sisa dalam hari
		countNextDays := 1
		// jika sisa waktu lebih dari 8 jam
		if remainingSLA > workingHour {
			countNextDays = (remainingSLA / workingHour) // jumlah hari yang akan dilewati
		}
		// fmt.Println("countNextDays >>>>>>>>>>>>>>>>", countNextDays)
		// fmt.Println("totalTime >>>>>>>>>>>>>>>>>>>>>>", totalTime)

		// jika waktu terakhir melebihi hari ini
		if countNextDays > 0 {
			// tentukan hari terakhir dimulai dari besok
			deadline := tomorrow.AddDate(0, 0, countNextDays)
			// fmt.Println("Deadline Awal >>>>>>>>>>>>", deadline)
			// fmt.Println("publicHolidayMap >>>>>>>>>>>>>>>", publicHolidayMap)

			// tentukan hari besok
			start := now.AddDate(0, 0, 1)
			// lakukan perulangan dari besok sampai hari terakhir
			for d := start; d.After(deadline) == false; d = d.AddDate(0, 0, 1) {
				// fmt.Println("")
				// fmt.Println("Tangal >>>>>>>>>>", d)
				// fmt.Println("Sisa SLA >>>>>>>>", remainingSLA)

				day := d.Format("2006-01-02")
				// jika dlm perulangan hari terdapat public holiday tambahkan waktu sehari dalam detik
				if _, ok := publicHolidayMap[day]; ok {
					// deadline ditambah agar perulangan terus berlanjut hingga mendapat hari terakhir
					deadline = deadline.AddDate(0, 0, 1)
					totalTime += dayInSecond
				} else if remainingSLA > workingHour { // jika sisa waktu lebih dari lama jam kerja maka tambah deadline dan totaltime sehari
					remainingSLA -= workingHour
					totalTime += dayInSecond

					// jika deadline sudah tercapai akan tetapi sisa waktu sla lebih dari sehari kerja maka deadline ditambah 1 hari
					if !deadline.After(d) && remainingSLA > workingHour {
						deadline = deadline.AddDate(0, 0, 1)
					}
				}

				// fmt.Println("total delay >>>>>>>", totalTime)
				// fmt.Println("deadline >>>>>>>>>>", deadline)
				// fmt.Println("")
			}

			// waktu tambahan pada saat hari deadline tergantung dimulai nya jam kerja
			startInSecond := 3600 * startWorkingHour
			totalTime += (startInSecond + remainingSLA)
		}

		// ditambah sisa waktu hari ini karena di callback service akan menghitung dari time.now
		totalTime += int(tomorrow.Sub(now).Seconds())
	} else {
		// jika waktu proses habis hari ini
		totalTime = processTime
	}

	// fmt.Println("totalTime >>>>>>>>>>>>>>>>>>", totalTime)
	// fmt.Println("Cancel Time >>>>>>>>>", now.Add(time.Second*time.Duration(totalTime)))

	return totalTime
}

func TimeToStringLocalDate(d time.Time) string {
	local_loc, _ := time.LoadLocation("Asia/Jakarta")
	localTime := d.In(local_loc)
	return localTime.Format("2006-01-02")
}

func DelayCalculationNormal(hours float32, publicHolidayMap map[string]string) int {
	// fmt.Println("sla hour >>>>>>>>>>", hours)
	processTime := int(hours * 3600)
	dayInSecond := 86400

	now := time.Now().In(LocationJakarta())
	yyyy, mm, dd := now.Date()
	tomorrow := time.Date(yyyy, mm, dd+1, 0, 0, 0, 0, now.Location())
	remainingSLA := 0 // sisa process time
	totalTime := 0    // total waktu delay dalam hitungan detik
	var additionalTime float64

	// Jika hari ini libur
	if _, ok := publicHolidayMap[now.Format("2006-01-02")]; ok {
		additionalTime = tomorrow.Sub(now).Seconds()
		remainingSLA = processTime
		// remainingSLA = processTime + int(additionalTime) // sisa waktu tidak dipotong
	} else {
		todaysCut := tomorrow.Sub(now).Seconds() // waktu yang di gunakan untuk memotong process time didapat dari (sekarang hingga besok)
		remainingSLA = processTime - int(todaysCut)
		totalTime += int(todaysCut)
	}

	// jika tersisa waktu proses di hari berikutnya
	if remainingSLA > 0 {
		// hitung sisa dalam hari
		countNextDays := 1
		// jika sisa waktu lebih dari 24 jam
		if remainingSLA > dayInSecond {
			countNextDays = (remainingSLA / dayInSecond) // jumlah hari yang akan dilewati
		}

		// jika waktu terakhir melebihi hari ini
		if countNextDays > 0 {
			// tentukan hari terakhir dimulai dari besok
			deadline := tomorrow.AddDate(0, 0, countNextDays)

			// tentukan hari besok
			start := now.AddDate(0, 0, 1)
			// lakukan perulangan dari besok sampai hari terakhir
			for d := start; d.After(deadline) == false; d = d.AddDate(0, 0, 1) {
				day := d.Format("2006-01-02")
				// jika dlm perulangan hari terdapat public holiday tambahkan waktu sehari dalam detik
				if _, ok := publicHolidayMap[day]; ok {
					// deadline ditambah agar perulangan terus berlanjut hingga mendapat hari terakhir
					deadline = deadline.AddDate(0, 0, 1)
					totalTime += dayInSecond
				} else if remainingSLA > dayInSecond { // jika sisa waktu lebih dari lama jam kerja maka tambah deadline dan totaltime sehari
					remainingSLA -= dayInSecond
					totalTime += dayInSecond
					deadline = deadline.AddDate(0, 0, 1)
				} else {
					totalTime += remainingSLA
					break
				}
			}

			totalTime += int(additionalTime) // ditambah sisa waktu libur hari pertama
		}
	} else {
		// jika waktu proses habis hari ini
		totalTime = processTime
	}

	// fmt.Println("totalTime >>>>>>>>>>>>>>>>>>", totalTime/3600)
	// fmt.Println("Cancel Time >>>>>>>>>", now.Add(time.Second*time.Duration(totalTime)))

	return totalTime
}
