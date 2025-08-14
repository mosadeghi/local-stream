import { useEffect, useState } from 'react'
import { useParams, Link } from 'react-router-dom'
import { fetchMovieById, type Movie } from '../services/movieService'
import VideoPlayer from '../components/VideoPlayer'
import placeholder from "../assets/placeholder.jpg"

export default function MoviePage() {
  const params = useParams<{ id: string }>()
  const id = Number(params.id)
  const [movie, setMovie] = useState<Movie | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    if (!id) { setError('Invalid ID'); setLoading(false); return }
    let ignore = false
    fetchMovieById(id)
      .then(m => { if (!ignore) setMovie(m) })
      .catch(err => setError(err.message))
      .finally(() => setLoading(false))
    return () => { ignore = true }
  }, [id])

  if (loading) return <p className="meta">Loading…</p>
  if (error) return <p className="meta">Error: {error}</p>
  if (!movie) return <p className="meta">Not found.</p>

  const poster = movie.posterPath ? `/static/${movie.posterPath}` : '/static/posters/placeholder.jpg'

  return (
    <div>
      <Link to="/" className="button" style={{ marginBottom: 16 }}>← Back</Link>

      <div className="hstack" style={{ gap: 24, alignItems: 'flex-start', marginBottom: 16 }}>
        <img
          src={poster}
          alt={movie.title}
          width={220} height={330}
          style={{ borderRadius: 8, objectFit: 'cover' }}
          onError={(e) => {
            e.currentTarget.src = placeholder
          }}
        />

        <div>
          <h1 style={{ margin: '0 0 8px 0' }}>{movie.title}</h1>
          <div className="meta"><strong>Director:</strong> {movie.director || '—'}</div>
          <div className="meta"><strong>Year:</strong> {movie.year || '—'}</div>
          <p style={{ marginTop: 12 }}>{movie.summary || 'No summary.'}</p>
        </div>
      </div>

      <VideoPlayer src={`/stream/${movie.id}`} />
    </div>
  )
}
