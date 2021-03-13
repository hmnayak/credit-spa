
export async function getCustomersApi(http) {
    const params = {
      method: "GET",
    };
    return await http("/api/customers", params);
}
  

export async function createCustomer(http, id, name, email, phonenumber, gstin, showError , showSuccess) {
    const data = {
        customerid : id,
        name : name,
        email: email,
        phone : phonenumber,
        gstin: gstin
    }

    const params = {
        method: "PUT",
        body: JSON.stringify(data),
    };

    return await http("/api/customers", params);
}