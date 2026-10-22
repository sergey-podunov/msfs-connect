package simconnect

import "fmt"

// HRESULTError is a failed SimConnect call. Code is the raw HRESULT; known codes are named in the message.
type HRESULTError struct {
	Op   string
	Code uint32
}

var hresultNames = map[uint32]string{
	0x80004005: "E_FAIL",
	0x80070057: "E_INVALIDARG",
	0x8007000E: "E_OUTOFMEMORY",
	0x80070006: "E_HANDLE",
	// Seen from GetNextDispatch when the simulator process is killed: it exits without sending QUIT.
	0xC000014B: "STATUS_PIPE_BROKEN, the simulator is gone",
}

func (e *HRESULTError) Error() string {
	if name, ok := hresultNames[e.Code]; ok {
		return fmt.Sprintf("%s failed: %s (HRESULT=0x%08X)", e.Op, name, e.Code)
	}
	return fmt.Sprintf("%s failed HRESULT=0x%08X", e.Op, e.Code)
}
