import { BrowserRouter, Route, Routes } from "react-router-dom";
import ROUTERS from "./constants/route-path";
import Home from "./pages/home";
import Mine from "./pages/mine";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path={ROUTERS.HOME} element={<Home />} />
          <Route path={ROUTERS.MINE} element={<Mine />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
