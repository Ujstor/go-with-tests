package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, url []string) map[string]bool {
	resault := make(map[string]bool)

	for _, url := range url {
		resault[url] = wc(url)
	}
	return resault
}