import { getToken } from "../services/authsvc";

export function fetchFn(setLoading) {
  return async function(url, params) {
    try {
      setLoading(true);

      const token = await getToken();
      params.headers = Object.assign(params.headers || {}, {
        "Authorization": token,
      });
  
      return fetch(url, params).then((res) => {
        if (res.status == 401) {
          window.location.href = '/login';
        }
        return res;
      });
    } catch(err) {
      if (err == 'nouser') {
        window.location.href = '/login';
      }
    } finally {
      setLoading(false);
    }
  }
}