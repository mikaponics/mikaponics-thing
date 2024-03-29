package internal // github.com/mikaponics/mikapod-iam/internal

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/mikaponics/mikaponics-thing/api"

    "github.com/mikaponics/mikaponics-thing/internal/services"
	"github.com/mikaponics/mikaponics-thing/internal/controllers"
	"github.com/mikaponics/mikaponics-thing/internal/models"
)

type MikaponicsThing struct {
	thingAddress string
	dal *models.DataAccessLayer
	iam *services.IAMClient
	grpcServer *grpc.Server
}

// Function will construct the Mikapod IAM application.
func InitMikaponicsThing(dbHost, dbPort, dbUser, dbPassword, dbName, thingAddress string, iamAddress string) (*MikaponicsThing) {

	// Initialize and connect our database layer for the entire application.
    dal := models.InitDataAccessLayer(dbHost, dbPort, dbUser, dbPassword, dbName)

    // Create our app's models if they haven't been created previously.
    dal.CreateThingTable(false)
	dal.CreateSensorTable(false)
	dal.CreateTimeSeriesDatumTable(false)

    // Initialize our IAM client.
	iam := services.InitIAMClient(iamAddress)

	// Create our application instance.
 	return &MikaponicsThing{
		thingAddress: thingAddress,
		dal: dal,
		iam: iam,
		grpcServer: nil,
	}
}

// Function will consume the main runtime loop and run the business logic
// of the Mikapod IAM application.
func (app *MikaponicsThing) RunMainRuntimeLoop() {
	// Open a TCP server to the specified localhost and environment variable
    // specified port number.
    lis, err := net.Listen("tcp", app.thingAddress)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Initialize our gRPC server using our TCP server.
    grpcServer := grpc.NewServer( grpc.UnaryInterceptor(unaryInterceptor), ) // EX: Added `UnaryInterceptor`.

    // Save reference to our application state.
    app.grpcServer = grpcServer

    // For debugging purposes only.
    log.Printf("gRPC server is running.")

    // Block the main runtime loop for accepting and processing gRPC requests.
    pb.RegisterMikaponicsThingServer(grpcServer, &controllers.MikaponicsThingServer{
        // DEVELOPERS NOTE:
        // We want to attach to every gRPC call the following variables...
        DAL: app.dal,
		IAM: app.iam,
    })
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (app *MikaponicsThing) StopMainRuntimeLoop() {
	// Finish any RPC communication taking place at the moment before
    // shutting down the gRPC server.
    app.grpcServer.GracefulStop()

	// Shutdown our connection with our database.
	app.dal.Shutdown()
}
