package grpc

import (
	"context"
	"log"
	pb "praisindo/proto/shorturl_service/v1"
	"praisindo/service"
)

// UserHandler is used to implement UnimplementedshorturlServiceServer
type ShortUrlHandler struct {
	pb.UnimplementedShortUrlServiceServer
	shorturlService service.IShortUrlService
}

// NewUserHandler membuat instance baru dari UserHandler
func NewShortUrlHandler(shorturlService service.IShortUrlService) *ShortUrlHandler {
	return &ShortUrlHandler{
		shorturlService: shorturlService,
	}
}

func (u *ShortUrlHandler) CreateShortUrl(ctx context.Context, req *pb.CreateShortUrlRequest) (*pb.CreateShortResponse, error) {
	CreateShortUrl, err := u.shorturlService.CreateShortUrl(req.UrlLong)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateShortResponse{
		UrlShort: CreateShortUrl.UrlShort,
		UrlLong:  CreateShortUrl.UrlLong,
	}, nil
}

func (u *ShortUrlHandler) GetShortUrl(ctx context.Context, req *pb.GetShortRequest) (*pb.GetShortResponse, error) {
	shorturl, err := u.shorturlService.GetShortUrl(req.UrlShort)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.GetShortResponse{
		UrlLong:  shorturl.UrlLong,
		UrlShort: shorturl.UrlShort,
	}, nil
}
