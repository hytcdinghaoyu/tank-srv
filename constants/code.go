package constants

const (
	Success = 0

	ServerIsFull = 1000

	DuplicatedLogin = 1001      // 重复登陆
    BattleIdNotExist = 2001     // 不存在的battleId
    UserNotLogin = 2002         // 用户未登录
    UserAlreadyJoin = 2003      // 用户已经加入了战场
    UserNotJoin = 2004          // 用户未加入战场
    JoinNotAvailable = 2005     // 没有合适的房间

)