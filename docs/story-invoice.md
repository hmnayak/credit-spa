# Create/Update/Show/List Invoice

As a businessman

I want to create invoices for my sales

So that I can track sales for GST purpose


### APIs
POST /api/invoices
PUT /api/invoices/<invoice_id>
GET /api/invoices/<invoive_id>
GET /api/invoices?page=3

    {
        date
        customer_id
        lineitems [{
            item_id
            quantity
            price
        }]
    }

### Notes
- All API/URL conventions must be consistent across app
- Must provide option to create customer on the fly or choose existing customer
- Must provide option to create item on the fly or choose existing item
