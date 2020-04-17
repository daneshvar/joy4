package format

import (
	"github.com/daneshvar/joy4/av/avutil"
	"github.com/daneshvar/joy4/format/mp4"
	"github.com/daneshvar/joy4/format/raw"
	"github.com/daneshvar/joy4/format/rtsp"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(raw.Handler)
}
