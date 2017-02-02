package esiV2

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetCharactersCharacterIdCalendarEventId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		eventId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "duration_in_minutes" : 60,
  "event_date" : "2016-06-26T21:00:00Z",
  "event_id" : 1386435,
  "event_response" : "Undecided",
  "event_text" : "o7: The EVE Online Show features latest developer news, fast paced action, community overviews and a lot more with CCP Guard and CCP Mimic. Join the thrilling o7 live broadcast at 20:00 EVE time (=UTC) on <a href=\"http://www.twitch.tv/ccp\">EVE TV</a>. Don't miss it!",
  "importance" : 1,
  "owner_id" : 1,
  "owner_name" : "EVE System",
  "owner_type_id" : 0,
  "title" : "o7 The EVE Online Show"
}`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	localV, err = processParameters(eventId, vars["event_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	eventId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage int32 
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PutCharactersCharacterIdCalendarEventId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		eventId int32
		response string
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["character_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	localV, err = processParameters(eventId, vars["event_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	eventId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(response, r.Form.Get("response"))
	if err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage int32 
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	} 

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


