package basedef

type STUserDetailResp struct {
	User STUser `json:"user"`
}

type STAffiliatesHighlightedLabel struct {
}
type STUrls struct {
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"expanded_url"`
	URL         string `json:"url"`
	Indices     []int  `json:"indices"`
}
type STDescription struct {
	Urls []STUrls `json:"urls"`
}
type STURL struct {
	Urls []STUrls `json:"urls"`
}
type STEntities struct {
	Description STDescription `json:"description"`
	URL         STURL         `json:"url"`
}
type STLegacy struct {
	CreatedAt               string     `json:"created_at"`
	DefaultProfile          bool       `json:"default_profile"`
	DefaultProfileImage     bool       `json:"default_profile_image"`
	Description             string     `json:"description"`
	Entities                STEntities `json:"entities"`
	FastFollowersCount      int        `json:"fast_followers_count"`
	FavouritesCount         int        `json:"favourites_count"`
	FollowersCount          int        `json:"followers_count"`
	FriendsCount            int        `json:"friends_count"`
	HasCustomTimelines      bool       `json:"has_custom_timelines"`
	IsTranslator            bool       `json:"is_translator"`
	ListedCount             int        `json:"listed_count"`
	Location                string     `json:"location"`
	MediaCount              int        `json:"media_count"`
	Name                    string     `json:"name"`
	NormalFollowersCount    int        `json:"normal_followers_count"`
	PinnedTweetIdsStr       []string   `json:"pinned_tweet_ids_str"`
	PossiblySensitive       bool       `json:"possibly_sensitive"`
	ProfileBannerURL        string     `json:"profile_banner_url"`
	ProfileImageURLHTTPS    string     `json:"profile_image_url_https"`
	ProfileInterstitialType string     `json:"profile_interstitial_type"`
	ScreenName              string     `json:"screen_name"`
	StatusesCount           int        `json:"statuses_count"`
	TranslatorType          string     `json:"translator_type"`
	URL                     string     `json:"url"`
	Verified                bool       `json:"verified"`
	WithheldInCountries     []any      `json:"withheld_in_countries"`
}
type STTipjarSettings struct {
	IsEnabled      bool   `json:"is_enabled"`
	BandcampHandle string `json:"bandcamp_handle"`
	BitcoinHandle  string `json:"bitcoin_handle"`
	CashAppHandle  string `json:"cash_app_handle"`
	EthereumHandle string `json:"ethereum_handle"`
	GofundmeHandle string `json:"gofundme_handle"`
	PatreonHandle  string `json:"patreon_handle"`
	PayPalHandle   string `json:"pay_pal_handle"`
	VenmoHandle    string `json:"venmo_handle"`
}
type STLegacyExtendedProfile struct {
}
type STRef struct {
	URL     string `json:"url"`
	URLType string `json:"url_type"`
}
type STDescEntities struct {
	FromIndex int   `json:"from_index"`
	ToIndex   int   `json:"to_index"`
	Ref       STRef `json:"ref"`
}
type STReasonDescription struct {
	Text     string           `json:"text"`
	Entities []STDescEntities `json:"entities"`
}
type Reason struct {
	Description       STReasonDescription `json:"description"`
	VerifiedSinceMsec string              `json:"verified_since_msec"`
}
type VerificationInfo struct {
	IsIdentityVerified bool   `json:"is_identity_verified"`
	Reason             Reason `json:"reason"`
}
type HighlightsInfo struct {
	CanHighlightTweets bool   `json:"can_highlight_tweets"`
	HighlightedTweets  string `json:"highlighted_tweets"`
}
type BusinessAccount struct {
}
type STResult struct {
	Typename                        string                       `json:"__typename"`
	ID                              string                       `json:"id"`
	RestID                          string                       `json:"rest_id"`
	AffiliatesHighlightedLabel      STAffiliatesHighlightedLabel `json:"affiliates_highlighted_label"`
	ParodyCommentaryFanLabel        string                       `json:"parody_commentary_fan_label"`
	IsBlueVerified                  bool                         `json:"is_blue_verified"`
	ProfileImageShape               string                       `json:"profile_image_shape"`
	Legacy                          STLegacy                     `json:"legacy"`
	TipjarSettings                  STTipjarSettings             `json:"tipjar_settings"`
	LegacyExtendedProfile           STLegacyExtendedProfile      `json:"legacy_extended_profile"`
	IsProfileTranslatable           bool                         `json:"is_profile_translatable"`
	HasHiddenSubscriptionsOnProfile bool                         `json:"has_hidden_subscriptions_on_profile"`
	VerificationInfo                VerificationInfo             `json:"verification_info"`
	HighlightsInfo                  HighlightsInfo               `json:"highlights_info"`
	UserSeedTweetCount              int                          `json:"user_seed_tweet_count"`
	BusinessAccount                 BusinessAccount              `json:"business_account"`
	CreatorSubscriptionsCount       int                          `json:"creator_subscriptions_count"`
}
type STUser struct {
	Result STResult `json:"result"`
}
