import { Route, Routes } from "react-router-dom"
import NormalLayout from "./layouts/NormalLayout"
import HomePage from "./pages/HomePage"
import HistoryPage from "./pages/HistoryPage"
import ComparisonPage from "./pages/ComparisonPage"

const App = () => {
  return (
    <Routes>
      <Route element={<NormalLayout />}>
        <Route index element={<HomePage />} />
        <Route path="/history" element={<HistoryPage />} />
        <Route path="/comparison/:report1/:report2" element={<ComparisonPage />} />
      </Route>
    </Routes>
  )
}

export default App
