package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a deferred function (which will always be run in the event of a panic as Go unwinds the stack).
		defer func() {
			// Use the builtin recover function to check if there has been a panic or not.
			if err := recover(); err != nil {
				// If there was a panic, set a "Connection: close" header on the response.
				// This acts as a trigger to make Go's HTTP server automatically close the current connection after a response has been sent.
				w.Header().Set("Connection", "close")

				// The value returned by recover() has the type interface{}, so we use fmt.Errorf() to normalize it into an error and call our serverErrorResponse() helper.
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimit(next http.Handler) http.Handler {
	// Any code here will run only once, when we wrap something with the middleware.
	// Define a client struct to hold the rate limiter and last seen time for each client.
	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)

	// Launch a background goroutine which removes old entries from the clients map once every minute.
	go func() {
		for {
			time.Sleep(time.Minute)

			// Lock the mutex to prevent any rate limiter checks from happening while the cleanup is taking place.
			mu.Lock()

			// Loop through all clients. If they haven't been seen within the last three minutes, delete the corresponding entry from the map.
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			// Importantly, unlock the mutex when the cleanup is complete.
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Any code here will run for every request that the middleware handles.
		// Only carry out the check if rate limiting is enabled.
		if app.config.limiter.enabled {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

			mu.Lock()

			if _, found := clients[ip]; !found {
				// Create and add a new client struct to the map if it doesn't already exist.
				clients[ip] = &client{
					limiter: rate.NewLimiter(rate.Limit(app.config.limiter.rps), app.config.limiter.burst),
				}
			}

			// Update the last seen time for the client.
			clients[ip].lastSeen = time.Now()

			// Call limiter.Allow() to see if the request is permitted, and if it's not,
			// then we call the rateLimitExceededResponse() helper to return a 429 Too Many Requests response.
			if !clients[ip].limiter.Allow() {
				mu.Unlock()
				app.rateLimitExceededResponse(w, r)
				return
			}

			mu.Unlock()
		}

		next.ServeHTTP(w, r)
	})
}
