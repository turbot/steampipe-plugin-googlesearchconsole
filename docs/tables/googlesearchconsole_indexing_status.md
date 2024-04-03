---
title: "Steampipe Table: googlesearchconsole_indexing_status - Query indexing status of all web pages using SQL"
description: "Explore the indexing status of web pages in Google's index, including live inspection results, using SQL queries."
---

# Table: googlesearchconsole_indexing_status - Query indexing status of all web pages using SQL

Results of index status inspection for either the live page or the version in Google's index, depending on whether you requested a live inspection or not

## Table Usage Guide

The `googlesearchconsole_indexing_status` table allows users to analyze the index status of pages on their site, including details on coverage state, indexing state, and referring URLs. It's essential for understanding how well your web pages are represented in Google Search.

**Required fields:**
  - `site_url`: The URL of the property as defined in Search Console. **Examples:** `http://www.example.com/` for a URL-prefix property, or `sc-domain:example.com` for a Domain property
  - `sitemap_url`: The URL of the sitemap that was submitted to Google Search Console. **Example:** `https://www.example.com/sitemap.xml`

## Examples

### Basic indexing status info
Retrieve basic details about the indexing status of pages, including the page's location (URL), coverage state, indexing state, and URLs referring to the page within the sitemap.

```sql+postgres
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

```sql+sqlite
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
Identify URLs within a sitemap that are not indexed by Google. This query helps in pinpointing pages that might need further optimization or review to meet Google's indexing requirements.

```sql+postgres
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

```sql+sqlite
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

### Get page count by indexing status
This query provides a count of pages grouped by their coverage state. It's useful for assessing the overall indexing health of your site and identifying potential areas for improvement.

```sql+postgres
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

```sql+sqlite
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