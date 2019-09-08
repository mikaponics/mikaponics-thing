/**
 *  THIS IS A SAMPLE FILE WHICH WORKS AND IS FOR TESTING PURPOSES.
 *  - https://github.com/kelseyhightower/grpc-hello-service
 *  - https://davidsbond.github.io/2019/06/14/creating-grpc-interceptors-in-go.html
 *
 */

package internal // github.com/mikaponics/mikaponics-thing/internal

import (
	"context"
	"log"
	// "net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)


// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	log.Printf("%v", token) // For debugging purposes only

	// // Perform the token validation here. For the sake of this example, the code
	// // here forgoes any of the usual OAuth2 token validation and instead checks
	// // for a token matching an arbitrary string.
	// return token == "bearer 1234567890"
	return true //TODO: IMPLEMENT FIX
}


func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	m, err := handler(ctx, req)
	if err != nil {
		log.Fatalf("RPC failed to serve: %v", err)
	}
	return m, err
}
