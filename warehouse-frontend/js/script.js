const API_URL = "http://localhost:8080/products"; // Update with your actual API base URL

// Fetch and display products
async function fetchProducts() {
    try {
        const response = await fetch(API_URL);
        const products = await response.json();

        const productTable = document.getElementById("productTable");
        productTable.innerHTML = ""; // Clear the table before re-adding rows

        products.forEach(product => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${product.ProductID}</td>
                <td>${product.Name}</td>
                <td>${product.Description}</td>
                <td>${product.Price}</td>
                <td>${product.Quantity}</td>
                <td>
                    <button onclick="editProduct(${product.ProductID})">Edit</button>
                    <button onclick="deleteProduct(${product.ProductID})">Delete</button>
                </td>
            `;
            productTable.appendChild(row);
        });
    } catch (error) {
        console.error("Error fetching products:", error);
    }
}

// Add or update a product
async function saveProduct(event) {
    event.preventDefault();
    const productID = document.getElementById("productID").value;
    const name = document.getElementById("name").value;
    const description = document.getElementById("description").value;
    const price = document.getElementById("price").value;
    const quantity = document.getElementById("quantity").value;

    const method = productID ? "PUT" : "POST"; // Determine whether to create or update
    const url = productID ? `${API_URL}/${productID}` : API_URL;

    try {
        await fetch(url, {
            method: method,
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                Name: name,
                Description: description,
                Price: parseFloat(price),
                Quantity: parseInt(quantity),
            }),
        });

        document.getElementById("productForm").reset(); // Clear form after submission
        fetchProducts(); // Reload products
    } catch (error) {
        console.error("Error saving product:", error);
    }
}

// Edit product
function editProduct(productID) {
    const row = document.querySelector(`tr:has(td:first-child:contains("${productID}"))`);

    document.getElementById("productID").value = productID;
    document.getElementById("name").value = row.cells[1].innerHTML;
    document.getElementById("description").value = row.cells[2].innerHTML;
    document.getElementById("price").value = row.cells[3].innerHTML;
    document.getElementById("quantity").value = row.cells[4].innerHTML;
}

// Delete product
async function deleteProduct(productID) {
    try {
        await fetch(`${API_URL}/${productID}`, {
            method: "DELETE",
        });
        fetchProducts(); // Reload products
    } catch (error) {
        console.error("Error deleting product:", error);
    }
}

// Initialize
document.getElementById("productForm").addEventListener("submit", saveProduct);
fetchProducts(); // Initial load of products