package cmd

import (
	"bytes"
	"embed"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/spf13/cobra"
)

var fs embed.FS

var rootCmd = &cobra.Command{
	Use:   "dora",
	Short: "dora is a library to play dora.mp3",
	Long:  `Dora is a traditional Chinese percussion instruments of Buddhist origin. It plays a solemn sound.`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		isSilent, _ := cmd.Flags().GetBool("silent")
		if isSilent {
			println("🔔 Done")
		} else {
			runDora()
		}
	},
}

func runDora() {
	f, err := fs.ReadFile(filepath.Join("resources", "dora.mp3"))
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(ioutil.NopCloser(bytes.NewReader(f)))
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func Execute(embedFs embed.FS) {
	fs = embedFs
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("silent", "s", false, "Mute dora")
}
