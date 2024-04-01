package gsc

import (
	"context"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"google.golang.org/api/pagespeedonline/v5"
	"google.golang.org/api/searchconsole/v1"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// getSearchConsoleSessionConfig returns the client options for the searchconsole service
func getSearchConsoleSessionConfig(ctx context.Context, d *plugin.QueryData) ([]option.ClientOption, error) {
	opts := []option.ClientOption{}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var credentialContent string
	gscConfig := GetConfig(d.Connection)

	if gscConfig.Credentials != nil {
		credentialContent = *gscConfig.Credentials
	}

	// If credential path provided, use domain-wide delegation
	if credentialContent != "" {
		ts, err := getSearchConsoleTokenSource(ctx, d)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(ts))
		return opts, nil
	}

	return nil, nil
}

// getSearchConsoleTokenSource returns the token source for the searchconsole service
func getSearchConsoleTokenSource(ctx context.Context, d *plugin.QueryData) (oauth2.TokenSource, error) {

	cacheKey := "gsc.token_source"
	if ts, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return ts.(oauth2.TokenSource), nil
	}

	gscConfig := GetConfig(d.Connection)

	credentialContent, err := pathOrContents(*gscConfig.Credentials)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON([]byte(credentialContent), searchconsole.WebmastersReadonlyScope)
	if err != nil {
		plugin.Logger(ctx).Error("getTokenSource", "Unable to parse service account key file to config: %v", err)
	}

	ts := config.TokenSource(ctx)

	return ts, nil
}

// getPagespeedSessionConfig returns the client options for the pagespeed service
func getPagespeedSessionConfig(ctx context.Context, d *plugin.QueryData) ([]option.ClientOption, error) {
	opts := []option.ClientOption{}

	// Get credential file path, and user to impersonate from config (if mentioned)
	var credentialContent string
	gscConfig := GetConfig(d.Connection)

	if gscConfig.Credentials != nil {
		credentialContent = *gscConfig.Credentials
	}

	// If credential path provided, use domain-wide delegation
	if credentialContent != "" {
		ts, err := getPagespeedTokenSource(ctx, d)
		if err != nil {
			return nil, err
		}
		opts = append(opts, option.WithTokenSource(ts))
		return opts, nil
	}

	return nil, nil
}

// getPagespeedTokenSource returns the token source for the pagespeed service
func getPagespeedTokenSource(ctx context.Context, d *plugin.QueryData) (oauth2.TokenSource, error) {

	cacheKey := "pagespeed.token_source"
	if ts, ok := d.ConnectionCache.Get(ctx, cacheKey); ok {
		return ts.(oauth2.TokenSource), nil
	}

	gscConfig := GetConfig(d.Connection)

	credentialContent, err := pathOrContents(*gscConfig.Credentials)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON([]byte(credentialContent), pagespeedonline.OpenIDScope)
	if err != nil {
		plugin.Logger(ctx).Error("getTokenSource", "Unable to parse service account key file to config: %v", err)
	}

	ts := config.TokenSource(ctx)

	return ts, nil
}

// getPageIndexingStatusService returns the indexing status of a page
func getPageIndexingStatusService(ctx context.Context, d *plugin.QueryData, pageURL string, siteUrl string) (*searchconsole.UrlInspectionResult, error) {
	// Create client
	opts, err := getSearchConsoleSessionConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := searchconsole.NewService(ctx, opts...)
	if err != nil {
		plugin.Logger(ctx).Error("listGSCSites", "connection_error", err)
		return nil, err
	}

	req := searchconsole.InspectUrlIndexRequest{
		InspectionUrl: pageURL,
		SiteUrl:       siteUrl,
	}
	resp, err := svc.UrlInspection.Index.Inspect(&req).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp.InspectionResult, nil
}

// getPagespeedAnalysisService returns the pagespeed analysis of a page
func getPagespeedAnalysisService(ctx context.Context, d *plugin.QueryData, pageURL string, strategy string) (*pagespeedonline.PagespeedApiPagespeedResponseV5, error) {
	// Create client
	opts, err := getPagespeedSessionConfig(ctx, d)
	if err != nil {
		return nil, err
	}

	// Create service
	svc, err := pagespeedonline.NewService(ctx, opts...)
	if err != nil {
		plugin.Logger(ctx).Error("listGSCSites", "connection_error", err)
		return nil, err
	}

	resp, err := svc.Pagespeedapi.Runpagespeed(pageURL).Strategy(strings.ToUpper(strategy)).Fields("id,loadingExperience,analysisUTCTimestamp").Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
