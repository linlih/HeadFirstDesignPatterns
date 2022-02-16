package display

import (
	"fmt"
	"observer/station"
)

type CurrentConditionDisplay struct {
	temperature float32
	humidity    float32
}

func NewCurrentConditionDisplay(weatherData *station.WeatherData) *CurrentConditionDisplay {
	var currentConditionDisplay CurrentConditionDisplay
	weatherData.RegisterObserver(&currentConditionDisplay)
	return &currentConditionDisplay
}

func (c *CurrentConditionDisplay) Update(temp, humidity, pressure float32) {
	c.temperature = temp
	c.humidity = humidity
	c.Display()
}

func (c *CurrentConditionDisplay) Display() {
	fmt.Printf("Current conditions: %.2f degrees and %.2f %% humidity\n", c.temperature, c.humidity)
}
