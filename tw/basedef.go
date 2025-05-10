package tw

const (
	TW_QUERY_TYPE_LATEST string = "Latest"
	TW_QUERY_TYPE_TOP    string = "Top"
)

type STUserInfo struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	UserName        string `json:"userName"`
	Location        string `json:"location"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Protected       bool   `json:"protected"`
	IsVerified      bool   `json:"isVerified"`
	IsBlueVerified  bool   `json:"isBlueVerified"`
	Followers       int    `json:"followers"`
	Following       int    `json:"following"`
	FavouritesCount int    `json:"favouritesCount"`
	StatusesCount   int    `json:"statusesCount"`
	MediaCount      int    `json:"mediaCount"`
	CreatedAt       string `json:"createdAt"`
	CoverPicture    string `json:"coverPicture"`
	ProfilePicture  string `json:"profilePicture"`
	CanDm           bool   `json:"canDm"`
	IsAutomated     bool   `json:"isAutomated"`
	AutomatedBy     string `json:"automatedBy"`
}

type STTweetInfo struct {
	Type              string   `json:"type"`
	ID                string   `json:"id"`
	URL               string   `json:"url"`
	TwitterURL        string   `json:"twitterUrl"`
	Text              string   `json:"text"`
	Source            string   `json:"source"`
	RetweetCount      int      `json:"retweetCount"`
	ReplyCount        int      `json:"replyCount"`
	LikeCount         int      `json:"likeCount"`
	QuoteCount        int      `json:"quoteCount"`
	ViewCount         int      `json:"viewCount"`
	CreatedAt         string   `json:"createdAt"`
	Lang              string   `json:"lang"`
	BookmarkCount     int      `json:"bookmarkCount"`
	IsReply           bool     `json:"isReply"`
	InReplyToID       string   `json:"inReplyToId"`
	ConversationID    string   `json:"conversationId"`
	InReplyToUserID   string   `json:"inReplyToUserId"`
	InReplyToUsername string   `json:"inReplyToUsername"`
	Author            STAuthor `json:"author"`
}

type STAuthor struct {
	Type               string   `json:"type"`
	UserName           string   `json:"userName"`
	URL                string   `json:"url"`
	TwitterURL         string   `json:"twitterUrl"`
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	IsVerified         bool     `json:"isVerified"`
	IsBlueVerified     bool     `json:"isBlueVerified"`
	ProfilePicture     string   `json:"profilePicture"`
	CoverPicture       string   `json:"coverPicture"`
	Description        string   `json:"description"`
	Location           string   `json:"location"`
	Followers          int      `json:"followers"`
	Following          int      `json:"following"`
	Status             string   `json:"status"`
	CanDm              bool     `json:"canDm"`
	CanMediaTag        bool     `json:"canMediaTag"`
	CreatedAt          string   `json:"createdAt"`
	FastFollowersCount int      `json:"fastFollowersCount"`
	FavouritesCount    int      `json:"favouritesCount"`
	HasCustomTimelines bool     `json:"hasCustomTimelines"`
	IsTranslator       bool     `json:"isTranslator"`
	MediaCount         int      `json:"mediaCount"`
	StatusesCount      int      `json:"statusesCount"`
	PossiblySensitive  bool     `json:"possiblySensitive"`
	PinnedTweetIds     []string `json:"pinnedTweetIds"`
	IsAutomated        bool     `json:"isAutomated"`
}

type STFollowingInfo struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	ScreenName           string `json:"screen_name"`
	Location             string `json:"location"`
	URL                  any    `json:"url"`
	Description          string `json:"description"`
	Email                any    `json:"email"`
	Protected            bool   `json:"protected"`
	Verified             bool   `json:"verified"`
	FollowersCount       int    `json:"followers_count"`
	FollowingCount       int    `json:"following_count"`
	FriendsCount         int    `json:"friends_count"`
	FavouritesCount      int    `json:"favourites_count"`
	StatusesCount        int    `json:"statuses_count"`
	MediaTweetsCount     int    `json:"media_tweets_count"`
	CreatedAt            string `json:"created_at"`
	ProfileBannerURL     string `json:"profile_banner_url"`
	ProfileImageURLHTTPS string `json:"profile_image_url_https"`
	CanDm                bool   `json:"can_dm"`
}
