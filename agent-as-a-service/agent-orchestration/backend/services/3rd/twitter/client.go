package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/oauth1"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
)

const TWITTER_OAUTH_TOKEN_URL = "https://api.twitter.com/2/oauth2/token"

type Client struct {
	ConsumerKey, ConsumerSecret, AccessToken, AccessSecret string
	OauthClientId                                          string
	OauthClientSecret                                      string
	RedirectUri                                            string
	RedirectUriForInternal                                 string
	AppTokenForInternal                                    string
	AppToken                                               string
	user                                                   *User
}

func NewClient(appToken, consumerKey, consumerSecret, accessToken, accessSecret, oauthClientId, oauthClientSecret, redirectUri string) *Client {
	return &Client{
		AppToken:          appToken,
		ConsumerKey:       consumerKey,
		ConsumerSecret:    consumerSecret,
		AccessToken:       accessToken,
		AccessSecret:      accessSecret,
		OauthClientId:     oauthClientId,
		OauthClientSecret: oauthClientSecret,
		RedirectUri:       redirectUri,
	}
}

func (c *Client) postJSON(apiURL string, headers map[string]string, jsonObject map[string]interface{}, result interface{}) error {
	data := url.Values{}
	for key, val := range jsonObject {
		data.Set(key, fmt.Sprintf("%v", val))
	}
	encodedData := data.Encode()

	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(encodedData))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.SetBasicAuth(c.OauthClientId, c.OauthClientSecret)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *Client) postJSONWithKey(apiURL, oauthClientId, oauthClientSecret string, headers map[string]string, jsonObject map[string]interface{}, result interface{}) error {
	data := url.Values{}
	for key, val := range jsonObject {
		data.Set(key, fmt.Sprintf("%v", val))
	}
	encodedData := data.Encode()

	req, err := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(encodedData))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.SetBasicAuth(oauthClientId, oauthClientSecret)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *Client) getJSONWithAccessToken(accessToken, url string, headers map[string]string, result interface{}) (int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, accessToken))
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed request: %v", err)
	}
	if resp.StatusCode >= 300 {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, err.Error())
		}
		return resp.StatusCode, fmt.Errorf("http response bad status %d %s", resp.StatusCode, string(bodyBytes))
	}
	if result != nil {
		return resp.StatusCode, json.NewDecoder(resp.Body).Decode(result)
	}
	return resp.StatusCode, nil
}

type TwitterUser struct {
	TwitterID       string
	TwitterUsername string
	Name            string
	ProfileUrl      string `gorm:"type:text collate utf8mb4_unicode_ci"`
	FollowersCount  uint
	FollowingsCount uint
	IsBlueVerified  bool
	CreatedAt       time.Time
}

func (c *Client) GetTwitterUserInfoByToken(username string) (*TwitterUser, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/by/username/%s?user.fields=public_metrics,verified,profile_image_url,created_at,verified_type", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.AppToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info, status: %s, body: %s", resp.Status, body)
	}

	// Parse response into TwitterUser struct
	var responseData struct {
		Data struct {
			ID              string    `json:"id"`
			Username        string    `json:"username"`
			Name            string    `json:"name"`
			ProfileImageUrl string    `json:"profile_image_url"`
			Verified        bool      `json:"verified"`
			VerifiedType    string    `json:"verified_type"`
			CreatedAt       time.Time `json:"created_at"`
			PublicMetrics   struct {
				FollowersCount uint `json:"followers_count"`
				FollowingCount uint `json:"following_count"`
			} `json:"public_metrics"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	// Determine if the user is "blue verified"
	isBlueVerified := responseData.Data.VerifiedType == "blue"

	// Map response data to TwitterUser struct
	user := &TwitterUser{
		TwitterID:       responseData.Data.ID,
		TwitterUsername: responseData.Data.Username,
		Name:            responseData.Data.Name,
		ProfileUrl:      responseData.Data.ProfileImageUrl,
		FollowersCount:  responseData.Data.PublicMetrics.FollowersCount,
		FollowingsCount: responseData.Data.PublicMetrics.FollowingCount,
		IsBlueVerified:  isBlueVerified,
		CreatedAt:       responseData.Data.CreatedAt,
	}

	fmt.Println("User info retrieved successfully")
	return user, nil
}

func (c *Client) UploadImage(imageURL string, additionalOwners []string) (string, error) {
	if imageURL == "" {
		return "", nil
	}
	response, err := http.Get(imageURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	part, err := writer.CreateFormFile("media", filepath.Base(imageURL))
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, response.Body)
	if err != nil {
		return "", err
	}
	if len(additionalOwners) > 0 {
		owners := strings.Join(additionalOwners, ",")
		if err := writer.WriteField("additional_owners", owners); err != nil {
			return "", fmt.Errorf("failed to add additional_owners: %w", err)
		}
	}

	writer.Close()

	uploadURL := "https://upload.twitter.com/1.1/media/upload.json"
	req, err := http.NewRequest("POST", uploadURL, &requestBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	mediaID, ok := result["media_id_string"].(string)
	if !ok {
		return "", errs.ErrBadRequest
	}
	return mediaID, nil
}

func (c *Client) GetTwitterUserInfoID(userTwitterID string) (*TwitterUser, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s?user.fields=public_metrics,verified,profile_image_url,created_at,verified_type", userTwitterID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.AppToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info, status: %s, body: %s", resp.Status, body)
	}

	// Parse response into TwitterUser struct
	var responseData struct {
		Data struct {
			ID              string    `json:"id"`
			Username        string    `json:"username"`
			Name            string    `json:"name"`
			ProfileImageUrl string    `json:"profile_image_url"`
			Verified        bool      `json:"verified"`
			VerifiedType    string    `json:"verified_type"`
			CreatedAt       time.Time `json:"created_at"`
			PublicMetrics   struct {
				FollowersCount uint `json:"followers_count"`
				FollowingCount uint `json:"following_count"`
			} `json:"public_metrics"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	// Determine if the user is "blue verified"
	isBlueVerified := responseData.Data.VerifiedType == "blue"

	// Map response data to TwitterUser struct
	user := &TwitterUser{
		TwitterID:       responseData.Data.ID,
		TwitterUsername: responseData.Data.Username,
		Name:            responseData.Data.Name,
		ProfileUrl:      responseData.Data.ProfileImageUrl,
		FollowersCount:  responseData.Data.PublicMetrics.FollowersCount,
		FollowingsCount: responseData.Data.PublicMetrics.FollowingCount,
		IsBlueVerified:  isBlueVerified,
		CreatedAt:       responseData.Data.CreatedAt,
	}

	fmt.Println("User info retrieved successfully")
	return user, nil
}

func (c *Client) PostTweet(tweetContent, imageURL string, additionalOwners []string) error {
	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterApiUrl := "https://api.twitter.com/2/tweets"
	mediaID, err := c.UploadImage(imageURL, additionalOwners)
	if err != nil {
		return err
	}
	body := map[string]interface{}{
		"text": tweetContent,
	}
	if mediaID != "" {
		body["media"] = map[string]interface{}{
			"media_ids": []string{mediaID},
		}
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", twitterApiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		return nil
	} else {
		return errs.ErrBadRequest
	}
}

func (c *Client) ReplyToTweet(tweetID, replyContent, imageURL string, additionalOwners []string) error {
	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterApiUrl := "https://api.twitter.com/2/tweets"
	mediaID, err := c.UploadImage(imageURL, additionalOwners)
	if err != nil {
		return err
	}
	body := map[string]interface{}{
		"text": replyContent,
		"reply": map[string]interface{}{
			"in_reply_to_tweet_id": tweetID,
		},
	}
	if mediaID != "" {
		body["media"] = map[string]interface{}{
			"media_ids": []string{mediaID},
		}
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", twitterApiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return errs.ErrBadRequest
	}
	return nil
}

type TwitterTokenResponse struct {
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func (c *Client) GetTwitterOAuthToken(code string, callbackUrl, address, agentID string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	redirectUri := fmt.Sprintf(`%s?callback=%s&address=%s&agent_id=%s`, c.RedirectUri, callbackUrl, address, agentID)
	postData := map[string]interface{}{
		"client_id":     c.OauthClientId,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	err := c.postJSON(
		TWITTER_OAUTH_TOKEN_URL,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTwitterOAuthTokenWithKey(clientID, clientSecret, code string, callbackUrl, address, agentID string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	redirectUri := fmt.Sprintf(`%s?callback=%s&address=%s&agent_id=%s&client_id=%s`, c.RedirectUri, callbackUrl, address, agentID, clientID)
	postData := map[string]interface{}{
		"client_id":     c.OauthClientId,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	err := c.postJSONWithKey(
		TWITTER_OAUTH_TOKEN_URL, clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) TwitterOauthCallbackForInternalData(clientID, clientSecret, code, callbackUrl string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	redirectUri := fmt.Sprintf(`%s?callback=%s`, fmt.Sprintf("%s/internal", c.RedirectUri), callbackUrl)
	postData := map[string]interface{}{
		"client_id":     clientID,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	err := c.postJSONWithKey(
		TWITTER_OAUTH_TOKEN_URL, clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) TwitterOauthCallbackForSampleApp(clientID, clientSecret, code, redirectUri string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	postData := map[string]interface{}{
		"client_id":     clientID,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}
	err := c.postJSONWithKey(
		TWITTER_OAUTH_TOKEN_URL, clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTwitterOAuthTokenWithKeyForRelink(clientID, clientSecret, code string, callbackUrl, address string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	redirectUri := fmt.Sprintf(`%s?callback=%s&address=%s&client_id=%s`, c.RedirectUri, callbackUrl, address, clientID)
	postData := map[string]interface{}{
		"client_id":     c.OauthClientId,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	err := c.postJSONWithKey(
		TWITTER_OAUTH_TOKEN_URL, clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTwitterOAuthTokenWithKeyForCreateAgent(clientID, clientSecret, code string, callbackUrl, address string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	redirectUri := fmt.Sprintf(`%s?callback=%s&address=%s&agent_id=1&client_id=%s`, c.RedirectUri, callbackUrl, address, clientID)
	postData := map[string]interface{}{
		"client_id":     c.OauthClientId,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	err := c.postJSONWithKey(
		TWITTER_OAUTH_TOKEN_URL, clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTwitterOAuthTokenWithKeyForVideoReward(clientID, clientSecret, code string, callbackUrl, address string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	redirectUri := fmt.Sprintf(`%s?callback=%s&address=%s&agent_id=2&client_id=%s`, c.RedirectUri, callbackUrl, address, clientID)
	postData := map[string]interface{}{
		"client_id":     c.OauthClientId,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	err := c.postJSONWithKey(
		TWITTER_OAUTH_TOKEN_URL, clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTwitterOAuthTokenWithKeyForDeveloper(clientID, clientSecret, code string, callbackUrl, address string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	redirectUri := fmt.Sprintf(`%s?callback=%s&address=%s&agent_id=0&client_id=%s`, c.RedirectUri, callbackUrl, address, clientID)
	postData := map[string]interface{}{
		"client_id":     c.OauthClientId,
		"code_verifier": "challenge",
		"redirect_uri":  redirectUri,
		"grant_type":    "authorization_code",
		"code":          code,
	}

	err := c.postJSONWithKey(
		TWITTER_OAUTH_TOKEN_URL, clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type TwitterMeResp struct {
	UserObj UserObj `json:"data"`
}

func (c *Client) GetTwitterMe(accessToken string) (*UserObj, error) {
	var resp TwitterMeResp
	_, err := c.getJSONWithAccessToken(
		accessToken,
		"https://api.twitter.com/2/users/me?user.fields=profile_image_url,public_metrics,verified,description,entities",
		map[string]string{},
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp.UserObj, nil
}

func (c *Client) GetTwitterAccessToken(refreshToken string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	postData := map[string]interface{}{
		"client_id":     c.OauthClientId,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	}

	err := c.postJSON(
		"https://api.twitter.com/2/oauth2/token",
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetTwitterAccessTokenWithKey(clientID, clientSecret, refreshToken string) (*TwitterTokenResponse, error) {
	var resp TwitterTokenResponse
	postData := map[string]interface{}{
		"client_id":     clientID,
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
	}

	err := c.postJSONWithKey(
		"https://api.twitter.com/2/oauth2/token",
		clientID, clientSecret,
		map[string]string{},
		postData,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type Mention struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
}
type TweetResp struct {
	ID              string    `json:"id"`
	Text            string    `json:"text"`
	AuthorID        string    `json:"author_id"`
	AuthorUsername  string    `json:"author_username,omitempty"`
	AuthorName      string    `json:"author_name,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	Language        string    `json:"lang"`
	LikeCount       int       `json:"like_count"`
	RetweetCount    int       `json:"retweet_count"`
	ReplyCount      int       `json:"reply_count"`
	QuoteCount      int       `json:"quote_count"`
	ImpressionCount int       `json:"impression_count"`
	InReplyToUserID string    `json:"in_reply_to_user_id,omitempty"`
	FollowersCount  uint      `json:"followers_count,omitempty"`

	// Reply-related fields
	IsReply          bool   `json:"is_reply"`
	InReplyToTweetID string `json:"in_reply_to_tweet_id,omitempty"`
	OriginalText     string `json:"original_text,omitempty"`

	// Retweet-related fields
	IsRetweet     bool   `json:"is_retweet"`
	RepostTweetID string `json:"repost_tweet_id,omitempty"`
	RepostText    string `json:"repost_text,omitempty"`

	// Quote-related fields
	IsQuote      bool   `json:"is_quote"`
	QuoteTweetID string `json:"quote_tweet_id,omitempty"`
	QuoteText    string `json:"quote_text,omitempty"`

	// New field for mentions
	Mentions []Mention `json:"mentions,omitempty"`

	ParentTweetID        string
	ParentAuthorUsername string `json:"parent_author_username,omitempty"`
	ParentAuthorID       string `json:"parent_author_id,omitempty"`
}

func (m *TweetResp) IsMention(username string) bool {
	for _, v := range m.Mentions {
		if strings.EqualFold(v.Username, username) {
			return true
		}
	}
	return false
}

func (c *Client) GetTweetDetails(tweetIDs []string) (map[string]*TweetResp, error) {
	lookupURL := "https://api.twitter.com/2/tweets"
	params := url.Values{}
	params.Add("ids", strings.Join(tweetIDs, ","))
	params.Add("tweet.fields", "id,text,author_id,created_at,lang,public_metrics,in_reply_to_user_id,referenced_tweets,entities")
	params.Add("expansions", "author_id")                     // Include author_id expansion
	params.Add("user.fields", "username,name,public_metrics") // Fetch username, name, and public_metrics

	req, err := http.NewRequest("GET", lookupURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AppToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch tweets, status: %s", resp.Status)
	}

	var response struct {
		Data []struct {
			ID            string `json:"id"`
			Text          string `json:"text"`
			AuthorID      string `json:"author_id"`
			CreatedAt     string `json:"created_at"`
			Lang          string `json:"lang"`
			PublicMetrics struct {
				LikeCount       int `json:"like_count"`
				RetweetCount    int `json:"retweet_count"`
				ReplyCount      int `json:"reply_count"`
				QuoteCount      int `json:"quote_count"`
				ImpressionCount int `json:"impression_count,omitempty"`
			} `json:"public_metrics"`
			InReplyToUserID  string `json:"in_reply_to_user_id,omitempty"`
			ReferencedTweets []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"referenced_tweets,omitempty"`
			Entities struct {
				Mentions []struct {
					Username string `json:"username"`
					ID       string `json:"id"`
				} `json:"mentions,omitempty"`
			} `json:"entities,omitempty"`
		} `json:"data"`
		Includes struct {
			Users []struct {
				ID            string `json:"id"`
				Username      string `json:"username"`
				Name          string `json:"name"`
				PublicMetrics struct {
					FollowersCount int `json:"followers_count"`
				} `json:"public_metrics"`
			} `json:"users"`
		} `json:"includes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	// Create a map to access user information by ID
	usersMap := make(map[string]struct {
		Username       string
		Name           string
		FollowersCount int
	})
	for _, user := range response.Includes.Users {
		usersMap[user.ID] = struct {
			Username       string
			Name           string
			FollowersCount int
		}{
			Username:       user.Username,
			Name:           user.Name,
			FollowersCount: user.PublicMetrics.FollowersCount, // Capture followers_count
		}
	}

	// Collect referenced tweet IDs and initialize Tweet structs
	referencedTweetIDs := map[string]bool{}
	tweets := make([]TweetResp, len(response.Data))

	for i, t := range response.Data {
		createdAt, err := time.Parse(time.RFC3339, t.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error parsing created_at: %w", err)
		}

		// Parse mentions
		mentions := make([]Mention, len(t.Entities.Mentions))
		for j, mention := range t.Entities.Mentions {
			mentions[j] = Mention{
				Username: mention.Username,
				UserID:   mention.ID,
			}
		}

		// Get author information from the users map
		authorInfo := usersMap[t.AuthorID]

		tweet := TweetResp{
			ID:              t.ID,
			Text:            t.Text,
			AuthorID:        t.AuthorID,
			AuthorUsername:  authorInfo.Username,
			AuthorName:      authorInfo.Name,
			FollowersCount:  uint(authorInfo.FollowersCount), // Add followers_count
			CreatedAt:       createdAt,
			Language:        t.Lang,
			LikeCount:       t.PublicMetrics.LikeCount,
			RetweetCount:    t.PublicMetrics.RetweetCount,
			ReplyCount:      t.PublicMetrics.ReplyCount,
			QuoteCount:      t.PublicMetrics.QuoteCount,
			ImpressionCount: t.PublicMetrics.ImpressionCount,
			InReplyToUserID: t.InReplyToUserID,
			IsReply:         false,
			IsRetweet:       false,
			IsQuote:         false,
			Mentions:        mentions,
		}

		// Determine if tweet is a reply, retweet, or quote and collect IDs
		for _, refTweet := range t.ReferencedTweets {
			switch refTweet.Type {
			case "replied_to":
				tweet.IsReply = true
				tweet.InReplyToTweetID = refTweet.ID
				referencedTweetIDs[refTweet.ID] = true
				tweet.ParentTweetID = refTweet.ID
			case "retweeted":
				tweet.IsRetweet = true
				tweet.RepostTweetID = refTweet.ID
				referencedTweetIDs[refTweet.ID] = true
				tweet.ParentTweetID = refTweet.ID
			case "quoted":
				tweet.IsQuote = true
				tweet.QuoteTweetID = refTweet.ID
				referencedTweetIDs[refTweet.ID] = true
				tweet.ParentTweetID = refTweet.ID
			}
		}

		tweets[i] = tweet
	}

	// Fetch referenced tweets if needed and add their texts
	originalTweets, err := c.FetchReferencedTweetByIDs(referencedTweetIDs)
	if err != nil {
		return nil, err
	}

	for i, tweet := range tweets {
		if tweet.IsReply {
			tmp := originalTweets[tweet.InReplyToTweetID]
			if tmp != nil {
				tweets[i].OriginalText = tmp.Text
				tweets[i].ParentAuthorUsername = tmp.AuthorUsername
				tweets[i].ParentAuthorID = tmp.AuthorID
			}
		}
		if tweet.IsRetweet {
			tmp := originalTweets[tweet.RepostTweetID]
			if tmp != nil {
				tweets[i].RepostText = tmp.Text
				tweets[i].ParentAuthorUsername = tmp.AuthorUsername
				tweets[i].ParentAuthorID = tmp.AuthorID
			}

		}
		if tweet.IsQuote {
			tmp := originalTweets[tweet.QuoteTweetID]
			if tmp != nil {
				tweets[i].QuoteText = tmp.Text
				tweets[i].ParentAuthorUsername = tmp.AuthorUsername
				tweets[i].ParentAuthorID = tmp.AuthorID
			}
		}
	}

	result := map[string]*TweetResp{}
	for _, v := range tweets {
		result[v.ID] = &v
	}
	return result, nil
}
func (c *Client) FetchReferencedTweetByIDs(tweetIDMap map[string]bool) (map[string]*TweetResp, error) {
	// Convert map keys to slice of IDs
	tweetIDs := []string{}
	for id := range tweetIDMap {
		tweetIDs = append(tweetIDs, id)
	}

	// Initialize result map
	result := make(map[string]*TweetResp)

	// Batch processing (max 100 per batch)
	const batchSize = 100
	for i := 0; i < len(tweetIDs); i += batchSize {
		end := i + batchSize
		if end > len(tweetIDs) {
			end = len(tweetIDs)
		}
		batch := tweetIDs[i:end]

		// Fetch tweets for the current batch
		tweetMap, err := c._fetchReferencedTweetByIDs(batch)
		if err != nil {
			return nil, err // Return error immediately if any batch fails
		}

		// Merge current batch result into the final result map
		for id, tweet := range tweetMap {
			result[id] = tweet
		}
	}

	return result, nil
}

func (c *Client) _fetchReferencedTweetByIDs(tweetIDs []string) (map[string]*TweetResp, error) {

	lookupURL := "https://api.twitter.com/2/tweets"
	params := url.Values{}
	params.Add("ids", strings.Join(tweetIDs, ","))
	params.Add("tweet.fields", "id,text,author_id")
	params.Add("expansions", "author_id")      // Include author_id expansion
	params.Add("user.fields", "username,name") // Include public_metrics to get followers_count

	req, err := http.NewRequest("GET", lookupURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AppToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch referenced tweets, status: %s", resp.Status)
	}

	var response struct {
		Data []struct {
			ID       string `json:"id"`
			Text     string `json:"text"`
			AuthorID string `json:"author_id"`
		} `json:"data"`
		Includes struct {
			Users []struct {
				ID       string `json:"id"`
				Username string `json:"username"`
				Name     string `json:"name"`
			} `json:"users"`
		} `json:"includes"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	// Create a map to access user information by ID
	usersMap := make(map[string]struct {
		Username string
		Name     string
	})
	for _, user := range response.Includes.Users {
		usersMap[user.ID] = struct {
			Username string
			Name     string
		}{
			Username: user.Username,
			Name:     user.Name,
		}
	}

	// Map each tweet ID to its text
	tweetMap := map[string]*TweetResp{}
	for _, tweet := range response.Data {
		// Get author information from the users map
		authorInfo := usersMap[tweet.AuthorID]

		tweetMap[tweet.ID] = &TweetResp{
			ID:             tweet.ID,
			Text:           tweet.Text,
			AuthorID:       tweet.AuthorID,
			AuthorUsername: authorInfo.Username,
		}
	}

	return tweetMap, nil
}

type TweetFullTextResponse struct {
	Data struct {
		ID                  string   `json:"id"`
		Text                string   `json:"text"`
		EditHistoryTweetIDs []string `json:"edit_history_tweet_ids"`
		NoteTweet           struct {
			Text string `json:"text"`
		} `json:"note_tweet"`
	} `json:"data"`
}

func (c *Client) GetTweetFullText(tweetID string) (*TweetFullTextResponse, error) {
	lookupURL := fmt.Sprintf("https://api.x.com/2/tweets/%s", tweetID)
	params := url.Values{}
	params.Add("expansions", "entities.mentions.username,entities.note.mentions.username")
	params.Add("tweet.fields", "text,note_tweet")

	req, err := http.NewRequest("GET", lookupURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.AppToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch tweet, status: %s", resp.Status)
	}

	var tweetResponse TweetFullTextResponse
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tweetResponse)
	if err != nil {
		return nil, err
	}

	return &tweetResponse, nil
}

func (c *Client) CheckIfTweetIsTruncated(tweetText string) bool {
	re := regexp.MustCompile(`https://t.co/[a-zA-Z0-9]+`)
	return re.MatchString(tweetText) && len(tweetText) > 270
}

func (c *Client) GetTweetsByUser(userID string, cursor string) (map[string]*TweetResp, string, error) {
	baseURL := "https://api.twitter.com/2/users/%s/tweets"
	url := fmt.Sprintf(baseURL, userID) +
		"?max_results=100" +
		"&tweet.fields=public_metrics,referenced_tweets,in_reply_to_user_id,text" +
		"&expansions=referenced_tweets.id"
	if cursor != "" {
		url += "&pagination_token=" + cursor
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+c.AppToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, "", nil
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, "", nil
	}

	tweetIDs := c.extractTweetIDs(result)
	nextToken := c.extractNextToken(result)
	mapTweet, err := c.GetTweetDetails(tweetIDs)
	if err != nil {
		return nil, "", nil
	}
	return mapTweet, nextToken, nil
}

func (c *Client) extractTweetIDs(data map[string]interface{}) []string {
	var tweetIDs []string

	if tweets, exists := data["data"].([]interface{}); exists {
		for _, t := range tweets {
			tweet := t.(map[string]interface{})
			if id, ok := tweet["id"].(string); ok {
				tweetIDs = append(tweetIDs, id)
			}
		}
	}

	return tweetIDs
}

func (c *Client) extractNextToken(data map[string]interface{}) string {
	if meta, exists := data["meta"].(map[string]interface{}); exists {
		if nextToken, ok := meta["next_token"].(string); ok {
			return nextToken
		}
	}
	return ""
}

func (c *Client) PostTweetWithVideo(tweetContent, videoURL string, additionalOwners []string) error {
	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterApiUrl := "https://api.twitter.com/2/tweets"
	mediaID, err := c.UploadVideo(videoURL, additionalOwners)
	if err != nil {
		return err
	}
	body := map[string]interface{}{
		"text": tweetContent,
	}
	if mediaID != "" {
		body["media"] = map[string]interface{}{
			"media_ids": []string{mediaID},
		}
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", twitterApiUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		return nil
	} else {
		return errs.ErrBadRequest
	}
}
