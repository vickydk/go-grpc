package constants

const (
	OrderCancel      int = 99
	PaymentPending   int = 1
	AuthOrder        int = 2
	WaitingForPickUp int = 3
	Dispatched       int = 4
	Delivered        int = 5
	Return           int = 6
	StockPending     int = 7

	UpdateAirwayBill int = 98

	//Logaction
	LogCreate         = "create"
	LogMoveOrder      = "moveorder"
	LogReturn         = "return"
	LogReorder        = "reorder"
	LogPickByCUstomer = "pickbycustomer"
	LogAssignCourier  = "assigncourier"
	LogOutOfDelivery  = "outofdelivery"
	LogInTransit      = "intransti"
	LogAuth           = "auth"
	LogWFP            = "wfp"

	//substatus Authorized
	PickingPending  int = 1
	InPicking       int = 2
	PickingComplete int = 3

	//substatus waiting for pickup
	CustToPickUp    int = 1
	WaitForShipment int = 2
	ReadyToPickUp   int = 3

	//substatus Dispatched
	OutForDelivery int = 1
	InTransit      int = 2

	//substatus Return
	ReturnInitiated int = 1
	ReturnReceived  int = 2
	ReturnOnProcess int = 3
	ReturnClosed    int = 4

	//status Payment
	UnPaid     int = 1
	Paid       int = 2
	CancelPaid int = 99

	//Order Type
	Pickup   int = 1
	Delivery int = 2

	//FIlter by
	OrdeDate      int = 1
	ShipDate      int = 2
	ReturnDate    int = 3
	DeliveredDate int = 4
	DispatchDate  int = 5

	//Date Time
	LayTimeISO     = "2006-01-02"
	LayDateTimeISO = "2006-01-02 15:04:05"
	LayTimeUS      = "January 2, 2006"
	ANSIC          = "Mon Jan _2 15:04:05 2006"
	UnixDate       = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate       = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822         = "02 Jan 06 15:04 MST"
	RFC822Z        = "02 Jan 06 15:04 -0700"
	RFC850         = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123        = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z       = "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339        = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano    = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen        = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
