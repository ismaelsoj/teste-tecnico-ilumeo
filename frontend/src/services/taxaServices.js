export const fetchTaxaConversao = async ({ start, end, origin }) => {
  const params = new URLSearchParams();
  if (start) params.append("start", start);
  if (end) params.append("end", end);
  if (origin) params.append("origin", origin);

  const res = await fetch(`/api/v1/taxa-conversao?${params.toString()}`);
  if (!res.ok) throw new Error("Erro ao buscar taxa de conversão");
  return res.json();
};
