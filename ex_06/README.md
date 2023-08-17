# Exercise 06: Simple E-commerce

## 1. Requirement
- Product Management - JWT Auth
```
- POST /products: Create a new product. It receives product details as JSON input.
- PUT /products/{product_id}: Update a product's details. It receives updated product details as JSON input.
- DELETE /products/{product_id}: Delete a product by its ID. GET /products: Retrieve a list of all products.
```

- Shopping Cart - Without Auth
```
- POST /cart/add: Add items to the cart. It receives a product ID and quantity as JSON input.
- DELETE /cart/remove: Remove items from the cart. It receives a product ID as JSON input.
- POST /cart/checkout: Checkout and clear the cart. It returns a receipt with the total price.
```

## 2. Install
- Import `ex_06.postman_collection.json` into Postman
- Start server
```
make server
```

- Test APIs in `ex_06.postman_collection.json`