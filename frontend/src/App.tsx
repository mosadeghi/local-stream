import { Route, Routes, Link } from 'react-router-dom'
import HomePage from './pages/HomePage'
import MoviePage from './pages/MoviePage'
import NotFound from './pages/NotFound'

export default function App() {
  return (
    <div className="app">
      <header className="header">
        <Link to="/" className="brand">🎬 LocalStream</Link>
        <nav className="nav">
          <a href="/admin" className="nav-link">Admin</a>{/* server-rendered admin */}
        </nav>
      </header>

      <main className="container">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/movie/:id" element={<MoviePage />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </main>

      <footer className="footer">LocalStream © 2025</footer>
    </div>
  )
}
