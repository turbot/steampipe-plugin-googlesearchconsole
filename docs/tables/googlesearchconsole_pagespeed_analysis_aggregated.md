---
title: "Steampipe Table: googlesearchconsole_pagespeed_analysis_aggregated - Query PageSpeed analysis on the site using SQL"
description: "Allows users to query PageSpeed analysis on the site, including details about each metric."
---

# Table: googlesearchconsole_pagespeed_analysis_aggregated - Query PageSpeed analysis on the site using SQL

PageSpeed Insights (PSI) reports on the user experience of a page on both mobile and desktop devices, and provides suggestions on how that page may be improved.

## Table Usage Guide

The `googlesearchconsole_pagespeed_analysis_aggregated` table runs PageSpeed analysis on the site, and returns aggregated PageSpeed scores.

**Required fields:**
  - `site_url`: The URL of the property as defined in Search Console. **Examples:** `http://www.example.com/` for a URL-prefix property, or `sc-domain:example.com` for a Domain property

## Examples

### Basic pagespeed analysis info
This query fetches the overall loading experience and individual PageSpeed metrics for a site, providing insights into the user experience and performance of the site.

```sql+postgres
select
  site_url,
  overall_loading_experience,
  analysis_utc_timestamp,
  cls,
  ttfb,
  fcp,
  lcp,
  fid,
  inp
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

```sql+sqlite
select
  site_url,
  overall_loading_experience,
  analysis_utc_timestamp,
  cls,
  ttfb,
  fcp,
  lcp,
  fid,
  inp
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

### Get aggregated Cumulative Layout Shift (CLS) for a site
This SQL query fetches the Cumulative Layout Shift (CLS) scores, percentile rankings, and score distribution for your site, essential for assessing visual stability and improving user experience.

```sql+postgres
select
  site_url,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

```sql+sqlite
select
  site_url,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

### Get aggregated Cumulative Layout Shift (CLS) for a site in Mobile
Optimize mobile user experience with this query that retrieves mobile-specific CLS data for your site, highlighting the need for stable content on mobile devices.

```sql+postgres
select
  site_url,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/'
  and strategy = 'MOBILE';
```

```sql+sqlite
select
  site_url,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/'
  and strategy = 'MOBILE';
```

### Get aggregated First Contentful Paint (FCP) for a site
This query provides First Contentful Paint (FCP) metrics for your site, key to understanding and enhancing perceived page load speed, a critical factor in SEO and user satisfaction.

```sql+postgres
select
  site_url,
  fcp,
  fcp_percentile,
  fcp_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

```sql+sqlite
select
  site_url,
  fcp,
  fcp_percentile,
  fcp_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

### Get aggregated First Input Delay (FID) for a site
Identify and improve interactivity on your site by fetching First Input Delay (FID) scores, a vital metric for enhancing responsiveness and user engagement.

```sql+postgres
select
  site_url,
  fid,
  fid_percentile,
  fid_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  lsite_urloc = 'https://example.io/';
```

```sql+sqlite
select
  site_url,
  fid,
  fid_percentile,
  fid_bucket_range
from
  googlesearchconsole_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```