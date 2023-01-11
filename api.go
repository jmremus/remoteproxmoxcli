package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var ticket string
var csrf string

// ignore cert
var tr = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
var jar, _ = cookiejar.New(nil)
var client = &http.Client{Transport: tr, Jar: jar}

func handleError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// Send an HTTP request and return the body
func makeRequest(getpost string, url string, data io.Reader, auth bool) []byte {
	req, err := http.NewRequest(getpost, url, data)
	handleError(err)

	if auth == true {
		req.Header.Set("CSRFPreventionToken", csrf)
		req.AddCookie(&http.Cookie{Name: "PVEAuthCookie", Value: ticket})
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	handleError(err)
	body, err := io.ReadAll(resp.Body)
	handleError(err)

	return body
}

func getTicket() string {
	passwd := "username=" + USERNAME + "&password=" + PASSWORD
	url := HOST + "/api2/json/access/ticket"
	body := makeRequest("POST", url, bytes.NewBufferString(passwd), true)

	var ticketJson interface{}
	json.Unmarshal(body, &ticketJson)
	ticketMap := ticketJson.(map[string]interface{})
	dataMap := ticketMap["data"].(map[string]interface{})
	ticket = dataMap["ticket"].(string)
	csrf = dataMap["CSRFPreventionToken"].(string)

	return ticket
}

func getNodes() []NodeInfo {
	url := HOST + "/api2/json/nodes/"
	body := makeRequest("GET", url, nil, true)

	var data NodeData
	json.Unmarshal(body, &data)

	return data.Data
}

func listVMs() []VmInfo {
	url := HOST + "/api2/json/nodes/" + NODE + "/qemu/"
	body := makeRequest("GET", url, nil, true)

	var data VmData
	json.Unmarshal(body, &data)

	return data.Data
}

func toggleVM(cmd string, vmid string) []VmInfo {
	url := HOST + "/api2/json/nodes/" + NODE + "/qemu/" + vmid + "/status/" + cmd
	body := makeRequest("POST", url, nil, true)

	var data VmData
	json.Unmarshal(body, &data)

	return data.Data
}
