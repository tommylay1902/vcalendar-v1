package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/gordonklaus/portaudio"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/tommylay1902/vcalendar/client"
	"github.com/tommylay1902/vcalendar/seed"
	"github.com/tommylay1902/vcalendar/wavwriter"
)

func main() {
	gc := client.InitializeClientGC()
	seed.SeedGCOperations()
	IgnoreAudioWarnings()
	qc := client.InitializeQdrantClient()

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	fmt.Println("Recording. Type 'exit' and press Enter to stop.")

	wavFmtChunk := wavwriter.Initialize(3, 16000, 16, 1) // 16-bit , 16kHz

	// Initialize PortAudio
	portaudio.Initialize()
	defer portaudio.Terminate()

	// Audio buffer
	in := make([]int16, 2048) // Larger buffer for better performance
	stream, err := portaudio.OpenDefaultStream(1, 0, float64(wavFmtChunk.SampleRate), len(in), in)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	defer stream.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// WebSocket connection to Vosk
	c, _, err := websocket.Dial(ctx, "ws://localhost:2700", nil)
	chk(err)
	defer c.Close(websocket.StatusNormalClosure, "")

	// Send configuration to Vosk
	config := map[string]any{
		"config": map[string]any{
			"sample_rate": 16000.0, // Vosk expects 16kHz
		},
	}

	writeCtx, wsCancel := context.WithTimeout(ctx, 100*time.Millisecond)
	err = wsjson.Write(writeCtx, c, config)
	chk(err)
	defer wsCancel()

	// Setup stop channel
	stopChan := make(chan struct{})
	doneChan := make(chan struct{})

	go ReadUserInput(stopChan)

	messageChan := make(chan any)
	errorChan := make(chan error)

	vc := InitVoskCommunication(ctx, c, stream, in)

	go vc.WriteWebsocket(messageChan, errorChan, doneChan)

	fmt.Println("Recording started...")
	// need to break up the logic in this, too many moving parts
	// go vc.RecordAudio(w, gc, qc,
	// 	messageChan, errorChan, stopChan)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		vc.RecordAudio(w, gc, qc, messageChan, errorChan, stopChan)
	}()

	// Wait for recording to finish
	<-stopChan
	wg.Wait()
	// Send EOF to Vosk
	if err := wsjson.Write(ctx, c, map[string]any{"eof": 1}); err != nil {
		log.Printf("Error sending EOF: %v", err)
	}

}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
