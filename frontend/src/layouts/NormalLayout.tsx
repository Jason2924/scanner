import { Outlet } from "react-router-dom"
import Header from "../components/public/Header"
import bgImage from "../assets/cloud.jpg"
import { Box } from "@radix-ui/themes"

const NormalLayout = () => {
  return (
    <div className="w-full flex flex-col min-h-screen">
      <Header />
      <main className={`relative flex-1 bg-no-repeat bg-cover bg-bottom`} style={{backgroundImage: `url(${bgImage})`}}>
        <Box className="size-full absolute bg-gray-700 opacity-30" />
        <Box className="relative z-10">
          <Outlet />
        </Box>
      </main>
    </div>
  )
}

export default NormalLayout
