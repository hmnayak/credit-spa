import { getUserToken } from "../services/authsvc";

export function aboutInfoApi() {
  const params = {
    method: "GET",
    headers: {
      "Content-Type": "application/text",
      "Authorization" :  getUserToken()
    },
  };
  return fetch("/api/ping", params);
}

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

  let response =  await fetch("/api/customers", params)
  .catch((error) => {
    showError(error)
  });
  if(response.ok){
    showSuccess();
  }
  else {
    showError(response.status);
  }
}

export function getCustomers() {
  const params = {
    method: "GET",
    headers: {
      "Content-Type": "application/text",
      "Authorization" :  getUserToken()
    },
  };
  return fetch("/api/customers", params);
}