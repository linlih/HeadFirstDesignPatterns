package station

type ISubject interface {
	RegisterObserver(o IObserver)
	RemoveObserver(o IObserver)
	NotifyObserver()
}
