import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

const Filtros = ({ filtros, setFiltros, onFiltrar }) => {
  return (
    <>
      <h2 className="text-xl font-semibold">Filtros</h2>
      <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <Input
          type="datetime-local"
          value={filtros.start}
          onChange={(e) => setFiltros((prev) => ({ ...prev, start: e.target.value }))}
        />
        <Input
          type="datetime-local"
          value={filtros.end}
          onChange={(e) => setFiltros((prev) => ({ ...prev, end: e.target.value }))}
        />
        <Input
          value={filtros.origin}
          onChange={(e) => setFiltros((prev) => ({ ...prev, origin: e.target.value }))}
          placeholder="Canal"
        />
      </div>
      <Button onClick={onFiltrar}>Aplicar Filtros</Button>
    </>
  );
};

export default Filtros;
