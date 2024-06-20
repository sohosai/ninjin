package slack

import (
	"fmt"
	"ninjin/util/cls"

	slackgo "github.com/nlopes/slack"
)

func (sl SlackUtil) AttachMessageInfo(msg *cls.Message, data map[string]interface{}) error {
	msg.Content = data["text"].(string)
	msg.Slack_ID = data["ts"].(string)
	msg.ChannelName = sl.GetChannelNameByID(data["channel"].(string))
	msg.FileURL = ""
	if files, ok := data["files"].([]interface{}); ok {
		for _, file := range files {
			if fileMap, ok := file.(map[string]interface{}); ok {
				if thumbs, ok := fileMap["thumb_360"]; ok {
					msg.FileURL = thumbs.(string)
				}
				if other, ok := fileMap["url_private"]; ok {
					msg.FileURL = other.(string)
				}
			}
		}
	}
	msg.FileName = ""
	if files, ok := data["files"].([]interface{}); ok {
		for _, file := range files {
			if fileMap, ok := file.(map[string]interface{}); ok {
				if name, ok := fileMap["name"]; ok {
					msg.FileName = name.(string)
				}
			}
		}
	}

	return nil
}

func (sl SlackUtil) GetChannelNameByID(channelID string) string {
	api := slackgo.New(sl.SLACK_API_TOKEN)
	channelinfo, err := api.GetConversationInfo(channelID, false)
	if err != nil {
		fmt.Println("error GetChannelInfo : ", err)
		return ""
	}
	return channelinfo.Name
}