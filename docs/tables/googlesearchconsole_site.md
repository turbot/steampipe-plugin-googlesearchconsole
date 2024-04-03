---
title: "Steampipe Table: googlesearchconsole_site - Query Search Console Properties using SQL"
description: "Access detailed information and permission levels of Google Search Console properties through SQL queries."
---

# Table: googlesearchconsole_site - Query Search Console Properties with SQL

Search Console properties are the websites that users add to their Google Search Console account to monitor and manage their presence on Google Search.

## Table Usage Guide

The `googlesearchconsole_site` table lists the user's Google Search Console properties, including URLs and permission levels. It's beneficial for site administrators and SEO professionals who need to oversee site verification status and access permissions.

## Examples

### Basic search console site info
This query provides a comprehensive overview of all properties in the user's Google Search Console, displaying complete information including site URLs and permission levels.

```sql+postgres
select
  *
from
  googlesearchconsole_site;
```
  
```sql+sqlite
select
  *
from
  googlesearchconsole_site;
```

### List the sites a user owns
Filter properties to list only those where the user is designated as the site owner, indicating full control over the property.

```sql+postgres
select
  *
from
  googlesearchconsole_site
where
  permission_level = 'siteOwner';
```

```sql+sqlite
select
  *
from
  googlesearchconsole_site
where
  permission_level = 'siteOwner';
```

### Get the site details for a specific site
Fetch detailed information for a specific site by its URL. This is particularly useful for directly accessing the details of a single property.

```sql+postgres
select
  *
from
  googlesearchconsole_site
where
  site_url = 'https://www.example.com/';
```

```sql+sqlite
select
  *
from
  googlesearchconsole_site
where
  site_url = 'https://www.example.com/';
```

### Get the site details for a specific domain
Get details for a domain-level property in Search Console, useful for managing domain properties that encompass subdomains and protocols.

```sql+postgres
select
  *
from
  googlesearchconsole_site
where
  site_url = 'sc-domain:example.com';
```

```sql+sqlite
select
  *
from
  googlesearchconsole_site
where
  site_url = 'sc-domain:example.com';
```