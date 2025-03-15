package service

import (
	"crypto/rand"
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	shortCodeLen  = 6
	maxRetries    = 10
	retryInterval = 100 * time.Millisecond
)

var (
	mu        sync.Mutex
	usedCodes = make(map[string]bool)
)

func generateShortCode() string {
	var shortCode string
	var retries int

	for retries < maxRetries {
		bytes := make([]byte, shortCodeLen)
		_, err := rand.Read(bytes)
		if err != nil {
			fmt.Println("Error generating random bytes:", err)
			return ""
		}

		for i, b := range bytes {
			bytes[i] = letterBytes[b%byte(len(letterBytes))]
		}

		shortCode = strings.ToLower(string(bytes))

		mu.Lock()
		if !usedCodes[shortCode] {
			usedCodes[shortCode] = true
			mu.Unlock()
			break
		}
		mu.Unlock()

		time.Sleep(retryInterval)
		retries++
	}

	if retries == maxRetries {
		fmt.Println("Failed to generate a unique short code after", maxRetries, "retries")
		return ""
	}

	return shortCode
}
