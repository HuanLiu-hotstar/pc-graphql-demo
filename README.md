# GQLgen usage

## create project using gqlgen

- we can follow the gqlgen README.md file https://github.com/99designs/gqlgen/blob/master/README.md
- Detail doc about project setup and implemention https://hotstar.atlassian.net/wiki/spaces/~875200411/pages/3206185078/Use+gqlgen+to+build+GraphQL+server


## Error handler


```graphql
# GraphQL schema example
#
# https://gqlgen.com/getting-started/
# extend type Cms @key(fields: "content_id"){
#   content_id: ID!
#   playback_set: [PlaybackSet!]!
# }
type Playback {
  content_id: String!
  match:  Boolean!
	drm_class: String! 
	download_drm_class: String! 
  playback_set: [PlaybackSet!]!
}

type PlaybackSet{
  PlaybackURL: String!
	PlaybackCDNType:String!
}

# input type
input Input {
  content_id: String!
}
type SuccessResult {
  playback: Playback!
}

type FailureResult {
  err_code: Int!
  err_msg: String!
}
union PlaybackResult = SuccessResult | FailureResult
type Query {
  get_playback_urls(input: Input): PlaybackResult!
}


```

## query with error response

```graphql


query {
	get_playback_urls(input:{content_id:"4"}){
		__typename
		... on SuccessResult {
			playback{
				match
				content_id
				playback_set {
					PlaybackURL
					PlaybackCDNType
				}
			}
		}
		... on FailureResult {
			err_code 
			err_msg
		}
	}
}

## response with json format
{
  "data": {
    "get_playback_urls": {
      "__typename": "FailureResult",
      "err_code": 1,
      "err_msg": "not found:4 "
    }
  }
}

```