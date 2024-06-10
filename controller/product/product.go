package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	//"strconv"

	"uas_webser/database"
	//"github.com/gorilla/mux"
	"uas_webser/model/product"

	"github.com/gorilla/mux"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query("SELECT * FROM products")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []product.Product
    for rows.Next() {
        var c product.Product
        if err := rows.Scan(&c.ID,&c.Name,&c.Description,&c.Price,&c.Created_at); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        products = append(products, c)
    }
    
    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
    var pc product.Product
    if err := json.NewDecoder(r.Body).Decode(&pc); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Prepare the SQL statement for inserting a new course
    query := `
        INSERT INTO products ( name, description,price,created_at) 
        VALUES (?,?,?,NOW())`

    // Execute the SQL statement
    res, err := database.DB.Exec(query, pc.Name, pc.Description, pc.Price)
    if err != nil {
        http.Error(w, "Failed to insert artist: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Get the last inserted ID
    id, err := res.LastInsertId()
    if err != nil {
        http.Error(w, "Failed to retrieve last insert ID: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the newly created ID in the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Product added successfully",
        "id":      id,
    })
}


func PutProduct(w http.ResponseWriter, r *http.Request) {
    // Ambil ID dari URL
    vars := mux.Vars(r)
    idStr, ok := vars["id"]
    if !ok {
        http.Error(w, "ID not provided", http.StatusBadRequest)
        return
    }
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    // Decode JSON body
    var pc product.Product
    if err := json.NewDecoder(r.Body).Decode(&pc); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Prepare the SQL statement for updating the category admin
    query := `
        UPDATE products 
        SET name=?, description=?, price=?
        WHERE id=?`

    // Execute the SQL statement
    result, err := database.DB.Exec(query, pc.Name, pc.Description, pc.Price,id)
    if err != nil {
        http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Get the number of rows affected
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, "Failed to retrieve affected rows: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Check if any rows were updated
    if rowsAffected == 0 {
        http.Error(w, "No rows were updated", http.StatusNotFound)
        return
    }

    // Return success message
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Product updated successfully",
    })
}


func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    // Extract ID from URL
    vars := mux.Vars(r)
    idStr, ok := vars["id"]
    if !ok {
        http.Error(w, "ID not provided", http.StatusBadRequest)
        return
    }
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    // Prepare the SQL statement for deleting a category admin
    query := `
        DELETE FROM products
        WHERE id = ?`

    // Execute the SQL statement
    result, err := database.DB.Exec(query, id)
    if err != nil {
        http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Check if any rows were affected
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        http.Error(w, "Failed to retrieve affected rows: "+err.Error(), http.StatusInternalServerError)
        return
    }

    if rowsAffected == 0 {
        http.Error(w, "No rows were deleted", http.StatusNotFound)
        return
    }

    // Return the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Product deleted successfully",
    })
}

