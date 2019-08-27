package entity

type BattleUser struct {
	Player
	Mass      float32 //质量，决定碰撞效果
	PosX      float32
	PoxY      float32
	AccX      float32
	AccY      float32
	Rotation  float32
	SpeedX    float32
	SpeedY    float32
	TopSpeed  int     //最大加速度
	AccPower  float32 //加速度系数
	Size      int     //角色尺寸
	Life      float32
	MaxLife   int
	SkillID   int
	SkillData string
}

type Ball struct {
	BallID  string
	PosX    float32
	PosY    float32
	SpeedX  float32
	SpeedY  float32
	Explode int
}
