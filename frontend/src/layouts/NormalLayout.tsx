import { Outlet } from "react-router-dom"
import Header from "../components/public/Header"

const NormalLayout = () => {
  return (
    <div className="w-full flex flex-col min-h-screen">
      <Header />
      <main className="container mx-auto py-10 flex-1">
        <Outlet />
      </main>
    </div>
  )
}

export default NormalLayout
