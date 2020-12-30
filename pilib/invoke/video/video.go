package video

import (
	"context"
	"os/exec"

	"picamera-go/pilib/model"
)

const (
	raspivid = "raspivid"
	// -----------------raspivid-----------------//
	// --width,    -w        Set image width <size> [64,1920]
	//  --height,    -h        Set image height <size> [64,1080]
	// --bitrate,    -b        Set bitrate. default 10Mbits/s=10000000, max is 25000000.
	// --output,    -o        Output filename <filename>,if "-", output to stdout
	// --listen,    -l  output to network, tcp://192.168.1.2:1234 or udp://192.168.1.2:1234,simple: raspivid -l -o tcp://0.0.0.0:3333
	// --verbose,    -v        Output verbose information during run
	// --timeout,    -t        Time before the camera takes picture and shuts down,default is 5 second, 0 is forever.
	// --demo,    -d        Run a demo mode <milliseconds>
	// --framerate,    -fps        Specify the frames per second to record, 2<=fps<=30,will be change
	// --penc,    -e        Display preview image after encoding,will be change
	// --qp,    -qp        Set quantisation parameter
	// --profile,    -pf        Specify H264 profile to use for encoding, baseline,main,high
	// --level,    -lev         Specifies the H264 encoder level to use for encoding. Options are 4, 4.1, and 4.2.
	// --inline,    -ih        Insert PPS, SPS headers .Forces the stream to include PPS and SPS headers on every I-frame. Needed for certain streaming cases e.g. Apple HLS. These headers are small, so don't greatly increase the file size.
	// --spstimings,    -stm        Insert timing information into the SPS block.
	// --timed,    -td        Do timed switches between capture and pause.simple:raspivid -o test.h264 -t 25000 -timed 2500,5000
	// --keypress,    -k        Toggle between record and pause on Enter keypress
	// --signal,    -s        Toggle between record and pause according to SIGUSR1
	// --split,    -sp      When in a signal or keypress mode, each time recording is restarted, a new file is created.
	// --initial,    -i        Define initial state on startup,Define whether the camera will start paused or will immediately start recording. Options are record or pause. Note that if you are using a simple timeout, and initial is set to pause, no output will be recorded.
	// --segment,    -sg        Segment the stream into multiple files
	// --circular,    -c        Select circular buffer mode. All encoded data is stored in a circular buffer until a trigger is activated, then the buffer is saved.
	// --vectors,    -x         Turns on output of motion vectors from the H264 encoder to the specified file name.
	// --flush,    -fl          Forces a flush of output data buffers as soon as video data is written. This bypasses any OS caching of written data, and can decrease latency.
	// --save-pts,    -pts      Saves timestamp information to the specified file. Useful as an imput file to mkvmerge.
	// --codec,    -cd          Specifies the encoder codec to use. Options are H264 and MJPEG. H264 can encode up to 1080p, whereas MJPEG can encode upto the sensor size, but at decreased framerates due to the higher processing and storage requirements.
	// --wrap,    -wr        Set the maximum value for segment number, So if set to 4, in the segment example above, the files produced will be video0001.h264, video0002.h264, video0003.h264, and video0004.h264. Once video0004.h264 is recorded, the count will reset to 1, and video0001.h264 will be overwritten.
	// --start,    -sn        Set the initial segment number ,When outputting segments, this is the initial segment number, giving the ability to resume a previous recording from a given segment. The default value is 1.
	// --raw,    -r         Specify the output file name for any raw data files requested.
	// --raw-format,    -rf         Specify the raw format to be used if raw output requested. Options as yuv, rgb, and grey. grey simply saves the Y channel of the YUV image.
)

type Observer struct {
	model.Observer
	calfunc context.CancelFunc
	osCmd   *exec.Cmd
}
type Option func(*Observer)
type Mode int

func NewObserver(opts ...Option) *Observer {
	v := &Observer{

	}
	for _, opt := range opts {
		opt(v)
	}
	return v
}

func (v *Observer) Start() error {
	return nil
}

func (v *Observer) Stop() error {
	return nil
}

func (v *Observer) Shoot() error {
	return nil
}

func (v *Observer) AddSub(s model.Subscriber) {

}
