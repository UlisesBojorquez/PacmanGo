package pacman

import (
	"log"

	pacmansounds "github.com/UlisesBojorquez/PacmanGo/sounds"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

type sounds struct {
	audioContext   *audio.Context
	sirenPlayer    *audio.Player
	wailPlayer     *audio.Player
	eatGhostPlayer *audio.Player
	deathPlayer    *audio.Player
	entrancePlayer *audio.Player
	applausePlayer *audio.Player
	on             bool
}

const (
	sampleRate = 44100
)

func newSounds() *sounds {
	audioContext, err := audio.NewContext(sampleRate)
	if err != nil {
		log.Fatal(err)
	}
	s := &sounds{
		audioContext: audioContext,
	}

	s.sirenPlayer = s.newPlayer(pacmansounds.Siren_wav)
	s.wailPlayer = s.newPlayer(pacmansounds.Wail_wav)
	s.eatGhostPlayer = s.newPlayer(pacmansounds.EatGhost_wav)
	s.deathPlayer = s.newPlayer(pacmansounds.Death_wav)
	s.entrancePlayer = s.newPlayer(pacmansounds.Beginning_wav)
	s.applausePlayer = s.newPlayer(pacmansounds.Applause_wav)

	if s.on {
		s.sirenPlayer.SetVolume(0)
		s.wailPlayer.SetVolume(0)
		s.eatGhostPlayer.SetVolume(0)
		s.deathPlayer.SetVolume(0)
		s.entrancePlayer.SetVolume(0)
		s.applausePlayer.SetVolume(0)
	} else {
		s.sirenPlayer.SetVolume(0.2)
		s.wailPlayer.SetVolume(0.05)
		s.eatGhostPlayer.SetVolume(0.05)
		s.deathPlayer.SetVolume(0.05)
		s.entrancePlayer.SetVolume(0.2)
		s.applausePlayer.SetVolume(0.2)
	}
	return s
}

func (s *sounds) newPlayer(b []byte) *audio.Player {
	p, err := audio.NewPlayer(s.audioContext, s.load(b))
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func (s *sounds) turnOff() {
	s.on = false
	s.sirenPlayer.SetVolume(0.2)
	s.wailPlayer.SetVolume(0.05)
	s.eatGhostPlayer.SetVolume(0.05)
	s.deathPlayer.SetVolume(0.05)
}

func (s *sounds) load(b []byte) *wav.Stream {
	stream, err := wav.Decode(s.audioContext, audio.BytesReadSeekCloser(b))
	if err != nil {
		log.Fatal(err)
	}
	return stream
}

func (s *sounds) playSiren() {
	if s.canPlaySiren() && !s.sirenPlayer.IsPlaying() {
		s.sirenPlayer.Rewind()
		s.sirenPlayer.Play()
	}
}

func (s *sounds) canPlaySiren() bool {
	d := [...]*audio.Player{
		s.wailPlayer,
		s.deathPlayer,
		s.entrancePlayer,
		s.applausePlayer,
	}

	for _, v := range d {
		if v.IsPlaying() {
			return false
		}
	}
	return true
}

func (s *sounds) playWail() {
	s.sirenPlayer.Pause()
	s.wailPlayer.Rewind()
	s.wailPlayer.Play()
}

func (s *sounds) playEeatGhost() {
	if !s.eatGhostPlayer.IsPlaying() {
		s.eatGhostPlayer.Rewind()
		s.eatGhostPlayer.Play()
	}
}

func (s *sounds) playDeath() {
	s.sirenPlayer.Pause()
	if !s.deathPlayer.IsPlaying() {
		s.deathPlayer.Rewind()
		s.deathPlayer.Play()
	}
}

func (s *sounds) pause() {
	d := [...]*audio.Player{
		s.wailPlayer,
		s.deathPlayer,
		s.sirenPlayer,
		s.eatGhostPlayer,
		s.entrancePlayer,
		s.applausePlayer,
	}

	for _, v := range d {
		v.Pause()
	}
}

func (s *sounds) playEntrance() {
	s.sirenPlayer.Pause()
	s.entrancePlayer.Rewind()
	s.entrancePlayer.Play()
}

func (s *sounds) playApplause() {
	s.sirenPlayer.Pause()
	if s.applausePlayer.IsPlaying() {
		return
	}
	s.applausePlayer.Rewind()
	s.applausePlayer.Play()
}
