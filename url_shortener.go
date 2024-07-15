package main

import (
	"fmt"
	"math"
	_ "strconv"
	"time"
)

type Data struct {
	time int64
	url  string
}

type UrlShortener struct {
	counter int
	number  int
	data    map[string]Data
	urlMap  map[string]string
}

func NewUrlShortener() *UrlShortener {
	minNumber := math.Pow(10, SIZE)
	maxNumber := math.Pow(10, SIZE+1)
	randomNumber := random(int(minNumber), int(maxNumber))
	return &UrlShortener{counter: 0, number: randomNumber, data: make(map[string]Data), urlMap: make(map[string]string)}
}

func (shortener *UrlShortener) GetShortenUrl(url string) string {
	prevURL := shortener.urlMap[url]
	if data, ok := shortener.data[prevURL]; ok {
		startTime := data.time
		endTime := time.Now().Unix()
		tenMinutes := int64(10 * 60 * 60)

		data.time = time.Now().Unix()

		if endTime-startTime < tenMinutes {
			return prevURL
		}
	}

	urls := make([]string, 0, len(shortener.data))

	for key, data := range shortener.data {
		startTime := data.time
		endTime := time.Now().Unix()
		tenMinutes := int64(10 * 60 * 60)

		if endTime-startTime >= tenMinutes {
			urls = append(urls, key)
		}
	}

	for _, key := range urls {
		delete(shortener.data, key)
	}

	shortener.counter++
	var result string = ""
	numbers := shortener.number
	for i := 0; i < SIZE; i++ {
		digit := numbers % 10
		result += string(STRINGS[digit])
		numbers = numbers / 10
	}

	shortener.number += 1

	data := Data{url: url, time: time.Now().Unix()}

	shortener.data[result] = data
	shortener.urlMap[url] = result

	return result
}

func (shortener *UrlShortener) GetRedirectUrl(url string) (string, error) {
	if data, ok := shortener.data[url]; ok {
		return data.url, nil
	}

	return "", fmt.Errorf("There is no route:%s", url)
}
