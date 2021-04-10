
# Create/Update/Show item

As a businessman

I want to manage items in my inventory

So that I can refer them in my invoices

### API
POST /api/items

    {
        name
        type
        hsn
        sac
        gst
        igst
    }

### Notes
- All API/URL conventions must be consistent across app
- Identifier strategy must be consistent across app
- UX must be consistent across app
- Anything in current implementation must be consistent with whatever was done before

# List items

As a businessman

I want to list all my items

So that I can find the item I'm looking for

### API
GET /api/items?page=3

    [{
        name
        type
        hsn
        sac
        gst
        igst
    }]

### Notes
- List must be paginated
- Pagination mechanism must be consistent across app
