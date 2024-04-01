---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/gsc.svg"
brand_color: "#1EA362"
display_name: "Google Search Console"
short_name: "gsc"
description: "Steampipe plugin for query data from Google Search Console (GSC)."
og_description: "Query Google Search Console (GSC) with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/gsc-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Google Search Console + Steampipe

[Google Search Console](https://search.google.com/search-console) is a free service offered by Google that helps you monitor, maintain, and troubleshoot your site's presence in Google Search results.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

Query all the sites you are an owner of:

```sql
select
  site_url,
  permission_level
from
  gsc_site;
```

```
+-------------------------+------------------+
| site_url                | permission_level |
+-------------------------+------------------+
| https://example.io/     | siteOwner        |
| https://hub.example.io/ | siteOwner        |
+-------------------------+------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/gsc/tables)**

## Get started

### Install

Download and install the latest Google Search Console plugin:

```shell
steampipe plugin install gsc
```

### Credentials

| Item        | Description |
| :---------- | :---------- |
| Credentials | Follow this [guide](https://developers.google.com/search/apis/indexing-api/v3/prereqs) from Google. By the end of it, you should have a project on Google Cloud with the Indexing API enabled, a service account with the `Owner` permission on your sites. |
| APIs | 1. Go to the [Google API Console](https://console.cloud.google.com/apis/dashboard). <br/> 2. Select the project that contains your credentials. <br/> 3. Click `Enable APIs and Services`. <br/> 4. Enable both the `Google Search Console API` and `PageSpeed Insights API`.
| Radius      | Each connection represents a single Google Cloud service account and can be used to query data from multiple Google Search Console properties. |
| Resolution  | Credentials from the JSON file specified by the `credentials` parameter in your Steampipe config. |

### Configuration

Installing the latest gsc plugin will create a config file (`~/.steampipe/config/gsc.spc`) with a single connection named `gsc`:

```hcl
connection "gsc" {
  plugin = "gsc"

  # You should have a project on Google Cloud with the Indexing API enabled, a service account with the `Owner` permission on your sites.
  # The path to the Google Cloud credentials file of your sevice account.
  # credentials = "/path/to/credentials.json"
}
```
