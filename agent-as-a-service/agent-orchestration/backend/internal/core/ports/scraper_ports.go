package ports

import (
	"context"
	"sync"
)

type IScraper interface {
	FetchLinksFromDomain(ctx context.Context, w *sync.WaitGroup) chan string
	FetchLinksFromDomainByRod(ctx context.Context, w *sync.WaitGroup) chan string
	ContentHtmlByUrl(ctx context.Context, rawUrl string) (string, error)
}
