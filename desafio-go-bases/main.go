package main

import (
	"desafio-go-bases/internal/tickets"
	"fmt"
	"sync"
)

func main() {
	// Parámetros predefinidos
	country := "China"
	period := "tarde"
	destination := "Indonesia"	

	var wg sync.WaitGroup//lo uso para sincronizar la ejecución de las goroutines
	wg.Add(3)//le indico que voy a esperar 3 goroutine
	
	/////////////////// Requerimiento 1 ///////////
	go func() {
		defer wg.Done()//le indico que la goroutine ha terminado 
		
		total, err := tickets.GetTotalTickets(country)
		fmt.Println("************** Requerimiento 1 ******************")
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println("*************************************************")
			return
		}
		
		fmt.Printf("Total de tickets para %s: %d\n", country, total)
		fmt.Println("*************************************************")
	}()

	/////////////////// Requerimiento 2 ///////////	
	go func() {
		defer wg.Done()
		
		result, err := tickets.GetCountByPeriod(period)
		fmt.Println("*************** Requerimiento 2 ******************")
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println("*************************************************")
			return
		}
		
		fmt.Printf("Pasajeros que viajan en la %s: %d\n", period, result)
		fmt.Println("*************************************************")
	}()

	/////////////////// Requerimiento 3 ///////////	
	go func() {
		defer wg.Done()
		totalTravelers, err := tickets.TotalTravelers()
		
		if err != nil {
			fmt.Printf("Error al obtener total de viajeros: %v\n", err)
			return
		}
		
		destPercentage, err := tickets.AverageDestination(destination, totalTravelers)
		fmt.Println("*************** Requerimiento 3 ******************")
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println("*************************************************")
			return
		}
		
		fmt.Printf("Porcentaje de tickets para el país %s: %.2f%%\n", destination, destPercentage)
		fmt.Println("*************************************************")
	}()

	wg.Wait()//espero que las goroutines terminen
	
	/*
	/////////////////////////Usando canales para la comunicación///////////////

	var waitGroup sync.WaitGroup
	resultChan := make(chan string, 3)//podria usar un canal para cada requerimiento y otra función para procesar los resultados, pero para está práctica no lo vi tan necesario
	errorChan := make(chan string, 3)

	waitGroup.Add(3)

	/////////////////// Requerimiento 1 ///////////
	go func() {
		defer waitGroup.Done()
		total, err := tickets.GetTotalTickets(country)
		if err != nil {
			errorChan <- fmt.Sprintf("Requerimiento 1: \n%v", err)
			return
		}
		resultChan <- fmt.Sprintf("Requerimiento 1: \nTotal de tickets para %s: %d", country, total)
	}()

	/////////////////// Requerimiento 2 ///////////	
	go func() {
		defer waitGroup.Done()
		result, err := tickets.GetCountByPeriod(period)
		if err != nil {
			errorChan <- fmt.Sprintf("Requerimiento 2: \n%v", err)
			return
		}
		resultChan <- fmt.Sprintf("Requerimiento 2: \nPasajeros que viajan en la %s: %d", period, result)
	}()

	/////////////////// Requerimiento 3 ///////////	
	go func() {
		defer waitGroup.Done()
		totalTravelers, err := tickets.TotalTravelers()
		if err != nil {
			errorChan <- fmt.Sprintf("Requerimiento 3: \nError al obtener total de viajeros: %v", err)
			return
		}

		destPercentage, err := tickets.AverageDestination(destination, totalTravelers)
		if err != nil {
			errorChan <- fmt.Sprintf("Requerimiento 3: \n%v", err)
			return
		}
		resultChan <- fmt.Sprintf("Requerimiento 3:\nPorcentaje de tickets para el país %s: %.2f%%", destination, destPercentage)
	}()

	go func() {
		waitGroup.Wait()
		close(resultChan)
		close(errorChan)
	}()

	// Recorro y muestro resultados
	for result := range resultChan {
		fmt.Println("*************************************************")
		fmt.Println(result)
	}

	// Recorro y muestro los errores si los hay
	for err := range errorChan {
		fmt.Println("*************************************************")
		fmt.Printf("Error: %v\n", err)
	}	*/





/*
result, err := tickets.GetCountAllPeriod()
if err != nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println("********************Obtengo todos los periodos******************")
fmt.Println(result)
fmt.Println("****************************************************************")
*/


}



