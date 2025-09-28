package util

import (
	"fmt"
	"sort"
)

const max = 5

func main() {
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 1, 1))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 1, 1, 1))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 1, 1, 5))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 1, 1, 4))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 4, 1, 4))
	fmt.Printf("%v\n\n", SplitAWB(3, 1, 1, 1, 4))
	fmt.Printf("%v\n\n", SplitAWB(3, 1, 1, 2, 4))
	fmt.Printf("%v\n\n", SplitAWB(3, 1, 1, 2, 4))
	fmt.Printf("%v\n\n", SplitAWB(1, 3, 3, 3, 1))
	fmt.Printf("%v\n\n", SplitAWB(4, 4, 1, 1, 1))
	fmt.Printf("%v\n\n", SplitAWB(4, 4, 3, 1, 3))
	fmt.Printf("%v\n\n", SplitAWB(2, 1, 3, 1))
	fmt.Printf("%v\n\n", SplitAWB(2, 1))
	fmt.Printf("%v\n\n", SplitAWB(3, 4, 3, 1, 1, 3, 5))
	fmt.Printf("%v\n\n", SplitAWB(2, 2, 2, 2))
	fmt.Printf("%v\n\n", SplitAWB(1, 2, 3, 4))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 1, 4, 3))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 2, 4, 3))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 2, 1, 2))
	fmt.Printf("%v\n\n", SplitAWB(1, 1, 2, 4, 2, 3, 2, 3))
	fmt.Printf("%v\n\n", SplitAWB(1, 5, 1, 2, 4, 2, 3, 2, 3, 1, 2, 4, 2))
	fmt.Printf("%v\n\n", SplitAWB(1, 4, 3, 4))
}

func SplitAWB(rawInput ...int) [][]int {

	// sort desc
	sort.Slice(rawInput, func(i, j int) bool {
		return rawInput[i] > rawInput[j]
	})

	// fmt.Println(rawInput)

	mapWeight := map[int]int{}
	sliceWeight := make([]int, 0)

	for _, r := range rawInput {
		_, exist := mapWeight[r]
		if exist {
			mapWeight[r]++
		} else {
			mapWeight[r] = 1
			sliceWeight = append(sliceWeight, r)
		}
	}

	// fmt.Println(mapWeight)
	// fmt.Println(sliceWeight)

	// =======================================================

	// kita punya list of keranjang
	listOfKeranjang := make([][]int, 0)

	// kita akan mulai dari index ke-0
	index := 0

	// siapkan keranjang kosong
	keranjang := make([]int, 0)

	for index < len(sliceWeight) {

		if len(mapWeight) == 0 {
			break
		}

		// kita mulai dari index
		theWeight := sliceWeight[index]

		for {

			// ada gak di map?
			_, exist := mapWeight[theWeight]

			// ada di map
			if exist {

				// hitung total berat dalam keranjang
				totalWeight := getTotalWeight(keranjang)
				if totalWeight+theWeight <= max {
					totalWeight += theWeight

				} else {

					// masukkan ke list of keranjang
					listOfKeranjang = append(listOfKeranjang, keranjang)

					// siapkan keranjang kosong berikutnya
					keranjang = make([]int, 0)
					break
				}

				// masukkan dalam keranjang
				keranjang = append(keranjang, theWeight)

				// setelah masuk keranjang, jangan lupa kurangi dari map
				mapWeight[theWeight]--

				// kalo mapnya sudah 0 quantitynya,
				if mapWeight[theWeight] == 0 {

					// maka hapus mapnya
					delete(mapWeight, theWeight)

					_, exist := mapWeight[sliceWeight[index]]
					if !exist {
						index++
					}

					if len(mapWeight) == 0 {
						// masukkan ke list of keranjang
						listOfKeranjang = append(listOfKeranjang, keranjang)
						break
					}

				}

				if totalWeight == max {

					// masukkan ke list of keranjang
					listOfKeranjang = append(listOfKeranjang, keranjang)

					// siapkan keranjang kosong berikutnya
					keranjang = make([]int, 0)

				} else {

					theWeight = max - totalWeight

				}

			} else //
			// gak ada di map
			{
				theWeight--
				if theWeight == 0 {
					break
				}
			}

		}

	}

	return listOfKeranjang

}

func getTotalWeight(keranjang []int) int {
	total := 0
	for _, x := range keranjang {
		total += x
	}
	return total
}
