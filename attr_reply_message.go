package radgo

import (
	"errors"
)

const (
	ArmEmptyUserName     	= "10001"
	ArmEmptyUserPassword 	= "10002"
	ArmNoUserMac         	= "10003"
	ArmNoChapChallenge   	= "10004"
	ArmNoNasPortId       	= "10005"
	ArmNoNasIdentifier   	= "10006"
	ArmNoNasIpAddress    	= "10007"
	ArmErrInnerService   	= "10008"

	ArmErrUserDomain       	= "11001"
	ArmBadUserNameFormat   	= "11002"
	ArmErrUserNamePassword 	= "11003"
	ArmExpireUserPassword  	= "11004"

	ArmLockedUser       	= "12001"
	ArmIllegalDev       	= "12002"
	ArmNeedAuth         	= "12003"
	ArmTooMoreOnlineDev 	= "12004"
)

var ErrUnknowReplyMessage	= errors.New("Unknow Reply Message")

var armMap = map[string]bool {
	ArmEmptyUserName     	: true,
	ArmEmptyUserPassword 	: true,
	ArmNoUserMac         	: true,
	ArmNoChapChallenge   	: true,
	ArmNoNasPortId       	: true,
	ArmNoNasIdentifier   	: true,
	ArmNoNasIpAddress    	: true,
	ArmErrInnerService   	: true,

	ArmErrUserDomain       	: true,
	ArmBadUserNameFormat   	: true,
	ArmErrUserNamePassword 	: true,
	ArmExpireUserPassword  	: true,

	ArmLockedUser       	: true,
	ArmIllegalDev       	: true,
	ArmNeedAuth         	: true,
	ArmTooMoreOnlineDev 	: true,
}

func IsGoodReplyMessage(msg []byte) bool {
	_, ok := armMap[string(msg)]
	
	return ok
}