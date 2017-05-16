package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/go-martini/martini"

	"net/http"
	"net/url"
)

var baseURL = "https://translate.google.com/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob"

func main() {

	m := martini.Classic()

	m.Get("/speech/:text", func(params martini.Params, w http.ResponseWriter, r *http.Request) {
		text := params["text"]
		speech, _ := Speak(text, "no")

		w.Header().Set("Content-Type", "audio/mpeg")

		speech.WriteTo(w)
	})
	m.RunOnAddr(":8080")
	m.Run()
}

type Speech struct {
	bytes.Buffer
}

func Speak(text, language string) (*Speech, error) {
	req := fmt.Sprintf(baseURL, url.QueryEscape(text), url.QueryEscape(language))
	res, err := http.Get(req)
	if err != nil {
		return nil, err
	}

	speech := &Speech{}                                          // It will be returned as response
	if _, err := io.Copy(&speech.Buffer, res.Body); err != nil { //Read response body and copy it to buffer,also сheck for errors
		return nil, err
	}

	return speech, nil
}
