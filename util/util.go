package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	LineBreak     = "\n"
	SessionCookie = "session"
	SessionEnvVar = "aoc_session_cookie"
)

func HttpRequestWithCookie(url string, method string, cookieName string, cookieVal string) (b []byte, err error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: cookieName, Value: cookieVal})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Failed to get input data from AoC webserver")
		fmt.Println(fmt.Sprintf("Is the environment variable %s set?", SessionEnvVar))
		err = errors.New(url +
			"\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}
