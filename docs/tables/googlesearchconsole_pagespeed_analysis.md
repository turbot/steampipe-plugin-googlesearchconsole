---
title: "Steampipe Table: googlesearchconsole_pagespeed_analysis - Query PageSpeed analysis per page on the site using SQL"
description: "Allows users to query PageSpeed analysis per page on the site, including details about each metric."
---

# Table: googlesearchconsole_pagespeed_analysis - Query PageSpeed analysis per page on the site using SQL

PageSpeed Insights (PSI) reports on the user experience of a page on both mobile and desktop devices, and provides suggestions on how that page may be improved.

## Table Usage Guide

The `googlesearchconsole_pagespeed_analysis` table allows users to analyze PageSpeed metrics for each page on their site, and compare the performance of pages across different devices.

**Required fields:**
  - `sitemap_url`: The URL of the sitemap that was submitted to Google Search Console. **Example:** `https://www.example.com/sitemap.xml`

## Examples

### Basic pagespeed analysis info
Retrieve essential PageSpeed insights for your site from its sitemap, including metrics like CLS, TTFB, FCP, LCP, FID, and INP, crucial for optimizing loading times and user experience.

```sql+postgres
select
  sitemap_url,
  loc,
  overall_loading_experience,
  analysis_utc_timestamp,
  cls,
  ttfb,
  fcp,
  lcp,
  fid,
  inp
from
  googlesearchconsole_pagespeed_analysis
where
  sitemap_url = 'https://example.io/sitemap-0.xml';
```

```sql+sqlite
select
  sitemap_url,
  loc,
  overall_loading_experience,
  analysis_utc_timestamp,
  cls,
  ttfb,
  fcp,
  lcp,
  fid,
  inp
from
  googlesearchconsole_pagespeed_analysis
where
  sitemap_url = 'https://example.io/sitemap-0.xml';
```

### Get Cumulative Layout Shift (CLS) for a specific page
Fetch detailed CLS data for https://example.io/, including scores, percentiles, and bucket ranges, to improve visual stability and enhance the user's visual experience on your page.

```sql+postgres
select
  loc,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/';
```

```sql+sqlite
select
  loc,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/';
```

### Get Cumulative Layout Shift (CLS) for a specific page in Mobile
Optimize mobile user experience by analyzing mobile-specific CLS data for https://example.io/, crucial for maintaining content stability on mobile devices.

```sql+postgres
select
  loc,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/'
  and strategy = 'MOBILE';
```

```sql+sqlite
select
  loc,
  cls,
  cls_percentile,
  cls_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/'
  and strategy = 'MOBILE';
```

### Get First Contentful Paint (FCP) for a specific page
This query provides insights into the First Contentful Paint (FCP) for https://example.io/, a key metric for evaluating and enhancing the perceived speed of content rendering on a page.

```sql+postgres
select
  loc,
  fcp,
  fcp_percentile,
  fcp_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/';
```

```sql+sqlite
select
  loc,
  fcp,
  fcp_percentile,
  fcp_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/';
```

### Get First Input Delay (FID) for a specific page
Improve page interactivity with detailed First Input Delay (FID) metrics for https://example.io/, essential for assessing and enhancing user interaction responsiveness.

```sql+postgres
select
  loc,
  fid,
  fid_percentile,
  fid_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/';
```

```sql+sqlite
select
  loc,
  fid,
  fid_percentile,
  fid_bucket_range
from
  googlesearchconsole_pagespeed_analysis
where
  loc = 'https://example.io/';
```