package utils

import (
	"io"
	"os"
	"os/exec"
	"time"
)

var FFpcmToMp3 = []string{"-hide_banner", "-loglevel", "info", "-acodec", "pcm_s16le", "-ac", "1", "-f", "s16le", "-sample_fmt", "s16", "-ar", "44100", "-i", "pipe:0", "-f", "mp3", "-c:a", "libmp3lame", "-ar", "44100", "-ac", "1", "-b:a", "64k", "-reservoir", "0", "-q:a", "0", "pipe:1"}
var FFfile = []string{"-hide_banner", "-loglevel", "error", "-i"}
var FFtoMP3 = []string{"-vn", "-ar", "44100", "-ac", "1", "-b:a", "64k", "-f", "mp3", "pipe:1"}

//var FFrawToWav = []string{"-hide_banner", "-loglevel", "error", "-i", "pipe:0", "-acodec", "pcm_s16le", "-ac", "1", "-f", "s16le", "-sample_fmt", "s16", "-ar", "16000", "pipe:1"}

func FFconvertChunk(command []string, inChan *chan []byte, outChan chan []byte, outChunkSize int, timeout time.Duration) error {
	cmd := exec.Command("ffmpeg", command...)

	cmd.Stderr = os.Stderr
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	// pipe input channel to ffmpeg
	go func() {
		timer := time.NewTimer(timeout)
		buf := []byte{}
		for {
			select {
			case data := <-*inChan:
				if data == nil {
					if len(buf) > 0 {
						in.Write(buf)
					}
					in.Close()
					return
				} else {
					buf = append(buf, data...)
					if len(buf) >= 4096 {
						in.Write(buf)
						buf = nil
					}
					timer.Reset(timeout)
				}
			case <-timer.C:
				in.Close()
				return
			}
		}
	}()
	// pipe ffmpeg to output channel
	go func() {
		for {
			data := make([]byte, outChunkSize)
			n, e := out.Read(data)
			if e == io.EOF {
				break
			} else if n > 0 {
				outChan <- data[0:n]
			}
		}
		outChan <- nil
	}()
	go func() {
		cmd.Wait()
		*inChan = nil
	}()
	return nil
}

func FFconvertFile(command []string, stopChan *chan []byte, ch chan []byte, chunkSize int) error {
	cmd := exec.Command("ffmpeg", command...)

	cmd.Stderr = os.Stderr
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		for {
			buf := make([]byte, chunkSize)
			n, e := out.Read(buf)
			if e == io.EOF {
				break
			}
			ch <- buf[0:n]
		}
		ch <- nil
	}()
	go func() {
		<-*stopChan
		cmd.Process.Kill()
	}()
	go func() {
		cmd.Wait()
		*stopChan = nil
	}()
	return nil
}
