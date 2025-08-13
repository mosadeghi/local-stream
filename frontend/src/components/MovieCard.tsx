import { Link } from 'react-router-dom'
import type { Movie } from '../services/movieService'

type Props = { movie: Movie }
export default function MovieCard({ movie }: Props) {
  const poster = movie.posterPath ? `/static/${movie.posterPath}` : '/static/posters/placeholder.jpg'
  return (
    <Link to={`/movie/${movie.id}`} className="card" title={movie.title}>
      <img src={poster} alt={movie.title} onError={(e) => {
        (e.currentTarget as HTMLImageElement).src = '/static/posters/placeholder.jpg'
      }} />
      <div className="content">
        <div className="hstack" style={{ justifyContent: 'space-between' }}>
          <strong>{movie.title || 'Untitled'}</strong>
          <span className="meta">{movie.year || ''}</span>
        </div>
        <div className="meta">{movie.director}</div>
      </div>
    </Link>
  )
}
