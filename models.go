package smsaero

import "time"

type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type AuthResponse struct {
	Status
	Data interface{} `json:"data"`
}

type SendResponse struct {
	Status
	Data interface{} `json:"data"`
}

type SmsMessage struct {
	ID           int64   `json:"id"`
	From         string  `json:"from"`
	Number       string  `json:"number"`
	Text         string  `json:"text"`
	Status       int64   `json:"status"`
	ExtendStatus string  `json:"extendStatus"`
	Channel      string  `json:"channel"`
	Cost         float64 `json:"cost"`
	DateCreate   int64   `json:"dateCreate"`
	DateSend     int64   `json:"dateSend"`
}

type MessageRequest struct {
	Numbers     []string   // numbers array Обязательно (на выбор)	Номера телефонов
	Sign        string     // sign	string	Обязательно	Подпись отправителя
	Text        string     // text	string	Обязательно	Текст сообщения
	Channel     string     // channel	string	Обязательно	Канал отправки
	DateSend    *time.Time // dateSend	integer	Не обязательно	Дата для отложенной отправки сообщения (в формате unixtime)
	CallbackUrl string     // callbackUrl	string	Не обязательно	url для отправки статуса сообщения в формате http://your.site или https://your.site (в ответ система ждет ста
}
