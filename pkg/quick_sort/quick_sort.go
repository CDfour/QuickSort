package quicksort

// Алгоритм быстрой сортировки
func QuickSort(array []int) {
	max := len(array) - 1
	mid := (array)[max/2+1]

	i, j := 0, max

	for i <= j {
		for (array)[i] < mid {
			i++
		}
		for (array)[j] > mid {
			j--
		}

		if i <= j {

			if (array)[i] == (array)[j] {
				i++
				j--
				continue
			}

			temp := (array)[i]
			(array)[i] = (array)[j]
			(array)[j] = temp

			i++
			j--
		}
	}

	if j > 0 {
		newArray := (array)[:j+1]
		QuickSort(newArray)
	}
	if i < max {
		newArray := (array)[i:]
		QuickSort(newArray)
	}
}
