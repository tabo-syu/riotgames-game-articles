package riotgamespatchnotes

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestLOL(t *testing.T) {
	for _, locale := range LOLLocales {
		t.Run(fmt.Sprintf("Test locale: %s", locale), func(t *testing.T) {
			notes, err := LOL(context.Background(), locale)
			if err != nil {
				t.Errorf("LOL() error = %v", err)

				return
			}

			for _, note := range notes {
				log.Printf(note.Title)
			}
		})

		time.Sleep(1 * time.Second)
	}
}

func TestTFT(t *testing.T) {
	for _, locale := range LOLLocales {
		t.Run(fmt.Sprintf("Test locale: %s", locale), func(t *testing.T) {
			notes, err := TFT(context.Background(), locale)
			if err != nil {
				t.Errorf("TFT() error = %v", err)

				return
			}

			for _, note := range notes {
				log.Printf(note.Title)
			}
		})

		time.Sleep(1 * time.Second)
	}
}
