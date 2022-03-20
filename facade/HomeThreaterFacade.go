package main

import "fmt"

type HomeTheaterFacade struct {
	amp       *Amplifier
	tuner     *Tuner
	dvd       *DvdPlayer
	cd        *CdPlayer
	projector *Projector
	lights    *TheaterLights
	screen    *Screen
	popper    *PopcornPopper
}

func NewHomeTheaterFacade(amp *Amplifier, tuner *Tuner, dvd *DvdPlayer, cd *CdPlayer, projector *Projector, lights *TheaterLights, screen *Screen, popper *PopcornPopper) *HomeTheaterFacade {
	return &HomeTheaterFacade{amp: amp, tuner: tuner, dvd: dvd, cd: cd, projector: projector, lights: lights, screen: screen, popper: popper}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie ...")
	h.popper.On()
	h.popper.Pop()
	h.lights.Dim(10)
	h.screen.Down()
	h.projector.On()
	h.projector.WideScreenMode()
	h.amp.On()
	h.amp.SetDvd(*h.dvd)
	h.amp.SetSurroundSound()
	h.amp.SetVolume(5)
	h.dvd.On()
	h.dvd.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.popper.Off()
	h.lights.On()
	h.screen.Up()
	h.projector.Off()
	h.amp.Off()
	h.dvd.Stop()
	h.dvd.Eject()
	h.dvd.Off()
}

func (h *HomeTheaterFacade) ListenToCd(cdTitle string) {
	fmt.Println("Get ready for an audiopile experence...")
	h.lights.On()
	h.amp.On()
	h.amp.SetVolume(5)
	h.amp.SetCd(*h.cd)
	h.amp.SetStereoSound()
	h.cd.On()
	h.cd.Play(cdTitle)
}

func (h *HomeTheaterFacade) EndCd() {
	fmt.Println("Shutting down CD...")
	h.amp.Off()
	h.amp.SetCd(*h.cd)
	h.cd.Eject()
	h.cd.Off()
}

func (h *HomeTheaterFacade) ListenToRadio(frequency float32) {
	fmt.Println("Tuning in the airwaves...")
	h.tuner.On()
	h.tuner.SetFrequency(frequency)
	h.amp.On()
	h.amp.SetVolume(5)
	h.amp.SetTuner(*h.tuner)
}

func (h *HomeTheaterFacade) EndRadio() {
	fmt.Println("Shutting down the tuner...")
	h.tuner.Off()
	h.amp.Off()
}
