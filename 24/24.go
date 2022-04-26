/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами
x,y и конструктором.

*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	name string
	x, y float64
}

type Plane struct {
	plane  [10][10]string
	points map[string]*Point
}

func newPoint(x, y float64, name string) *Point {
	return &Point{name, x, y}
}

func newPlane() *Plane {
	return &Plane{
		plane:  fillPlane(),
		points: make(map[string]*Point)}
}

func fillPlane() [10][10]string {
	pl := [10][10]string{}
	for i := 0; i < len(pl); i++ {
		for j := 0; j < len(pl[i]); j++ {
			pl[i][j] = " ○ "
		}
	}
	return pl
}

func (pl *Plane) findDistance(p1, p2 *Point) (float64, error) {
	if i, ok := pl.points[p1.name]; ok {
		if j, ok := pl.points[p2.name]; ok {
			return math.Sqrt((i.x-j.x)*(i.x-j.x) + (i.y-j.y)*(i.y-j.y)), nil
		}
		return 0, fmt.Errorf("Point 2 don't exists\n")
	}
	return 0, fmt.Errorf("Point 1 don't exists\n")
}

func (pl *Plane) findAllDistances(p *Point) map[string]float64 {
	g := make(map[string]float64)
	for _, v := range pl.points {
		value, err := pl.findDistance(p, v)
		if err != nil {
			fmt.Println(err)
		}
		g[fmt.Sprintf("destance between %s and %s", p.name, v.name)] = value
	}
	return g
}

func (pl *Plane) addPoints(p ...*Point) error {
	if len(p) == 0 {
		return fmt.Errorf("Add at least 1 point\n")
	} else {
		for _, point := range p {
			if _, ok := pl.points[point.name]; ok {
				return fmt.Errorf("Point `%s` already exists\n", point.name)
			} else {
				if point.x <= 10 && point.y <= 10 {
					pl.plane[int(point.y-1)][int(point.x-1)] = " ● "
				}
				pl.points[point.name] = point
			}
		}
	}
	return nil
}

func (pl *Plane) visualise() {
	for i := len(pl.plane) - 1; i >= 0; i-- {
		fmt.Printf("%v\n", pl.plane[i])
	}
}

func main() {
	MyPlane := newPlane()
	A := newPoint(2, 7, "A")
	B := newPoint(7, 9, "B")
	MyPlane.addPoints(A, B)
	d, _ := MyPlane.findDistance(A, B)
	fmt.Printf("%.2f\n", d)
	MyPlane.visualise()
}
