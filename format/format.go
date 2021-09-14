package format

import (
	"github.com/zhangm168520/vdk/av/avutil"
	"github.com/zhangm168520/vdk/format/aac"
	"github.com/zhangm168520/vdk/format/flv"
	"github.com/zhangm168520/vdk/format/mp4"
	"github.com/zhangm168520/vdk/format/rtmp"
	"github.com/zhangm168520/vdk/format/rtsp"
	"github.com/zhangm168520/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
