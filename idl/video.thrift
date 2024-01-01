namespace go video

struct BaseResp {
    1: i64 code,
    2: string msg
}

struct Video {
    1: i64 id,
    2: i64 author_id,
    3: string play_url,
    4: string cover_url,
    5: i64 favorite_count,
    6: i64 comment_count,
    8: string title
}

struct FeedRequest {
    1: optional i64 latest_time,
}

struct FeedResponse {
    1: BaseResp base,
    2: i64 next_time,
    3: list<Video> video_list
}

struct PublishVideoRequest {
    1: string title,      // 视频标题
    2: i64 author_id
}

struct PublishVideoResponse {
    1: BaseResp base
}

struct GetVideoListRequest {
    1: i64 author_id
}

struct GetVideoListResponse {
    1: BaseResp base,
    2: list<Video> video_list
}

service VideoService {
    FeedResponse Feed(1: FeedRequest request),
    PublishVideoResponse PublishVideo(1: PublishVideoRequest request),
    GetVideoListResponse GetVideoList(1: GetVideoListRequest request)
}
