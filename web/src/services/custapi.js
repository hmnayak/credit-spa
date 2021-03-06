export async function getCustomersApi(http) {
    const params = {
      method: "GET",
    };
    return await http("/api/customers", params);
}

export async function getCustomersPaginated(http, pageToken) {
    const params = {
      method: "GET",
    };
    return await http(`/api/customers?page=${pageToken}`, params);
}

export async function getCustomerApi(http, id) {
    const params = {
      method: "GET",
    };
    return await http("/api/customers/" + id, params);
}

export async function upsertCustomer(http, id, name, email, phonenumber, gstin) {
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