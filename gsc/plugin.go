package gsc

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-gsc",
		DefaultTransform: transform.FromCamel().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"gsc_indexing_status":               tableGSCIndexingStatus(ctx),
			"gsc_pagespeed_analysis":            tableGSCPagespeedAnalysis(ctx),
			"gsc_pagespeed_analysis_aggregated": tableGSCPagespeedAnalysisAggregated(ctx),
			"gsc_site":                          tableGSCSite(ctx),
			"gsc_sitemap":                       tableGSCSitemap(ctx),
		},
	}
	return p
}
