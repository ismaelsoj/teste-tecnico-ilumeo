import React from "react";
import { Routes, Route } from "react-router-dom";
import Dashboard from "@/pages/Dashboard";

const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Dashboard />} />
      {/* outras rotas futuras podem ir aqui */}
    </Routes>
  );
};

export default App;
