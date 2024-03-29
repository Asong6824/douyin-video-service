// Code generated by Kitex v0.8.0. DO NOT EDIT.

package videoservice

import (
	"context"
	video "github.com/Asong6824/douyin-video-service/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Feed(ctx context.Context, request *video.FeedRequest, callOptions ...callopt.Option) (r *video.FeedResponse, err error)
	PublishVideo(ctx context.Context, request *video.PublishVideoRequest, callOptions ...callopt.Option) (r *video.PublishVideoResponse, err error)
	GetVideoList(ctx context.Context, request *video.GetVideoListRequest, callOptions ...callopt.Option) (r *video.GetVideoListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) Feed(ctx context.Context, request *video.FeedRequest, callOptions ...callopt.Option) (r *video.FeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, request)
}

func (p *kVideoServiceClient) PublishVideo(ctx context.Context, request *video.PublishVideoRequest, callOptions ...callopt.Option) (r *video.PublishVideoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishVideo(ctx, request)
}

func (p *kVideoServiceClient) GetVideoList(ctx context.Context, request *video.GetVideoListRequest, callOptions ...callopt.Option) (r *video.GetVideoListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideoList(ctx, request)
}
