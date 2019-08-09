package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	pb "timeservices"

	"google.golang.org/grpc"
)

func greet(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello from %s!\n", hostname)
	query := r.URL.Query()
	if query.Get("cookies") != "" {
		fmt.Fprintf(w, "Cookies:  %s\n", r.Cookies())
	}
	if query.Get("headers") != "" {
		fmt.Fprintf(w, "Headers:  %s\n", r.Header)
	}
	if query.Get("time") != "" {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Did not connect grpc: %v", err)
		}
		defer conn.Close()
		client := pb.NewUTCTimeServicesClient(conn)
		timeresp, err := client.GetUTCTime(context.Background(), &pb.UTCTimeRequest{})
		if err != nil {
			log.Fatalf("Couldn't get server time: %v", err)
		}
		fmt.Fprintln(w, timeresp.Time)
	}
}

func main() {
	http.HandleFunc("/", greet)
	fmt.Println("Listening on port 8080 for connections...")
	http.ListenAndServe(":8080", nil)
}
