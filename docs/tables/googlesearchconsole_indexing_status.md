# Table: googlesearchconsole_indexing_status

Results of index status inspection for either the live page or the version in Google's index, depending on whether you requested a live inspection or not

Required fields:
  - site_url: The URL of the property as defined in Search Console. **Examples:** `http://www.example.com/` for a URL-prefix property, or `sc-domain:example.com` for a Domain property
  - sitemap_url: The URL of the sitemap that was submitted to Google Search Console. **Example:** `https://www.example.com/sitemap.xml`

## Examples

### Basic info

```sql
select
  loc,
  coverage_state,
  indexing_state,
  referring_urls
from
  googlesearchconsole_indexing_status
where
  sitemap_url = 'https://example.io/sitemap-0.xml' 
  and site_url = 'https://example.io/';
```

### List unindexed URLs

```sql
select
  loc,
  coverage_state,
  indexing_state
from
  googlesearchconsole_indexing_status
where
  sitemap_url = 'https://example.io/sitemap-0.xml' 
  and site_url = 'https://example.io/'
  and coverage_state <> 'Submitted and indexed';
```

### List overall status of the pages

```sql
select
  coverage_state,
  count(*) as page_count
from
  googlesearchconsole_indexing_status
where
  sitemap_url = 'https://example.io/sitemap-0.xml' 
  and site_url = 'https://example.io/'
group by
  coverage_state;
```