package raft

import (
	"testing"
)

func TestElectionBasic(t *testing.T) {
	h := NewHarness(t, 3)
	defer h.Shutdown()
	h.CheckSingLeLeader()
}

func TestElectionLeaderDisconnect(t *testing.T) {
	h := NewHarness(t, 3)
	defer h.Shutdown()

	origLeaderId, origTime := h.CheckSingLeLeader()

	//t.Logf("origLeaderId:%d", origLeaderId)

	h.DisconnectPeer(origLeaderId)
	sleepMs(350)

	newLeaderId, newTerm := h.CheckSingLeLeader()
	if newLeaderId == origLeaderId {
		t.Errorf("want new leader to be different from oring leader")
	}

	if newTerm <= origTime {
		t.Errorf("want newTerm <= origTerm,got %d and %d", newTerm, origLeaderId)
	}
}
