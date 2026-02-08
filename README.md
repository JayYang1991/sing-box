> Sponsored by [Warp](https://go.warp.dev/sing-box), built for coding with multiple AI agents

<a href="https://go.warp.dev/sing-box">
<img alt="Warp sponsorship" width="400" src="https://github.com/warpdotdev/brand-assets/raw/refs/heads/main/Github/Sponsor/Warp-Github-LG-02.png">
</a>

---

# sing-box

The universal proxy platform.

[![Packaging status](https://repology.org/badge/vertical-allrepos/sing-box.svg)](https://repology.org/project/sing-box/versions)

## Documentation

https://sing-box.sagernet.org

## Custom Enhancements

This version includes several optimizations for proxy group management:

### 1. URLTest Proactive Re-testing
A new configuration option `max_failed` is added to `urltest` outbound groups.
- **Behavior**: When a node fails consecutively for `max_failed` times, it proactively triggers a URL test to switch to a healthier node.
- **Default**: 10 (if not specified).
- **Log Level**: Warnings are logged when the threshold is reached or when a selection change occurs.

### 2. Group Management CLI
Added a new command to query the currently selected node in proxy groups.
- **Requirement**: `experimental.cache_file` must be configured in your `config.json`.
- **Usage**:
  ```bash
  sing-box group list -c /path/to/config.json
  ```

### 3. Enhanced Logging
Path switching and selection changes are now logged at the `warning` level for better visibility.

## License

```
Copyright (C) 2022 by nekohasekai <contact-sagernet@sekai.icu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

In addition, no derivative work may use the name or imply association
with this application without prior consent.
```