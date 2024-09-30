import http from "k6/http";
import { check, group } from "k6";

// Base URL for the API
const baseUrl = "http://192.168.3.37:8080";

const serverNames = ["coupon", "order", "product", "stock", "user"];

export default function() {
    for (let i = 0; i < serverNames.length; i++) {
        let name = serverNames[i];
        group(`list ${name}`, function() {
            let response = http.get(baseUrl+`/api/v1/${name}/list?limit=5`);
            let body = JSON.parse(response.body)
            let ok = (response.status === 200) && (body.code === 0)
            check(response, {
                "status is ok": ok,
            });
        });
    }
}
