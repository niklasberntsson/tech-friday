# Single match

## Description

A single match is a match that is played between two players.

## Models

``` go
type Player struct {
    ID string `json:"id"`
    Name string `json:"name"`
}

type Team struct {
    ID      string   `json:"id"`
    Players []Player `json:"players"`
}

type Set struct {
    ID string `json:"id"`
    Team1Score int `json:"team1Score"`
    Team2Score int `json:"team2Score"`
}

type Match struct {
    ID string `json:"id"`
    Team1 Team `json:"team1"`
    Team2 Team `json:"team2"`
    Config MatchConfig `json:"config"`
    Status MatchStatus `json:"status"`
    StartTime time.Time `json:"startTime"`
    EndTime time.Time `json:"endTime"`
    Sets []Set `json:"sets"`
}

type MatchConfig struct {
    MaxSets int `json:"maxSets"` // e.g. 1, 3, 5
    TimeBased bool `json:"timeBased"`
    TieBreak bool `json:"tieBreak"`
}
```


## Rules

- A single match can be either regular sets or time based.
- The sets does not have to be completed to end the match.
- Tie break can be allowed but not required.
- A match can be started with a predefined score.

## Endpoints

GET /match/{id} // Get a single match by id
POST /match // Create a new single match
PUT /match/{id} // Update a single match by id
DELETE /match/{id} // Delete a single match by id

## Considerations

* There will be more types of matches in the future. **This needs to be considered when designing the API.**

* The score update endpoint will be used for both regular sets and time based matches.
* A match can be private (only visible to participants) or public .