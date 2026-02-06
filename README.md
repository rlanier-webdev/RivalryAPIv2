# API Documentation
------------
## Introduction
------------
This API provides endpoints to retrieve information about the historic City-Poly game.

## Try it in Postman
------------
[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.postman.com/thrtn85/main/api/2451bd7d-f19d-4b66-9317-a605ced2ba0c/entity?action=share&creator=13028315)

## Endpoints
------------
### Get All Games
- **URL**: `/api/games`
- **Method**: `GET`
- **Description**: Returns all games in JSON format.
- **Response**: 
  - Status Code: `200 OK`
  - Body: JSON array of games.

### Get Game by ID
- **URL**: `/api/games/:id`
- **Method**: `GET`
- **Description**: Returns the game with the specified ID.
- **Parameters**:
  - `id` (int): The ID of the game to retrieve.
- **Response**:
  - Status Code: 
    - `200 OK` if the game is found.
    - `404 Not Found` if the game with the specified ID is not found.
  - Body: JSON object representing the game.

### Get Games by Home Team
- **URL**: `/api/games/home/:name`
- **Method**: `GET`
- **Description**: Returns all games where the specified team plays as the home team.
- **Parameters**:
  - `name` (string): Name of the home team to search for.
- **Response**:
  - Status Code:
    - `200 OK` if games involving the team are found.
    - `404 Not Found` if no games involving the team are found.
  - Body: JSON array of games involving the specified team.

### Get Games by Away Team
- **URL**: `/api/games/away/:name`
- **Method**: `GET`
- **Description**: Returns all games where the specified team plays as the away team.
- **Parameters**:
  - `name` (string): Name of the away team to search for.
- **Response**:
  - Status Code:
    - `200 OK` if games involving the team are found.
    - `404 Not Found` if no games involving the team are found.
  - Body: JSON array of games involving the specified team.

### Get Games by Year
- **URL**: `/api/games/year/:year`
- **Method**: `GET`
- **Description**: Returns all games played in the specified year.
- **Parameters**:
  - `year` (int): The year to retrieve games for.
- **Response**:
  - Status Code:
    - `200 OK` if games for the specified year are found.
    - `404 Not Found` if no games are found for the specified year.
  - Body: JSON array of games played in the specified year.

### Get All Teams
- **URL**: `/api/teams`
- **Method**: `GET`
- **Description**: Returns all unique team names.
- **Response**:
  - Status Code: `200 OK`
  - Body: JSON object with teams array.

## Data Structure
------------
### Game
Represents a sports game.

- `ID` (int): The unique identifier of the game.
- `HomeTeam` (Team): The home team participating in the game.
- `AwayTeam` (Team): The away team participating in the game.
- `Date` (string): The date of the game in the format "YYYY-MM-DD".
- `Score` (Score): The score of the game.
- `Notes` (string): Additional notes about the game.
