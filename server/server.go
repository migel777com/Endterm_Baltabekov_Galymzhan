package main

import (
	"CalculatorService/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type Server struct {
	proto.UnimplementedCalcServiceServer
}


func (s *Server) PrimeDecomposition(req *proto.PrimeDecomposeRequest,stream proto.CalcService_PrimeDecompositionServer)  error{
	number := int(req.Number)
	for number%2 == 0 {
		res := &proto.PrimeDecomposeResponse{Number: 2}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending stream responses: %v", err.Error())
		}
		number = number / 2
	}


	for i := 3; i*i <= number; i = i + 2 {
		// while i divides n, append i and divide n
		for number%i == 0 {
			res := &proto.PrimeDecomposeResponse{Number: int32(i)}
			if err := stream.Send(res); err != nil {
				log.Fatalf("error while sending stream responses: %v", err.Error())
			}
			number = number / i
		}
	}

	if number > 2 {
		res := &proto.PrimeDecomposeResponse{Number: int32(number),}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error while sending stream responses: %v", err.Error())
		}
	}

	return nil
}

func (s *Server) AverageCalculator(stream proto.CalcService_AverageCalculatorServer) error {
	var result float64
	var sum float64
	var devider float64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&proto.AverageRespond{
				Number: result,
			})

		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		tmp := req.GetNumber()
		sum+= tmp
		devider++

		result = sum/devider
	}

}



func main() {
	l, err := net.Listen("tcp", "0.0.0.0:59751")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCalcServiceServer(s, &Server{})
	log.Println("Server is running on port:59751")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
