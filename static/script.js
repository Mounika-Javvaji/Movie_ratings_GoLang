async function fetchMovie() {
    const title = document.getElementById('movieTitle').value;
    const response = await fetch(`/api/movie?title=${title}`);
    const movie = await response.json();
    displayMovieDetails(movie);
}

function displayMovieDetails(movie) {
    const movieDetails = document.getElementById('movieDetails');
    movieDetails.innerHTML = `
        <h2>${movie.Title} (${movie.Year})</h2>
        <img src="${movie.Poster}" alt="Poster">
        <p><strong>Director:</strong> ${movie.Director}</p>
        <p><strong>Actors:</strong> ${movie.Actors}</p>
        <p><strong>Plot:</strong> ${movie.Plot}</p>
        <p><strong>IMDB Rating:</strong> ${movie.Ratings.find(r => r.Source === "Internet Movie Database")?.Value || "N/A"}</p>
        <p><strong>Rotten Tomatoes:</strong> ${movie.Ratings.find(r => r.Source === "Rotten Tomatoes")?.Value || "N/A"}</p>
        <p><strong>Metacritic:</strong> ${movie.Metascore}</p>
    `;
}

function showAbout() {
    const movieDetails = document.getElementById('movieDetails');
    movieDetails.innerHTML = `
        <h2>About This Application</h2>
        <p>This application allows you to search for movies and retrieve detailed information including the title, year, director, actors, plot, and various ratings.</p>
    `;
}

function moviesList() {
    const movieDetails = document.getElementById('movieDetails');
    movieDetails.innerHTML = `
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('animation')">Animation</button>
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('classic')">Classic</button>
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('comedy')">Comedy</button>
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('drama')">Drama</button>
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('horror')">Horror</button>
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('family')">Family</button>
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('mystery')">Mystery</button>
        <button type="button" class="button-header" onclick="fetchMoviesByGenre('western')">Western</button>
    `;
}

async function fetchMoviesByGenre(genre) {
    const response = await fetch(`/api/movies/${genre}`);
    const movies = await response.json();
    displayMoviesList(movies);
}

function displayMoviesList(movies) {
    const movieDetails = document.getElementById('movieDetails');
    movieDetails.innerHTML = `
            ${movies.map(movie => `
                <div class="movie-item">
                    <h3>${movie.Title}</h3>
                </div>
            `).join('')}
    `;
}
