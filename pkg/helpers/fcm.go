package helpers

import (
	"github.com/appleboy/go-fcm"
	"github.com/jolebo/e-canteen-cashier-api/config"
)

/**
* Data: conf["data"].(map[string]interface{})
* Notificatoon
**/
func SendFCM(sentTo []string, conf map[string]interface{}) *fcm.Response {
	msg := &fcm.Message{
		RegistrationIDs: sentTo,
		Data:            conf["data"].(map[string]interface{}),
		Priority:        "high",
		Notification: &fcm.Notification{
			Title:       conf["title"].(string),
			Body:        conf["body"].(string),
			ChannelID:   "ecanteen_notif_1",
			Icon:        "ic_stat_onesignal_default",
			ClickAction: "FCM_PLUGIN_ACTIVITY",
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient(config.GetEnv("FCM_SERVER_KEY"))
	if err != nil {
		PanicIfError(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		PanicIfError(err)
	}
	return response
}
