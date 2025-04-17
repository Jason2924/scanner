import { Route, Routes } from "react-router-dom"
import NormalLayout from "./layouts/NormalLayout"
import HomePage from "./pages/HomePage"

const App = () => {
  return (
    <Routes>
      <Route element={<NormalLayout />}>
        <Route index element={<HomePage />} />
      </Route>
    </Routes>
  )
}

export default App
