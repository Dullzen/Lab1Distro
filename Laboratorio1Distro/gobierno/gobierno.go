package main

import (
	"context"
	"encoding/csv"
	"log"
	"net"
	"os"
	"strconv"

	pb "gobierno/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

// Cazarrecompensa representa un cazarrecompensas con su reputación
type Cazarrecompensa struct {
	Nombre           string
	Reputacion       int
	EntregasMarina   int // Contador de entregas a la Marina
	EntregasSubmundo int // Contador de entregas al Submundo
}

type server struct {
	pb.UnimplementedGobiernoServiceServer
	piratas          []*pb.Pirata                // Lista de piratas cargada en memoria
	cazarrecompensas map[string]*Cazarrecompensa // Mapa de cazarrecompensas por nombre
	submundoCounter  int                         // Contador de entregas al Submundo
	marinaCounter    int                         // Contador de entregas a la Marina
	FlujoIlegalAlto  bool                        // Indicador de si el flujo ilegal es alto
	marinaClient     pb.MarinaServiceClient      // Cliente para enviar alertas a la Marina
}

func (s *server) ObtenerListaPiratas(ctx context.Context, in *pb.Empty) (*pb.ListaPiratas, error) {

	return &pb.ListaPiratas{Piratas: s.piratas}, nil
}

func (s *server) ObtenerEstadoGlobal(ctx context.Context, req *pb.Empty) (*pb.EstadoGlobalResponse, error) {
	return &pb.EstadoGlobalResponse{
		MarinaCounter:   int32(s.marinaCounter),
		SubmundoCounter: int32(s.submundoCounter),
	}, nil
}

func (s *server) ReporteDeEstado(ctx context.Context, req *pb.EstadoRequest) (*pb.ResultadoOperacion, error) {
	// Buscar al pirata en la lista
	for _, pirata := range s.piratas {
		if pirata.Id == req.Pirata.Id {
			if req.NuevoEstado == "Perdido" {
				// Cambiar el estado del pirata a "Buscado"
				pirata.Estado = "Buscado"
				// Reducir la reputación del cazarrecompensas
				cazarrecompensa, exists := s.cazarrecompensas[req.IdCazarrecompensas]

				if !exists {
					cazarrecompensa = &Cazarrecompensa{
						Nombre:           req.IdCazarrecompensas,
						Reputacion:       50,
						EntregasMarina:   0,
						EntregasSubmundo: 0,
					}
					s.cazarrecompensas[req.IdCazarrecompensas] = cazarrecompensa
					log.Printf("Gobierno: %s se registra como cazarrecompensas.\n", req.IdCazarrecompensas)
				}

				cazarrecompensa.Reputacion -= 10
				if cazarrecompensa.Reputacion < 0 {
					cazarrecompensa.Reputacion = 0 // Evitar que la reputación sea negativa
				}
				log.Printf("Gobierno: Reputación de %s reducida a %d por perder al pirata %s\n", req.IdCazarrecompensas, cazarrecompensa.Reputacion, pirata.Nombre)

			} else if req.NuevoEstado == "Capturado" {
				log.Printf("-------------------------------------------------")
				log.Printf("-------------------------------------------------")

				pirata.Estado = req.NuevoEstado
				log.Printf("Gobierno: Estado del pirata %s actualizado a '%s'", pirata.Nombre, req.NuevoEstado)
			} else {

				pirata.Estado = req.NuevoEstado
				log.Printf("Gobierno: Estado del pirata %s actualizado a '%s'", pirata.Nombre, req.NuevoEstado)
			}
			break
		}
	}

	return &pb.ResultadoOperacion{Exito: true, Mensaje: "Estado actualizado"}, nil
}

func (s *server) ConfirmarEntrega(ctx context.Context, req *pb.ConfirmarEntregaRequest) (*pb.ResultadoOperacion, error) {
	log.Printf("Gobierno: Confirmada entrega de %s por %s al destino %s\n", req.Pirata.Nombre, req.IdCazarrecompensas, req.Destino)

	// Actualizar la reputación del cazarrecompensa
	cazarrecompensa, exists := s.cazarrecompensas[req.IdCazarrecompensas]
	if !exists {
		// Si no existe, agregarlo con reputación base de 50
		cazarrecompensa = &Cazarrecompensa{
			Nombre:           req.IdCazarrecompensas,
			Reputacion:       50,
			EntregasMarina:   0,
			EntregasSubmundo: 0,
		}
		s.cazarrecompensas[req.IdCazarrecompensas] = cazarrecompensa
	}

	// Modificar la reputación y los contadores según el destino
	if req.Destino == "Marina" {
		cazarrecompensa.Reputacion += 10
		cazarrecompensa.EntregasMarina++ // Incrementar el contador de entregas a la Marina
		s.marinaCounter++                // Incrementar el contador global de entregas a la Marina
	} else if req.Destino == "Submundo" {
		cazarrecompensa.Reputacion -= 10
		if cazarrecompensa.Reputacion < 0 {
			cazarrecompensa.Reputacion = 0 // Evitar que la reputación sea negativa
		}
		cazarrecompensa.EntregasSubmundo++ // Incrementar el contador de entregas al Submundo
		s.submundoCounter++                // Incrementar el contador global de entregas al Submundo
	}

	// Actualizar el estado del pirata
	for _, pirata := range s.piratas {
		if pirata.Id == req.Pirata.Id {
			if req.Destino == "Marina" {
				pirata.Estado = "En Marina"
			} else if req.Destino == "Submundo" {
				pirata.Estado = "En Submundo"
			}
			log.Printf("Gobierno: Estado del pirata %s actualizado a '%s'\n", pirata.Nombre, pirata.Estado)
			break
		}
	}

	// Calcular el porcentaje de entregas al Submundo
	totalEntregas := s.marinaCounter + s.submundoCounter
	porcentajeSubmundo := float64(s.submundoCounter) / float64(totalEntregas) * 100

	// Verificar si se debe enviar una alerta o registrar un mensaje de flujo normal
	if totalEntregas > 4 {
		if porcentajeSubmundo > 40 {
			if !s.FlujoIlegalAlto {
				log.Println("(*****Gobierno: Alto tráfico ilegal detectado. Alerta a la Marina.*****)")

				_, err := s.marinaClient.AlertaTraficoIlegal(ctx, &pb.Empty{})
				if err != nil {
					log.Printf("Error al enviar alerta a la Marina: %v", err)
				} else {
					log.Println("Gobierno: Alerta enviada a la Marina con éxito.")
					s.FlujoIlegalAlto = true // Marcar que el flujo ilegal es alto
				}
			}
		} else {
			if s.FlujoIlegalAlto {
				log.Println("(*****Gobierno: El flujo de entregas ha vuelto a la normalidad.*****)")
				s.FlujoIlegalAlto = false // Marcar que el flujo está normal
			}
		}
	}

	log.Printf("Gobierno: Reputación de %s actualizada a %d\n", cazarrecompensa.Nombre, cazarrecompensa.Reputacion)

	return &pb.ResultadoOperacion{Exito: true, Mensaje: "Entrega confirmada"}, nil
}

func (s *server) ConsultarReputacion(ctx context.Context, req *pb.ReputacionRequest) (*pb.ReputacionResponse, error) {
	cazarrecompensa, exists := s.cazarrecompensas[req.IdCazarrecompensas]
	if !exists {
		// Si no existe, crearlo con valores iniciales
		cazarrecompensa = &Cazarrecompensa{
			Nombre:           req.IdCazarrecompensas,
			Reputacion:       50,
			EntregasMarina:   0,
			EntregasSubmundo: 0,
		}
		s.cazarrecompensas[req.IdCazarrecompensas] = cazarrecompensa
	}

	return &pb.ReputacionResponse{
		Reputacion:       int32(cazarrecompensa.Reputacion),
		EntregasMarina:   int32(cazarrecompensa.EntregasMarina),
		EntregasSubmundo: int32(cazarrecompensa.EntregasSubmundo),
	}, nil
}

func cargarPiratasDesdeCSV(nombreArchivo string) ([]*pb.Pirata, error) {
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

	var piratas []*pb.Pirata
	for _, record := range records[1:] { // Saltar la cabecera
		recompensa, _ := strconv.Atoi(record[2])
		piratas = append(piratas, &pb.Pirata{
			Id:           record[0],
			Nombre:       record[1],
			Recompensa:   int32(recompensa),
			Peligrosidad: record[3],
			Estado:       record[4],
		})
	}

	return piratas, nil
}

func main() {
	// Cargar la lista de piratas al iniciar el servidor
	piratas, err := cargarPiratasDesdeCSV("piratas.csv")
	if err != nil {
		log.Fatalf("Error al cargar piratas desde el archivo CSV: %v", err)
	}

	// Conectar al servicio Marina
	connMarina, err := grpc.Dial("marina:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al servicio Marina: %v", err)
	}
	defer connMarina.Close()

	marinaClient := pb.NewMarinaServiceClient(connMarina)

	// Crear el servidor con la lista de piratas cargada y un mapa vacío de cazarrecompensas
	s := &server{
		piratas:          piratas,
		cazarrecompensas: make(map[string]*Cazarrecompensa),
		submundoCounter:  0,
		marinaCounter:    0,
		FlujoIlegalAlto:  false,
		marinaClient:     marinaClient,
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Fallo al escuchar: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGobiernoServiceServer(grpcServer, s)
	log.Println("Gobierno corriendo en puerto 50051")
	log.Println(" ")
	log.Println(" ")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al iniciar servidor: %v", err)
	}
}
