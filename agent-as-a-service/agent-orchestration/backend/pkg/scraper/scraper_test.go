package scraper

import (
	"context"
	"sync"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	ctx     context.Context
	scraper ports.IScraper
}

func TestRun(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	logger.NewLogger("agents-ai-api", "test", "", true)
	s.ctx = context.Background()
}

func (s *Suite) TestContentHtmlByUrl() {
	var err error

	mainW := &sync.WaitGroup{}

	mainW.Add(1)
	s.scraper, err = NewScraper("https://www.ethdenver.com/", 5)
	if err != nil {
		panic(err)
	}

	outChan := s.scraper.FetchLinksFromDomainByRod(s.ctx, mainW)
	for l := range outChan {
		spew.Dump("url: " + l)
	}
	mainW.Wait()
}
