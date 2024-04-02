# Table: googlesearchconsole_sitemap

Lists the sitemaps-entries submitted for all sites.

## Examples

### Basic info

```sql
select
  site_url,
  path,
  type,
  c ->> 'submitted' as total_urls
from
  googlesearchconsole_sitemap,
  jsonb_array_elements(contents) as c;
```

### List sitemaps for a specific site

```sql
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

### List sitemaps for a specific domain

```sql
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

### List sitemaps with errors

```sql
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

### List index sitemaps

```sql
select
  site_url,
  path
from
  googlesearchconsole_sitemap
where
  is_sitemaps_index;
```