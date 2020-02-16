package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faiface/beep"
	"github.com/sirupsen/logrus"

	"github.com/MarinX/keylogger"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var resampler *beep.Resampler
var ctrl *beep.Ctrl
var vol *effects.Volume
var playing bool

func main() {

	var eatSignal = make(chan os.Signal)

	go func() {
		sig := <-eatSignal
		fmt.Printf("eating sig: %+v", sig)
	}()

	signal.Notify(eatSignal, syscall.SIGTERM)
	signal.Notify(eatSignal, syscall.SIGINT)
	signal.Notify(eatSignal, syscall.SIGQUIT)
	signal.Notify(eatSignal, syscall.SIGUSR1)
	signal.Notify(eatSignal, syscall.SIGUSR2)
	signal.Notify(eatSignal, syscall.SIGSTOP)
	signal.Notify(eatSignal, syscall.SIGTSTP)

	// find keyboard device, does not require a root permission
	keyboard := keylogger.FindKeyboardDevice()

	// check if we found a path to keyboard
	if len(keyboard) <= 0 {
		fmt.Printf("No keyboard found...\n")
		return
	}

	fmt.Printf("Found a keyboard at %s\n", keyboard)
	//this requires root
	k, err := keylogger.New(keyboard)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer k.Close()

	f, err := os.Open("sophia-mar-mar.mp3")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer f.Close()
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	vol := &effects.Volume{Streamer: resampler, Base: 2}

	speaker.Play(vol)
	speaker.Lock()
	counter := 0
	playing = false

	for {
		events := k.Read()

		for e := range events {

			switch e.Type {
			// check the input_event.go for more events
			case keylogger.EvKey:
				if e.KeyPress() {
					counter++
					//logrus.Println("[event] press key", e.KeyString())
					logrus.Println("num keys ", counter)
				}
				if e.KeyRelease() {
					counter--
					//logrus.Println("[event] release key", e.KeyString())
					logrus.Println("num keys ", counter)
				}
				break
			case keylogger.EvPwr:
				if e.KeyPress() {
					//logrus.Println("[event] press key", e.KeyString())
					//logrus.Println("num keys ", counter)
				}
				if e.KeyRelease() {
					//logrus.Println("[event] release key", e.KeyString())
					//logrus.Println("num keys ", counter)
				}
				break
			}
			//fmt.Printf("number of keys currently pressed %d\n", counter)
			if counter > 0 && !playing {
				speaker.Unlock()
				playing = true
				logrus.Println("[playing] %b", playing)
			} else if counter <= 0 && playing {
				speaker.Lock()
				playing = false
				logrus.Println("[playing] %b", playing)
			}
		}
	}
}
