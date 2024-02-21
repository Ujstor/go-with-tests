package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, url []string) map[string]bool {
	resault := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range url {

		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(url); i++ {
		r := <-resultChannel
		resault[r.string] = r.bool
	}

	return resault
}