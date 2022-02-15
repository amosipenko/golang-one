package sliceSort

func SortByInserts(slice []int) {
	// Сортируем: начиная со 2 элемента сравниваем с предыдущим и сдвигаем его влево
	for i := 1; i < len(slice); i++ {
		for j := i; j > 0 && slice[j-1] > slice[j]; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
	}
}

func SortByBubble(slice []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
	}
}
