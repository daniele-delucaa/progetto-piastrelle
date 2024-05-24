package main

import (
	"fmt"
	"strconv"
	"strings"
)

type piastrella struct {
	x, y int
}

type colore struct {
	coloree   string
	intensita int
}

type piano struct {
	piastrelle map[piastrella]colore
	regole     *[]regola_
}

type regola_ struct {
	addendi   []colore
	risultato string
	consumo   int
}

func main() {

}

func esegui(p piano, s string) {
	arr := strings.Split(s, " ")
	switch arr[0] {
	case "C":
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		intensity, _ := strconv.Atoi(arr[4])
		colora(p, cX, cY, arr[3], intensity)
	case "S":
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		spegni(p, cX, cY)
	case "?":
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		stato(p, cX, cY)
	case "b":
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		blocco(p, cX, cY)
	case "B":
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		bloccoOmog(p, cX, cY)
	case "p":
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		propaga(p, cX, cY)
	case "P":
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		propagaBlocco(p, cX, cY)
	}
}

/*func gestioneInput(arr []string) (int, int, string, int) {
	if len(arr) == 1 {
		return 0, 0, "", 0
	}
	if len(arr) < 4 {
		cX, _ := strconv.Atoi(arr[1])
		cY, _ := strconv.Atoi(arr[2])
		return cX, cY, "", 0
	}
	cX, _ := strconv.Atoi(arr[1])
	cY, _ := strconv.Atoi(arr[2])
	intensity, _ := strconv.Atoi(arr[4])
	return cX, cY, arr[3], intensity
}*/

func colora(p piano, x int, y int, alpha string, i int) {
	var piast piastrella = piastrella{x, y}
	p.piastrelle[piast] = colore{alpha, i}
}

func spegni(p piano, x int, y int) {
	var piast piastrella = piastrella{x, y}
	delete(p.piastrelle, piast)
}

func regola(p piano, r string) {
	arr := strings.Split(r, " ")
	var nuovaRegola regola_
	nuovaRegola.risultato = arr[0]
}

func stato(p piano, x int, y int) (string, int) {
	piast, ok := p.piastrelle[piastrella{x, y}]
	if ok {
		fmt.Println(piast.intensita, piast.intensita)
	}
	return piast.coloree, piast.intensita
}
