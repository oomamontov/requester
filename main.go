package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	url = "https://google.com/"
)

var (
	delay  = 34300 * time.Millisecond // Абсолютно случайное значение
	client = http.DefaultClient
)

func Log(message string) {
	fmt.Printf("{%v}: %s\n", time.Now().Format("02.01.2006 15:04:05.000000000"), message)
}

func Loop() {
	start := time.Now()
	Log("making request")
	resp, err := client.Get(url)
	if err != nil {
		Log(fmt.Sprintf("got error: %w", err))
		return
	}
	defer resp.Body.Close()

	if _, err := io.Copy(io.Discard, resp.Body); err != nil { // Вычитываем все данные, чтобы можно было переиспользовать соединение
		Log(fmt.Sprintf("omfg: %w", err)) // Этого никогда не должно случиться
		panic(err) // На всякий случай убиваем приложуху
	}

	Log(fmt.Sprintf("request done in %v", time.Now().Sub(start)))
}

func main() {
	ticker := time.NewTicker(delay)
	go Loop()
	for range ticker.C {
		go Loop()
	}
}
