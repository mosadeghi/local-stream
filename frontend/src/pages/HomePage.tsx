import { useEffect, useMemo, useState } from 'react'
import { fetchMovies, type Movie } from '../services/movieService'
import MovieList from '../components/MovieList'

export default function HomePage() {
  const [movies, setMovies] = useState<Movie[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [q, setQ] = useState('')

  useEffect(() => {
    let ignore = false
    fetchMovies()
      .then(list => { if (!ignore) setMovies(list) })
      .catch(err => setError(err.message))
      .finally(() => setLoading(false))
    return () => { ignore = true }
  }, [])

  const filtered = useMemo(() => {
    const s = q.trim().toLowerCase()
    if (!s) return movies
    return movies.filter(m =>
      (m.title || '').toLowerCase().includes(s) ||
      (m.director || '').toLowerCase().includes(s) ||
      String(m.year || '').includes(s)
    )
  }, [movies, q])

  if (loading) return <p className="meta">Loading…</p>
  if (error) return <p className="meta">Error: {error}</p>

  return (
    <>
      <div className="hstack" style={{ marginBottom: 16 }}>
        <input className="input" placeholder="Search title, director, or year…" value={q} onChange={e => setQ(e.target.value)} />
      </div>
      <MovieList movies={filtered} />
    </>
  )
}
