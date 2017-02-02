package esiLegacy

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

var _ time.Time
var _ = mux.NewRouter

func GetMarketsPrices(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "adjusted_price" : 306988.09,
  "average_price" : 306292.67,
  "type_id" : 32772
} ]`
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

func GetMarketsRegionIdHistory(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		regionId int32
		typeId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "average" : 5.25,
  "date" : "2015-05-01",
  "highest" : 5.27,
  "lowest" : 5.11,
  "order_count" : 2267,
  "volume" : 16276782035
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(regionId, vars["region_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	regionId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	localV, err = processParameters(typeId, r.Form.Get("type_id"))
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

func GetMarketsRegionIdOrders(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		regionId int32
		orderType string
		typeId int32
		page int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "duration" : 90,
  "is_buy_order" : false,
  "issued" : "2016-09-03T05:12:25Z",
  "location_id" : 60005599,
  "min_volume" : 1,
  "order_id" : 4623824223,
  "price" : 9.9,
  "range" : "region",
  "type_id" : 34,
  "volume_remain" : 1296000,
  "volume_total" : 2000000
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(regionId, vars["region_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	regionId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("typeId") != "" {
		localV, err = processParameters(typeId, r.Form.Get("type_id"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		typeId = localV.(int32)
	}
	localV, err = processParameters(orderType, r.Form.Get("order_type"))
	if err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("page") != "" {
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
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

func GetMarketsStructuresStructureId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		structureId int64
		page int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "duration" : 90,
  "is_buy_order" : false,
  "issued" : "2016-09-03T05:12:25Z",
  "location_id" : 60005599,
  "min_volume" : 1,
  "order_id" : 4623824223,
  "price" : 9.9,
  "range" : "region",
  "type_id" : 34,
  "volume_remain" : 1296000,
  "volume_total" : 2000000
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(structureId, vars["structure_id"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	structureId = localV.(int64)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("page") != "" {
		localV, err = processParameters(page, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		page = localV.(int32)
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


