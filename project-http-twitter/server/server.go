package server

import (
  "net/http"
  "log"
  "encoding/json"
  "io"
  "time"
  "fmt"
  "sync"
)

type TweetRepository interface {
  // Add tweets to db and return the tweet id
  AddTweet(tw tweet) (int, error)
  Tweets() ([]tweet, error)
}

type TweetMemoryRepo struct {
  tweets []tweet
  poottu  sync.RWMutex
}

type tweetsList struct {
	Tweets []tweet `json:"tweets"`
}

func (this *TweetMemoryRepo) AddTweet(tw tweet) (int, error)  {
  this.poottu.RLock()
  defer this.poottu.RUnlock()

  this.tweets = append(this.tweets, tw)
  return len(this.tweets), nil
}

func (this *TweetMemoryRepo) Tweets() ([]tweet, error)  {
  this.poottu.RLock()
  defer this.poottu.RUnlock()

  return this.tweets, nil
}

type Server struct {
  TwRepo TweetRepository
}



type tweet struct {
  Message string `json:"message"`
  Location string `json:"location"`
}

type tweetCount struct {
  ID int
}

func (s Server) TweetsHandler(w http.ResponseWriter, r *http.Request)  {

  start := time.Now()

	defer func() {
		fmt.Printf("%s %s %s\n", r.Method, r.URL, time.Since(start))
	}()

   if r.Method == http.MethodPost {
     s.AddToTweet(w,r)
   } else if r.Method == http.MethodGet {
     s.Tweets(w,r)
   }

}

func (s Server) Tweets(w http.ResponseWriter, r *http.Request)  {

  twts, err := s.TwRepo.Tweets()

  if err != nil {
    log.Println("Failed to read tweets:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
  }


  TwList := tweetsList{Tweets:twts}

  payload, err := json.Marshal(TwList)
  if err != nil {
  	log.Println("Failed to marshal:", err)
  	w.WriteHeader(http.StatusInternalServerError)
  	return
  }

  w.Write(payload)

}

func (s *Server) AddToTweet(w http.ResponseWriter, r *http.Request)  {

  body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

  t := tweet{}

  if err := json.Unmarshal(body, &t); err != nil {
  	log.Println("Failed to unmarshal payload:", err)
  	w.WriteHeader(http.StatusBadRequest)
  	return
  }

  id, err := s.TwRepo.AddTweet(t)

  if err != nil {
		log.Println("Failed to add tweet:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

  tc := tweetCount{ID:id}

  payload, err := json.Marshal(tc)
  if err != nil {
  	log.Println("Failed to marshal:", err)
  	w.WriteHeader(http.StatusInternalServerError)
  	return
  }

  w.Write(payload)

}

func Tweets(w http.ResponseWriter, r *http.Request)  {

  body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("Failed to read body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

  t := tweet{}

  if err := json.Unmarshal(body, &t); err != nil {
  	log.Println("Failed to unmarshal payload:", err)
  	w.WriteHeader(http.StatusBadRequest)
  	return
  }

  fmt.Printf("Tweet: `%s` from %s\n", t.Message, t.Location)


  tc := tweetCount{ID:0}

  payload, err := json.Marshal(tc)
  if err != nil {
  	log.Println("Failed to marshal:", err)
  	w.WriteHeader(http.StatusInternalServerError)
  	return
  }

  w.Write(payload)
}
