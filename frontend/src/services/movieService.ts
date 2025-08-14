export type Movie = {
  id: number
  title: string
  year: number
  director: string
  summary: string
  posterPath: string | null
}

function normalizeMovie(apiMovie: any): Movie {
  return {
    id: apiMovie.ID,
    title: apiMovie.Title,
    year: apiMovie.Year,
    director: apiMovie.Director,
    summary: apiMovie.Summary,
    posterPath: apiMovie.PosterPath,
  }
}

// GET /api/v1/movies  -> { movies: Movie[] }
export async function fetchMovies(): Promise<Movie[]> {
  const res = await fetch('/api/v1/movies')
  if (!res.ok) throw new Error(`Failed to fetch movies: ${res.status}`)
  const data = await res.json()
  return (data?.movies ?? []).map(normalizeMovie)
}

// GET /api/v1/movies/:id  -> { movie: Movie }
export async function fetchMovieById(id: number): Promise<Movie> {
  const res = await fetch(`/api/v1/movies/${id}`)
  if (!res.ok) throw new Error(`Movie ${id} not found`)
  const data = await res.json()
  return normalizeMovie(data?.movie)
}
