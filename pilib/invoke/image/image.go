package image

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/micmonay/keybd_event"
	"github.com/pkg/errors"
	"picamera-go/pilib/key"

	"picamera-go/pilib/model"
)

const (
	raspistill = "raspistill"
	// ---------------raspistill---------------//
	// --width,    -w
	defaultWidth = "640" // Set image width <size>
	// --height,    -h
	defaultHeight = "480" // Set image height <size>
	// --quality,    -q
	defaultQuality = "75" // Set JPEG quality <0 to 100>
	// --timeout,    -t
	defaultTimeout = "5000" // Time before the camera takes picture and shuts down
	// --timelapse,    -tl
	defaultTimelapse = "1000" //  time-lapse mode ticker
	// --raw,    -r
	addRaw bool = false // Add raw Bayer data to JPEG metadata
	// --encoding,    -e        Encoding to use for output file
	defaultEncoding string = "jpg" // support jpg，bmp，gif，png
	// --output,    -o // if filename is '-',will output to stdout
	defaultOutput string = "pi_image_%04d." + defaultEncoding
	// --keypress,    -k        Keypress mode

	defaultMode Mode   = Once
	defaultPath string = "./image/"
	// --verbose,    -v
	// --latest,    -l
	// --framestart,    -fs
	// --datetime,    -dt
	// --timestamp,    -ts
	// --thumb,    -th        Set thumbnail parameters (x:y:quality) defaults are a size of 64x48 at quality 35. can set --thumb none.
	// --demo,    -d        Run a demo mode <milliseconds>
	// --restart,    -rs
	// --exif,    -x        EXIF tag to apply to captures (format as 'key=value'). can set --exif none.
	// --gpsdexif,    -gps      gps location. need libgps.so
	// --fullpreview,    -fp        Full preview mode. max fps = 15fps. dev feature.
	// --signal,    -s        Signal mode

	// -----------------raspiyuv-----------------//
	// --rgb,    -rgb        Save uncompressed data as RGB888
	// --luma,    -y
	// --bgr,    -bgr
)

type Observer struct {
	model.Observer
	width     string
	height    string
	quality   string
	obMode    Mode
	obDir     string
	encoding  string
	output    string
	timeout   string
	timelapse string
	calfunc   context.CancelFunc
	osCmd     *exec.Cmd
}

type Option func(*Observer)
type Mode int

const (
	Once   Mode = 1
	Manual Mode = 2
	Auto   Mode = 3
)

func OptionMode(m Mode) func(*Observer) {
	return func(ob *Observer) {
		ob.obMode = m
	}
}

func OptionDir(path string) func(*Observer) {
	return func(ob *Observer) {
		ob.obDir = path
	}
}

func OptionSize(width, height int) func(*Observer) {
	return func(ob *Observer) {
		if width < 64 || height < 64 || width > 1920 || height > 1080 {
			panic(errors.Errorf("size err.width:%d,height:%d", width, height))
		}
		ob.width = fmt.Sprint(width)
		ob.height = fmt.Sprint(height)
	}
}
func OptionQuality(quality int) func(*Observer) {
	return func(ob *Observer) {
		if quality < 0 || quality > 100 {
			panic(errors.Errorf("quality must >=0 || <=100, val:%d", quality))
		}
		ob.quality = fmt.Sprint(quality)
	}
}

func OptionTimeout(timeout int) func(*Observer) {
	return func(ob *Observer) {
		ob.timeout = fmt.Sprint(timeout)
	}
}

func OptionTimelapse(timelapse int) func(*Observer) {
	return func(ob *Observer) {
		ob.timelapse = fmt.Sprint(timelapse)
	}
}

// new image observer
func NewObserver(opts ...Option) *Observer {
	i := &Observer{
		width:     defaultWidth,
		height:    defaultHeight,
		quality:   defaultQuality,
		encoding:  defaultEncoding,
		output:    defaultOutput,
		timeout:   defaultTimeout,
		timelapse: defaultTimelapse,
		obMode:    defaultMode,
		obDir:     defaultPath,
	}
	// options
	for _, opt := range opts {
		opt(i)
	}
	if err := os.MkdirAll(i.obDir, os.ModePerm); err != nil {
		panic(err)
	}
	return i
}
func (i *Observer) buildArgs() []string {
	var args []string
	if i.width != "" {
		args = append(args, "-w", i.width)
	}
	if i.height != "" {
		args = append(args, "-h", i.height)
	}
	if i.quality != "" {
		args = append(args, "-q", i.quality)
	}
	if i.encoding != "" {
		args = append(args, "-e", i.encoding)
	}
	if i.output != "" {
		args = append(args, "-o", i.obDir+i.output)
	}
	switch i.obMode {
	case Once:
		if i.timeout != "" {
			args = append(args, "-t", i.timeout)
		}
	case Manual:
		args = append(args, "-t", "0")
		args = append(args, "-k")
	case Auto:
		args = append(args, "-tl", i.timelapse)
	}
	return args
}

func (i *Observer) Start() error {
	if err := i.run(); err != nil {
		return err
	}
	return nil
}

func (i *Observer) run() error {
	if i.obMode == Once {
		return errors.New("observer is not manual or auto mode.")
	}
	ctx, calfunc := context.WithCancel(context.Background())
	i.calfunc = calfunc
	i.osCmd = exec.CommandContext(ctx, raspistill, i.buildArgs()...)
	if err := i.osCmd.Start(); err != nil {
		return err
	}
	return nil
}

func (i *Observer) runOnce() error {
	if i.obMode != Once {
		return errors.New("observer is not once mode.")
	}
	i.osCmd = exec.CommandContext(context.Background(), raspistill, i.buildArgs()...)
	if err := i.osCmd.Run(); err != nil {
		return err
	}
	return nil
}

func (i *Observer) Shoot() error {
	switch i.obMode {
	case Once:
		return i.runOnce()
	case Manual:
		key.Key().SetKeys(keybd_event.VK_ENTER)
		if err := key.Key().Press(); err != nil {
			return err
		}
		time.Sleep(10 * time.Millisecond)
		if err := key.Key().Release(); err != nil {
			return err
		}
		return nil
	default:
		return nil
	}
}

func (i *Observer) Stop() error {
	i.calfunc()
	if err := i.osCmd.Wait(); err != nil {
		return err
	}
	return nil
}

func (i *Observer) AddSub(s model.Subscriber) {

}
