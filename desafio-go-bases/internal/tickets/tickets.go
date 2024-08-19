package tickets

import (

	//"io"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	Id     string
	Nombre string
	Email  string
	Pais   string
	Hora   string
	Precio float64
}

// GetTickets lee el archivo CSV y retorna un slice de Tickets
// Para utilizar en el resto de las funciones, ya que se repite y es mas facil para realizar comparaciones
func GetTickets() ([]Ticket, error) {

	file, err := os.Open("./tickets.csv")
	if err != nil {
		//fmt.Println(err)
		panic(fmt.Sprintf("Error fatal: No se puede abrir el archivo de tickets: %v", err))
		//return nil, fmt.Errorf("error al abrir el archivo: %v", err)
	}
	defer file.Close()

	result, err := os.ReadFile("./tickets.csv")
	if err != nil {
		//fmt.Println(err)
		return nil, fmt.Errorf("error al abrir el archivo: %v", err)

	}

	var tickets2 []Ticket
	// Dividir los datos por líneas
	lines := strings.Split(string(result), "\n")
	//Itero sobre cada línea y creo un ticket
	for _, line := range lines {
		if line == "" {
			continue // Saltar líneas vacías
		}

		record := strings.Split(line, ",")
		value, err := strconv.ParseFloat(record[5], 64)
		if err != nil {
			continue // Ignorar líneas con precio inválido

		}

		ticket := Ticket{
			Id:     record[0],
			Nombre: record[1],
			Email:  record[2],
			Pais:   record[3],
			Hora:   record[4],
			Precio: value,
		}
		tickets2 = append(tickets2, ticket)

	}
	return tickets2, nil

}

// //////////////////////////////////////////////////////
// ////////////// Requerimiento 1 ///////////////////////
// //////////////////////////////////////////////////////
func GetTotalTickets(destination string) (int, error) {

	if destination == "" {
		return 0, fmt.Errorf("El parámetro 'destination' no puede estar vacío")
	}

	tickets, err := GetTickets()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener tickets: %v", err)
	}

	total := 0
	for _, tickets := range tickets {
		if tickets.Pais == destination {
			total++
		}
	}

	if total == 0 {
		return 0, fmt.Errorf("No se encontraron tickets para el pais: %s", destination)
	} else {
		//fmt.Println(total)
		return total, nil
	}
}

////////////////////////////////////////////////////////
//////////////// Requerimiento 2 obteniendo todos los periodos, para pruebas///////////////////////
////////////////////////////////////////////////////////

func GetPeriod(time string) (int, error) {
	//extraigo la hora para poder comparar el periodo
	hora, err := strconv.Atoi(strings.Split(time, ":")[0])
	if err != nil {
		return 0, fmt.Errorf("Error parsing hora: %v", err)
	}

	switch {
	case hora >= 0 && hora < 6:
		return 0, nil //madrugada
	case hora >= 6 && hora < 12:
		return 1, nil //mañana
	case hora >= 12 && hora < 20:
		return 2, nil //tarde
	case hora >= 20 && hora <= 23:
		return 3, nil //noche
	default:
		return 4, fmt.Errorf("Hora invalida: %d", hora)
	}

}

func GetCountAllPeriod() (string, error) {

	tickets, err := GetTickets()
	if err != nil {
		return "", fmt.Errorf("Error al obtener tickets: %v", err)
	}

	var madrugada, manana, tarde, noche int

	for _, ticket := range tickets {

		period, err := GetPeriod(ticket.Hora)
		if err != nil {
			return "", fmt.Errorf("Error al obtener periodo: %v", err)
		}

		switch period {
		case 0:
			madrugada++
		case 1:
			manana++
		case 2:
			tarde++
		case 3:
			noche++
		}
	}

	periods := fmt.Sprintf("Pasajeros que viajan en la madrugada: %d\nPasajeros que viajan en la mañana: %d\nPasajeros que viajan en la tarde: %d\nPasajeros que viajan en la noche: %d",
		madrugada, manana, tarde, noche)

	return periods, nil
}

////////////////////////////////////////////////////////
//////////////// Requerimiento 2 pasandole el periodo es el que se usa en el main///////////////////////
////////////////////////////////////////////////////////

func GetPeriod2(time string) (string, error) {

	//extraigo la hora para poder comparar el periodo
	hora, err := strconv.Atoi(strings.Split(time, ":")[0])
	if err != nil {
		return "", fmt.Errorf("Error parsing hora: %v", err)
	}

	switch {
	case hora >= 0 && hora < 6:
		return "madrugada", nil //madrugada
	case hora >= 6 && hora < 12:
		return "mañana", nil //mañana
	case hora >= 12 && hora < 20:
		return "tarde", nil //tarde
	case hora >= 20 && hora <= 23:
		return "noche", nil //noche
	default:
		return "", fmt.Errorf("Hora invalida: %d", hora)
	}

}

func GetCountByPeriod(periodo string) (int, error) {

	if periodo == "" {
		return 0, fmt.Errorf("El parámetro 'periodo' no puede estar vacío")
	}

	//lo paso a minuscula para evitar error en la comparaciones
	periodo = strings.ToLower(periodo)

	// Valido que el periodo sea válido
	switch periodo {
	case "madrugada", "mañana", "tarde", "noche":

	default:
		return 0, fmt.Errorf("Periodo %s no válido. Debe ser 'madrugada', 'mañana', 'tarde' o 'noche'", periodo)
	}

	tickets, err := GetTickets()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener tickets: %v", err)
	}

	var contador int

	for _, ticket := range tickets {

		period, err := GetPeriod2(ticket.Hora)

		if err != nil {
			fmt.Println("Error al procesar el registro:", err)
			continue
		}
		if periodo == period {
			contador++
		}

	}

	return contador, nil

}

/* ejemplo 3
func AverageDestination(destination string, total int) (int, error) {}
*/
////////////////////////////////////////////////////////
//////////////// Requerimiento 3 ///////////////////////
////////////////////////////////////////////////////////
//cuenta el número total de viajeros, era mas facil hacer el len(tickets) en AverageDestination pero la letra lo pedia asi
func TotalTravelers() (int, error) {
	tickets, err := GetTickets()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener tickets: %v", err)
	}

	var total int

	total = len(tickets)
	//Devuelvo el número total de tickets
	return total, nil
}

func AverageDestination(destination string, total int) (float64, error) {

	if destination == "" {
		return 0, fmt.Errorf("El parámetro 'destination' no puede estar vacío")
	}

	tickets, err := GetTickets()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener tickets: %v", err)
	}

	var totalPasajeros int

	for _, ticket := range tickets {

		if ticket.Pais == destination {
			totalPasajeros++
		}
	}

	if totalPasajeros == 0 {
		return 0, fmt.Errorf("No se encontraron tickets para el país %s", destination)
	}

	average := float64(totalPasajeros) / float64(total) * 100

	return average, nil

}
