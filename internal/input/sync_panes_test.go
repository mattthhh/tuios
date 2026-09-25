package input

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/Gaurav-Gosain/tuios/internal/app"
	"github.com/Gaurav-Gosain/tuios/internal/config"
	"github.com/Gaurav-Gosain/tuios/internal/terminal"
)

func TestToggleSyncPanesKeybinding(t *testing.T) {
	cfg := config.DefaultConfig()
	o, _ := osWithFocusedPane(t, cfg, app.WindowManagementMode)

	o, _ = HandleKeyPress(tea.KeyPressMsg{Code: 'y', Text: "y"}, o)
	if !o.SyncPanes {
		t.Fatal("sync panes was not enabled by the default window-mode keybinding")
	}

	o, _ = HandleKeyPress(tea.KeyPressMsg{Code: 'y', Text: "y"}, o)
	if o.SyncPanes {
		t.Fatal("sync panes was not disabled on second key press")
	}
}

func TestSyncPanesMirrorsTerminalInputToAllOpenWindows(t *testing.T) {
	cfg := config.DefaultConfig()
	o := app.NewOS(app.OSOptions{
		UserConfig:      cfg,
		KeybindRegistry: config.NewKeybindRegistry(cfg),
	})
	o.Mode = app.TerminalMode
	o.CurrentWorkspace = 1

	var got0, got1, got2 []byte
	o.Windows = []*terminal.Window{
		{ID: "w0", Workspace: 1, DaemonMode: true, DaemonWriteFunc: func(b []byte) error { got0 = append(got0, b...); return nil }},
		{ID: "w1", Workspace: 1, DaemonMode: true, DaemonWriteFunc: func(b []byte) error { got1 = append(got1, b...); return nil }},
		{ID: "w2", Workspace: 1, DaemonMode: true, DaemonWriteFunc: func(b []byte) error { got2 = append(got2, b...); return nil }},
	}
	o.FocusedWindow = 0
	o.SyncPanes = true

	o, _ = HandleKeyPress(tea.KeyPressMsg{Code: 'a', Text: "a"}, o)
	if string(got0) != "a" || string(got1) != "a" || string(got2) != "a" {
		t.Fatalf("typed key was not mirrored to all panes: focused=%q peer1=%q peer2=%q", got0, got1, got2)
	}
}
