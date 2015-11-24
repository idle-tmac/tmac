package common

/*      struct dating    */

type DatingPushMsg struct {
    Datingid int64  
	Uid int64     
	Text string 
}

type DatingMsg struct {
	Uid int64
    Time int64
    Text string
}

type DatingPullMsg struct {
    Datingid int64 
    Uid int64     
    Time int64 
}
