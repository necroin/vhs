package remote_desktop_handlers

import (
	"bytes"
	"encoding/json"
	"io"
	plugins_core "vhs/src/plugins/core"
	remote_desktop_input "vhs/src/plugins/remote_desktop/handlers/input"
)

type Coords struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type CoordsDelta Coords

type MouseEvent struct {
	Type   string      `json:"type"`
	Coords Coords      `json:"coords"`
	Scroll CoordsDelta `json:"scroll_delta"`
}

type KeyboardEvent struct {
	Type    string `json:"type"`
	Keycode uint16 `json:"keycode"`
}

func MouseEventHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	event := &MouseEvent{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(event); err != nil {
		return err
	}

	coords := event.Coords

	if event.Type == "leftDown" {
		remote_desktop_input.MouseMove(coords.X, coords.Y)
		remote_desktop_input.MouseLeftDown()
	}

	if event.Type == "leftUp" {
		remote_desktop_input.MouseLeftUp()
	}

	if event.Type == "move" {
		remote_desktop_input.MouseMove(coords.X, coords.Y)
	}

	if event.Type == "scroll" {
		remote_desktop_input.MouseWheel(coords.X, coords.Y, event.Scroll.Y)
	}

	return nil
}

func KeyboardEventHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	event := &KeyboardEvent{}
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(event); err != nil {
		return err
	}

	switch event.Type {
	case "keydown":
		remote_desktop_input.KeyboardDown(event.Keycode)
	case "keyup":
		remote_desktop_input.KeyboardUp(event.Keycode)
	}

	return nil
}
