package display

import (
	"fmt"
	"observer/station"
)

type HeatIndexDisplay struct {
	heatIndex float32
}

func NewHeatIndexDisplay(weatherData *station.WeatherData) *HeatIndexDisplay {
	var heatIndexDisplay HeatIndexDisplay
	weatherData.RegisterObserver(&heatIndexDisplay)
	return &heatIndexDisplay
}

func (h *HeatIndexDisplay) computeHeatIndex(t, rh float32) float32 {
	return (16.923 + (0.185212 * t) + (5.37941 * rh) - (0.100254 * t * rh) +
		(0.00941695 * (t * t)) + (0.00728898 * (rh * rh)) +
		(0.000345372 * (t * t * rh)) - (0.000814971 * (t * rh * rh)) +
		(0.0000102102 * (t * t * rh * rh)) - (0.000038646 * (t * t * t)) + (0.0000291583 *
		(rh * rh * rh)) + (0.00000142721 * (t * t * t * rh)) +
		(0.000000197483 * (t * rh * rh * rh)) - (0.0000000218429 * (t * t * t * rh * rh)) +
		0.000000000843296*(t*t*rh*rh*rh)) -
		(0.0000000000481975 * (t * t * t * rh * rh * rh))
}

func (h *HeatIndexDisplay) Update(temp, humidity, pressure float32) {
	h.heatIndex = h.computeHeatIndex(temp, humidity)
	h.Display()
}

func (h *HeatIndexDisplay) Display() {
	fmt.Printf("Heat index is : %.2f\n", h.heatIndex)
}
