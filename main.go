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
	var rules []regola_
	var punt *[]regola_ = &rules
	m := make(map[piastrella]colore)
	p := piano{m, punt}

	for scanner.Scan() {
		l := scanner.Text()
		//fmt.Println(l)
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
		val, _ := blocco(p, cX, cY)
		fmt.Println(val)
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
	case "r":
		regola(p, s)
	case "s":
		stampa(p)
	case "q":
		os.Exit(0)
	}
}

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
				fmt.Print(rule.addendi[i].intensita, " ", rule.addendi[i].coloree, " ")
			}
			fmt.Println()
		}
		fmt.Println(")")
	}
}

// serve restituire la slice di piastrelle (slice che contiene le piastrelle di un blocco) per poi usarla per propagaBlocco
func blocco(p piano, x, y int) (int, []piastrella) {
	var inizio colore
	var ok bool
	var intensitaTotale int
	var sliceBlocco []piastrella
	if inizio, ok = p.piastrelle[piastrella{x, y}]; !ok {
		return 0, nil
	}

	intensitaTotale += inizio.intensita
	visitati := make(map[piastrella]bool)

	sliceBlocco = append(sliceBlocco, piastrella{x, y})
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
				sliceBlocco = append(sliceBlocco, adiacenti[i])
				visitati[adiacenti[i]] = true
				coda.Enqueue(adiacenti[i])
			}
		}
	}
	return intensitaTotale, sliceBlocco
}

func bloccoOmog(p piano, x, y int) {
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
				if val.coloree == inizio.coloree {
					intensitaTotale += val.intensita
					visitati[adiacenti[i]] = true
					coda.Enqueue(adiacenti[i])
				}
			}
		}
	}
	fmt.Println(intensitaTotale)
}

/*func cercaAdiacenti(p piano, piast piastrella) []piastrella {
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
}*/

func cercaAdiacenti(p piano, piast piastrella) []piastrella {
	var circonvicine []piastrella

	// genera combinazioni di coordinate possibili per la piastrella adiacente a quella in input
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if _, ok := p.piastrelle[piastrella{piast.x + i, piast.y + j}]; ok {
				circonvicine = append(circonvicine, piastrella{piast.x + i, piast.y + j})
			}
		}
	}
	return circonvicine
}

func propaga(p piano, x, y int) {
	colori := propagaGenerico(p, x, y)
	coloraPiastrelle(p, colori)
}

// serve mappa per contare quantita di piastrelle con un certo colore?
func propagaGenerico(p piano, x, y int) map[piastrella]regola_ {
	quantitaColori := make(map[string]int) // mappa che conta i colori delle piastrelle adiacenti a quella in input
	coloriRisultati := make(map[piastrella]regola_)
	//piast := p.piastrelle[piastrella{x, y}]
	var flag bool
	adiacenti := cercaAdiacenti(p, piastrella{x, y})

	for _, piastSingola := range adiacenti {
		val := p.piastrelle[piastSingola]
		col := val.coloree
		quantitaColori[col]++
	}
	for _, rule := range *p.regole {
		for _, str := range rule.addendi {
			arr := strings.Split(str.coloree, " ")
			// v = quantita sulla regola
			v := str.intensita
			// c = quantita sulla mappa
			if c, ok := quantitaColori[arr[0]]; ok && c >= v {
				flag = true
			} else {
				flag = false
				break
			}
		}
		if flag == true {
			coloriRisultati[piastrella{x, y}] = rule
			rule.consumo++
			break
		}
	}
	return coloriRisultati
}

func propagaBlocco(p piano, x, y int) {
	_, slc := blocco(p, x, y)
	coloriRisultati := make(map[piastrella]regola_)

	for i := range slc {
		coloriRisultati = propagaGenerico(p, slc[i].x, slc[i].y)
		if len(coloriRisultati) > 0 {
			coloraPiastrelle(p, coloriRisultati)
		}
	}
	// fmt.Println(slc)
}

func coloraPiastrelle(p piano, coloriRisultati map[piastrella]regola_) {
	var coloreRisultato string
	for piast, _ := range coloriRisultati {
		_, ok := p.piastrelle[piast]
		coloreRisultato = coloriRisultati[piast].risultato
		cc := p.piastrelle[piast].intensita
		if ok {
			colora(p, piast.x, piast.y, coloreRisultato, cc)
		} else {
			colora(p, piast.x, piast.y, coloreRisultato, 1)
		}
	}
}
