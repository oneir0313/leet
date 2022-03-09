package problems

import (
	"fmt"
	"strconv"
)

type Problem struct{}

func (p Problem) Default() {
	fmt.Println(getDesensitizedRoomName(6802))
}

var gameAlertMap = map[string]string{
	"620":  "DZPK",
	"360":  "DZPK",
	"370":  "DZPK",
	"720":  "28G",
	"830":  "QZNN",
	"220":  "ZJH",
	"860":  "SG",
	"900":  "LH",
	"600":  "21D",
	"880":  "HLHB",
	"870":  "TBNN",
	"610":  "DDZ",
	"730":  "QZPJ",
	"230":  "JSZJH",
	"630":  "SSS",
	"380":  "XYWZ",
	"390":  "SLM",
	"910":  "BJL",
	"920":  "SLWH",
	"890":  "K3ZQZNN",
	"930":  "BRNN",
	"950":  "HHDZ",
	"740":  "ERMJ",
	"550":  "QZTB",
	"680":  "SDDZ",
	"1950": "WRZJH",
	"650":  "XLCH",
	"940":  "JSYS",
	"8120": "XZDD",
	"1350": "XYZP",
	"8130": "PDK",
	"3930": "TB",
	"510":  "HBBY",
	"1960": "Jackpotbenz",
	"8150": "K4ZQZNN",
	"8180": "BSXXL",
	"8160": "LZNN",
	"8210": "BYB",
}

func getDesensitizedRoomName(roomID int) string {
	roomNum := strconv.Itoa(roomID)
	gameNum := roomNum[0 : len(roomNum)-1]
	desName, ok := gameAlertMap[gameNum]
	if ok {
		return fmt.Sprintf("%s型號%s", desName, roomNum[len(roomNum)-1:])
	}
	return "未知型號"
}
