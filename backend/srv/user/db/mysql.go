//package db// connect returns SQL database connection from the pool
//import (
//	"database/sql"
//	"fmt"
//	"github.com/golang/protobuf/ptypes"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//)
//type toDoServiceServer struct {
//	db *sql.DB
//}
//func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
//	c, err := s.db.Conn(ctx)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
//	}
//	return c, nil
//}
//
//c, err := s.connect(ctx)
//if err != nil {
//return nil, err
//}
//defer c.Close()
//
//// insert ToDo entity data
//res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
//req.ToDo.Title, req.ToDo.Description, reminder)
//if err != nil {
//return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
//}
//
//
//package v1
//
//import (
//"context"
//"database/sql"
//"fmt"
//"time"
//
//"github.com/golang/protobuf/ptypes"
//"google.golang.org/grpc/codes"
//"google.golang.org/grpc/status"
//
//"github.com/amsokol/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
//)
//
//const (
//	// apiVersion is version of API is provided by server
//	apiVersion = "v1"
//)
//
//// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
//type toDoServiceServer struct {
//	db *sql.DB
//}
//
//// NewToDoServiceServer creates ToDo service
//func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer {
//	return &toDoServiceServer{db: db}
//}
//
//// checkAPI checks if the API version requested by client is supported by server
//func (s *toDoServiceServer) checkAPI(api string) error {
//	// API version is "" means use current version of the service
//	if len(api) > 0 {
//		if apiVersion != api {
//			return status.Errorf(codes.Unimplemented,
//				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
//		}
//	}
//	return nil
//}
//
//// connect returns SQL database connection from the pool
//func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
//	c, err := s.db.Conn(ctx)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to connect to database-> "+err.Error())
//	}
//	return c, nil
//}
//
//// Create new todo task
//func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
//	// check if the API version requested by client is supported by server
//	if err := s.checkAPI(req.Api); err != nil {
//		return nil, err
//	}
//
//	// get SQL connection from pool
//	c, err := s.connect(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, "reminder field has invalid format-> "+err.Error())
//	}
//
//	// insert ToDo entity data
//	res, err := c.ExecContext(ctx, "INSERT INTO ToDo(`Title`, `Description`, `Reminder`) VALUES(?, ?, ?)",
//		req.ToDo.Title, req.ToDo.Description, reminder)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to insert into ToDo-> "+err.Error())
//	}
//
//	// get ID of creates ToDo
//	id, err := res.LastInsertId()
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to retrieve id for created ToDo-> "+err.Error())
//	}
//
//	return &v1.CreateResponse{
//		Api: apiVersion,
//		Id:  id,
//	}, nil
//}
//
//// Read todo task
//func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
//	// check if the API version requested by client is supported by server
//	if err := s.checkAPI(req.Api); err != nil {
//		return nil, err
//	}
//
//	// get SQL connection from pool
//	c, err := s.connect(ctx)
//	if err != nil {
//		return nil, err
//	}
//	defer c.Close()
//
//	// query ToDo by ID
//	rows, err := c.QueryContext(ctx, "SELECT `ID`, `Title`, `Description`, `Reminder` FROM ToDo WHERE `ID`=?",
//		req.Id)
//	if err != nil {
//		return nil, status.Error(codes.Unknown, "failed to select from ToDo-> "+err.Error())
//	}
//	defer rows.Close()
//
//	if !rows.Next() {
//		if err := rows.Err(); err != nil {
//			return nil, status.Error(codes.Unknown, "failed to retrieve data from ToDo-> "+err.Error())
//		}
//		return nil, status.Error(codes.NotFound, fmt.Sprintf("ToDo with ID='%d' is not found",
//			req.Id))
//	}
//
//	type Config struct {
//		// gRPC server start parameters section
//		// gRPC is TCP port to listen by gRPC server
//		GRPCPort string
//
//		// DB Datastore parameters section
//		// DatastoreDBHost is host of database
//		DatastoreDBHost string
//		// DatastoreDBUser is username to connect to database
//		DatastoreDBUser string
//		// DatastoreDBPassword password to connect to database
//		DatastoreDBPassword string
//		// DatastoreDBSchema is schema of database
//		DatastoreDBSchema string
//	}


// get configuration
//var cfg Config
//flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
//flag.StringVar(&cfg.DatastoreDBHost, "db-host", "", "Database host")
//flag.StringVar(&cfg.DatastoreDBUser, "db-user", "", "Database user")
//flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "", "Database password")
//flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "", "Database schema")
//flag.Parse()

/// add MySQL driver specific parameter to parse date/time
//// Drop it for another database
//param := "parseTime=true"
//
//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
//cfg.DatastoreDBUser,
//cfg.DatastoreDBPassword,
//cfg.DatastoreDBHost,
//cfg.DatastoreDBSchema,
//param)
//db, err := sql.Open("mysql", dsn)
//if err != nil {
//return fmt.Errorf("failed to open database: %v", err)
//}
//defer db.Close()