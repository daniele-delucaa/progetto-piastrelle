package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	var rules *[]regola_
	//var punt *[]regola_ = &rules
	m := make(map[piastrella]colore)
	p := piano{m, rules}

	for scanner.Scan() {
		l := scanner.Text()
		esegui(p, l)
	}
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
		/*case "B":
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
			propagaBlocco(p, cX, cY)*/
	case "r":
		regola(p, s)
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
	nuovaRegola.risultato = arr[1]
	addendo := colore{}

	for i := 2; i < len(arr); i++ {
		if i%2 == 0 {
			addendo.intensita, _ = strconv.Atoi(arr[i])
		} else {
			addendo.coloree = arr[i]
			nuovaRegola.addendi = append(nuovaRegola.addendi, addendo)
		}
	}
	*p.regole = append(*p.regole, nuovaRegola)
}

/*func stato(p piano, x int, y int) (string, int) {
	piast, ok := p.piastrelle[piastrella{x, y}]
	if ok {
		fmt.Println(piast.coloree, piast.intensita)
	}
	return piast.coloree, piast.intensita
}*/

func stato(p piano, x int, y int) (string, int) {
	var piast colore
	var ok bool
	if piast, ok = p.piastrelle[piastrella{x, y}]; ok {
		fmt.Println(piast.coloree, piast.intensita)
	}
	return piast.coloree, piast.intensita
}

func stampa(p piano) {
	if len(*(p).regole) > 0 {
		fmt.Println("(")
		for _, rule := range *p.regole {
			fmt.Print(rule.risultato, ": ")
			for i := 0; i < len(rule.addendi); i++ {
				fmt.Print(rule.addendi[i].intensita, " ", rule.addendi[i].coloree)
			}
			fmt.Println()
		}
		fmt.Println(")")
	}
}

func blocco(p piano, x, y int) {
	var inizio colore
	var ok bool
	var intensitaTotale int
	if inizio, ok = p.piastrelle[piastrella{x, y}]; !ok {
		fmt.Println(intensitaTotale)
		return
	}

	intensitaTotale += inizio.intensita
	visitati := make(map[piastrella]bool)

	coda := queue{}
	coda.Enqueue(piastrella{x, y})

	visitati[piastrella{x, y}] = true
	for coda.Len() != 0 {
		piast, _ := coda.Dequeue()

		adiacenti := cercaAdiacenti(p, piast)

		for i := 0; i < len(adiacenti); i++ {
			if _, ok := visitati[adiacenti[i]]; !ok {
				val := p.piastrelle[adiacenti[i]]
				intensitaTotale += val.intensita
				visitati[adiacenti[i]] = true
				coda.Enqueue(adiacenti[i])
			}
		}
	}

}

func cercaAdiacenti(p piano, piast piastrella) []piastrella {
	var circonvicine []piastrella

	// combinazioni di coordinate possibili per una piastrella adiacente a quella in input
	arrX := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	arrY := []int{1, 1, 1, 0, -1, -1, -1, 0}

	for i := 0; i < len(arrX); i++ {
		if _, ok := p.piastrelle[piastrella{piast.x + arrX[i], piast.y + arrY[i]}]; ok {
			circonvicine = append(circonvicine, piastrella{piast.x + arrX[i], piast.y + arrY[i]})
		}
	}
	return circonvicine
}
