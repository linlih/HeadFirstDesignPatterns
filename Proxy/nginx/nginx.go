package nginx

type Nginx struct {
	application      *Application
	maxAllowedRequst int
	rateLimiter      map[string]int
}

func NewNginxServer() *Nginx {
	return &Nginx{
		application:      &Application{},
		maxAllowedRequst: 2,
		rateLimiter:      make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	alllowed := n.checkRateLimiting(url)
	if !alllowed {
		return 403, "Not Allowed"
	}
	return n.application.HandleRequest(url, method)
}

func (n *Nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequst {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}
