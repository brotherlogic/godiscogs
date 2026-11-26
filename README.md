# Godiscogs

Golang interface to discogs, with some refinements for my system. Also some bonus bits and pieces.

[![Coverage Status](https://coveralls.io/repos/github/brotherlogic/go-discogs/badge.svg?branch=master)](https://coveralls.io/github/brotherlogic/go-discogs?branch=master)

### Recent Updates

- **Database Search**: Added `Search(ctx, SearchParams)` (`GET /database/search`) to `DiscogsRetriever` with `SearchParams` and `SearchResult` types, `Authorization: Discogs token=...` header auth, context-aware concurrency-safe throttling, and gRPC status error mapping.
- **In-Place Marketplace Listing Updates**: Added `UpdateSaleListing` to `DiscogsRetriever` to update existing Discogs marketplace listings in-place (`release_id`, `condition`, `sleeve_condition`, `price`, `status`, and `comments`).
- **Release Field Enhancements**: Added support for `blocked_from_sale` (bool) and `artists_sort` (string) fields on `Release` objects.
- **Marketplace Listing Enhancements**: Added support for appending custom `notes` to Discogs marketplace item listings via `SellRecord` (`comments` param mapping).
