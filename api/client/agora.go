package client

import (
	"fmt"
	"math/rand"
	"net/http"
	CONSTANT "shivamtiwari/github.com/agora/constant"
	UTIL "shivamtiwari/github.com/agora/util"
	"time"
)

func GenerateAgoraToken(w http.ResponseWriter, r *http.Request) {

	var response = make(map[string]interface{})

	var roleStr, agora_token, uidStr, channelName string

	uidStr = generateRandomID()

	channelName = r.FormValue("channel_id")
	if r.FormValue("type") == "1" {
		roleStr = CONSTANT.RolePublisher
	} else if r.FormValue("type") == "2" {
		roleStr = CONSTANT.RoleSubscriber
	} else {
		roleStr = "attended"
	}

	//uidStr = generateRandomID()
	// For demonstration purposes the expiry time is set to 7200 seconds = 2 hours. This shows you the automatic token renew actions of the client.
	expireTimeInSeconds := uint32(7200)
	// Get current timestamp.
	currentTimestamp := uint32(time.Now().UTC().Unix())
	// Timestamp when the token expires.
	expireTimestamp := currentTimestamp + expireTimeInSeconds

	token, err := UTIL.GenerateAgoraRTCToken(channelName, roleStr, uidStr, expireTimestamp)
	if err != nil {
		UTIL.SetReponse(w, "", "", CONSTANT.ShowDialog, response)
		return
	}
	agora_token = token

	response["token"] = agora_token
	response["UID"] = uidStr

	fmt.Println("status" + " " + CONSTANT.StatusCodeOk)

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.ShowDialog, response)

}

func generateRandomID() string {
	const randomIDdigits = "123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = randomIDdigits[rand.Intn(len(randomIDdigits))]
	}
	return string(b)
}
