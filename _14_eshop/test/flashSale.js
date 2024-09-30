import http from "k6/http";
import { check, group } from "k6";

// Base URL for the API
const baseUrl = "http://192.168.3.37:8080";

function flashSale() {
    let submitOrderRequest = {
        userID: 1,
        productID: 1,
        amount: 88
    };

    let requestBody = JSON.stringify(submitOrderRequest);
    let response = http.post(baseUrl+"/api/v1/flashSale", requestBody, {
        headers: { "Content-Type": "application/json" },
    });

    let body = JSON.parse(response.body)
    let ok = (response.status === 200) && (body.code === 0)
    check(response, {
        "status is ok": ok,
    });
}

export default function() {
    group("test flash sale", () => {
        flashSale();
    });
}
