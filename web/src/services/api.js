import { getUserToken } from "../services/authsvc";
import { getToken } from "../services/authsvc";

export function fetchFn(setLoading) {
  return async function(url, params) {
    try {
      setLoading(true);

      const token = await getToken();
      params.headers = Object.assign(params.headers || {}, {
        "Authorization": token,
      });
  
      return fetch(url, params).then(res => {
        if (res.status == 401) {
          window.location.href = '/login';
        }
        return res.json();
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

// todo: move to separate file
export async function createCustomer(id, name, email, phonenumber, gstin, showError , showSuccess) {
  const data = {
    customerid : id,
    name : name,
    email: email,
    phone : phonenumber,
    gstin: gstin
  }

  const params = {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      "Authorization" :  getUserToken()
    },
    body: JSON.stringify(data),
  };

  let response =  await fetch("/api/customers", params).catch(err => showError(err) );
  if(response.ok) {
    showSuccess();
  } else {
    showError(response.status);
  }
}

// todo: move to separate file
export function getCustomers() {
  const params = {
    method: "GET",
    headers: {
      "Authorization" :  getUserToken()
    },
  };
  return fetch("/api/customers", params);
}