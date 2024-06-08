***Movie API Server***
This is a simple API server written in Go that fetches movie information from the OMDB API and a sample movie API by genre. It also serves static files from the ./static directory.

***Features***
Fetch movie details by title from the OMDB API.
Fetch movies by genre from a sample movie API.
Serve static files.

***Endpoints***
***Fetch Movie by Title****
URL: /api/movie

Method: GET

Query Parameters:

title (required): The title of the movie.
Response:
Returns the details of the movie in JSON format.

***Fetch Movies by Genre****
URL: /api/movies/{genre}

Method: GET

***Path Parameters:****

{genre} (required): The genre of the movies to fetch.
Response:

Returns a list of movies in the specified genre in JSON format.

***Static Files***
Static files are served from the ./static directory. You can place your HTML, CSS, and JavaScript files in this directory.
***How to Run***
1)Clone the repository.
2)Replace the apiKey constant in the code with your OMDB API key.
3)Build and run the server:
    go build -o movie-api-server
    ./movie-api-server

The server will start on port 8080.
