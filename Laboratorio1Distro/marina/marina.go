package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "marina/proto/grpc-server/proto" // Usar el mismo proto compartido

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMarinaServiceServer
	gobiernoClient    pb.GobiernoServiceClient // Cliente para Gobierno usando el mismo paquete
	piratasEntregados map[string]bool          // Mapa para rastrear piratas entregados
	factor            float64                  // Factor de reducción de recompensa
}

func (s *server) EntregarPirataMarina(ctx context.Context, req *pb.EntregaRequest) (*pb.ResultadoOperacion, error) {
	log.Printf("Marina: Recibido pirata %s de %s\n", req.Pirata.Nombre, req.IdCazarrecompensas)

	// Verificar si el pirata ya fue entregado
	if s.piratasEntregados[req.Pirata.Id] {
		log.Printf("Marina: El pirata %s ya fue entregado previamente. Rechazando la entrega.", req.Pirata.Nombre)
		return &pb.ResultadoOperacion{
			Exito:   false,
			Mensaje: "El pirata ya fue entregado previamente",
		}, nil
	}

	// Consultar la reputación del cazarrecompensas
	reputacionResp, err := s.gobiernoClient.ConsultarReputacion(ctx, &pb.ReputacionRequest{
		IdCazarrecompensas: req.IdCazarrecompensas,
	})
	if err != nil {
		log.Printf("Error al consultar la reputación del cazarrecompensas: %v", err)
		return &pb.ResultadoOperacion{
			Exito:   false,
			Mensaje: "Error al consultar reputación",
		}, nil
	}

	// Verificar si la reputación es suficiente
	if reputacionResp.Reputacion <= 30 {
		// Hacer un 50/50 aleatorio para decidir si se permite la entrega
		if rand.Intn(2) == 0 { // 50% de probabilidad de rechazo
			log.Printf("Marina: Reputación de %s es %d. Se ha decidido rechazar esta vez.", req.IdCazarrecompensas, reputacionResp.Reputacion)
			return &pb.ResultadoOperacion{
				Exito:   false,
				Mensaje: "0", // Recompensa es 0
			}, nil
		}
		log.Printf("Marina: Reputación de %s es %d. Se ha decidido aceptar esta vez.", req.IdCazarrecompensas, reputacionResp.Reputacion)
	}

	// Calcular el porcentaje de entregas al Submundo
	totalEntregas := reputacionResp.EntregasMarina + reputacionResp.EntregasSubmundo
	porcentajeSubmundo := 0.0
	if totalEntregas > 0 {
		porcentajeSubmundo = float64(reputacionResp.EntregasSubmundo) / float64(totalEntregas) * 100
	}

	// Reducir la recompensa si el porcentaje de entregas al Submundo es mayor al 30%
	recompensa := req.Pirata.Recompensa
	if porcentajeSubmundo > 40 {
		log.Printf("Marina: Porcentaje de entregas al Submundo es mayor al 40%%. Recompensa reducida en un 50%%.")
		recompensa = recompensa / 2
	}

	// Aplicar el factor de reducción
	recompensa = int32(float64(recompensa) * s.factor)

	// Notificar al Gobierno sobre la entrega
	_, err = s.gobiernoClient.ConfirmarEntrega(ctx, &pb.ConfirmarEntregaRequest{
		IdCazarrecompensas: req.IdCazarrecompensas,
		Pirata:             req.Pirata,
		Destino:            "Marina",
	})

	if err != nil {
		log.Printf("Error al notificar al Gobierno: %v", err)
		return &pb.ResultadoOperacion{
			Exito:   false,
			Mensaje: "Error al confirmar entrega",
		}, nil
	}

	// Agregar el pirata a la lista de entregados
	s.piratasEntregados[req.Pirata.Id] = true
	log.Printf("Marina: El pirata %s ha sido entregado exitosamente.", req.Pirata.Nombre)

	// Usar el campo Recompensa del pirata como mensaje de éxito
	return &pb.ResultadoOperacion{
		Exito:   true,
		Mensaje: strconv.Itoa(int(recompensa)), // Convertir la recompensa a string
	}, nil
}

func (s *server) AlertaTraficoIlegal(ctx context.Context, req *pb.Empty) (*pb.ResultadoOperacion, error) {
	log.Println("Marina: Recibida alerta de alto tráfico ilegal del Gobierno.")
	return &pb.ResultadoOperacion{
		Exito:   true,
		Mensaje: "Alerta recibida. Tomando medidas.",
	}, nil
}

func main() {
	// Conectar al servicio Gobierno
	connGob, err := grpc.Dial("gobierno:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al servicio Gobierno: %v", err)
	}
	defer connGob.Close()

	gobiernoClient := pb.NewGobiernoServiceClient(connGob) // Crear cliente usando el mismo paquete

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Fallo al escuchar: %v", err)
	}

	// Crear el servidor
	s := &server{
		gobiernoClient:    gobiernoClient,
		piratasEntregados: make(map[string]bool), // Inicializar el mapa
		factor:            1.0,                   // Inicialmente no hay reducción
	}

	// Rutina que verifica el porcentaje de piratas "Buscado"
	go func() {
		for {
			time.Sleep(5 * time.Second) // Ejecutar cada 5 segundos

			// Obtener la lista de piratas del Gobierno
			resp, err := s.gobiernoClient.ObtenerListaPiratas(context.Background(), &pb.Empty{})
			if err != nil {
				log.Printf("Error al obtener la lista de piratas: %v", err)
				continue
			}

			// Calcular el porcentaje de piratas con estado "Buscado"
			totalPiratas := len(resp.Piratas)
			if totalPiratas == 0 {
				s.factor = 1.0 // No hay piratas, no aplicar reducción
				continue
			}

			buscados := 0
			for _, pirata := range resp.Piratas {
				if pirata.Estado == "Buscado" {
					buscados++
				}
			}

			porcentajeBuscados := float64(buscados) / float64(totalPiratas) * 100

			// Si el porcentaje cae por debajo del 50%, aplicar reducción
			if porcentajeBuscados < 20 {
				s.factor = 0.6 // Reducir la recompensa en un 30%
			} else if porcentajeBuscados < 40 {
				s.factor = 0.8 // Reducir la recompensa en un 30%
			} else {
				s.factor = 1.0 // No aplicar reducción
			}
		}
	}()

	// Iniciar el servidor gRPC
	grpcServer := grpc.NewServer()
	pb.RegisterMarinaServiceServer(grpcServer, s)
	log.Println("Marina corriendo en puerto 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fallo al iniciar servidor: %v", err)
	}
}
