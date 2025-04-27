package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "submundo/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSubmundoServiceServer
	gobiernoClient pb.GobiernoServiceClient // Cliente para Gobierno
}

func (s *server) EntregarPirataSubmundo(ctx context.Context, req *pb.VentaRequest) (*pb.ResultadoOperacion, error) {
	log.Printf("Submundo: Recibido pirata %s de %s\n", req.Pirata.Nombre, req.IdCazarrecompensas)

	// Determinar si ocurre un fraude
	rand.Seed(time.Now().UnixNano()) // Semilla para el generador de números aleatorios
	if rand.Intn(100) < 35 {         // 35% de probabilidad de fraude
		return &pb.ResultadoOperacion{
			Exito:   false,
			Mensaje: "0", // Recompensa es 0 debido al fraude
		}, nil
	}

	// Calcular la recompensa entre 100% y 150%
	recompensaFactor := 1.0 + (rand.Float64() * 0.5) // Genera un número entre 1.0 y 1.5
	recompensa := int(float64(req.Pirata.Recompensa) * recompensaFactor)

	// Notificar al Gobierno sobre la entrega
	_, err := s.gobiernoClient.ConfirmarEntrega(ctx, &pb.ConfirmarEntregaRequest{
		IdCazarrecompensas: req.IdCazarrecompensas,
		Pirata:             req.Pirata,
		Destino:            "Submundo",
	})
	if err != nil {
		log.Printf("Error al notificar al Gobierno: %v", err)
	}

	// Retornar la recompensa calculada
	return &pb.ResultadoOperacion{
		Exito:   true,
		Mensaje: strconv.Itoa(recompensa), // Convertir la recompensa a string
	}, nil
}

func conectarGRPC(address string, maxRetries int, retryDelay time.Duration) (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error

	for i := 0; i < maxRetries; i++ {
		conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
		if err == nil {
			return conn, nil // Conexión exitosa
		}

		log.Printf("Intento %d: No se pudo conectar a %s: %v", i+1, address, err)
		time.Sleep(retryDelay) // Esperar antes de reintentar
	}

	return nil, fmt.Errorf("no se pudo conectar a %s después de %d intentos", address, maxRetries)
}

func main() {
	// Crear el servidor con el cliente para Gobierno (inicialmente vacío)
	s := &server{}

	// Iniciar el servidor gRPC en una goroutine
	go func() {
		lis, err := net.Listen("tcp", ":50053")
		if err != nil {
			log.Fatalf("Fallo al escuchar: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterSubmundoServiceServer(grpcServer, s)

		log.Println("Submundo corriendo en puerto 50053")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Fallo al iniciar servidor: %v", err)
		}
	}()

	// Conectar al servicio Gobierno con intentos y timeout
	connGob, err := conectarGRPC("10.35.168.64:50051", 5, 2*time.Second)
	if err != nil {
		log.Fatalf("No se pudo conectar al servicio Gobierno: %v", err)
	}
	defer connGob.Close()

	// Asignar el cliente de Gobierno al servidor
	s.gobiernoClient = pb.NewGobiernoServiceClient(connGob)

	// Bloquear el programa para que no termine
	select {}
}
