package input

import "github.com/Gaurav-Gosain/tuios/internal/app"

// forwardTerminalInputPeers mirrors terminal input to non-focused panes selected
// for multi-input.
func forwardTerminalInputPeers(o *app.OS, input []byte) {
	if len(input) == 0 || o == nil {
		return
	}
	for idx, w := range o.Windows {
		if w == nil || idx == o.FocusedWindow {
			continue
		}
		if o.SyncPanes || (len(o.MultifocusSet) > 0 && o.MultifocusSet[w.ID]) {
			_ = w.SendInput(input)
		}
	}
}
