
export async function pingApi(http) {
  const params = {
    method: "GET",
  };
  return await http("/api/ping", params);
}
