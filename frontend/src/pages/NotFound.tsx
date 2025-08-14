import { Link } from 'react-router-dom'
export default function NotFound() {
  return (
    <div>
      <h2>404 – Not Found</h2>
      <p className="meta">The page you requested does not exist.</p>
      <Link className="button" to="/">Go Home</Link>
    </div>
  )
}
