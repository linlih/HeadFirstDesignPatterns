package station

type IObserver interface {
	Update(temp, humidity, pressure float32)
}
