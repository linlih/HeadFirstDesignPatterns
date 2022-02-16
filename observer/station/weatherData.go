package station

type WeatherData struct {
	observers   []IObserver
	temperature float32
	humidity    float32
	pressure    float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) RegisterObserver(o IObserver) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o IObserver) {
	var idx int
	for idx = 0; idx < len(w.observers); idx++ {
		if w.observers[idx] == o {
			break
		}
	}
	w.observers = append(w.observers[:idx], w.observers[idx+1:]...)
}

func (w *WeatherData) NotifyObserver() {
	for i := 0; i < len(w.observers); i++ {
		w.observers[i].Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementChanged() {
	w.NotifyObserver()
}

func (w *WeatherData) SetMeasurements(temperature, humidity, pressure float32) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementChanged()
}

func (w *WeatherData) GetTemperature() float32 {
	return w.temperature
}

func (w *WeatherData) GetHumidity() float32 {
	return w.humidity
}

func (w *WeatherData) GetPressure() float32 {
	return w.pressure
}
