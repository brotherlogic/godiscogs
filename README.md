# Godiscogs

Golang interface to discogs, with some refinements for my system. Also some bonus bits and pieces.

[![Coverage Status](https://coveralls.io/repos/github/brotherlogic/go-discogs/badge.svg?branch=master)](https://coveralls.io/github/brotherlogic/go-discogs?branch=master)

### Recent Updates

- **Release Field Enhancements**: Added support for `blocked_from_sale` (bool) and `artists_sort` (string) fields on `Release` objects.
- **Marketplace Listing Enhancements**: Added support for appending custom `notes` to Discogs marketplace item listings via `SellRecord` (`comments` param mapping).
