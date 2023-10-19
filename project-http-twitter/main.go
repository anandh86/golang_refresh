package main


import (
  "net/http"
  "log"
  "encoding/json"
  "twitter/server"
  "bytes"
  "time"
  "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  "github.com/go-chi/httprate"
)

func main() {
	// Your solution goes here. Good luck!
  repo := server.TweetMemoryRepo{}
  s := server.Server{&repo}

  //-----
  r := chi.NewRouter()

	r.Use(middleware.Logger)

	rateLimit := httprate.LimitByIP(10, 1*time.Minute)
	r.With(rateLimit).Post("/tweets", s.AddToTweet)

	r.Get("/tweets", s.Tweets)

	go spamTweets()

	log.Fatal(http.ListenAndServe(":8080", r))
}

type tmpTweet struct {
  Message string `json:"message"`
  Location string `json:"location"`
}

func spamTweets()  {

  tTweet := tmpTweet{Message:"ass",Location:"ass"}

  payload, err := json.Marshal(tTweet)
  if err != nil {
  	log.Println("Failed to marshal:", err)
  	return
  }

  //w.Write(payload)

  for {

		// Send HTTP POST request
    url := "http://localhost:8080/tweets"

    resp, err1 := http.Post(url, "application/json", bytes.NewBuffer(payload))
    if err1 != nil {
    	return
    }
		// Close response body
    defer resp.Body.Close()


  }

}
