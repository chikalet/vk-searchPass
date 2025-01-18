// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// in := bufio.NewReader(os.Stdin)
// 	// out := bufio.NewWriter(os.Stdout)
// 	// defer out.Flush()

// 	// array := []string{}
// 	// in.ReadString('\n')
// 	// matrix, _ := in.ReadString('\n')

// 	// matrixArray := strings.Fields(matrix)
// 	// numArray := make([]int, len(matrixArray))
// 	// for i, s := range matrixArray {
// 	// 	num, _ := strconv.Atoi(s)
// 	// 	numArray[i] = num
// 	// }
// 	// fmt.Println(numArray[0])
// 	// fmt.Println(matrixArray)

// 	// for i := 0; i < numArray[0]-1; i++ {
// 	// 	inn := bufio.NewReader(os.Stdin)
// 	// 	var meanings string
// 	// 	fmt.Fscan(inn, &meanings)
// 	// 	//meanings, _ := in.ReadString('\n')
// 	// 	meanings = strings.TrimSpace(meanings)
// 	// 	array[i] = meanings
// 	// }
// 	// fmt.Println(array)

// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	line, _ := in.ReadString('\n')
// 	matrix := strings.Fields(line)
// 	numArray := make([]int, len(matrix))
// 	for i, s := range matrix {
// 		num, _ := strconv.Atoi(s)
// 		numArray[i] = num
// 	}
// 	fmt.Println(numArray[0])
// 	fmt.Println(matrix)
// 	array := make([]string, len(numArray))
// 	for i := 0; i < numArray[0]; i++ {
// 		meanings, _ := in.ReadString('\n')
// 		meanings = strings.TrimSpace(meanings)
// 		array[i] = meanings
// 	}
// 	initCoord, _ := in.ReadString('\n')
// 	finCoord, _ := in.ReadString('\n')
// 	fmt.Fprintln(out, "Это NUMARRAY:", numArray)
// 	fmt.Fprintln(out, "Вы ввели:", line)
// 	fmt.Fprintln(out, "Вы сохранили", array[0])
// 	fmt.Fprintln(out, "Длина сохранения равна", len(array))
// 	fmt.Fprintln(out, "Начальные координаты", initCoord)
// 	fmt.Fprintln(out, "Финишные координаты", finCoord)

// 	result := Validation(numArray, initCoord, finCoord, array)

// 	if result == "OK" {
// 		SearchPass(numArray, initCoord, finCoord, array)
// 	}

// }

// func SearchPass(numArray []int, initCoord, finCoord string, array []string) string {

// 	return "NO"
// }

// func Validation(numArray []int, initCoord, finCoord string, array []string) string {
// 	initCoordArray := strings.Fields(initCoord)
// 	finCoordArray := strings.Fields(finCoord)
// 	sortArray := make([]int, len(array))

// 	if len(numArray) != 2 {
// 		return "NO"
// 	} else if len(initCoordArray) != 2 || len(finCoordArray) != 2 {
// 		fmt.Println("введены неправильные конечные точки")
// 		return "NO"
// 	}

// 	for _, v := range initCoordArray {
// 		matrix, _ := strconv.Atoi(v)
// 		if matrix < 0 {
// 			fmt.Println("координата не может быть отрицательной")
// 			return "NO"
// 		}
// 	}
// 	for _, v := range finCoordArray {
// 		matrix, _ := strconv.Atoi(v)
// 		if matrix < 0 {
// 			fmt.Println("координата не может быть отрицательной")
// 			return "NO"
// 		}
// 	}
// 	for _, v := range array {
// 		matrix := strings.Fields(v)
// 		if len(matrix) != len(sortArray) {
// 			fmt.Println("Размер строки не совпадает с ожидаемым")
// 			return "NO"
// 		}
// 		for i, s := range matrix {
// 			sortirated, err := strconv.Atoi(s)
// 			if err != nil {
// 				fmt.Println("Некорректный ввод: значения должны быть числами")
// 				return "NO"
// 			}
// 			sortArray[i] = sortirated
// 		}
// 	}
// 	rows, cols := numArray[0], numArray[1]
// 	initRow, _ := strconv.Atoi(initCoordArray[0])
// 	initCol, _ := strconv.Atoi(initCoordArray[1])
// 	finRow, _ := strconv.Atoi(finCoordArray[0])
// 	finCol, _ := strconv.Atoi(finCoordArray[1])

// 	if initRow < 0 || initRow >= rows || initCol < 0 || initCol >= cols {
// 		fmt.Println("Начальная точка выходит за пределы лабиринта")
// 		return "NO"
// 	}
// 	if finRow < 0 || finRow >= rows || finCol < 0 || finCol >= cols {
// 		fmt.Println("Конечная точка выходит за пределы лабиринта")
// 		return "NO"
// 	}

//		return "OK"
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	row, col int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Считываем размеры лабиринта
	line, _ := in.ReadString('\n')
	matrix := strings.Fields(line)
	numArray := make([]int, len(matrix))
	for i, s := range matrix {
		num, _ := strconv.Atoi(s)
		numArray[i] = num
	}

	// Считываем сам лабиринт
	array := make([]string, numArray[0])
	for i := 0; i < numArray[0]; i++ {
		meanings, _ := in.ReadString('\n')
		meanings = strings.TrimSpace(meanings)
		array[i] = meanings
	}

	// Считываем начальную и конечную точки
	initCoord, _ := in.ReadString('\n')
	finCoord, _ := in.ReadString('\n')

	// Валидация входных данных
	result := Validation(numArray, initCoord, finCoord, array)

	if result == "OK" {
		// Поиск пути
		path := SearchPass(numArray, initCoord, finCoord, array)
		if path == "" {
			fmt.Fprintln(out, "NO PATH")
		} else {
			fmt.Fprintln(out, path)
		}
	}
}

// SearchPass выполняет поиск минимального пути (алгоритм BFS)
func SearchPass(numArray []int, initCoord, finCoord string, array []string) string {
	rows, cols := numArray[0], numArray[1]
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, cols)
		line := strings.Fields(array[i])
		for j, val := range line {
			grid[i][j], _ = strconv.Atoi(val)
		}
	}

	// Конвертация координат
	init := strings.Fields(initCoord)
	fin := strings.Fields(finCoord)
	initRow, _ := strconv.Atoi(init[0])
	initCol, _ := strconv.Atoi(init[1])
	finRow, _ := strconv.Atoi(fin[0])
	finCol, _ := strconv.Atoi(fin[1])

	// Проверка стартовой и финишной точки
	if grid[initRow][initCol] == 0 || grid[finRow][finCol] == 0 {
		return ""
	}

	// BFS для поиска пути
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	queue := []struct {
		row, col int
		path     []string
	}{
		{
			row:  initRow,
			col:  initCol,
			path: []string{fmt.Sprintf("%d %d", initRow, initCol)},
		},
	}
	visited[initRow][initCol] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.row == finRow && current.col == finCol {
			return strings.Join(current.path, "\n") + "\n."
		}

		for _, d := range directions {
			newRow, newCol := current.row+d.row, current.col+d.col
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols &&
				grid[newRow][newCol] != 0 && !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				newPath := append([]string{}, current.path...)
				newPath = append(newPath, fmt.Sprintf("%d %d", newRow, newCol))
				queue = append(queue, struct {
					row, col int
					path     []string
				}{
					row:  newRow,
					col:  newCol,
					path: newPath,
				})
			}
		}
	}

	return "" // Путь не найден
}

// Validation проверяет корректность входных данных
func Validation(numArray []int, initCoord, finCoord string, array []string) string {
	initCoordArray := strings.Fields(initCoord)
	finCoordArray := strings.Fields(finCoord)

	if len(numArray) != 2 {
		return "NO"
	}
	if len(initCoordArray) != 2 || len(finCoordArray) != 2 {
		fmt.Println("введены неправильные конечные точки")
		return "NO"
	}

	rows, cols := numArray[0], numArray[1]
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		line := strings.Fields(array[i])
		if len(line) != cols {
			fmt.Println("Неверный формат строки лабиринта")
			return "NO"
		}
		grid[i] = make([]int, cols)
		for j, s := range line {
			val, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Некорректное значение в лабиринте")
				return "NO"
			}
			grid[i][j] = val
		}
	}

	initRow, _ := strconv.Atoi(initCoordArray[0])
	initCol, _ := strconv.Atoi(initCoordArray[1])
	finRow, _ := strconv.Atoi(finCoordArray[0])
	finCol, _ := strconv.Atoi(finCoordArray[1])

	if initRow < 0 || initRow >= rows || initCol < 0 || initCol >= cols {
		fmt.Println("Начальная точка выходит за пределы лабиринта")
		return "NO"
	}
	if finRow < 0 || finRow >= rows || finCol < 0 || finCol >= cols {
		fmt.Println("Конечная точка выходит за пределы лабиринта")
		return "NO"
	}
	if grid[initRow][initCol] == 0 {
		fmt.Println("Начальная точка не может быть стеной")
		return "NO"
	}
	if grid[finRow][finCol] == 0 {
		fmt.Println("Конечная точка не может быть стеной")
		return "NO"
	}

	return "OK"
}
