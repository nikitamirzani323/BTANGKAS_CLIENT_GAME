package controller

import (
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/btangkas-client/config"
)

type responsechecktoken struct {
	Client_status    int    `json:"status"`
	Client_idcompany string `json:"client_idcompany"`
	Client_name      string `json:"client_name"`
	Client_username  string `json:"client_username"`
	Client_credit    int    `json:"client_credit"`
}
type responsesavetransaksi struct {
	Client_status      int    `json:"status"`
	Client_message     string `json:"message"`
	Client_idtransaksi string `json:"idtransaksi"`
	Client_card_game   string `json:"card_game"`
}
type responsedefault struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Record  interface{} `json:"record"`
}
type responseerror struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var PATH string = config.GetPathAPI()

func CheckToken(c *fiber.Ctx) error {
	type payload_checktoken struct {
		Client_token string `json:"client_token" `
	}
	hostname := c.Hostname()
	client := new(payload_checktoken)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsechecktoken{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"token": client.Client_token,
		}).
		Post(PATH + "api/checktoken")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsechecktoken)
	if result.Client_status == 200 {
		return c.JSON(fiber.Map{
			"status":          result.Client_status,
			"client_name":     result.Client_name,
			"client_username": result.Client_username,
			"client_credit":   result.Client_credit,
			"client_company":  result.Client_idcompany,
			"time":            time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func SaveTransaksi(c *fiber.Ctx) error {
	type payload_savetransaksi struct {
		Transaksi_company  string `json:"transaksi_company" `
		Transaksi_username string `json:"transaksi_username" `
		Transaksi_roundbet int    `json:"transaksi_roundbet" `
		Transaksi_bet      int    `json:"transaksi_bet" `
		Transaksi_cbefore  int    `json:"transaksi_cbefore" `
		Transaksi_cafter   int    `json:"transaksi_cafter" `
		Transaksi_win      int    `json:"transaksi_win" `
		Transaksi_idpoin   int    `json:"transaksi_idpoin" `
		Transaksi_status   string `json:"transaksi_status" `
	}
	hostname := c.Hostname()
	client := new(payload_savetransaksi)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsesavetransaksi{}).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"transaksi_company":  client.Transaksi_company,
			"transaksi_username": client.Transaksi_username,
			"transaksi_roundbet": client.Transaksi_roundbet,
			"transaksi_bet":      client.Transaksi_bet,
			"transaksi_cbefore":  client.Transaksi_cbefore,
			"transaksi_cafter":   client.Transaksi_cafter,
			"transaksi_win":      client.Transaksi_win,
			"transaksi_idpoin":   client.Transaksi_idpoin,
			"transaksi_status":   client.Transaksi_status,
		}).
		Post(PATH + "api/savetransaksi")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsesavetransaksi)
	if result.Client_status == 200 {
		return c.JSON(fiber.Map{
			"status":             result.Client_status,
			"client_idtransaksi": result.Client_idtransaksi,
			"client_cardgame":    result.Client_card_game,
			"time":               time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
