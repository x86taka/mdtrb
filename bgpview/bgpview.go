package bgpview

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetASNFromIP(addr string) string {
	httpclient := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.bgpview.io/ip/"+addr, nil)
	resp, err := httpclient.Do(req)
	if err != nil {
		log.Println("Errored when sending request to the server")
		return ""
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	data := &BGPViewIP{}
	if err := json.Unmarshal(resp_body, data); err != nil {
		log.Println("Error BGPView json Unmarshal")
	}
	return data.Data.Prefixes[0].Asn.Name
}

func GetASNName(asn string) string {
	httpclient := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.bgpview.io/asn/"+asn, nil)
	resp, err := httpclient.Do(req)
	if err != nil {
		log.Println("Errored when sending request to the server")
		return ""
	}
	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	data := &GETASN{}
	if err := json.Unmarshal(resp_body, data); err != nil {
		log.Println("Error BGPView GetASN json Unmarshal")
	}
	return fmt.Sprintf("%s\t%s", asn, data.Data.Name)
}
