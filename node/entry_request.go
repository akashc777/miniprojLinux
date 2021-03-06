package node

type EntryRequest struct {
	CmdID        int64  `json:"cmd_id"`
	Term         int64  `json:"term"`
	LeaderID     string `json:"leader_id"`
	PrevLogIndex int64  `json:"prev_log_index"`
	PrevLogTerm  int64  `json:"prev_log_term"`
	Data         []byte `json:"data"`
	Stat    map[string]Health  `json:"stat"`
	State	string	`json:"state"`
}

type EntryResponse struct {
	Term    int64   `json:"term"`
	Success bool    `json:"success"`
  Stat    Health  `json:"stat"`
}

func newEntryRequest(cmdID int64, leaderID string, term int64, prevLogIndex int64, prevLogTerm int64, data []byte, h map[string]Health, s string) EntryRequest {
	return EntryRequest{
		CmdID:        cmdID,
		LeaderID:     leaderID,
		Term:         term,
		PrevLogIndex: prevLogIndex,
		PrevLogTerm:  prevLogTerm,
		Data:         data,
		Stat:	h,
		State: s,
	}
}
