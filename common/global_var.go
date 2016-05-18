package common

const (
        DATINGTALK = iota
        SINGLETALK
        TEAMTALK
)

const (
        USERNOEXIST = iota + 10000
        USERWRONGPASSWD
        USERPERFECT
	USERLESSINFO
)

const (
        PASSWDDIFFERROR = iota + 20000
        PASSWDLENGTHERROR
        PASSWDOK
	PASSWDINSERTERROR
)
