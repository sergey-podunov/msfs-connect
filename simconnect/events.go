package simconnect

// Flags of the Pause_EX1 system event, in SIMCONNECT_RECV_EVENT.DwData.
const (
	PAUSE_STATE_FLAG_OFF              = 0
	PAUSE_STATE_FLAG_PAUSE            = 1 // full pause
	PAUSE_STATE_FLAG_PAUSE_WITH_SOUND = 2
	PAUSE_STATE_FLAG_ACTIVE_PAUSE     = 4
	PAUSE_STATE_FLAG_SIM_PAUSE        = 8 // the simulation is paused, but not the time
)
