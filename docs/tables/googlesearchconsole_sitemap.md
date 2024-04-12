---
title: "Steampipe Table: googlesearchconsole_sitemap - Query Search Console sitemap submissions with SQL"
description: "Query and analyze sitemap submissions across all sites in Google Search Console, including details on types, submitted URLs, and errors."
---

# Table: googlesearchconsole_sitemap - Query Search Console sitemap submissions with SQL

Search Console sitemaps are XML files that list the URLs of a site, allowing webmasters to inform search engines about the pages on their site that are available for crawling. Sitemaps help search engines discover and index pages more efficiently.

## Table Usage Guide

This `googlesearchconsole_sitemap` table is instrumental for webmasters and SEO professionals to analyze the sitemaps they have submitted to Google Search Console. It helps in monitoring the health and status of these sitemaps, identifying submission errors, and understanding the scope of URLs covered.

## Examples

### Basic search console sitemap info
Retrieve basic information about all submitted sitemaps, including the site URL, sitemap path, type, and total number of URLs submitted within each sitemap. This query is ideal for getting an overview of sitemap submissions.

```sql+postgres
select
  site_url,
  path,
  type,
  c ->> 'submitted' as total_urls
from
  googlesearchconsole_sitemap,
  jsonb_array_elements(contents) as c;
```

```sql+sqlite
select
  site_url,
  path,
  type,
  json_extract(content.value, '$.submitted') as total_urls
from
  googlesearchconsole_sitemap,
  json_each(contents) as content;
```

### List sitemaps for a specific site
Filter the sitemaps for a particular site, providing focused insights on the sitemap submissions for that site. This can be useful for site-specific audits and optimization efforts.

```sql+postgres
select
  site_url,
  path,
  type,
  c ->> 'submitted' as total_urls
from
  googlesearchconsole_sitemap,
  jsonb_array_elements(contents) as c
where
  site_url = 'https://www.example.com/';
```

```sql+sqlite
select
  site_url,
  path,
  type,
  json_extract(content.value, '$.submitted') as total_urls
from
  googlesearchconsole_sitemap,
  json_each(contents) as content
where
  site_url = 'https://www.example.com/';
```

### List sitemaps for a specific domain
View the sitemap submissions for a domain-level property in Google Search Console. This query is particularly useful for managing sitemaps across an entire domain, including all subdomains and protocols.

```sql+postgres
select
  site_url,
  path,
  type,
  c ->> 'submitted' as total_urls
from
  googlesearchconsole_sitemap,
  jsonb_array_elements(contents) as c
where
  site_url = 'sc-domain:example.com';
```

```sql+sqlite
select
  site_url,
  path,
  type,
  json_extract(content.value, '$.submitted') as total_urls
from
  googlesearchconsole_sitemap,
  json_each(contents) as content
where
  site_url = 'sc-domain:example.com';
```

### List sitemaps with errors
List all sitemaps that have errors, including the total number of URLs submitted and the count of errors. This query helps in identifying sitemaps that need attention and potentially re-submission after correcting the errors.

```sql+postgres
select
  site_url,
  path,
  is_pending,
  c ->> 'submitted' as total_urls,
  errors as error_count
from
  googlesearchconsole_sitemap,
  jsonb_array_elements(contents) as c
where
  errors > 0;
```

```sql+sqlite
select
  site_url,
  path,
  is_pending,
  json_extract(content.value, '$.submitted') as total_urls,
  errors as error_count
from
  googlesearchconsole_sitemap,
  json_each(contents) as content
where
  errors > 0;
```

### List index sitemaps
Identify all index sitemaps, which are sitemaps that contain other sitemaps. This is useful for large sites that need to organize their URLs into multiple sitemaps for efficient management.

```sql+postgres
select
  site_url,
  path
from
  googlesearchconsole_sitemap
where
  is_sitemaps_index;
```

```sql+sqlite
select
  site_url,
  path
from
  googlesearchconsole_sitemap
where
  is_sitemaps_index = 1;
```