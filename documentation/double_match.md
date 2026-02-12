# Double match

## Description

A double match is a match that is played between two teams of two players each.

## Models

Models are defined in single_match.md


## Rules

- A double match can be either regular sets or time based.
- The sets does not have to be completed to end the match.
- Tie break can be allowed but not required.
- A match can be started with a predefined score.

## Endpoints

Same endpoints as single match but with the following differences:

- The match type is double match
- The team size is 2

## Considerations

* There will be more types of matches in the future. **This needs to be considered when designing the API.**

* The score update endpoint will be used for both regular sets and time based matches.
* A match can be private (only visible to participants) or public (visible to all users in the facility).
