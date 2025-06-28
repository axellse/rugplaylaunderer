package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func makeRequest(method string, url string, body any, token string) ([]byte, error) {
  var bodyrd io.Reader = nil
  if bodystr, ok := body.(string); ok {
    bodyrd = strings.NewReader(bodystr)
  } else if body != nil {
    ba, jerr := json.Marshal(body)
    if jerr != nil {return nil, jerr}

    bodyrd = bytes.NewReader(ba)
  }
  
  req, err := http.NewRequest("POST", "https://rugplay.com/api/rewards/claim", bodyrd)
  req.AddCookie(&http.Cookie{
    Name:  "__Secure-better-auth.session_token",
    Value: token,
  })
  //chrome-like headers
  req.Header.Add("content-type", "application/json")
  req.Header.Add("sec-fetch-site", "same-origin")
  req.Header.Add("sec-fetch-mode", "cors")
  req.Header.Add("sec-fetch-dest", "empty")
  req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
  req.Header.Add("sec-ch-ua-mobile", "?0")
  req.Header.Add("sec-ch-ua", "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"Google Chrome\";v=\"138\"")
  req.Header.Add("priority", "u=1, i")
  req.Header.Add("accept-language", "en-US,en;q=0.9")
  req.Header.Add("accept", "*/*")

  if err != nil {return nil, err}
  resp, err := http.DefaultClient.Do(req)
  if err != nil {return nil, err}

  ba, err := io.ReadAll(resp.Body)
  if err != nil {return nil, err}

  return ba, nil
}