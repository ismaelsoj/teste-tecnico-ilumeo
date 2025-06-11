import React, { useEffect, useState } from "react";
import { fetchTaxaConversao } from "../services/taxaServices";
import Filtros from "../components/Filtros";
import Grafico from "../components/Grafico";
import { Card, CardContent } from "../components/ui/card";

const Dashboard = () => {
  const [data, setData] = useState([]);
  const [filtros, setFiltros] = useState({ start: "", end: "", origin: "" });

  const aplicarFiltros = async () => {
    const res = await fetchTaxaConversao(filtros);
    setData(res);
  };

  useEffect(() => {
    aplicarFiltros();
  }, []);

  return (
    <div className="p-6 space-y-6">
      <Card>
        <CardContent className="p-6 space-y-4">
          <Filtros filtros={filtros} setFiltros={setFiltros} onFiltrar={aplicarFiltros} />
        </CardContent>
      </Card>

      <Card>
        <CardContent className="p-6">
          <Grafico data={data} />
        </CardContent>
      </Card>
    </div>
  );
};

export default Dashboard;

