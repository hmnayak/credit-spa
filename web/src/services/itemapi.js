export async function upsertItem(http, id, name, type, hsn, sac, gst, igst) {
    const data = {
        itemid : id,
        name : name,
        type: type,
        hsn : hsn,
        sac: sac,
        gst: gst,
        igst: igst
    }

    const params = {
        method: "PUT",
        body: JSON.stringify(data),
    };
    return await http("/api/items", params);
}

export async function getItemsPaginated(http, pageToken) {
    const params = {
      method: "GET",
    };
    return await http(`/api/items?page=${pageToken}`, params);
}