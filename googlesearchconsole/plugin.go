package googlesearchconsole

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-googlesearchconsole",
		DefaultTransform: transform.FromCamel().NullIfZero(),
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			"googlesearchconsole_indexing_status":               tableGSCIndexingStatus(ctx),
			"googlesearchconsole_pagespeed_analysis":            tableGSCPagespeedAnalysis(ctx),
			"googlesearchconsole_pagespeed_analysis_aggregated": tableGSCPagespeedAnalysisAggregated(ctx),
			"googlesearchconsole_site":                          tableGSCSite(ctx),
			"googlesearchconsole_sitemap":                       tableGSCSitemap(ctx),
		},
	}
	return p
}
