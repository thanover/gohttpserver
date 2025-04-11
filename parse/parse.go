package parse

import logger "github.com/thanover/gologger"

type ParsedMsg struct {
	Headers string
	Content string
	Trailer string
}

var AppLogger logger.Logger = logger.GetLogger()

func ParseMsg(msg []byte) (parsedMsg ParsedMsg) {
	lines := getLines(msg)
	AppLogger.Info(string(lines))
	return ParsedMsg{
		Headers: "",
		Content: "",
		Trailer: "",
	}
}

func getLines(msg []byte) (lines []byte) {
	return msg

}
