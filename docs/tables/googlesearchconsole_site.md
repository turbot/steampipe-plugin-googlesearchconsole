# Table: googlesearchconsole_site

Lists the user's Search Console properties, contains permission level information about a Search Console site

## Examples

### Basic info

```sql
select
  *
from
  googlesearchconsole_site;
```

### List the sites a user owns

```sql
select
  *
from
  googlesearchconsole_site
where
  permission_level = 'siteOwner';
```

### Get the site details for a specific site

```sql
select
  site_url
from
  googlesearchconsole_site
where
  site_url = 'https://www.example.com/';
```

### Get the site details for a specific domain

```sql
select
  site_url
from
  googlesearchconsole_site
where
  site_url = 'sc-domain:example.com';
```