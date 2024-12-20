// Function to get input value
function getInputValue(inputId) {
    return document.getElementById(inputId).value;
}

// Function to create request data
function createRequestData(messageText) {
    return { message: messageText };
}

// Function to post data to the API
function postToApi(url, requestData) {
    return fetch(url, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(requestData)
    });
}

// Function to display server response
function displayResponse(targetId, responseText) {
    document.getElementById(targetId).innerText = responseText;
}

// Main handler that integrates all logic
function sendData() {
    const messageText = getInputValue("message"); // Get input value
    const requestData = createRequestData(messageText); // Create request payload

    postToApi("http://localhost:8080/post", requestData)
        .then(response => response.json()) // Parse JSON response
        .then(data => displayResponse("response", JSON.stringify(data, null, 2))) // Display response
        .catch(error => displayResponse("response", "Error: " + error)); // Display error
}