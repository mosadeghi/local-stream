import type { Movie } from '../services/movieService'
import MovieCard from './MovieCard'

type Props = { movies: Movie[] }
export default function MovieList({ movies }: Props) {
  if (!movies.length) return <p className="meta">No movies found.</p>
  return (
    <div className="grid">
      {movies.map(m => <MovieCard key={m.id} movie={m} />)}
    </div>
  )
}
