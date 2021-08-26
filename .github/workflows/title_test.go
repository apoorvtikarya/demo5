package workflows

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func Test_Tittle(t *testing.T) {

	//var title string
    //_, err := fmt.Scanln(&title)
	//if err != nil {
	//	return
	//}
	//err = os.Setenv("TITLE", title)
	err := os.Setenv("TITLE","feat(dot/rpc): implement chain_subscribeAllHeads RPC" )
	if err != nil {
		return
	}
	//os.Setenv("BODY", "2")
	fmt.Println("TITLE:", os.Getenv("TITLE"))
	//fmt.Println("BODY:", os.Getenv("BODY"))
	//var match, _ = regexp.MatchString(".+\\(.+\\)\\:.+",os.Getenv("TITLE"))
	var match, _ = regexp.MatchString("",os.Getenv("TITLE"))

	fmt.Println(match)
	//for _, e := range os.Environ() {
	//	pair := strings.SplitN(e, "=", 2)
	//	fmt.Println(pair[0])
	//}

}
//var match1, _ = regexp.MatchString("(?i)[##\\sChanges\\s\\[-]+.+ ##\\sTests\\s\\-.+ ##\\sIssues\\s\\-.+ ##\\sPrimary Reviewer\\-.+]", result)
	//fmt.Println(match1)
	//var match1, _ = regexp.MatchString("(?i)[\\##\\s+Changes+\\s+\\-\\.+\\##\\s+Tests\\s+\\-\\.+\\##\\s+Issues\\s+\\-\\.+\\##\\s+Primary Reviewer\\s+\\-\\.+]",result)
	//fmt.Println(match1)