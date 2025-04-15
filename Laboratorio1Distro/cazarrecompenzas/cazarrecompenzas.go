package main

import (
	"context"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"

	pb "cazarrecompenzas/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

// Cazarrecompensa representa un cazarrecompensas con sus atributos
type Cazarrecompensa struct {
	Nombre string
	Dinero int
}

func leerCazarrecompensasDesdeCSV(nombreArchivo string) ([]Cazarrecompensa, error) {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var cazarrecompensas []Cazarrecompensa
	for _, record := range records[1:] { // Saltar la cabecera
		cazarrecompensas = append(cazarrecompensas, Cazarrecompensa{
			Nombre: record[0],
			Dinero: 0, // Inicializar dinero en 0
		})
	}

	return cazarrecompensas, nil
}

// Función para entregar al Submundo
func entregarAlSubmundo(ctx context.Context, sub pb.SubmundoServiceClient, cazarrecompensa Cazarrecompensa, pirata *pb.Pirata, cazarrecompensas []Cazarrecompensa, x int) {
	resp, err := sub.EntregarPirataSubmundo(ctx, &pb.VentaRequest{
		IdCazarrecompensas: cazarrecompensa.Nombre,
		Pirata:             pirata,
	})
	if err != nil {
		log.Printf("Error al intentar entregar al Submundo: %v\n", err)
		return
	}

	// Verificar si resp es nil
	if resp == nil {
		log.Printf("Error: respuesta nula al intentar entregar al Submundo el pirata %s por %s.\n", pirata.Nombre, cazarrecompensa.Nombre)
		return
	}
	if resp.Mensaje == "0" {
		log.Printf("Fraude, %s no recibe nada por el pirata %s.\n", cazarrecompensa.Nombre, pirata.Nombre)
	} else {
		log.Printf("Pago del submundo a %s por %s: %s\n", cazarrecompensa.Nombre, pirata.Nombre, resp.Mensaje)
		cazarrecompensas[x].Dinero += int(pirata.Recompensa) // Aumentar dinero del cazarrecompensa
	}
}

// Función para entregar a la Marina
func entregarALaMarina(ctx context.Context, mar pb.MarinaServiceClient, sub pb.SubmundoServiceClient, gob pb.GobiernoServiceClient, cazarrecompensa Cazarrecompensa, pirata *pb.Pirata, cazarrecompensas []Cazarrecompensa, x int) {
	resp, _ := mar.EntregarPirataMarina(ctx, &pb.EntregaRequest{
		IdCazarrecompensas: cazarrecompensa.Nombre,
		Pirata:             pirata,
	})
	if resp.Mensaje == "0" {
		log.Printf("Reputación insuficiente para %s. Pirata %s no entregado a la Marina. Intentando con el Submundo...\n", cazarrecompensa.Nombre, pirata.Nombre)

		// Simular el transporte antes de intentar entregar al Submundo
		if !simularTransporte(ctx, gob, cazarrecompensa, pirata) {
			log.Printf("El pirata %s escapó durante el transporte. No se puede completar la entrega al Submundo.\n", pirata.Nombre)
			return // Salir de la función si el pirata escapa
		}

		// Intentar entregar al Submundo
		entregarAlSubmundo(ctx, sub, cazarrecompensa, pirata, cazarrecompensas, x)
	} else {
		log.Printf("Pago de la marina a %s por %s: %s\n", cazarrecompensa.Nombre, pirata.Nombre, resp.Mensaje)
		cazarrecompensas[x].Dinero += int(pirata.Recompensa) // Aumentar dinero del cazarrecompensa
	}
}

// Función para simular el transporte del pirata
func simularTransporte(ctx context.Context, gob pb.GobiernoServiceClient, cazarrecompensa Cazarrecompensa, pirata *pb.Pirata) bool {
	// Cambiar el estado del pirata a "En camino"
	gob.ReporteDeEstado(ctx, &pb.EstadoRequest{
		IdCazarrecompensas: cazarrecompensa.Nombre,
		Pirata:             pirata,
		NuevoEstado:        "En camino",
	})

	time.Sleep(1 * time.Second) // Espera
	log.Printf("Viajando.......\n")

	//Rescate de pirata
	if pirata.Recompensa > 200000000 {
		// Semilla para generar números aleatorios
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(100) + 1
		// Generar un número aleatorio del 1 al 100
		if r < 35 {
			log.Printf("El pirata %s ha sido rescatado por el Submundo.\n", pirata.Nombre)
			gob.ReporteDeEstado(ctx, &pb.EstadoRequest{
				IdCazarrecompensas: cazarrecompensa.Nombre,
				Pirata:             pirata,
				NuevoEstado:        "Perdido",
			})
			return false
		}

	}
	time.Sleep(1 * time.Second) // Espera
	log.Printf("Viajando.......\n")

	//Redada Marina
	reputacionResp, err := gob.ConsultarReputacion(ctx, &pb.ReputacionRequest{
		IdCazarrecompensas: cazarrecompensa.Nombre,
	})
	if err != nil {
		log.Printf("Error al consultar la reputación de %s: %v\n", cazarrecompensa.Nombre, err)
		return false // Salir si hay un error
	}

	// Calcular el porcentaje de entregas al Submundo del cazarecompensas
	totalEntregas := reputacionResp.EntregasMarina + reputacionResp.EntregasSubmundo
	porcentajeSubmundo_cazarrecompensas := 0.0
	if totalEntregas > 0 {
		porcentajeSubmundo_cazarrecompensas = (float64(reputacionResp.EntregasSubmundo) / float64(totalEntregas)) * 100
	}

	if porcentajeSubmundo_cazarrecompensas > 40 {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(100) + 1
		if r < 30 {
			log.Printf("El pirata %s ha sido confiscado por la Marina debido a la actividad ilegal del cazarrecompensa.\n", pirata.Nombre)
			gob.ReporteDeEstado(ctx, &pb.EstadoRequest{
				IdCazarrecompensas: cazarrecompensa.Nombre,
				Pirata:             pirata,
				NuevoEstado:        "En Marina",
			})
			return false
		}
	}

	// Calcular el porcentaje de entregas al Submundo global
	//Redada Marina (usando ObtenerEstadoGlobal)
	estadoGlobal, err := gob.ObtenerEstadoGlobal(ctx, &pb.Empty{})
	if err != nil {
		log.Printf("Error al obtener el estado global del flujo ilegal: %v\n", err)
		return false // Salir si hay un error
	}

	// Calcular el porcentaje de entregas al Submundo global
	totalEntregasGlobal := estadoGlobal.MarinaCounter + estadoGlobal.SubmundoCounter
	porcentajeSubmundoGlobal := 0.0
	if totalEntregasGlobal > 0 {
		porcentajeSubmundoGlobal = (float64(estadoGlobal.SubmundoCounter) / float64(totalEntregasGlobal)) * 100
	}

	if porcentajeSubmundoGlobal > 40 {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(100) + 1

		if pirata.Peligrosidad == "Alto" {
			if r < 50 {
				log.Printf("El pirata %s ha sido confiscado por la Marina debido al alto flujo ilegal global y la peligrosidad del pirata.\n", pirata.Nombre)
				gob.ReporteDeEstado(ctx, &pb.EstadoRequest{
					IdCazarrecompensas: cazarrecompensa.Nombre,
					Pirata:             pirata,
					NuevoEstado:        "En Marina",
				})
				return false
			}
		} else if r < 25 {
			log.Printf("El pirata %s ha sido confiscado por la Marina debido al alto flujo ilegal global.\n", pirata.Nombre)
			gob.ReporteDeEstado(ctx, &pb.EstadoRequest{
				IdCazarrecompensas: cazarrecompensa.Nombre,
				Pirata:             pirata,
				NuevoEstado:        "En Marina",
			})
			return false
		}
	}

	time.Sleep(1 * time.Second) // Espera
	log.Printf("Viajando.......\n")
	// Determinar la probabilidad de escape según la peligrosidad del pirata
	var probabilidadEscape float64
	if pirata.Recompensa < 100000000 {
		probabilidadEscape = 0.15 // 15% para baja peligrosidad
	} else if pirata.Recompensa < 200000000 {
		probabilidadEscape = 0.25 // 25% para media peligrosidad
	} else {
		probabilidadEscape = 0.45 // 45% para alta peligrosidad
	}

	// Calcular si el pirata escapa
	if rand.Float64() < probabilidadEscape {
		log.Printf("El pirata %s ha escapado\n", pirata.Nombre)
		// Cambiar el estado del pirata a "Escapado"
		gob.ReporteDeEstado(ctx, &pb.EstadoRequest{
			IdCazarrecompensas: cazarrecompensa.Nombre,
			Pirata:             pirata,
			NuevoEstado:        "Perdido",
		})
		return false // El pirata escapó
	}

	return true // El pirata no escapó
}

func main() {
	// Leer cazarrecompensas desde el archivo CSV
	cazarrecompensas, err := leerCazarrecompensasDesdeCSV("cazarrecompenzas.csv")
	if err != nil {
		log.Fatalf("Error al leer cazarrecompensas desde el archivo CSV: %v", err)
	}
	log.Printf("Cazarrecompensas cargados: %+v\n", cazarrecompensas)

	connGob, err := grpc.Dial("gobierno:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatalf("No se pudo conectar al servicio Gobierno: %v", err)
	}
	defer connGob.Close()

	connMar, _ := grpc.Dial("marina:50052", grpc.WithInsecure())
	connSub, _ := grpc.Dial("submundo:50053", grpc.WithInsecure())

	defer connMar.Close()
	defer connSub.Close()

	gob := pb.NewGobiernoServiceClient(connGob)
	mar := pb.NewMarinaServiceClient(connMar)
	sub := pb.NewSubmundoServiceClient(connSub)

	// Rutina que se ejecuta cada 6 segundos
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		// Seleccionar un cazarrecompensas aleatorio
		x := rand.Intn(len(cazarrecompensas))
		cazarrecompensa := cazarrecompensas[x]

		// Obtener la lista de piratas
		lista, err := gob.ObtenerListaPiratas(ctx, &pb.Empty{})
		if err != nil {
			log.Printf("Error al obtener lista de piratas: %v", err)
			cancel() // Cancelar el contexto antes de continuar
			continue
		}

		// Filtrar piratas con estado "Buscado"
		var piratasBuscados []*pb.Pirata
		for _, pirata := range lista.Piratas {
			if pirata.Estado == "Buscado" {
				piratasBuscados = append(piratasBuscados, pirata)
			}
		}

		if len(piratasBuscados) == 0 {
			log.Println("No hay piratas con estado 'Buscado'. Terminando el proceso.")
			cancel()   // Cancelar el contexto antes de salir
			os.Exit(0) // Terminar el programa
		}

		// Seleccionar un pirata aleatorio
		pirata := piratasBuscados[rand.Intn(len(piratasBuscados))]

		// Reportar el estado del pirata como "Capturado"
		gob.ReporteDeEstado(ctx, &pb.EstadoRequest{
			IdCazarrecompensas: cazarrecompensa.Nombre,
			Pirata:             pirata,
			NuevoEstado:        "Capturado",
		})

		// Obtener el estado global del flujo ilegal
		estadoGlobal, err := gob.ObtenerEstadoGlobal(ctx, &pb.Empty{})
		if err != nil {
			log.Printf("Error al obtener el estado global del flujo ilegal: %v", err)
			cancel()
			continue
		}

		// Calcular el porcentaje de entregas al Submundo
		totalEntregas := estadoGlobal.MarinaCounter + estadoGlobal.SubmundoCounter
		porcentajeSubmundo := 0.0
		if totalEntregas > 0 {
			porcentajeSubmundo = float64(estadoGlobal.SubmundoCounter) / float64(totalEntregas) * 100
		}

		// Decidir el destino basado en el estado global y la recompensa del pirata
		if porcentajeSubmundo < 30 || pirata.Recompensa > 300000000 {
			log.Printf("Cazarrecompensas %s decide ir al Submundo por el bajo flujo ilegal o la alta recompensa del pirata %s.\n", cazarrecompensa.Nombre, pirata.Nombre)
			// Simular el transporte del pirata
			if !simularTransporte(ctx, gob, cazarrecompensa, pirata) {

				cancel()
				continue // Pasar al siguiente ciclo
			}

			entregarAlSubmundo(ctx, sub, cazarrecompensa, pirata, cazarrecompensas, x)
		} else {
			log.Printf("Cazarrecompensas %s decide ir a la Marina debido al alto flujo ilegal y la baja recompensa del pirata %s.\n", cazarrecompensa.Nombre, pirata.Nombre)
			// Simular el transporte del pirata
			if !simularTransporte(ctx, gob, cazarrecompensa, pirata) {

				cancel()
				continue // Pasar al siguiente ciclo
			}

			entregarALaMarina(ctx, mar, sub, gob, cazarrecompensa, pirata, cazarrecompensas, x)
		}

		// Cancelar el contexto al final de la iteración
		cancel()
	}
}
