package services

var (
	Player      PlayerService
	Level       LevelService
	Room        RoomService
	Reservation ReservationService
	Challenge   ChallengeService
	Payment     PaymentService
	GameLog     GameLogService
)

// InitServices 初始化所有服务
func InitServices() {
	Player = NewPlayerService()
	Level = NewLevelService()
	Room = NewRoomService()
	Reservation = NewReservationService()
	Challenge = NewChallengeService()
	Payment = NewPaymentService()
	GameLog = NewGameLogService()
}