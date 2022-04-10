package nginx

type Server interface {
	HandleRequest(string, string) (int, string)
}
