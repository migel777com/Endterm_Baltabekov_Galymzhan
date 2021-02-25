package main

import (
	"CalculatorService/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:59751", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewCalcServiceClient(conn)
	getDecomposedNumbers(c)

}

func getDecomposedNumbers(c proto.CalcServiceClient) {
	ctx := context.Background()
	req := &proto.PrimeDecomposeRequest{Number: 120}
	stream, err := c.PrimeDecomposition(ctx, req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumDecompos RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// we've reached the end of the stream
				break LOOP
			}
			log.Fatalf("error while reciving from PrimeNumDecompos RPC %v", err)
		}
		log.Printf("response from PrimeNumDecompos:%v \n", res.Number)
	}

}


func getAverage(c proto.CalcServiceClient) {

	requests := []*proto.AverageRequest{
		{
			Number: 5,
		},
		{
			Number: 10,
		},
		{
			Number: 15,
		},
	}

	ctx := context.Background()
	stream, err := c.AverageCalculator(ctx)
	if err != nil {
		log.Fatalf("error while calling: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}
	fmt.Printf("Response: %v\n", res)
}


