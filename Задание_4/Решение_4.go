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

type Line struct {
	point1, point2 int
}

func (line *Line) Normalize() {
	if line.point2 > line.point1 {
		line.point2, line.point1 = line.point1, line.point2
	}
}

// C проверяет, выполняется ли условие a < b < c.
func C(a, b, c int) bool {
	if a < b && b < c {
		return true
	}
	return false
}

// Intersection пересечение двух линий.
func Intersection(line1, line2 Line) bool {
	line1.Normalize()
	line2.Normalize()
	p11 := line1.point1
	p12 := line1.point2
	p21 := line2.point1
	p22 := line2.point2
	if (C(p11, p21, p12) && C(p21, p12, p22)) || (C(p11, p22, p12) && C(p21, p11, p22)) {
		return true
	}
	return false
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

// AngSorted приводит точки угола в правильный формат.
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
	length := Line{angle.point1, angle.point3}
	radian := float64(length.Steps()) * math.Pi / float64(numberOfSides)
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
	//fmt.Println(Length(line1.Steps()), Length(line2.Steps()), math.Sin(angle.AngRad()))
	return (0.5) * Length(line1.Steps()) * Length(line2.Steps()) * math.Sin(angle.AngRad())
}

func IntersectionOfTriangles(tr1, tr2 Triangle) bool {
	tr1M := []int{tr1.point1, tr1.point2, tr1.point3}
	tr2M := []int{tr2.point1, tr2.point2, tr2.point3}
	for n, i := range tr1M {
		var i2 int
		if n == 2 {
			i2 = tr1M[0]
		} else {
			i2 = tr1M[n+1]
		}
		for m, j := range tr2M {
			var j2 int
			if m == 2 {
				j2 = tr2M[0]
			} else {
				j2 = tr1M[m+1]
			}
			if Intersection(Line{i, i2}, Line{j, j2}) || i == j {
				return true
			}
		}
	}
	return false
}

// numberOfSides — количество сторон правильного n-угольника.
var numberOfSides int

func main() {
	_, err := fmt.Scan(&numberOfSides)
	CheckError(err)
}
