# Getting messages

Source: [Khoros Community API v2 Docs](https://developer.khoros.com/khoroscommunitydevdocs/reference/community-api-v2-object-index)

There are 2 ways to fetch data from the API's:
1. Restful Endpoints
2. Search Endpoint

If you're looking for Lists, you'll want to use the `/search` endpoint with LiQL to construct your List Query.

# Restful Endpoints
We'll be using these when fetching a record from a specific collection. These are not useful for Lists.

Used when we make a `GET`/`UPDATE`/`CREATE` request on a record in a collection.

# Search Endpoint
We'll be using the `/search?q=` endpoint with a LiQL query parameter.

Used for `GET` operations and when returning Lists.
