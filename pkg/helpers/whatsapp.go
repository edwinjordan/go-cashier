package helpers

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/jolebo/e-canteen-cashier-api/config"
)

/* rapiwha */
func SendWA(ctx context.Context, conf map[string]interface{}) {
	urlString := "https://panel.rapiwha.com/send_message.php"
	req, _ := http.NewRequestWithContext(ctx, "GET", urlString, nil)
	q := req.URL.Query()
	q.Add("apikey", config.GetEnv("APIWA"))
	q.Add("number", conf["phonenumber"].(string))
	q.Add("text", conf["text"].(string))
	req.URL.RawQuery = q.Encode()
	_, err := http.DefaultClient.Do(req)
	PanicIfError(err)
}

/* zenziva */

func SendWhatsapp(ctx context.Context, conf map[string]interface{}) *http.Response {
	urlString := "https://console.zenziva.net/wareguler/api/sendWA/"
	data := url.Values{}
	data.Set("userkey", config.GetEnv("ZENZIVA_USER"))
	data.Set("passkey", config.GetEnv("ZENZIVA_PASS"))
	data.Set("to", conf["phonenumber"].(string))
	data.Set("message", conf["text"].(string))
	req, _ := http.NewRequestWithContext(ctx, "POST", urlString, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	PanicIfError(err)

	return res
}
