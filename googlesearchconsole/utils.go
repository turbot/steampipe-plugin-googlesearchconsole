package googlesearchconsole

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/mitchellh/go-homedir"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	sitemapper "github.com/yterajima/go-sitemap"
	"google.golang.org/api/pagespeedonline/v5"
	"google.golang.org/api/searchconsole/v1"
)

var (
	statusPerUrl            = make(map[string]*searchconsole.UrlInspectionResult)
	pagespeedAnalysisPerUrl = make(map[string]*pagespeedonline.PagespeedApiPagespeedResponseV5)
	mutex                   sync.Mutex
)

// Returns the content of given file, or the inline JSON credential as it is
func pathOrContents(poc string) (string, error) {
	if len(poc) == 0 {
		return poc, nil
	}

	path := poc
	if path[0] == '~' {
		var err error
		path, err = homedir.Expand(path)
		if err != nil {
			return path, err
		}
	}

	// Check for valid file path
	if _, err := os.Stat(path); err == nil {
		contents, err := os.ReadFile(path)
		if err != nil {
			return string(contents), err
		}
		return string(contents), nil
	}

	// Return error if content is a file path and the file doesn't exist
	if len(path) > 1 && (path[0] == '/' || path[0] == '\\') {
		return "", fmt.Errorf("%s: no such file or dir", path)
	}

	// Return the inline content
	return poc, nil
}

// createBatches divides the slice into smaller slices of the given size.
func createBatches(urls []sitemapper.URL, size int) [][]sitemapper.URL {
	var batches [][]sitemapper.URL
	for size < len(urls) {
		urls, batches = urls[size:], append(batches, urls[0:size:size])
	}
	batches = append(batches, urls)
	return batches
}

// processPageIndexingStatusBatch processes a batch of URLs concurrently.
func processPageIndexingStatusBatch(ctx context.Context, d *plugin.QueryData, siteUrl string, urls []sitemapper.URL, batchIndex int, wg *sync.WaitGroup) {
	var batchWG sync.WaitGroup
	batchWG.Add(len(urls))

	for _, url := range urls {
		go func(url sitemapper.URL) {
			defer batchWG.Done()
			status, err := getPageIndexingStatusService(ctx, d, url.Loc, siteUrl)
			if err != nil {
				plugin.Logger(ctx).Error("Error fetching status for %s: %v\n", url.Loc, err)
				return
			}

			result := status

			mutex.Lock()
			statusPerUrl[url.Loc] = result
			mutex.Unlock()
		}(url)
	}

	batchWG.Wait()
	wg.Done()

	plugin.Logger(ctx).Info("Batch %d complete\n", batchIndex+1)
}

// processPagespeedAnalysisBatch processes a batch of URLs concurrently.
func processPagespeedAnalysisBatch(ctx context.Context, d *plugin.QueryData, strategy string, urls []sitemapper.URL, batchIndex int, wg *sync.WaitGroup) {
	var batchWG sync.WaitGroup
	batchWG.Add(len(urls))

	for _, url := range urls {
		go func(url sitemapper.URL) {
			defer batchWG.Done()
			status, err := getPagespeedAnalysisService(ctx, d, url.Loc, strategy)
			if err != nil {
				plugin.Logger(ctx).Error("Error fetching status for %s: %v\n", url.Loc, err)
				// return
			}

			result := status

			mutex.Lock()
			pagespeedAnalysisPerUrl[url.Loc] = result
			mutex.Unlock()
		}(url)
	}

	batchWG.Wait()
	wg.Done()

	plugin.Logger(ctx).Info("Batch %d complete\n", batchIndex+1)
}
