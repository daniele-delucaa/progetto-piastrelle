De Luca Daniele 01576A
# *Relazione progetto "Piastrelle Digitali"*

## Modellazione e implementazione
Per affrontare il problema sono state implementate delle apposite strutture dati e algoritmi. 

### Implementazione del piano
La rappresentazione del **piano** doveva includere l'insieme delle **piastrelle** e delle **regole**. E' stato necessario quindi memorizzare tutte le piastrelle, per fare ciò si è utilizzata una mappa che ha come chiave una singola piastrella e come valore associato un tipo **colore**. Le piastrelle sono salvate nella mappa tramite le coordinate (x, y) nel piano. Ad ogni piastrella sono associati i dati relativi al **colore** e all'**intensità** di essa grazie al tipo colore, che presenta questi dati come campi.

```Go
type piano struct {
	piastrelle map[piastrella]colore
	regole     *[]regola_
}
type piastrella struct {
	x, y int
}

type colore struct {
	coloree   string
	intensita int
}
```

Le **regole** sono memorizzate tramite un puntatore all'indirizzo di una slice di regole. L' utilizzo del puntatore è utile per gestire dinamicamente la struttura dati, senza dover restituire la slice modificata dalle funzioni che operano su di essa. 

```Go
type regola_ struct {
	addendi   []colore
	risultato string
	consumo   int
}
```

Il tipo regola_ ha tre campi:
- **addendi**, rappresentati da una slice di colori, contengono i colori e le intensità associate di una regola
- **risultato**, cioè il colore finale se la regola di propagazione può essere utilizzata, rappresentato tramite una stringa
- **consumo**, un campo intero, che viene incrementato all'utilizzo della una regola di propagazione

## Implementazione funzioni principali

### Colora
```Go
func colora(p piano, x int, y int, alpha string, i int) {
	p.piastrelle[piastrella{x, y}] = colore{alpha, i}
}
```
La funziona colora ha come parametri il **piano**, le **coordinate** di una piastrella, il **colore** sotto forma di stringa, e l'**intensità**. La piastrella in input viene colorata grazie all'utilizzo della **mappa** che contiene le piastrelle nel piano, a questa vengono passate le coordinate e il colore e l'intensità desiderata.  
- **Complessità temporale**: l'accesso a una mappa ha costo **O(1)**
- **Complessità spaziale**: non viene allocato alcuno spazio, quindi abbiamo costo costante di **O(1)**

### Spegni
```Go
func spegni(p piano, x int, y int) {
	delete(p.piastrelle, piastrella{x, y})
}
```
Spegni ha come parametri il **piano** e le **coordinate** di una piastrella. Questa funzione permette di spegnere la piastrella passata in input tramite le coordinate, nel piano. Ciò viene fatto sotto forma di cancellazione della chiave nella mappa che contiene tutte le piastrelle. 
- **Complessità temporale**: l'operazione di **delete** ha tempo costante **O(1)**
- **Complessità spaziale**: non viene allocato alcuno spazio, quindi abbiamo costo costante di **O(1)**

### Regola
```Go
func regola(p piano, r string) {
	arr := strings.Split(r, " ")
	var nuovaRegola regola_
	nuovaRegola.risultato = arr[1]
	var addendo colore

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
```
La funzione regola permette di aggiungere una regola nel piano, ha come parametri il **piano** e una **stringa**, che contiene i dati della regola da aggiungere. La stringa ha questa forma: 
```
β k1 α1 k2 α2 · · · kn αn
```
Dove β è il **colore risultato** della regola, ki αi sono gli **addendi** della regola. Viene effettuata una **Split** sulla stringa, i suoi dati vengono salvati su una slice di stringhe, ignorando gli spazi. 
Viene effettuato **parsing** della slice, e grazie a questa operazione viene creata e aggiunta la regola nella slice di regole nel **piano**.
- **Complessità temporale**: la **Split** ha complessità **O(n)**, dove **n = numero caratteri stringa**. Abbiamo un **ciclo for** che itera sulla **slice**, che ha **k elementi**. Inoltre le restanti operazioni (assegnamenti di variabili e confronti hanno complessità costante **O(1)**). Nè risulta una complessità di **O(n) + O(k) = O(n)**, poichè **k <= n**.
- **Complessità spaziale**: Abbiamo due **variabili** che occupano spazio costante **O(1)**. La slice **addendi** cresce nell'ordine di **O(n) dove n è al max 8** dato che una regola non ha mai più di 8 addendi. L' aggiunta della regola occupa spazio costante **O(1)**. La **Split** ha complessità pari a **O(n)**.

### Stato
```Go
func stato(p piano, x int, y int) (string, int) {
	var piast colore
	var ok bool
	if piast, ok = p.piastrelle[piastrella{x, y}]; ok {
		fmt.Println(piast.coloree, piast.intensita)
	}
	return piast.coloree, piast.intensita
}
```
Stato prende come parametro il **piano** e le **coordinate** di una piastrella, **restituisce** e **stampa** il **colore** e l'**intensità** della piastrella in input se questa è accesa nel piano, altrimenti non verrà stampato nulla. Stato funziona correttamente grazie a un controllo di esistenza di una chiave in una mappa, in questo caso, passo alla mappa che contiene le piastrelle, le coordinate date come parametro in input e appunto nè controllo l'esistenza.
- **Complessità temporale**: assumiamo che i confronti, gli assegnamenti di variabili e la restituzione di valori abbiano tempo costante, quindi la complessità è **O(1)**
- **Complessità spaziale**: lo spazio utilizzato in questa funzione è nell'ordine di **O(1)**

### Stampa
```Go
func stampa(p piano) {
	if len(*(p).regole) > 0 {
		fmt.Println("(")
		for _, rule := range *p.regole {
			fmt.Print(rule.risultato, ":")
			for i := 0; i < len(rule.addendi); i++ {
				fmt.Print(" ", rule.addendi[i].intensita, " ", rule.addendi[i].coloree)
			}
			fmt.Println()
		}
		fmt.Println(")")
	}
}
```
Questa funzione permette di **stampare** tutte le **regole** nel piano, nell'ordine in cui si trovano. La stampa viene effettuata nel seguente formato:
```
(
r1
r2
..
.
rm
)
```
La funzione prima controlla se sono presenti delle regole da stampare, in caso affermativo, viene iterata la slice di regole, e per ognuna stampa i diversi addendi, ciò viene fatto con due **cicli for** annidati. 
- **Complessità temporale**: il primo ciclo ha complessità **O(n)** dove **n = numero regole nel piano**, il for interno effettua sempre al massimo **8 iterazioni** (perchè una regola ha al max 8 addendi), quindi ha complessità **O(k)** dove **k = 8**. 
Abbiamo quindi **O(n) * O(k) = O(n * k)**. 
- **Complessità spaziale**: assumiamo che i confronti, gli assegnamenti e le operazioni di stampa abbiano complessità costante di **O(1)**

### Blocco