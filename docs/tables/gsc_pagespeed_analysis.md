# Table: googlesearchconsole_pagespeed_analysis

Runs PageSpeed analysis on the site, and returns PageSpeed scores, and other information.

Required fields:
  - sitemap_url: The URL of the sitemap that was submitted to Google Search Console. **Example:** `https://www.example.com/sitemap.xml`

## Examples

### Basic info

```sql
select
  sitemap_url,
  loc,
  overall_loading_experience,
  analysis_utc_timestamp,
  cls,
  ttf,
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

```sql
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

```sql
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

```sql
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

```sql
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