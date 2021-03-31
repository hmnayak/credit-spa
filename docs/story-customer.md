# Create customer

As a businessman

I want to create a customer record

So that I can create invoices for them

### API
PUT /api/customers

    {
        customer_id
        name
        email
        phone_no
        gstin
    }

### UI
GET /customers/new

### Notes
- API must do upsert. If customer_id is absent, create new customer, else update existing customer
- for new customer, generate sequential customer_id as CUST-xxxx

### Notes
- Redirect to dummy /customers/<customer_id> on success

# Show customer

As a businessman

I want to see all details of a customer

So that I can seek any information regarding my customer

### API
Get /api/customers/<customer_id>

    {
        name
        email
        phone_no
        gstin
    }

### UI
GET /customers/<customer_id>

### Notes
- Show details within editable form. User must be able to edit and update record

# List customers

As a businessman

I want to list all my customers

So that I can find customers I'm looking for

### API
GET /api/customers

    [{
        name
        email
        phone_no
        gstin
    }]

### UI
GET /customers

### Notes
- Create left sidepane with menu option Customers pointing to /customers
- Provide link to open /customers/<customer_id> for each entry

# Paginate customers

As a businessman

I want to list all my customers by pages

So that I can see small list of customers at a time 

### API
GET /api/customers?page=3

    [{
        name
        email
        phone_no
        gstin
    }]

### UI
GET /customers?page=3

### Notes
- Make page size configurable
- Show page no, prev, next links on each page