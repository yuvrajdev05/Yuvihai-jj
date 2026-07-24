/*
 * ○ A high-performance engine for streaming music in Telegram voicechats.
 *
 * Copyright (C) 2026 Team Arc
 */

package modules

import (
	"time"
	"github.com/Laky-64/gologging"
	tg "github.com/amarnathcjd/gogram/telegram"

	"main/internal/config"
	"main/internal/core"
	"main/internal/database"
	"main/internal/locales"
	"main/internal/utils"
)

func init() {
	helpTexts["/start"] = `<i>Start the bot and show main menu.</i>`
}

func startHandler(m *tg.NewMessage) error {
	if m.ChatType() != tg.EntityUser {
		database.AddServedChat(m.ChannelID())
		m.Reply(
			F(m.ChannelID(), "start_group"),
		)
		return tg.ErrEndGroup
	}

	arg := m.Args()
	database.AddServedUser(m.ChannelID())

	if arg != "" {
		gologging.Info(
			"Got Start parameter: " + arg + " in ChatID: " + utils.IntToStr(
				m.ChannelID(),
			),
		)
	}

	switch arg {
	case "pm_help":
		gologging.Info("User requested help via start param")
		helpHandler(m)

		default:	             

         msg1, _ := m.Respond("🌟 Welcome To Inaya Music 🌟", nil)
         time.Sleep(1 * time.Second)
         if msg1 != nil {
	     msg1.Delete()
         }

         msg2, _ := m.Respond("🎵 Best Music Experience 🎵", nil)
         time.Sleep(1 * time.Second)
         if msg2 != nil {
	     msg2.Delete()
         }

         msg3, _ := m.Respond("💖 High Quality Streaming 💖", nil)
         time.Sleep(1 * time.Second)
         if msg3 != nil {
	     msg3.Delete()
         }

         msg4, _ := m.Respond(
	     `✨ Powered By <a href="https://t.me/x_yuvii">Yuvi</a> ✨`,
	     &tg.SendOptions{
		 ParseMode: "HTML",
	     },
         )
         time.Sleep(1 * time.Second)
         if msg4 != nil {
	    msg4.Delete()
        }


        _, _ = m.RespondSticker(
	    "CAACAgUAAxkBAAEg1G5qY4WlzbzpH0Np-r2IyhRtREm7_wAC7BYAAj-FsVWpaDfP5N7RAj0E",
	    nil,
        )

        time.Sleep(1 * time.Second)
		caption := F(m.ChannelID(), "start_private", locales.Arg{
			"user": utils.MentionHTML(m.Sender),
			"bot":  utils.MentionHTML(m.Client.Me()),
		})

		_, err := m.RespondMedia(&tg.InputMediaWebPage{
			URL:             config.StartImage,
			ForceLargeMedia: true,
		}, &tg.MediaOptions{
			Caption:     caption,
			NoForwards:  true,
			ReplyMarkup: core.GetStartMarkup(m.ChannelID()),
		})
		if err != nil {
			gologging.Error(
				"[start] InputMediaWebPage Reply failed: " + err.Error(),
			)

			_, err = m.RespondMedia(config.StartImage, &tg.MediaOptions{
				Caption:     caption,
				NoForwards:  true,
				ReplyMarkup: core.GetStartMarkup(m.ChannelID()),
			})
			if err != nil {
				gologging.Error(
					"[start] URL media reply failed: " + err.Error(),
				)

				_, err = m.Respond(caption, &tg.SendOptions{
					NoForwards:  true,
					ReplyMarkup: core.GetStartMarkup(m.ChannelID()),
				})
				return err
			}
		}
	}

	if config.LoggerID != 0 && isLoggerEnabled() {
		uName := "N/A"
		if m.Sender.Username != "" {
			uName = "@" + m.Sender.Username
		}
		msg := F(m.ChannelID(), "logger_bot_started", locales.Arg{
			"mention":       utils.MentionHTML(m.Sender),
			"user_id":       m.SenderID(),
			"user_username": uName,
		})
		_, err := m.Client.SendMessage(config.LoggerID, msg)
		if err != nil {
			gologging.Error(
				"Failed to send logger_bot_started msg, Err: " + err.Error(),
			)
		}
	}
	return tg.ErrEndGroup
}

func startCB(cb *tg.CallbackQuery) error {
	cb.Answer("")

	caption := F(cb.ChannelID(), "start_private", locales.Arg{
		"user": utils.MentionHTML(cb.Sender),
		"bot":  utils.MentionHTML(cb.Client.Me()),
	})

	sendOpt := &tg.SendOptions{
		ReplyMarkup: core.GetStartMarkup(cb.ChannelID()),
		NoForwards:  true,
	}

	if config.StartImage != "" {
		sendOpt.Media = &tg.InputMediaWebPage{
			URL:             config.StartImage,
			ForceLargeMedia: true,
		}
	}

	cb.Edit(caption, sendOpt)
	return tg.ErrEndGroup
}
