package util

import (
	CONFIG "shivamtiwari/github.com/agora/config"

	rtctokenbuilder "github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
)

func GenerateAgoraRTCToken(channelName string, roleStr string, uidStr string, expireTime uint32) (result string, err error) {

	// var appID, appCertificate string
	var role rtctokenbuilder.Role

	if roleStr == "publisher" {
		role = rtctokenbuilder.RolePublisher
	} else if roleStr == "subscriber" {
		role = rtctokenbuilder.RoleSubscriber
	} else {
		role = rtctokenbuilder.RoleAttendee
	}

	// appID = CONSTANT.AGORA_APP_ID
	// appCertificate = CONSTANT.AGORA_APP_CERTIFICATE
	result, err = rtctokenbuilder.BuildTokenWithUserAccount(CONFIG.AGORA_APP_ID, CONFIG.AGORA_APP_CERTIFICATE, channelName, uidStr, role, expireTime)

	return result, err

}
