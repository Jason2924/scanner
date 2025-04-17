import { Link } from "react-router-dom"

const Header = () => {
  return (
    <header className="bg-gradient-to-r from-rose-400 to-cyan-100 py-6">
      <div className="container mx-auto flex justify-between">
        <div className="text-3xl text-white font-bold tracking-tight">
          <Link to="/">Scanner</Link>
        </div>
        <div className="flex space-x-2">
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
          </ul>
        </div>
      </div>
    </header>
  )
}

export default Header
