package main

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

const Message = "ВНИМАНИЕ!\n" +
	"Переходи по [ссылке]() в новый маттермост\n" +
	"Это простанство заморожено."

func New() *Plugin {
	return &Plugin{}
}

type Plugin struct {
	plugin.MattermostPlugin
}

func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	channel, er1 := p.API.GetChannel(post.ChannelId)
	if er1 == nil {
		if channel.TeamId == "qy68iknxw7bopqqn1bbkswr4yy" {
			usr, er2 := p.API.GetUser(post.UserId)
			if er2 == nil {
				if usr.IsBot == false {
					post.Message = Message
				}
			}
		}
	}

	return post, ""
}
