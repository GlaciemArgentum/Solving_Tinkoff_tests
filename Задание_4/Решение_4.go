package main

import (
	"fmt"
	"math"
)

// CheckError проверяет наличие ошибки.
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// C проверяет равенство a >= b >= c.
func C(a, b, c float64) bool {
	if a >= b && b >= c {
		return true
	}
	return false
}

type Line struct {
	point1, point2 int
}

// Steps считает кол-во сторон между точками.
func (line *Line) Steps() int {
	steps := math.Abs(float64(line.point1 - line.point2))
	return int(math.Min(steps, float64(numberOfSides)-steps))
}

// Length считает длину отрезка по кол-ву отрезков.
func Length(steps int) float64 {
	if steps == 0 {
		return 0
	}
	if steps == 1 {
		return 1
	}
	a := Angle{1, steps, steps + 1}
	return ThCos(1, Length(steps-1), a.AngRad())
}

type Angle struct {
	point1, point2, point3 int
}

// AngSorted приводит точки угла в правильный формат.
func (angle *Angle) AngSorted() {
	if !((angle.point1 < angle.point2 && angle.point2 < angle.point3) || (angle.point1 > angle.point3 && angle.point2 < angle.point3) || (angle.point1 < angle.point2 && angle.point1 > angle.point3)) {
		angle.point1, angle.point3 = angle.point3, angle.point1
	}

	step := angle.point1 - 1
	angle.point1 = 1
	angle.point2 -= step
	angle.point3 -= step
	if angle.point2 < 1 {
		angle.point2 += numberOfSides
	}
	if angle.point3 < 1 {
		angle.point3 += numberOfSides
	}
}

// AngRad считает градусы угла.
func (angle *Angle) AngRad() float64 {
	angle.AngSorted()
	if angle.point1 == angle.point3 && angle.point3 != angle.point2 {
		return 0
	}
	if angle.point1 == angle.point2 || angle.point3 == angle.point2 {
		return math.Pi
	}

	steps := Line{angle.point1, angle.point3}
	radian := float64(steps.Steps()) * math.Pi / float64(numberOfSides)
	if float64(angle.point3-1) <= float64(numberOfSides)/2 {
		radian = math.Pi - radian
	}
	return radian
}

// ThCos — теорема косинусов.
func ThCos(a, b, vab float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2) - 2*a*b*math.Cos(vab))
}

type Triangle struct {
	point1, point2, point3 int
}

// Area считает площадь треугольника.
func (triangle *Triangle) Area() float64 {
	line1 := Line{triangle.point1, triangle.point2}
	line2 := Line{triangle.point2, triangle.point3}
	angle := Angle{triangle.point1, triangle.point2, triangle.point3}
	return (0.5) * Length(line1.Steps()) * Length(line2.Steps()) * math.Sin(angle.AngRad())
}

// Pr переприсваивает значения вершин треугольника.
func (triangle *Triangle) Pr(_, b, c float64) {
	triangle.point1, triangle.point2, triangle.point3 = 1, 1+int(c*float64(numberOfSides)/math.Pi), 1+numberOfSides-int(b*float64(numberOfSides)/math.Pi)
}

// TriSorted приводит точки угла в правильный формат.
func (triangle *Triangle) TriSorted() {
	ang1 := Angle{triangle.point2, triangle.point1, triangle.point3}
	ang2 := Angle{triangle.point1, triangle.point2, triangle.point3}
	ang3 := Angle{triangle.point1, triangle.point3, triangle.point2}

	rad1 := ang1.AngRad()
	rad2 := ang2.AngRad()
	rad3 := ang3.AngRad()

	switch {
	case C(rad1, rad2, rad3):
		triangle.Pr(rad1, rad2, rad3)
	case C(rad2, rad3, rad1):
		triangle.Pr(rad2, rad3, rad1)
	case C(rad3, rad1, rad2):
		triangle.Pr(rad3, rad1, rad2)
	case C(rad1, rad3, rad2):
		triangle.Pr(rad1, rad3, rad2)
	case C(rad2, rad1, rad3):
		triangle.Pr(rad2, rad1, rad3)
	case C(rad3, rad2, rad1):
		triangle.Pr(rad3, rad2, rad1)
	}
}

// CheckTypeOfTriangle проверяет, есть ли данный треугольник в массиве треугольников.
func (triangle *Triangle) CheckTypeOfTriangle(varOfTriangles []Triangle) bool {
	for _, i := range varOfTriangles {
		if *triangle == i {
			return true
		}
	}
	return false
}

// IntersectionOfTriangles проверяет, существует ли такое расположение треугольников в многоугольнике, при котором они не пересекаются.
func IntersectionOfTriangles(indexes []int) bool {
	sum := 0
	for _, i := range indexes {
		max := varOfTriangles[i].point1
		sum += numberOfSides - max + 1
		if sum > numberOfSides {
			return false
		}
	}
	return true
}

// AreaOfTriangles считает сумму площадей массива треугольников.
func AreaOfTriangles(indexes []int) float64 {
	sum := 0.0
	for _, i := range indexes {
		sum += varOfTriangles[i].Area()
	}
	return sum
}

func MaxFromI(n int) float64 {
	maxArea := 0.0
	list := make([]int, n, n)
	for {
		if area := AreaOfTriangles(list); IntersectionOfTriangles(list) && area > maxArea {
			maxArea = area
		}
		// Тут надо сделать нормальное переприсваивание для индексов массива треугольников.
		/*for i := n - 1; i >= 0; i-- {
			if list[i] < len(varOfTriangles) {
				list[i] += 1
			} else {

			}*/
		}
	}

	return maxArea
}

// numberOfSides — количество сторон правильного n-угольника.
var numberOfSides int
var varOfTriangles []Triangle

func main() {
	_, err := fmt.Scan(&numberOfSides)
	CheckError(err)

	// Эта часть кода определяет все возможные конфигурации треугольников, которые можно вписать в данный n-угольника.
	varOfTriangles = make([]Triangle, 0, numberOfSides-3)
	for i := 1; i <= numberOfSides; i++ {
		for j := i; j <= numberOfSides; j++ {
			for k := j; k <= numberOfSides; k++ {
				if i == j || i == k || j == k {
					continue
				}
				tr := Triangle{i, j, k}
				tr.TriSorted()
				if !tr.CheckTypeOfTriangle(varOfTriangles) {
					varOfTriangles = append(varOfTriangles, tr)
				}
			}
		}
	}

	maxCountOfTriangles := numberOfSides / 3
	maxArea := 0.0
	for i := 1; i <= maxCountOfTriangles; i++ {
		max := MaxFromI(i)
		if max > maxArea {
			maxArea = max
		}
	}

	fmt.Printf("%.6f\n", maxArea)
	fmt.Printf("%.6f\n", 3*math.Sin(math.Pi/3)-3/2*math.Sin(2*math.Pi/3))
}
