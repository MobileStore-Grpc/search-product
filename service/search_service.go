package service

import (
	"context"
	"errors"
	"log"

	"github.com/MobileStore-Grpc/search-product/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//MobileService provide mobile CURD operations
type SearchService struct {
	pb.UnimplementedMobileServiceServer
}

//NewMobileService returns a new mobile service
func NewSearchService() *SearchService {
	return &SearchService{}
}

//SearchLaptop is a server-streaming RPC to search for laptops
func (server *SearchService) SearchLaptop(req *pb.SearchLaptopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter: %v", filter)

	err := server.laptopStore.Search(
		stream.Context(),
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{Laptop: laptop}

			err := stream.Send(res)
			if err != nil {
				return err
			}
			log.Printf("send laptop with id: %s", laptop.GetId())
			return nil
		})

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
