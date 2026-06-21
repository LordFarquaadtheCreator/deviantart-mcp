package deviantart

import "time"

// User represents a DeviantArt user
type User struct {
	UserID    string    `json:"userid"`
	Username  string    `json:"username"`
	UserIcon  string    `json:"usericon"`
	Type      string    `json:"type"`
	IsWatching bool     `json:"is_watching,omitempty"`
	Details   UserDetails `json:"details,omitempty"`
	Geo       UserGeo   `json:"geo,omitempty"`
	Profile   UserProfile `json:"profile,omitempty"`
	Stats     UserStats `json:"stats,omitempty"`
}

// UserDetails contains user details
type UserDetails struct {
	Sex string `json:"sex,omitempty"`
	Age int    `json:"age,omitempty"`
}

// UserGeo contains user geographic information
type UserGeo struct {
	Country   string `json:"country,omitempty"`
	CountryID int    `json:"countryid,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
}

// UserProfile contains user profile information
type UserProfile struct {
	UserTitle  string      `json:"user_title,omitempty"`
	RealName   string      `json:"real_name,omitempty"`
	ProfilePic Image       `json:"profile_pic,omitempty"`
	CoverPhoto *Image      `json:"cover_photo,omitempty"`
	Tagline    string      `json:"tagline,omitempty"`
}

// UserStats contains user statistics
type UserStats struct {
	Watchers   int `json:"watchers,omitempty"`
	Friends    int `json:"friends,omitempty"`
	Deviations int `json:"deviations,omitempty"`
}

// Deviation represents a DeviantArt deviation
type Deviation struct {
	DeviationID      string           `json:"deviationid"`
	PrintID          string           `json:"printid,omitempty"`
	Author           User             `json:"author"`
	IsFavourited     bool             `json:"is_favourited"`
	IsMature         bool             `json:"is_mature"`
	IsDownloadable   bool             `json:"is_downloadable"`
	AllowComments    bool             `json:"allow_comments"`
	Title            string           `json:"title"`
	Body             string           `json:"body,omitempty"`
	Stats            DeviationStats   `json:"stats"`
	PublishedTime    time.Time        `json:"published_time"`
	Media            Media            `json:"media"`
	AllowsComments   bool             `json:"allows_comments,omitempty"`
	Preview          *Image           `json:"preview,omitempty"`
	Thumbnail        *Image           `json:"thumbnail,omitempty"`
	Video            *Video           `json:"video,omitempty"`
	DailyDeviation   *DailyDeviation  `json:"daily_deviation,omitempty"`
	IsDeleted        bool             `json:"is_deleted,omitempty"`
	IsBlocked        bool             `json:"is_blocked,omitempty"`
	DownloadFilesize int64            `json:"download_filesize,omitempty"`
	Tags             []string         `json:"tags,omitempty"`
	Category         string           `json:"category,omitempty"`
	Excerpt          string           `json:"excerpt,omitempty"`
	Suggestions      []string         `json:"suggestions,omitempty"`
	IsShareAvailable bool             `json:"is_share_available,omitempty"`
	IsEditable       bool             `json:"is_editable,omitempty"`
	HasComments      bool             `json:"has_comments,omitempty"`
	FileSize         int64            `json:"file_size,omitempty"`
	Resolution       Resolution       `json:"resolution,omitempty"`
	Submission       *Submission      `json:"submission,omitempty"`
	FullSize         *Image           `json:"full_size,omitempty"`
	FlashFile        *FlashFile       `json:"flash_file,omitempty"`
}

// DeviationStats contains deviation statistics
type DeviationStats struct {
	Comments    int `json:"comments"`
	Favourites  int `json:"favourites"`
	Views       int `json:"views"`
	Downloads   int `json:"downloads"`
}

// Media represents media information
type Media struct {
	BaseURI     string       `json:"base_uri"`
	PrettyName  string       `json:"pretty_name"`
	Types       []MediaType  `json:"types"`
	IsProcessed bool         `json:"is_processed"`
	Content     *Image       `json:"content,omitempty"`
	Token       string       `json:"token,omitempty"`
}

// MediaType represents a media type
type MediaType struct {
	T   string `json:"t"`
	CSS string `json:"css"`
}

// Image represents an image
type Image struct {
	Src           string `json:"src"`
	Height        int    `json:"height"`
	Width         int    `json:"width"`
	Transparency  bool   `json:"transparency"`
	Filesize      int64  `json:"filesize,omitempty"`
}

// Video represents video information
type Video struct {
	// Video-specific fields
}

// DailyDeviation represents daily deviation information
type DailyDeviation struct {
	Body    string   `json:"body"`
	Time    time.Time `json:"time"`
	Giver   User     `json:"giver"`
	Suggests []User  `json:"suggests"`
}

// Resolution represents image resolution
type Resolution struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Submission represents submission information
type Submission struct {
	CreationTime *time.Time     `json:"creation_time,omitempty"`
	Category     string         `json:"category,omitempty"`
	FileSize     int64          `json:"file_size,omitempty"`
	Resolution   Resolution     `json:"resolution,omitempty"`
	SubmittedWith *SubmittedWith `json:"submitted_with,omitempty"`
}

// SubmittedWith represents submission metadata
type SubmittedWith struct {
	App string `json:"app,omitempty"`
	URL string `json:"url,omitempty"`
}

// FlashFile represents flash file information
type FlashFile struct {
	// Flash-specific fields
}

// Comment represents a comment
type Comment struct {
	CommentID string     `json:"commentid"`
	ParentID  *string    `json:"parentid,omitempty"`
	Posted    time.Time  `json:"posted"`
	Hidden    string     `json:"hidden,omitempty"`
	Body      string     `json:"body"`
	User      User       `json:"user"`
	Replies   []Comment  `json:"replies,omitempty"`
	Stats      CommentStats `json:"stats"`
}

// CommentStats contains comment statistics
type CommentStats struct {
	Total  int `json:"total"`
	Hidden int `json:"hidden"`
}

// Status represents a status update
type Status struct {
	StatusID      string    `json:"statusid"`
	Body          string    `json:"body"`
	Timestamp     time.Time `json:"ts"`
	URL           string    `json:"url"`
	CommentsCount int       `json:"comments_count"`
	IsShare       bool      `json:"is_share"`
	IsEdited      bool      `json:"is_edited"`
	Mentions      []User    `json:"mentions,omitempty"`
	Media         []Media   `json:"media,omitempty"`
	Author        User      `json:"author"`
}

// Category represents a category
type Category struct {
	Catpath         string     `json:"catpath"`
	Title           string     `json:"title"`
	HasSubcategory  bool       `json:"has_subcategory"`
	Parents         []string   `json:"parents,omitempty"`
	Children        []Category `json:"children,omitempty"`
}

// Collection represents a collection folder
type Collection struct {
	CollectionID      string     `json:"collection_id"`
	Name              string     `json:"name"`
	Description       string     `json:"description,omitempty"`
	Size              int        `json:"size"`
	CoverDeviationID  *string    `json:"cover_deviation_id,omitempty"`
	CoverDeviation    *Deviation `json:"cover_deviation,omitempty"`
}

// GalleryFolder represents a gallery folder
type GalleryFolder struct {
	FolderID         string     `json:"folderid"`
	Name             string     `json:"name"`
	Parent           string     `json:"parent,omitempty"`
	Size             int        `json:"size"`
	CoverDeviationID string     `json:"cover_deviation_id,omitempty"`
	CoverDeviation   *Deviation `json:"cover_deviation,omitempty"`
}

// Pagination represents pagination information
type Pagination struct {
	HasMore      bool   `json:"has_more"`
	NextOffset   *int   `json:"next_offset,omitempty"`
	TotalResults int    `json:"total_results,omitempty"`
}

// PaginatedResponse wraps a paginated response
type PaginatedResponse[T any] struct {
	Pagination
	Results []T `json:"results"`
}

// Feedback represents feedback
type Feedback struct {
	FeedbackID  string     `json:"feedbackid"`
	Type        string     `json:"type"`
	Timestamp   time.Time  `json:"ts"`
	ByUser      User       `json:"by_user"`
	ByAnonymous bool       `json:"by_anonymous"`
	Message     string     `json:"message"`
	Deviation   *Deviation `json:"deviation,omitempty"`
	Comment     *Comment   `json:"comment,omitempty"`
	Status      *Status    `json:"status,omitempty"`
}

// Mention represents a mention
type Mention struct {
	MentionID string     `json:"mentionid"`
	Timestamp time.Time  `json:"ts"`
	ByUser    User       `json:"by_user"`
	Message   string     `json:"message"`
	Deviation *Deviation `json:"deviation,omitempty"`
	Comment   *Comment   `json:"comment,omitempty"`
	Status    *Status    `json:"status,omitempty"`
}

// Tier represents a subscription tier
type Tier struct {
	TierID      int    `json:"tierid"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
	Preview     Image  `json:"preview"`
	Color       string `json:"color"`
	IsAvailable bool   `json:"is_available"`
	Perks       []string `json:"perks"`
}

// UserTier represents a user's tier subscription
type UserTier struct {
	Tier     Tier       `json:"tier"`
	IsActive bool       `json:"is_active"`
	Start    *time.Time `json:"start,omitempty"`
	End      *time.Time `json:"end,omitempty"`
}

// ErrorResponse represents an API error
type ErrorResponse struct {
	Error             string `json:"error"`
	ErrorDescription  string `json:"error_description,omitempty"`
	Status            string `json:"status,omitempty"`
}

// Topic represents a hot topic
type Topic struct {
	Name    string `json:"name"`
	TopicID string `json:"topic_id"`
}

// Country represents a country
type Country struct {
	CountryID int    `json:"countryid"`
	Value     string `json:"value"`
}

// DeviationContent represents deviation content
type DeviationContent struct {
	DeviationID string   `json:"deviationid"`
	Download    bool     `json:"download"`
	Media       []Media  `json:"media"`
	HTML        string   `json:"html,omitempty"`
	CSS         string   `json:"css,omitempty"`
	CSSFonts    string   `json:"css_fonts,omitempty"`
}

// DeviationDownload represents download information
type DeviationDownload struct {
	Src      string `json:"src"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Filesize int64  `json:"filesize"`
}

// UserFriend represents a user friendship
type UserFriend struct {
	User            User       `json:"user"`
	IsFriend        bool       `json:"is_friend"`
	IsWatching      bool       `json:"is_watching,omitempty"`
	LastFriendTime  *time.Time `json:"last_friend_time,omitempty"`
	LastViewTime    *time.Time `json:"last_view_time,omitempty"`
}

// WhoFaved represents who favorited a deviation
type WhoFaved struct {
	DeviationID string     `json:"deviationid"`
	Title       string     `json:"title"`
	Time        time.Time  `json:"time"`
	Deviation   *Deviation `json:"deviation,omitempty"`
}

// DeviationMetadata represents extended deviation metadata
type DeviationMetadata struct {
	DeviationID string    `json:"deviationid"`
	Title       string    `json:"title"`
	Author      User      `json:"author"`
	IsFavourited bool     `json:"is_favourited"`
	IsMature     bool     `json:"is_mature"`
	IsDownloadable bool   `json:"is_downloadable"`
	AllowComments bool    `json:"allow_comments"`
	Stats        DeviationStats `json:"stats"`
	PublishedTime time.Time      `json:"published_time"`
	Media        Media     `json:"media"`
	Preview      *Image    `json:"preview,omitempty"`
	Thumbnail    *Image    `json:"thumbnail,omitempty"`
	Tags         []string  `json:"tags,omitempty"`
	Category     string    `json:"category,omitempty"`
	Excerpt      string    `json:"excerpt,omitempty"`
	IsShareAvailable bool   `json:"is_share_available,omitempty"`
	IsEditable   bool      `json:"is_editable,omitempty"`
	HasComments  bool      `json:"has_comments,omitempty"`
}

// DataCountries represents countries data
type DataCountries struct {
	Results []Country `json:"results"`
}

// DataText represents text-based data (privacy, submission, tos)
type DataText struct {
	Text string `json:"text"`
}

// FeedItem represents a feed item
type FeedItem struct {
	FeedID          string      `json:"feedid"`
	Type            string      `json:"type"`
	ByUser          User        `json:"by_user"`
	Deviations      []Deviation `json:"deviations,omitempty"`
	Status          *Status     `json:"status,omitempty"`
	Comment         *Comment    `json:"comment,omitempty"`
	CommentParent   *Comment    `json:"comment_parent,omitempty"`
	CommentDeviation *Deviation `json:"comment_deviation,omitempty"`
	CommentProfile  User        `json:"comment_profile,omitempty"`
	Timestamp       time.Time   `json:"ts"`
}

// FeedHome represents home feed
type FeedHome struct {
	Cursor  *string    `json:"cursor,omitempty"`
	HasMore bool       `json:"has_more"`
	Items   []FeedItem `json:"items"`
}

// FeedProfile represents profile feed
type FeedProfile struct {
	Cursor  *string    `json:"cursor,omitempty"`
	HasMore bool       `json:"has_more"`
	Items   []FeedItem `json:"items"`
}

// UserProfileResponse represents a user profile response
type UserProfileResponse struct {
	User       User           `json:"user"`
	ProfileURL string         `json:"profile_url"`
	Profile    ExtendedProfile `json:"profile"`
}

// ExtendedProfile contains extended profile information
type ExtendedProfile struct {
	UserTitle           string         `json:"user_title"`
	RealName            string         `json:"real_name"`
	ProfilePic          Image          `json:"profile_pic"`
	CoverPhoto          *Image         `json:"cover_photo"`
	Tagline             string         `json:"tagline"`
	Bio                 string         `json:"bio"`
	Website             string         `json:"website"`
	Location            string         `json:"location"`
	Joined              time.Time      `json:"joined"`
	Collections         []Collection   `json:"collections"`
	IsPremium           bool           `json:"is_premium"`
	IsArtist            bool           `json:"is_artist"`
	DefaultCollectionID int            `json:"default_collection_id"`
	ArtistLevel         string         `json:"artist_level"`
	ArtistSpecialty     string         `json:"artist_specialty"`
	HomePage            string         `json:"home_page"`
	LastStatus          *Status        `json:"last_status"`
}
