package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func SplitTextByCharLimit(text string, limit int) []string {
	var result []string
	words := strings.Fields(text)
	var line string
	for _, word := range words {
		if len(line)+len(word)+1 > limit {
			result = append(result, line)
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}
	if line != "" {
		result = append(result, line)
	}
	return result
}

func RemoveTrailingHashTag(line string) string {
	line = strings.TrimSpace(line)
	words := strings.Fields(line)
	var isOk bool
	hashtags := []string{}
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		if !strings.HasPrefix(word, "#") {
			isOk = true
		}
		if isOk {
		} else {
			hashtags = append(hashtags, word)
		}
	}
	for _, v := range hashtags {
		line = strings.TrimSpace(strings.TrimSuffix(line, v))
	}
	return line
}

// SplitTextBySentenceAndCharLimit splits text into lines by sentence boundaries and a character limit per line.
func SplitTextBySentenceAndCharLimitAndRemoveTrailingHashTag(originText string, limit int) []string {
	limit = limit - 1
	// Split the text into sentences based on the period "."
	originText = RemoveTrailingHashTag(strings.TrimSpace(originText))
	if len(originText) <= limit {
		if strings.HasPrefix(originText, "@") {
			originText = "." + originText
		}
		return []string{originText}
	}
	text := strings.TrimSuffix(originText, ".")
	strs := strings.Split(text, ". ")
	var sentences []string
	for _, str := range strs {
		sentences = append(sentences, strings.Split(str, ".\n")...)
	}
	var result []string
	var line string
	for _, sentence := range sentences {
		// Trim any leading or trailing spaces in the sentence
		sentence = strings.TrimSpace(sentence)
		if sentence == "" {
			continue
		}
		// If the sentence itself exceeds the limit, split it further by words
		if len(sentence) > limit {
			words := strings.Fields(sentence)
			var subLine string

			for _, word := range words {
				// If adding the word exceeds the limit, add the current subLine to result and start a new subLine
				if len(subLine)+len(word)+1 > limit {
					result = append(result, subLine)
					subLine = word
				} else {
					// Otherwise, add the word to the current subLine
					if subLine != "" {
						subLine += " "
					}
					subLine += word
				}
			}
			// Append the remaining subLine if it's not empty
			if subLine != "" {
				result = append(result, subLine)
			}
		} else {
			// If the sentence is within the limit, try to add it to the current line
			if len(line)+len(sentence)+2 > limit { // +2 for the period and space
				result = append(result, line+".")
				line = sentence
			} else {
				if line != "" {
					line += ". "
				}
				line += sentence
			}
		}
	}
	// Append the last line if there's any content left
	if line != "" {
		if strings.HasSuffix(originText, ".") {
			result = append(result, line+".")
		} else {
			result = append(result, line)
		}
	}
	if len(result) > 0 {
		if strings.HasPrefix(result[0], "@") {
			result[0] = "." + result[0]
		}
	}
	return result
}

func ReplyTweetByToken(bearerToken string, replyContent string, tweetID, mediaID string) (string, error) {
	url := "https://api.twitter.com/2/tweets"
	payload := map[string]interface{}{
		"text": replyContent,
		"reply": map[string]string{
			"in_reply_to_tweet_id": tweetID,
		},
	}
	if mediaID != "" {
		payload["media"] = map[string]interface{}{
			"media_ids": []string{mediaID},
		}
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error encoding payload: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to reply tweet, status: %s, body: %s", resp.Status, body)
	}
	// Parse response to extract the reply ID
	var responseData struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}
	fmt.Println("Reply posted successfully, ID:", responseData.Data.ID)
	return responseData.Data.ID, nil
}

func QuoteTweetByToken(bearerToken string, content string, quoteTweetId string) (string, error) {
	url := "https://api.twitter.com/2/tweets"
	payload := map[string]interface{}{
		"text":           content,
		"quote_tweet_id": quoteTweetId,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error encoding payload: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to reply tweet, status: %s, body: %s", resp.Status, body)
	}
	// Parse response to extract the reply ID
	var responseData struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}
	fmt.Println("Reply posted successfully, ID:", responseData.Data.ID)
	return responseData.Data.ID, nil
}

func RepostTweetByToken(bearerToken string, userID string, tweetID string) (string, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/retweets", userID)
	payload := map[string]string{
		"tweet_id": tweetID,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error encoding payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to repost tweet, status: %s, body: %s", resp.Status, body)
	}

	// Parse response to confirm repost success or get additional info if needed
	var responseData struct {
		Data struct {
			ID string `json:"rest_id"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	fmt.Println("Repost successful, ID:", responseData.Data.ID)
	return responseData.Data.ID, nil
}

func PostTweetByToken(accessToken, message string, replyID string) (string, error) {
	url := "https://api.twitter.com/2/tweets"
	payload := map[string]interface{}{
		"text": message,
	}
	if replyID != "" {
		payload["reply"] = map[string]string{
			"in_reply_to_tweet_id": replyID,
		}
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error encoding payload: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to post tweet, status: %s, body: %s", resp.Status, body)
	}
	// Parse response to extract the tweet ID
	var responseData struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}
	fmt.Println("Tweet posted successfully, ID:", responseData.Data.ID)
	return responseData.Data.ID, nil
}

type TwitterNotification struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

func GetTwitterNotifications(accessToken, userID string) (*TwitterNotification, error) {
	twitterApiUrl := fmt.Sprintf("https://api.twitter.com/2/users/%s/mentions?max_results=5", userID)
	req, err := http.NewRequest("GET", twitterApiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("StatusCode not 200")
	}
	var respData TwitterNotification
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

type TwitterUserMe struct {
	Data *struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		UserName        string `json:"username"`
		CreatedAt       string `json:"created_at"`
		Description     string `json:"description"`
		Location        string `json:"location"`
		PinnedTweetID   string `json:"pinned_tweet_id"`
		ProfileImageURL string `json:"profile_image_url"`
		Protected       bool   `json:"protected"`
		URL             string `json:"url"`
		Verified        bool   `json:"verified"`
	} `json:"data"`
}

func GetTwitterUserMe(accessToken string) (*TwitterUserMe, error) {
	twitterApiUrl := "https://api.twitter.com/2/users/me?user.fields=username,name,verified,verified_type,profile_image_url"
	req, err := http.NewRequest("GET", twitterApiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, nil
	}
	var respData TwitterUserMe
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

func TwitterFollowUserCreate(accessToken, twitterId string, targetTwitterId string) error {
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/following", twitterId)
	payload := map[string]string{
		"target_user_id": targetTwitterId,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error encoding payload: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to post tweet, status: %s, body: %s", resp.Status, body)
	}
	var responseData struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	return nil
}

// Response structure for liked tweets
type LikedTweetsResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Meta struct {
		NextToken string `json:"next_token"`
	} `json:"meta"`
}

// Function to fetch liked tweets
func GetLikedTweets(userID string, bearerToken string, nextToken string) (*LikedTweetsResponse, error) {
	baseURL := "https://api.twitter.com/2/users/" + userID + "/liked_tweets"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	// Add pagination if needed
	query := req.URL.Query()
	query.Add("max_results", "100")
	if nextToken != "" {
		query.Add("pagination_token", nextToken)
	}
	req.URL.RawQuery = query.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch liked tweets: %s", resp.Status)
	}

	var likedTweets LikedTweetsResponse
	if err := json.NewDecoder(resp.Body).Decode(&likedTweets); err != nil {
		return nil, err
	}

	return &likedTweets, nil
}

func PostTweetWithMediaByToken(accessToken, message string, mediaID string) (string, error) {
	url := "https://api.twitter.com/2/tweets"
	payload := map[string]interface{}{
		"text": message,
	}
	if mediaID != "" {
		payload["media"] = map[string]interface{}{
			"media_ids": []string{mediaID},
		}
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error encoding payload: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to post tweet, status: %s, body: %s", resp.Status, body)
	}
	// Parse response to extract the tweet ID
	var responseData struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}
	fmt.Println("Tweet posted successfully, ID:", responseData.Data.ID)
	return responseData.Data.ID, nil
}

func PostThreadTweetByToken(accessToken string, messages []string) (string, error) {
	if len(messages) == 0 {
		return "", nil
	}
	parentTweetID, err := PostTweetByToken(accessToken, messages[0], "")
	if err != nil {
		return "", err
	}
	if parentTweetID != "" {
		tweetID := parentTweetID
		for i := 1; i < len(messages); i++ {
			tweetID, err = PostTweetByToken(accessToken, fmt.Sprintf("#%d: %s", i, messages[i]), tweetID)
			if err != nil {
				return "", err
			}
			if tweetID == "" {
				return parentTweetID, nil
			}
			time.Sleep(1 * time.Second)
		}
	}
	return parentTweetID, err
}

// dm
func GetListDmRecent(accessToken string) error {

	return nil
}

func GetDmByID(accessToken string) error {

	return nil
}

func ReplyDm(accessToken string) error {

	return nil
}

//
