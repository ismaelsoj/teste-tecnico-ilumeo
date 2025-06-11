import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  LineElement,
  CategoryScale,
  LinearScale,
  PointElement,
  Tooltip,
  Legend,
} from "chart.js";

ChartJS.register(LineElement, CategoryScale, LinearScale, PointElement, Tooltip, Legend);

const Grafico = ({ data }) => {
  const chartData = {
    labels: data.map((d) => d.data.split("T")[0]),
    datasets: [
      {
        label: "Taxa de Conversão (%)",
        data: data.map((d) => d.taxa_conversao),
        borderColor: "#4f46e5",
        fill: false,
      },
    ],
  };

  return <Line data={chartData} />;
};

export default Grafico;
