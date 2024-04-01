# Table: gsc_pagespeed_analysis_aggregated

Runs PageSpeed analysis on the site, and returns aggregated PageSpeed scores.

Required fields:
  - site_url: The URL of the property as defined in Search Console. **Examples:** `http://www.example.com/` for a URL-prefix property, or `sc-domain:example.com` for a Domain property

## Examples

### Basic info

```sql
select
  site_url,
  overall_loading_experience,
  analysis_utc_timestamp,
  cls,
  ttf,
  fcp,
  lcp,
  fid,
  inp
from
  gsc_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

### Get aggregated Cumulative Layout Shift (CLS) for a site

```sql
select
  site_url,
  cls,
  cls_percentile,
  cls_bucket_range
from
  gsc_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

### Get aggregated Cumulative Layout Shift (CLS) for a site in Mobile

```sql
select
  site_url,
  cls,
  cls_percentile,
  cls_bucket_range
from
  gsc_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/'
  and strategy = 'MOBILE';
```

### Get aggregated First Contentful Paint (FCP) for a site

```sql
select
  site_url,
  fcp,
  fcp_percentile,
  fcp_bucket_range
from
  gsc_pagespeed_analysis_aggregated
where
  site_url = 'https://example.io/';
```

### Get aggregated First Input Delay (FID) for a site

```sql
select
  site_url,
  fid,
  fid_percentile,
  fid_bucket_range
from
  gsc_pagespeed_analysis_aggregated
where
  lsite_urloc = 'https://example.io/';
```