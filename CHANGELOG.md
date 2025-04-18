## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#65](https://github.com/turbot/steampipe-plugin-googlesearchconsole/pull/65))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#65](https://github.com/turbot/steampipe-plugin-googlesearchconsole/pull/65))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#24](https://github.com/turbot/steampipe-plugin-googlesearchconsole/pull/24))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#24](https://github.com/turbot/steampipe-plugin-googlesearchconsole/pull/24))

## v0.0.1 [2024-04-12]

_What's new?_

- New tables added
  - [googlesearchconsole_indexing_status](https://hub.steampipe.io/plugins/turbot/googlesearchconsole/tables/googlesearchconsole_indexing_status)
  - [googlesearchconsole_pagespeed_analysis_aggregated](https://hub.steampipe.io/plugins/turbot/googlesearchconsole/tables/googlesearchconsole_pagespeed_analysis_aggregated)
  - [googlesearchconsole_pagespeed_analysis](https://hub.steampipe.io/plugins/turbot/googlesearchconsole/tables/googlesearchconsole_pagespeed_analysis)
  - [googlesearchconsole_site](https://hub.steampipe.io/plugins/turbot/googlesearchconsole/tables/googlesearchconsole_site)
  - [googlesearchconsole_sitemaps](https://hub.steampipe.io/plugins/turbot/googlesearchconsole/tables/googlesearchconsole_sitemaps)
