package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/png"
	"log"

	sc "github.com/kbinani/screenshot"
)

//Takes a screenshot of the first screen and returns the base64 encoded image
func getScreenShot(message Message, rm *ResponseMessage) []byte {

	n := sc.NumActiveDisplays()

	if n == 0 {
		return nil
	}

	bounds := sc.GetDisplayBounds(0)
	image, err := sc.CaptureRect(bounds)
	buf := new(bytes.Buffer)

	if err != nil {
		log.Println(err)
		return nil
	}

	png.Encode(buf, image)

	screenshot := &ScreenShot{
		ResponseMessage: rm,
		Content:         base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
	res, err := json.Marshal(screenshot)
	if err != nil {
		return nil
	}
	return res

}
