package main

import (
				"io/ioutil"
				"net/http"
				"fmt"
				"encoding/json"
				"errors"
)

type ApiResponse struct {
				Amount        string `json:"amount"`
				Block         int    `json:"block"`
				Blockhash     string `json:"blockhash"`
				Blocktime     int    `json:"blocktime"`
				Category      string `json:"category"`
				Confirmations int    `json:"confirmations"`
				Creationtxid  string `json:"creationtxid"`
				Data          string `json:"data"`
				Divisible     bool   `json:"divisible"`
				Ecosystem     string `json:"ecosystem"`
				Fee           string `json:"fee"`
				Fixedissuance bool   `json:"fixedissuance"`
				Flags         struct {
				} `json:"flags"`
				Freezingenabled bool `json:"freezingenabled"`
				Ismine          bool `json:"ismine"`
				Issuances       []struct {
								Grant string `json:"grant"`
								Txid  string `json:"txid"`
				} `json:"issuances"`
				Issuer          string      `json:"issuer"`
				Managedissuance bool        `json:"managedissuance"`
				Name            string      `json:"name"`
				Positioninblock int         `json:"positioninblock"`
				Propertyid      int         `json:"propertyid"`
				Propertyname    string      `json:"propertyname"`
				Propertytype    string      `json:"propertytype"`
				Rdata           interface{} `json:"rdata"`
				Registered      bool        `json:"registered"`
				Sendingaddress  string      `json:"sendingaddress"`
				Subcategory     string      `json:"subcategory"`
				Totaltokens     string      `json:"totaltokens"`
				Txid            string      `json:"txid"`
				Type            string      `json:"type"`
				TypeInt         int         `json:"type_int"`
				URL             string      `json:"url"`
				Valid           bool        `json:"valid"`
				Version         int         `json:"version"`
}

func getSupply() (string, error) {
				resp, err := http.Get("https://api.omniexplorer.info/v1/property/31")

				if err != nil {
								// handle error
				}
				defer resp.Body.Close()
				if resp.StatusCode == http.StatusOK {
								bodyBytes, err2 := ioutil.ReadAll(resp.Body)
								//bodyString := string(bodyBytes)
								if err2 == nil {
												//fmt.Printf(bodyString)
								}
							  respData := &ApiResponse{}
								err := json.Unmarshal(bodyBytes, respData)
								fmt.Println(err)
								fmt.Println(respData.Totaltokens)
								return respData.Totaltokens, nil
				}
				return "", errors.New("api error")
}

