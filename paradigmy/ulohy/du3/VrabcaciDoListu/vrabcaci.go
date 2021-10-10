// go get github.com/gonutz/prototype/...
// zmenit riadok 502 na: 			int32(style & ^w32.WS_OVERLAPPEDWINDOW),

package main

import (
	"github.com/gonutz/prototype/draw"
	"math"
	"math/rand"
	"sort"
	"strconv"
)
var windowWidth, windowHeight = //1024, 800
 								  1920, 1080

// pocet vrabcakov
var N = 500

type Vrabcak struct {
	x     int; y int	// suradnice vracaka
	dx    int; dy int	// smer letu vrabcaka
	size  int
	color draw.Color
}

func main() {
	vrabcaci := make([]*Vrabcak, 0)
	for i := 0; i < N; i++ {
		size := 3+rand.Intn(3)  // 3,4,5
		vrabcaci = append(vrabcaci,
						&Vrabcak{  // new Vrabcak
							x:    rand.Intn(windowWidth),
							y:    rand.Intn(windowHeight),
							dx:   rand.Intn(size),  // vacsi lieta rychlejsie
							dy:   rand.Intn(size),
							size: size,
							color:  draw.Color{
								(float32)(30+rand.Intn(200)),
								(float32)(30+rand.Intn(200)),
								(float32)(30+rand.Intn(200)),
								1},
						},
			)
	}
	loop := 0
	err := draw.RunWindow("Vrabčáci", windowWidth, windowHeight, func(window draw.Window) {
		// na ovladanie simulacie pouzivam numericku klavesnicu, ak take klavesy nemate, tak si to preprogramujte
		if window.WasKeyPressed(draw.KeyNumAdd) {  // Num+ vacsi pocet susedov, na ktorych reaguje
			kolkoNajblizsich++
		}
		if window.WasKeyPressed(draw.KeyNumSubtract) { // Num- mensi pocet susedov
			kolkoNajblizsich--
			if (kolkoNajblizsich <= 0) {
				kolkoNajblizsich = 1
			}
		}
		if window.WasKeyPressed(draw.KeyNumMultiply) { // Num* zvacsenie vzdialenosti
			horizont+=10
		}
		if window.WasKeyPressed(draw.KeyNumDivide) { // Num/ zmensenie vzdialenosti
			horizont-=10
			if (horizont <= 0) {
				horizont = 10
			}
		}
		window.FillRect(0, 0, windowWidth, windowHeight, draw.LightYellow)  // background
		loop++
		for _, vrabcak := range vrabcaci {
			okolitiVrabcaci := vrabcak.najblizsi(vrabcaci)
			//chaos(vrabcak, okolitiVrabcaci)
			separate(vrabcak, okolitiVrabcaci)
			align(vrabcak, okolitiVrabcaci)
			cohere(vrabcak, okolitiVrabcaci)
			wrapWindow(vrabcak)
		}
		kresli(vrabcaci, window)
		window.DrawScaledText("Loop:" + strconv.Itoa(loop) +
			", najblizsich:" +  strconv.Itoa(kolkoNajblizsich) +
			", horizont: " +  strconv.Itoa(horizont),
			10, 10, 1.0, draw.Black)
	})
	if err != nil {
		panic(err)
	}
}

var kolkoNajblizsich = 10
func (v *Vrabcak) najblizsi(vrabcaci []*Vrabcak) []Vrabcak {
	najblizsi := make([]Vrabcak, len(vrabcaci))
	for _, vrabcak := range vrabcaci {
		najblizsi = append(najblizsi, *vrabcak)
	}
	sort.SliceStable(najblizsi, func(i, j int) bool {
		return v.distance(najblizsi[i]) < v.distance(najblizsi[j])
	})
	return najblizsi[0:kolkoNajblizsich]
}

func (v *Vrabcak) distance(n Vrabcak) float64 {
	return math.Sqrt(float64((v.x - n.x)*(v.x - n.x) + (v.y - n.y)*(v.y - n.y)))
}

func wrapWindow(vrabcak *Vrabcak) {
	if vrabcak.x < 0 {
		vrabcak.x += windowWidth
	} else if vrabcak.x > windowWidth {
		vrabcak.x -= windowWidth
		vrabcak.x *=-1
	}
	if vrabcak.y < 0 {
		vrabcak.y += windowHeight
	} else if vrabcak.y > windowHeight {
		vrabcak.y -= windowHeight
		vrabcak.y *= -1
	}
}

func chaos(vrabcak *Vrabcak, najblizsi []Vrabcak) {
	x, y := 0, 0
	for _, v := range najblizsi {
			x += v.x-vrabcak.x
			y += v.y-vrabcak.y
	}
	vrabcak.x += rand.Intn(2*vrabcak.size+1) - vrabcak.size
	vrabcak.y += rand.Intn(2*vrabcak.size+1) - vrabcak.size
}

/**
  uvazuje len vrabcakov blizsich ako horizont od vrabcaka
  pohne sa smerom opacnym, ako je sucet vektorov k tymto vrabciakom
  rychlostou umernou tomuto suctu vektorov
 */
var horizont = 16
func separate(vrabcak *Vrabcak, najblizsi []Vrabcak) {
}

/**
  smeruje k priemernemu susedovi
  urobi priemer vektorov rychlosti najblizsich vrabciakov
  o tolko sa zmeni smerovy vektor aj poloha vrabciaka
 */
func align(v *Vrabcak, najblizsi []Vrabcak) {
}

func cohere(v *Vrabcak, najblizsi []Vrabcak) {
}

func kresli(vrabcaci []*Vrabcak, window draw.Window) {
	for _, vrabcak := range vrabcaci {
		window.FillEllipse(int(vrabcak.x-vrabcak.size), int(vrabcak.y-vrabcak.size), int(2*vrabcak.size), int(2*vrabcak.size), vrabcak.color)
		window.DrawLine(int(vrabcak.x-vrabcak.dx), int(vrabcak.y-vrabcak.dy), int(vrabcak.x), int(vrabcak.y), vrabcak.color)
	}
}
