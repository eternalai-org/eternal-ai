package scraper

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

var (
	tracerTagName string = "scraper-kb"
	chanSize      int64  = 100
)

type scraper struct {
	url      string
	maxDepth int
	browser  *rod.Browser
	domain   string
}

func NewScraper(url string, maxDepth int) (ports.IScraper, error) {
	s := &scraper{}
	s.maxDepth = maxDepth
	s.url = url
	domain, err := utils.ExtractDomainFromUrl(url)
	if err != nil {
		return nil, err
	}
	s.domain = domain

	path, _ := launcher.LookPath()
	u, err := launcher.New().Bin(path).Headless(true).Launch()
	if err != nil {
		return nil, err
	}
	s.browser = rod.New().ControlURL(u).MustConnect()
	return s, nil
}

func (s *scraper) ContentHtmlByUrl(ctx context.Context, rawUrl string) (string, error) {
	logger.Info(tracerTagName, "content_html_by_url", zap.Any("url", rawUrl))
	page := s.browser.MustPage(rawUrl)
	//defer s.browser.Close()
	page.MustWaitLoad()
	logger.Info(tracerTagName, "content_html_by_url#MustWaitLoad", zap.Any("url", rawUrl))

	i := 0
	for i <= 3 {
		page.MustEval(`() => window.scrollTo(0, document.body.scrollHeight)`)
		time.Sleep(1 * time.Second)
		i += 1
	}

	logger.Info(tracerTagName, "content_html_by_url#window.scrollTo", zap.Any("url", rawUrl))
	page.MustWaitLoad()
	page.MustWaitStable()
	page.MustEval(`() => document.querySelectorAll("[crossorigin]").forEach((el) => el.removeAttribute('crossorigin'))`)

	htmlStr, err := page.HTML()
	if err != nil {
		logger.Error(tracerTagName, "Failed to get HTML: %v", zap.Error(err))
		return "", err
	}

	htmlStr, err = utils.MinifyHTML(htmlStr)
	if err != nil {
		logger.Error(tracerTagName, "Failed to MinifyHTML: %v", zap.Error(err))
		return "", err
	}

	return htmlStr, nil
}

func (s *scraper) FetchLinksFromDomainByRod(ctx context.Context, w *sync.WaitGroup) chan string {
	logger.Info(tracerTagName, "fetch-all-url-from-domain", zap.Any("domain-url", s.url))
	outChan := make(chan string, chanSize)
	go func() {
		defer func() {
			w.Done()
			close(outChan)
		}()

		domainURL, err := url.Parse(s.url)
		if err != nil {
			logger.Error(tracerTagName, "invalid-url", zap.Error(err), zap.Any("url", s.url))
			return
		}
		domain := domainURL.Hostname()
		logger.Info(tracerTagName, "fetch-links-from-domain-by-rod", zap.Any("domain", domain), zap.Any("input-url", s.url))
		l := launcher.New().Headless(true)
		browser := rod.New().ControlURL(l.MustLaunch()).MustConnect()
		defer browser.MustClose()

		page := browser.MustPage(s.url)
		page.MustWaitLoad()

		var (
			visited = make(map[string]bool)
			queue   = []string{s.url}
			depth   = 0
		)

		for len(queue) > 0 && depth <= s.maxDepth {
			nextQueue := []string{}
			for _, currentURL := range queue {
				if visited[currentURL] {
					continue
				}

				visited[currentURL] = true
				logger.Info(tracerTagName, "visiting", zap.Any("url", currentURL))

				page := browser.MustPage(currentURL)
				page.MustWaitStable()
				page.MustWaitLoad()

				links := page.MustElements("a[href]")
				// links = append(links, page.MustElement("img[src]"))
				for _, linkElement := range links {
					href, err := linkElement.Attribute("href")
					if err != nil {
						logger.Error(tracerTagName, "error getting href", zap.Error(err))
						continue
					}

					// if href == nil || *href == "" {
					// 	href, err = linkElement.Attribute("src")
					// 	if err != nil {
					// 		logger.Error(tracerTagName, "error getting href", zap.Error(err))
					// 		continue
					// 	}
					// }

					if href == nil || *href == "" {
						continue
					}

					absoluteURL, err := url.Parse(*href)
					if err != nil {
						absoluteURL, err = url.Parse(currentURL)
						if err != nil {
							logger.Error(tracerTagName, "error parsing current url", zap.Error(err))
							continue
						}

						relativeUrl, err := url.Parse(*href)
						if err != nil {
							logger.Error(tracerTagName, "error parsing relative url", zap.Error(err))
							continue
						}

						absoluteURL = absoluteURL.ResolveReference(relativeUrl)
					} else if !absoluteURL.IsAbs() {
						absoluteURL, err = url.Parse(currentURL)
						if err != nil {
							logger.Error(tracerTagName, "error parsing current url", zap.Error(err))
							continue
						}
						relativeUrl, err := url.Parse(*href)
						if err != nil {
							logger.Error(tracerTagName, "error parsing relative url", zap.Error(err))
							continue
						}
						absoluteURL = absoluteURL.ResolveReference(relativeUrl)
					}

					link := absoluteURL.String()
					link, _ = utils.RemoveFragment(link)
					if strings.HasSuffix(link, "/") {
						link = link[:len(link)-1]
					}

					if strings.HasSuffix(link, "?") {
						link = link[:len(link)-1]
					}
					// || strings.Contains(link, "cdn.prod.website-files.com")
					if strings.Contains(link, domain) {
						outChan <- link
						if !visited[link] {
							nextQueue = append(nextQueue, link)
						}
					} else {
						logger.Info(tracerTagName, "invalid-url", zap.Any("link", link))
					}
				}
			}
			queue = nextQueue
			depth++
		}
	}()
	return outChan
}

func (s *scraper) FetchLinksFromDomain(ctx context.Context, w *sync.WaitGroup) chan string {
	logger.Info(tracerTagName, "fetch-all-url-from-domain", zap.Any("domain-url", s.url))
	outChan := make(chan string, chanSize)
	go func() {
		defer w.Done()
		defer close(outChan)

		domain, _ := utils.ExtractDomainFromUrl(s.url)
		c := colly.NewCollector(
			colly.MaxDepth(s.maxDepth),
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
			colly.AllowedDomains(
				domain,
			),
		)

		c.OnRequest(func(r *colly.Request) {
			logger.Info(tracerTagName, "visiting", zap.Any("url", r.URL.String()))
		})

		c.OnError(func(r *colly.Response, err error) {
			logger.Error(tracerTagName, "scraper_error", zap.Error(err), zap.Any("request_url", r.Request.URL), zap.Any("failed with response", r))
		})

		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			if link == "" {
				return
			}

			link, err := utils.CreateValidLink(e.Request.URL.String(), link)
			if err != nil {
				return
			}

			link, _ = utils.RemoveFragment(link)
			if strings.HasSuffix(link, "/") {
				link = link[:len(link)-1]
			}

			if strings.HasSuffix(link, "?") {
				link = link[:len(link)-1]
			}

			if strings.Contains(link, domain) {
				outChan <- link
				c.Visit(e.Request.AbsoluteURL(link))
			} else {
				logger.Info(tracerTagName, "invalid-url", zap.Any("link", link))
			}
		})

		req := &colly.Request{}
		if err := c.Visit(req.AbsoluteURL(s.url)); err != nil {
			logger.Error(tracerTagName, fmt.Sprintf("visit website: %s has err", s.url), zap.Error(err))
			return
		}
	}()
	return outChan
}
