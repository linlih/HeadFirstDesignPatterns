package main

func main() {
	amp := NewAmplifier("Top-O-Line Amplifier")
	tuner := NewTuner("Top-O-Line AM/FM Tuner")
	dvd := NewDvdPlayer("Top-O-Line DVD Player")
	cd := NewCdPlayer("Top-O-Line CD Player")
	projector := NewProjector("Top-O-Line Projector")
	lights := NewTheaterLights("Theater Ceiling Lights")
	screen := NewScreen("Theater Screen")
	popper := NewPopcornPopper("Popcorn Popper")

	homeTheater := NewHomeTheaterFacade(amp, tuner, dvd, cd,
		projector, lights, screen, popper)

	homeTheater.WatchMovie("Raiders of the Lost Ark")
	homeTheater.EndMovie()

	homeTheater.ListenToCd("Micheal Jackson")
	homeTheater.EndCd()

	homeTheater.ListenToRadio(1234)
	homeTheater.EndRadio()

}
