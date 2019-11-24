package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/LunaNode/cloug/provider/lunanode"
	//"github.com/LunaNode/cloug/service/compute"
)

type LunaNode struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Region string `json:"region"`
	APIID  string `json:"api_id"`
	APIKEY string `json:"api_key"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	ln := LunaNode{
		Name:   "medienwerksalzburg",
		Type:   "lndynamic",
		Region: "roubaix",
		APIID:  os.Getenv("LN_ID"),
		APIKEY: os.Getenv("LN_KEY"),
	}

	b, _ := json.Marshal(ln)

	p, err := lunanode.LunaNodeFromJSON(b)

	if err != nil {

		fmt.Fprint(w, err)

	}

	p.CreateInstance()

}
